// Copyright 2019 The LUCI Authors.
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

package recorder

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/spanner"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/appstatus"
	"go.chromium.org/luci/resultdb/internal/recorder/chromium"
	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/tasks"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

// testResultBatchSizeMax is the maximum number of TestResults to include per transaction.
const testResultBatchSizeMax = 1000

var urlPrefixes = []string{"http://", "https://"}

// validateDeriveInvocationRequest returns an error if req is invalid.
func validateDeriveInvocationRequest(req *pb.DeriveInvocationRequest) error {
	if req.SwarmingTask == nil {
		return errors.Reason("swarming_task missing").Err()
	}

	if req.SwarmingTask.Hostname == "" {
		return errors.Reason("swarming_task.hostname missing").Err()
	}

	for _, prefix := range urlPrefixes {
		if strings.HasPrefix(req.SwarmingTask.Hostname, prefix) {
			return errors.Reason("swarming_task.hostname should not have prefix %q", prefix).Err()
		}
	}

	if req.SwarmingTask.Id == "" {
		return errors.Reason("swarming_task.id missing").Err()
	}

	return nil
}

// DeriveInvocation derives the invocation associated with the given swarming task.
//
// If the task is a dedup of another task, the invocation returned is the underlying one; otherwise,
// the invocation returned is associated with the swarming task itself.
func (s *recorderServer) DeriveInvocation(ctx context.Context, in *pb.DeriveInvocationRequest) (*pb.Invocation, error) {
	if err := validateDeriveInvocationRequest(in); err != nil {
		return nil, appstatus.BadRequest(err)
	}

	// Get the swarming service to use.
	swarmingURL := "https://" + in.SwarmingTask.Hostname
	swarmSvc, err := chromium.GetSwarmSvc(internal.HTTPClient(ctx), swarmingURL)
	if err != nil {
		return nil, errors.Annotate(err, "creating swarming client for %q", swarmingURL).Err()
	}

	// Get the swarming task, deduping if necessary.
	task, err := chromium.GetSwarmingTask(ctx, in.SwarmingTask.Id, swarmSvc)
	if err != nil {
		return nil, errors.Annotate(err, "getting swarming task %q on %q",
			in.SwarmingTask.Id, in.SwarmingTask.Hostname).Err()
	}
	if task, err = chromium.GetOriginTask(ctx, task, swarmSvc); err != nil {
		return nil, errors.Annotate(err, "getting origin for swarming task %q on %q",
			in.SwarmingTask.Id, in.SwarmingTask.Hostname).Err()
	}
	invID := chromium.GetInvocationID(task, in)

	client := span.Client(ctx)

	// Check if we even need to write this invocation: is it finalized?
	doWrite, err := shouldWriteInvocation(ctx, client.Single(), invID)
	switch {
	case err != nil:
		return nil, err
	case !doWrite:
		readTxn := client.ReadOnlyTransaction()
		defer readTxn.Close()
		return span.ReadInvocationFull(ctx, readTxn, invID)
	}

	// Otherwise, get the protos and prepare to write them to Spanner.
	logging.Infof(ctx, "Deriving task %q on %q", in.SwarmingTask.Id, in.SwarmingTask.Hostname)
	inv, results, err := chromium.DeriveProtosForWriting(ctx, task, in)
	if err != nil {
		return nil, errors.Annotate(err,
			"task %q on %q named %q", in.SwarmingTask.Id, in.SwarmingTask.Hostname, task.Name).Err()
	}
	if inv.FinalizeTime == nil {
		panic("missing inv.FinalizeTime")
	}
	inv.Deadline = inv.FinalizeTime

	// TODO(jchinlee): Validate invocation and results.

	// Write test results in batches concurrently, updating inv with the names of the invocations
	// that will be included.
	batchInvs, err := s.batchInsertTestResults(ctx, inv, results, testResultBatchSizeMax)
	if err != nil {
		return nil, err
	}
	inv.IncludedInvocations = batchInvs.Names()

	// Prepare mutations.
	ms := make([]*spanner.Mutation, 0, len(batchInvs)+2)
	ms = append(ms, span.InsertMap("Invocations", s.rowOfInvocation(ctx, inv, "", "")))
	for includedID := range batchInvs {
		ms = append(ms, span.InsertMap("IncludedInvocations", map[string]interface{}{
			"InvocationId":         invID,
			"IncludedInvocationId": includedID,
		}))
	}
	ms = append(ms, tasks.EnqueueBQExport(invID, s.DerivedInvBQTable, clock.Now(ctx).UTC()))

	_, err = span.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Check invocation state again.
		switch doWrite, err := shouldWriteInvocation(ctx, txn, invID); {
		case err != nil:
			return err
		case !doWrite:
			return nil
		default:
			return txn.BufferWrite(ms)
		}
	})

	return inv, err
}

func shouldWriteInvocation(ctx context.Context, txn span.Txn, id span.InvocationID) (bool, error) {
	state, err := span.ReadInvocationState(ctx, txn, id)
	s, _ := appstatus.Get(err)
	switch {
	case s.Code() == codes.NotFound:
		// No such invocation found means we may have to write it, so proceed.
		return true, nil

	case err != nil:
		return false, err

	case state != pb.Invocation_FINALIZED:
		return false, errors.Reason(
			"attempting to derive an existing non-finalized invocation").Err()
	}

	// The invocation exists and is finalized, so no need to write it.
	return false, nil
}

// batchInsertTestResults inserts the given TestResults in batches under container Invocations,
// returning container ids.
func (s *recorderServer) batchInsertTestResults(ctx context.Context, inv *pb.Invocation, trs []*pb.TestResult, batchSize int) (span.InvocationIDSet, error) {
	batches := batchTestResults(trs, batchSize)
	includedInvs := make(span.InvocationIDSet, len(batches))

	invID := span.MustParseInvocationName(inv.Name)
	eg, ctx := errgroup.WithContext(ctx)
	client := span.Client(ctx)
	for i, batch := range batches {
		i := i
		batch := batch

		batchID := batchInvocationID(invID, i)
		includedInvs.Add(batchID)

		eg.Go(func() error {
			muts := make([]*spanner.Mutation, 0, len(batch)+1)

			// Convert the container Invocation in the batch.
			batchInv := &pb.Invocation{
				Name:         batchID.Name(),
				State:        pb.Invocation_FINALIZED,
				CreateTime:   inv.CreateTime,
				FinalizeTime: inv.FinalizeTime,
				Deadline:     inv.Deadline,
			}
			muts = append(muts, span.InsertOrUpdateMap(
				"Invocations", s.rowOfInvocation(ctx, batchInv, "", "")),
			)

			// Convert the TestResults in the batch.
			for k, tr := range batch {
				muts = append(muts, insertOrUpdateTestResult(batchID, tr, k))
			}

			_, err := client.Apply(ctx, muts)
			return err
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return includedInvs, nil
}

// batchInvocationID returns an InvocationID for the Invocation containing the referenced batch.
func batchInvocationID(invID span.InvocationID, batchInd int) span.InvocationID {
	return span.InvocationID(fmt.Sprintf("%s::batch::%d", invID, batchInd))
}

// batchTestResults batches the given TestResults given the maximum batch size.
func batchTestResults(trs []*pb.TestResult, batchSize int) [][]*pb.TestResult {
	batches := make([][]*pb.TestResult, 0, len(trs)/batchSize+1)
	for len(trs) > 0 {
		end := batchSize
		if end > len(trs) {
			end = len(trs)
		}

		batches = append(batches, trs[:end])
		trs = trs[end:]
	}

	return batches
}
