// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package swarming

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/luci/gae/impl/memory"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	swarming "github.com/luci/luci-go/common/api/swarming/swarming/v1"
	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/errors"
	miloProto "github.com/luci/luci-go/common/proto/milo"
	"github.com/luci/luci-go/logdog/api/endpoints/coordinator/logs/v1"
	"github.com/luci/luci-go/logdog/api/logpb"
	"github.com/luci/luci-go/logdog/client/coordinator"
	"github.com/luci/luci-go/milo/api/resp"
	"github.com/luci/luci-go/milo/appengine/settings"
	"github.com/luci/luci-go/server/templates"
)

type testCase struct {
	name string

	swarmResult string
	swarmOutput string
	annotations string
}

func getTestCases() []*testCase {
	testCases := make(map[string]*testCase)

	// References "milo/appengine/swarming/testdata".
	testdata := filepath.Join("..", "swarming", "testdata")
	f, err := ioutil.ReadDir(testdata)
	if err != nil {
		panic(err)
	}
	for _, fi := range f {
		fileName := fi.Name()
		parts := strings.SplitN(fileName, ".", 2)

		name := parts[0]
		tc := testCases[name]
		if tc == nil {
			tc = &testCase{name: name}
			testCases[name] = tc
		}

		switch {
		case len(parts) == 1:
			tc.swarmOutput = fileName
		case parts[1] == "swarm":
			tc.swarmResult = fileName
		case parts[1] == "pb.txt":
			tc.annotations = fileName
		}
	}

	// Order test cases by name.
	names := make([]string, 0, len(testCases))
	for name := range testCases {
		names = append(names, name)
	}
	sort.Strings(names)

	results := make([]*testCase, len(names))
	for i, name := range names {
		results[i] = testCases[name]
	}

	return results
}

func (tc *testCase) getContent(name string) []byte {
	if name == "" {
		return nil
	}

	// ../swarming below assumes that
	// - this code is not executed by tests outside of this dir
	// - this dir is a sibling of frontend dir
	path := filepath.Join("..", "swarming", "testdata", name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("failed to read [%s]: %s", path, err))
	}
	return data
}

func (tc *testCase) getSwarmingResult() *swarming.SwarmingRpcsTaskResult {
	var sr swarming.SwarmingRpcsTaskResult
	data := tc.getContent(tc.swarmResult)
	if err := json.Unmarshal(data, &sr); err != nil {
		panic(fmt.Errorf("failed to unmarshal [%s]: %s", tc.swarmResult, err))
	}
	return &sr
}

func (tc *testCase) getSwarmingOutput() string {
	return string(tc.getContent(tc.swarmOutput))
}

func (tc *testCase) getAnnotation() *miloProto.Step {
	var step miloProto.Step
	data := tc.getContent(tc.annotations)
	if err := proto.UnmarshalText(string(data), &step); err != nil {
		panic(fmt.Errorf("failed to unmarshal text protobuf [%s]: %s", tc.annotations, err))
	}
	return &step
}

type debugSwarmingService struct {
	tc *testCase
}

func (svc debugSwarmingService) getHost() string { return "example.com" }

func (svc debugSwarmingService) getSwarmingResult(c context.Context, taskID string) (
	*swarming.SwarmingRpcsTaskResult, error) {

	return svc.tc.getSwarmingResult(), nil
}

func (svc debugSwarmingService) getSwarmingRequest(c context.Context, taskID string) (
	*swarming.SwarmingRpcsTaskRequest, error) {

	return nil, errors.New("not implemented")
}

func (svc debugSwarmingService) getTaskOutput(c context.Context, taskID string) (string, error) {
	return svc.tc.getSwarmingOutput(), nil
}

// TestableLog is a subclass of Log that interfaces with TestableHandler and
// includes sample test data.
type TestableLog struct{ Log }

// TestableBuild is a subclass of Build that interfaces with TestableHandler and
// includes sample test data.
type TestableBuild struct{ Build }

// TestData returns sample test data.
func (l TestableLog) TestData() []settings.TestBundle {
	return []settings.TestBundle{
		{
			Description: "Basic log",
			Data: templates.Args{
				"Log":    "This is the log",
				"Closed": true,
			},
		},
	}
}

// testLogDogClient is a minimal functional LogsClient implementation.
//
// It retains its latest input parameter and returns its configured err (if not
// nil) or resp.
type testLogDogClient struct {
	logdog.LogsClient

	req  interface{}
	resp interface{}
	err  error
}

func (tc *testLogDogClient) Tail(ctx context.Context, in *logdog.TailRequest, opts ...grpc.CallOption) (
	*logdog.GetResponse, error) {

	tc.req = in
	if tc.err != nil {
		return nil, tc.err
	}
	return tc.resp.(*logdog.GetResponse), nil
}

func logDogClientFunc(tc *testCase) func(context.Context, string) (*coordinator.Client, error) {
	return func(c context.Context, host string) (*coordinator.Client, error) {
		return &coordinator.Client{
			Host: "example.com",
			C: &testLogDogClient{
				resp: datagramGetResponse("testproject", "foo/bar", tc.getAnnotation()),
			},
		}, nil
	}
}

func datagramGetResponse(project, prefix string, msg proto.Message) *logdog.GetResponse {
	data, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return &logdog.GetResponse{
		Project: project,
		Desc: &logpb.LogStreamDescriptor{
			Prefix:      prefix,
			ContentType: miloProto.ContentTypeAnnotations,
			StreamType:  logpb.StreamType_DATAGRAM,
		},
		State: &logdog.LogStreamState{},
		Logs: []*logpb.LogEntry{
			{
				Content: &logpb.LogEntry_Datagram{
					Datagram: &logpb.Datagram{
						Data: data,
					},
				},
			},
		},
	}
}

// TestData returns sample test data.
func (b TestableBuild) TestData() []settings.TestBundle {
	basic := resp.MiloBuild{
		Summary: resp.BuildComponent{
			Label:    "Test swarming build",
			Status:   resp.Success,
			Started:  time.Date(2016, 1, 2, 15, 4, 5, 999999999, time.UTC),
			Finished: time.Date(2016, 1, 2, 15, 4, 6, 999999999, time.UTC),
			Duration: time.Second,
		},
	}
	results := []settings.TestBundle{
		{
			Description: "Basic successful build",
			Data:        templates.Args{"Build": basic},
		},
	}
	c := context.Background()
	c, _ = testclock.UseTime(c, time.Date(2016, time.March, 14, 11, 0, 0, 0, time.UTC))
	c = memory.Use(c)

	for _, tc := range getTestCases() {
		svc := debugSwarmingService{tc}
		bl := buildLoader{
			logDogClientFunc: logDogClientFunc(tc),
		}

		build, err := bl.swarmingBuildImpl(c, svc, "foo", tc.name)
		if err != nil {
			panic(fmt.Errorf("Error while processing %s: %s", tc.name, err))
		}
		results = append(results, settings.TestBundle{
			Description: tc.name,
			Data:        templates.Args{"Build": build},
		})
	}
	return results
}
