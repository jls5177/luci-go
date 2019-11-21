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

package span

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes"
	durpb "github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/grpcutil"

	"go.chromium.org/luci/resultdb/internal/pagination"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

// MustParseTestResultName retrieves the invocation ID, unescaped test path, and
// result ID.
// Panics if the name is invalid. Useful for situations when name was already
// validated.
func MustParseTestResultName(name string) (invID InvocationID, testPath, resultID string) {
	invIDStr, testPath, resultID, err := pbutil.ParseTestResultName(name)
	if err != nil {
		panic(err)
	}
	invID = InvocationID(invIDStr)
	return
}

// ReadTestResult reads specified TestResult within the transaction.
// If the TestResult does not exist, the returned error is annotated with
// NotFound GRPC code.
func ReadTestResult(ctx context.Context, txn Txn, name string) (*pb.TestResult, error) {
	invID, testPath, resultID := MustParseTestResultName(name)
	tr := &pb.TestResult{
		Name:     name,
		TestPath: testPath,
		ResultId: resultID,
		Expected: true,
	}

	var maybeUnexpected spanner.NullBool
	var micros int64
	var summaryMarkdown Snappy
	err := ReadRow(ctx, txn, "TestResults", invID.Key(testPath, resultID), map[string]interface{}{
		"Variant":         &tr.Variant,
		"IsUnexpected":    &maybeUnexpected,
		"Status":          &tr.Status,
		"SummaryMarkdown": &summaryMarkdown,
		"StartTime":       &tr.StartTime,
		"RunDurationUsec": &micros,
		"Tags":            &tr.Tags,
		"InputArtifacts":  &tr.InputArtifacts,
		"OutputArtifacts": &tr.OutputArtifacts,
	})
	switch {
	case spanner.ErrCode(err) == codes.NotFound:
		return nil, errors.Reason("%q not found", name).
			InternalReason("%s", err).
			Tag(grpcutil.NotFoundTag).
			Err()
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch %q", name).Err()
	}

	tr.SummaryMarkdown = string(summaryMarkdown)
	populateExpectedField(tr, maybeUnexpected)
	populateDurationField(tr, micros)
	return tr, nil
}

// TestResultQuery specifies test results to fetch.
type TestResultQuery struct {
	InvocationIDs []InvocationID
	Predicate     *pb.TestResultPredicate // Predicate.Invocation must be nil.
	PageSize      int                     // must be positive
	PageToken     string
}

