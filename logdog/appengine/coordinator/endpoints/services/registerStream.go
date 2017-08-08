// Copyright 2015 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"crypto/subtle"
	"time"

	"go.chromium.org/luci/common/clock"
	log "go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/logdog/api/endpoints/coordinator/services/v1"
	"go.chromium.org/luci/logdog/api/logpb"
	"go.chromium.org/luci/logdog/appengine/coordinator"
	"go.chromium.org/luci/logdog/appengine/coordinator/endpoints"
	"go.chromium.org/luci/logdog/appengine/coordinator/tasks"
	"go.chromium.org/luci/logdog/common/types"

	ds "go.chromium.org/gae/service/datastore"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func buildLogStreamState(ls *coordinator.LogStream, lst *coordinator.LogStreamState) *logdog.LogStreamState {
	st := logdog.LogStreamState{
		ProtoVersion:  ls.ProtoVersion,
		Secret:        lst.Secret,
		TerminalIndex: lst.TerminalIndex,
		Archived:      lst.ArchivalState().Archived(),
		Purged:        ls.Purged,
	}
	if !lst.Terminated() {
		st.TerminalIndex = -1
	}
	return &st
}

// RegisterStream is an idempotent stream state register operation.
//
// Successive operations will succeed if they have the correct secret for their
// registered stream, regardless of whether the contents of their request match
// the currently registered state.
func (s *server) RegisterStream(c context.Context, req *logdog.RegisterStreamRequest) (*logdog.RegisterStreamResponse, error) {
	var path types.StreamPath

	// Unmarshal the serialized protobuf.
	var desc logpb.LogStreamDescriptor
	switch req.ProtoVersion {
	case logpb.Version:
		if err := proto.Unmarshal(req.Desc, &desc); err != nil {
			log.Fields{
				log.ErrorKey:   err,
				"protoVersion": req.ProtoVersion,
			}.Errorf(c, "Failed to unmarshal descriptor protobuf.")
			return nil, grpcutil.Errf(codes.InvalidArgument, "Failed to unmarshal protobuf.")
		}

	default:
		log.Fields{
			"protoVersion": req.ProtoVersion,
		}.Errorf(c, "Unrecognized protobuf version.")
		return nil, grpcutil.Errf(codes.InvalidArgument, "Unrecognized protobuf version: %q", req.ProtoVersion)
	}

	path = desc.Path()
	logStreamID := coordinator.LogStreamID(path)
	log.Fields{
		"project":       req.Project,
		"path":          path,
		"prospectiveID": logStreamID,
		"terminalIndex": req.TerminalIndex,
	}.Infof(c, "Registration request for log stream.")

	if err := desc.Validate(true); err != nil {
		return nil, grpcutil.Errf(codes.InvalidArgument, "Invalid log stream descriptor: %s", err)
	}
	prefix, _ := path.Split()

	// Load our service and project configs.
	cfg, err := coordinator.GetServices(c).Config(c)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to load configuration.")
		return nil, grpcutil.Internal
	}

	pcfg, err := coordinator.CurrentProjectConfig(c)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to load current project configuration.")
		return nil, grpcutil.Internal
	}

	// Load our Prefix. It must be registered.
	pfx := &coordinator.LogPrefix{ID: coordinator.LogPrefixID(prefix)}
	if err := ds.Get(c, pfx); err != nil {
		log.Fields{
			log.ErrorKey: err,
			"id":         pfx.ID,
			"prefix":     prefix,
		}.Errorf(c, "Failed to load log stream prefix.")
		if err == ds.ErrNoSuchEntity {
			return nil, grpcutil.Errf(codes.FailedPrecondition, "prefix is not registered")
		}
		return nil, grpcutil.Internal
	}

	// If we're past prefix's expiration, reject this stream.
	//
	// If the prefix doesn't have an expiration, use its creation time and apply
	// the maximum expiration.
	expirationTime := pfx.Expiration
	if expirationTime.IsZero() {
		expiration := endpoints.MinDuration(cfg.Coordinator.PrefixExpiration, pcfg.PrefixExpiration)
		if expiration > 0 {
			expirationTime = pfx.Created.Add(expiration)
		}
	}
	if now := clock.Now(c); expirationTime.IsZero() || !now.Before(expirationTime) {
		log.Fields{
			"prefix":     pfx.Prefix,
			"expiration": expirationTime,
		}.Errorf(c, "The log stream Prefix has expired.")
		return nil, grpcutil.Errf(codes.FailedPrecondition, "prefix has expired")
	}

	// The prefix secret must match the request secret. If it does, we know this
	// is a legitimate registration attempt.
	if subtle.ConstantTimeCompare(pfx.Secret, req.Secret) != 1 {
		log.Errorf(c, "Request secret does not match prefix secret.")
		return nil, grpcutil.Errf(codes.InvalidArgument, "invalid secret")
	}

	// Check for registration, and that the prefix did not expire
	// (non-transactional).
	ls := &coordinator.LogStream{ID: logStreamID}
	lst := ls.State(c)

	if err := ds.Get(c, ls, lst); err != nil {
		if !anyNoSuchEntity(err) {
			log.WithError(err).Errorf(c, "Failed to check for log stream.")
			return nil, err
		}

		// The stream does not exist. Proceed with transactional registration.
		err = ds.RunInTransaction(c, func(c context.Context) error {
			// Load our state and stream (transactional).
			switch err := ds.Get(c, ls, lst); {
			case err == nil:
				// The stream is already registered.
				return nil

			case !anyNoSuchEntity(err):
				log.WithError(err).Errorf(c, "Failed to check for stream registration (transactional).")
				return err
			}

			// The stream is not yet registered.
			log.Infof(c, "Registering new log stream.")

			// Construct our LogStreamState.
			now := clock.Now(c).UTC()
			lst.Created = now
			lst.Updated = now
			lst.Secret = pfx.Secret // Copy Prefix Secret to reduce datastore Gets.

			// Construct our LogStream.
			ls.Created = now
			ls.ProtoVersion = req.ProtoVersion

			if err := ls.LoadDescriptor(&desc); err != nil {
				log.Fields{
					log.ErrorKey: err,
				}.Errorf(c, "Failed to load descriptor into LogStream.")
				return grpcutil.Errf(codes.InvalidArgument, "Failed to load descriptor.")
			}

			// If our registration request included a terminal index, terminate the
			// log stream state as well.
			if req.TerminalIndex >= 0 {
				log.Fields{
					"terminalIndex": req.TerminalIndex,
				}.Debugf(c, "Registration request included terminal index.")

				lst.TerminalIndex = req.TerminalIndex
				lst.TerminatedTime = now
			} else {
				lst.TerminalIndex = -1
			}

			if err := ds.Put(c, ls, lst); err != nil {
				log.Fields{
					log.ErrorKey: err,
				}.Errorf(c, "Failed to Put LogStream.")
				return grpcutil.Internal
			}

			// Add a named delayed task queue task to archive this stream if it's not
			// archived yet.
			//
			// If the registration did not include a terminal index, this will be our
			// pessimistic archival request, scheduled on registration to catch
			// streams that don't expire. This task will be replaced by the optimistic
			// archival task when/if the stream is terminated via TerminateStream.
			//
			// If the registration included a terminal index, apply our standard
			// parameters to the archival. Since TerminateStream will not be called,
			// this will be our formal optimistic archival task.
			params := standardArchivalParams(cfg, pcfg)

			var (
				delay       time.Duration
				archivalTag logdog.ArchiveDispatchTask_Tag
			)
			if req.TerminalIndex < 0 {
				// No terminal index, schedule pessimistic cleanup archival.
				delay = params.CompletePeriod
				archivalTag = logdog.ArchiveDispatchTask_EXPIRED

				// For cleanup, we instruct the archival to not wait any longer or
				// allow the stream time extra time to become complete (archive as-is).
				params.SettleDelay = 0
				params.CompletePeriod = 0

				log.Fields{
					"deadline": delay,
				}.Debugf(c, "Scheduling cleanup archival task.")
			} else {
				// Schedule this task to execute after our settle delay.
				delay = params.SettleDelay
				archivalTag = logdog.ArchiveDispatchTask_TERMINATED

				log.Fields{
					"settleDelay":    params.SettleDelay,
					"completePeriod": params.CompletePeriod,
					"scheduledAt":    delay,
				}.Debugf(c, "Scheduling archival task.")
			}

			if err := tasks.CreateArchivalTask(c, logStreamID, archivalTag, delay, params); err != nil {
				log.WithError(err).Errorf(c, "Failed to create archival task.")
				return grpcutil.Internal
			}

			return nil
		}, nil)
		if err != nil {
			log.Fields{
				log.ErrorKey: err,
			}.Errorf(c, "Failed to register LogStream.")
			return nil, err
		}
	}

	return &logdog.RegisterStreamResponse{
		Id:    string(ls.ID),
		State: buildLogStreamState(ls, lst),
	}, nil
}
