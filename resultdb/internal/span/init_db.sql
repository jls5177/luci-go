-- Copyright 2019 The LUCI Authors.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

--------------------------------------------------------------------------------
-- This script initializes a ResultDB Spanner database.

-- Stores the invocations.
-- This is the root table for much of the other data and tables, which define the
-- hierarchy (dependency graph, subsets of interest) for invocations.
CREATE TABLE Invocations (
  -- Identifies an invocation.
  -- Format: "${hex(sha256(user_provided_id)[:8])}:${user_provided_id}".
  InvocationId STRING(MAX) NOT NULL,

  -- A random value in [0, InvocationShards) where InvocationShards constant is
  -- defined in code.
  -- Used in global secondary indexes, to prevent hot spots.
  -- The maximum value of ShardId in Spanner can be determined by querying the
  -- very first row in InvocationsByExpiration index.
  ShardId INT64 NOT NULL,

  -- Invocation state, see InvocationState in invocation.proto
  State INT64 NOT NULL,

  -- Security realm this invocation belongs to.
  -- Used to enforce ACLs.
  Realm STRING(64) NOT NULL,

  -- When to delete the invocation from the table.
  InvocationExpirationTime TIMESTAMP NOT NULL,

  -- When to delete expected test results from this invocation.
  -- When expected results are removed, this column is set to NULL.
  ExpectedTestResultsExpirationTime TIMESTAMP,

  -- Secret token that a client must provide to mutate this or any of the
  -- interleaved rows.
  UpdateToken STRING(64),

  -- When the invocation was created.
  CreateTime TIMESTAMP NOT NULL,

  -- When the invocation was finalized.
  FinalizeTime TIMESTAMP,

  -- When to force invocation finalization with state INTERRUPTED.
  Deadline TIMESTAMP NOT NULL,

  -- Variant definition pairs common for all test results in this invocation.
  -- Each string is a colon-separated key-value pair.
  -- A particular test result can have additional key-value pairs.
  BaseTestVariant ARRAY<STRING(MAX)>,

  -- List of colon-separated key-value tags.
  -- Corresponds to Invocation.tags in invocation.proto.
  -- For each tag here, there is a row in the InvocationsByTag table.
  Tags ARRAY<STRING(MAX)>,

  -- Value of CreateInvocationRequest.request_id.
  -- Used to dedup invocation creation requests.
  CreateRequestId STRING(MAX),
) PRIMARY KEY (InvocationId);

-- Index of invocations by expiration time.
-- Used by a cron job that periodically removes expired invocations.
CREATE INDEX InvocationsByInvocationExpiration
  ON Invocations (ShardId DESC, InvocationExpirationTime, InvocationId);

-- Index of invocations by expected test result expiration.
-- Used by a cron job that periodically removes expected test results.
CREATE NULL_FILTERED INDEX InvocationsByExpectedTestResultsExpiration
  ON Invocations (ShardId DESC, ExpectedTestResultsExpirationTime, InvocationId);

-- Index of invocations by a tag.
CREATE TABLE InvocationsByTag (
  -- Format: "${sha256_hex(tag)}_${key}:${value}".
  TagId STRING(MAX) NOT NULL,

  -- FK to Invocations.InvocationId.
  InvocationId STRING(MAX) NOT NULL
) PRIMARY KEY (TagId, InvocationId);

-- Stores ids of invocations included in another invocation.
-- Interleaved in Invocations table.
CREATE TABLE IncludedInvocations (
  -- ID of the including invocation, the "source" node of the edge.
  InvocationId STRING(MAX) NOT NULL,

  -- ID of the included invocation, the "target" node of the edge.
  IncludedInvocationId STRING(MAX) NOT NULL
) PRIMARY KEY (InvocationId, IncludedInvocationId),
  INTERLEAVE IN PARENT Invocations ON DELETE CASCADE;

-- Stores test results. Interleaved in Invocations.
CREATE TABLE TestResults (
  -- ID of the parent Invocations row.
  InvocationId STRING(MAX) NOT NULL,

  -- Unique identifier of the test,
  -- see also TestResult.test_path in test_result.proto.
  TestPath STRING(MAX) NOT NULL,

  -- A suffix for PK to allow multiple test results for the same test path in
  -- a given invocation.
  -- Generated on the server.
  ResultId STRING(MAX) NOT NULL,

  -- Extra variant pairs, in addition to parent Invocation.BaseTestVariant.
  -- See also TestResult.extra_variant_pairs in test_result.proto.
  ExtraVariantPairs ARRAY<STRING(MAX)>,

  -- Last time this row was modified.
  -- Given that we only create and delete row, for an existing row this equals
  -- row creation time.
  CommitTimestamp TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),

  -- Whether the test status was unexpected
  -- MUST be either NULL or True, to keep null-filtered index below thin.
  IsUnexpected BOOL,

  -- Test status, see TestStatus in test_result.proto.
  Status INT64 NOT NULL,

  -- Snappy-encoded summary of the test result for humans, in Markdown.
  SummaryMarkdown BYTES(MAX),

  -- When the test execution started.
  StartTime TIMESTAMP,

  -- How long the test execution took, in microseconds.
  RunDurationUsec INT64,

  -- Tags associated with the test result, for example GTest-specific test
  -- status.
  Tags ARRAY<STRING(MAX)>,

  -- Input artifacts, see also TestResult.input_artifacts in test_result.proto.
  -- Encoded as resultdb.internal.Artifacts message.
  InputArtifacts BYTES(MAX),

  -- Output artifacts, see also TestResult.output_artifacts in test_result.proto.
  -- Encoded as resultdb.internal.Artifacts message.
  OutputArtifacts BYTES(MAX)
) PRIMARY KEY (InvocationId, TestPath, ResultId),
  INTERLEAVE IN PARENT Invocations ON DELETE CASCADE;

-- Unexpected test results for each invocation.
-- It is significantly smaller (<2%) than TestResult table and should be used
-- for most queries.
-- It includes TestPath to be able to find all unexpected test result with a
-- given test path or a test path prefix.
CREATE NULL_FILTERED INDEX UnexpectedTestResults
  ON TestResults (InvocationId, TestPath, IsUnexpected),
  INTERLEAVE IN Invocations;


-- Stores test exonerations, see TestExoneration in test_result.proto
CREATE TABLE TestExonerations (
  -- ID of the parent Invocations row.
  InvocationId STRING(MAX) NOT NULL,

  -- The exoneration applies only to test results with this exact test path.
  -- This is a foreign key to TestResults.TestPath column.
  TestPath STRING(MAX) NOT NULL,

  -- Server-generated exoneration ID.
  -- Uniquely identifies a test exoneration within an invocation.
  --
  -- Starts with "{hex(sha256(join(sorted('{p}\n' for p in Variant))))}:".
  -- The prefix can be used to reduce scanning for test exonerations for a
  -- particular test variant.
  ExonerationId STRING(MAX) NOT NULL,

  -- The exoneration applies only to test results with this exact test variant.
  Variant ARRAY<STRING(MAX)> NOT NULL,

  -- Explanation of the exoneration for humans, in Markdown.
  ExplanationMarkdown STRING(MAX)
) PRIMARY KEY (InvocationId, TestPath, ExonerationId),
  INTERLEAVE IN PARENT Invocations ON DELETE CASCADE;