// QueryTestResults reads test results matching the predicate.
// Returned test results from the same invocation are contiguous.
func QueryTestResults(ctx context.Context, txn *spanner.ReadOnlyTransaction, q TestResultQuery) (trs []*pb.TestResult, nextPageToken string, err error) {
	switch {
	case q.PageSize <= 0:
		panic("PageSize <= 0")
	case q.Predicate.GetInvocation() != nil:
		panic("q.Predicate.Invocation != nil")
	}

	// TODO(nodir): stop using squirrel.
	sql := squirrel.
		Select(
			"InvocationId",
			"TestPath",
			"ResultId",
			"Variant",
			"IsUnexpected",
			"Status",
			"SummaryMarkdown",
			"StartTime",
			"RunDurationUsec",
			"Tags",
			"InputArtifacts",
			"OutputArtifacts",
		).
		From("TestResults").
		Where("InvocationId IN UNNEST(@invIDs)").
		OrderBy("InvocationId", "TestPath", "ResultId").
		Limit(uint64(q.PageSize))

	queryParams := map[string]interface{}{
		"invIDs": q.InvocationIDs,
	}

	// Set start position if requested.
	queryParams["afterInvocationID"],
		queryParams["afterTestPath"],
		queryParams["afterResultID"],
		err = parseTestObjectPageToken(q.PageToken)
	if err != nil {
		return
	}
	sql = sql.Where(`(
			(InvocationId > @afterInvocationID)
			OR (InvocationId = @afterInvocationID AND TestPath > @afterTestPath)
			OR (InvocationId = @afterInvocationID AND TestPath = @afterTestPath AND ResultId > @afterResultID)
	)`)

	if q.Predicate.GetTestPath() != nil {
		// TODO(nodir): add support for q.Predicate.TestPath.
		return nil, "", grpcutil.Unimplemented
	}
	if q.Predicate.GetVariant() != nil {
		// TODO(nodir): add support for q.Predicate.Variant.
		return nil, "", grpcutil.Unimplemented
	}
	if q.Predicate.GetExpectancy() != pb.TestResultPredicate_ALL {
		// TODO(nodir): add support for q.Predicate.Expectancy.
		return nil, "", grpcutil.Unimplemented
	}

	sqlStr, _ := sql.MustSql()
	st := spanner.NewStatement(sqlStr)
	st.Params = queryParams

	trs = make([]*pb.TestResult, 0, q.PageSize)
	var summaryMarkdown Snappy
	var b Buffer
	err = query(ctx, txn, st, func(row *spanner.Row) error {
		var invID InvocationID
		var maybeUnexpected spanner.NullBool
		var micros int64
		tr := &pb.TestResult{}
		err = b.FromSpanner(row,
			&invID,
			&tr.TestPath,
			&tr.ResultId,
			&tr.Variant,
			&maybeUnexpected,
			&tr.Status,
			&summaryMarkdown,
			&tr.StartTime,
			&micros,
			&tr.Tags,
			&tr.InputArtifacts,
			&tr.OutputArtifacts,
		)
		if err != nil {
			return err
		}

		tr.Name = pbutil.TestResultName(string(invID), tr.TestPath, tr.ResultId)
		tr.SummaryMarkdown = string(summaryMarkdown)
		populateExpectedField(tr, maybeUnexpected)
		populateDurationField(tr, micros)

		trs = append(trs, tr)
		return nil
	})
	if err != nil {
		trs = nil
		return
	}

	// If we got pageSize results, then we haven't exhausted the collection and
	// need to return the next page token.
	if len(trs) == q.PageSize {
		last := trs[q.PageSize-1]
		invID, testPath, resultID := MustParseTestResultName(last.Name)
		nextPageToken = pagination.Token(string(invID), testPath, resultID)
	}
	return
}

func populateDurationField(tr *pb.TestResult, micros int64) {
	tr.Duration = FromMicros(micros)
}

func populateExpectedField(tr *pb.TestResult, maybeUnexpected spanner.NullBool) {
	tr.Expected = !maybeUnexpected.Valid || !maybeUnexpected.Bool
}

// ToMicros converts a duration.Duration proto to microseconds.
func ToMicros(d *durpb.Duration) int64 {
	return 1e6*d.Seconds + int64(1e-3*float64(d.Nanos))
}

// FromMicros converts microseconds to a duration.Duration proto.
func FromMicros(micros int64) *durpb.Duration {
	return ptypes.DurationProto(time.Duration(1e3 * micros))
}

// parseTestObjectPageToken parses the page token into invocation ID, test path
// and a test object id.
func parseTestObjectPageToken(pageToken string) (inv InvocationID, testPath, objID string, err error) {
	switch pos, tokErr := pagination.ParseToken(pageToken); {
	case tokErr != nil:
		err = encapsulatePageTokenError(tokErr)

	case pos == nil:

	case len(pos) != 3:
		err = encapsulatePageTokenError(errors.Reason("expected 3 position strings, got %q", pos).Err())

	default:
		inv = InvocationID(pos[0])
		testPath = pos[1]
		objID = pos[2]
	}

	return
}

// encapsulatePageTokenError returns a generic error message that a page token
// is invalid and records err as an internal error.
// The returned error is anontated with INVALID_ARUGMENT code.
func encapsulatePageTokenError(err error) error {
	return errors.Reason("invalid page_token").InternalReason("%s", err).Tag(grpcutil.InvalidArgumentTag).Err()
}