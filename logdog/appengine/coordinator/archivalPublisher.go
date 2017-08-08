// Copyright 2016 The LUCI Authors.
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

package coordinator

import (
	"time"

	"go.chromium.org/luci/common/gcloud/pubsub"
	log "go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/logdog/api/endpoints/coordinator/services/v1"

	gcps "cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/proto"

	"golang.org/x/net/context"
)

// ArchivalPublisher is capable of publishing archival requests.
type ArchivalPublisher interface {
	// Close shutdowns this publisher instance, releasing all its resources.
	Close() error

	// Publish publishes the supplied ArchiveTask.
	Publish(context.Context, *logdog.ArchiveTask) error

	// NewPublishIndex returns a new publish index. Each publish index is unique
	// within its request.
	NewPublishIndex() uint64
}

type pubsubArchivalPublisher struct {
	// publisher is the client used to publish messages.
	publisher pubsub.Publisher

	// publishIndexFunc is a function that will return a unique publish index
	// for this request.
	publishIndexFunc func() uint64
}

func (p *pubsubArchivalPublisher) Close() error { return nil }

func (p *pubsubArchivalPublisher) Publish(c context.Context, t *logdog.ArchiveTask) error {
	d, err := proto.Marshal(t)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to marshal task.")
		return err
	}

	// TODO: Route this through some system (e.g., task queue, tumble) that can
	// impose a dispatch delay for the settle period.
	msg := gcps.Message{
		Data: d,
	}

	return retry.Retry(c, retry.Default, func() error {
		log.Fields{
			"project": t.Project,
			"hash":    t.Id,
			"key":     t.Key,
		}.Infof(c, "Publishing archival message for stream.")

		_, err := p.publisher.Publish(c, &msg)
		return err
	}, func(err error, d time.Duration) {
		log.Fields{
			log.ErrorKey: err,
			"delay":      d,
		}.Warningf(c, "Failed to publish task. Retrying...")
	})
}

func (p *pubsubArchivalPublisher) NewPublishIndex() uint64 {
	return p.publishIndexFunc()
}
