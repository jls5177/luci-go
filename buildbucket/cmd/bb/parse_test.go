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

package main

import (
	"testing"

	buildbucketpb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestParseCL(t *testing.T) {
	Convey("ParseCL", t, func() {

		Convey("Perfect", func() {
			actual, err := parseCL("https://chromium-review.googlesource.com/c/infra/luci/luci-go/+/1541677/7")
			So(err, ShouldBeNil)
			So(actual, ShouldResembleProto, &buildbucketpb.GerritChange{
				Host:     "chromium-review.googlesource.com",
				Project:  "infra/luci/luci-go",
				Change:   1541677,
				Patchset: 7,
			})
		})

		Convey("Hash", func() {
			actual, err := parseCL("https://chromium-review.googlesource.com/#/c/infra/luci/luci-go/+/1541677/7")
			So(err, ShouldBeNil)
			So(actual, ShouldResembleProto, &buildbucketpb.GerritChange{
				Host:     "chromium-review.googlesource.com",
				Project:  "infra/luci/luci-go",
				Change:   1541677,
				Patchset: 7,
			})
		})

		Convey("File", func() {
			actual, err := parseCL("https://chromium-review.googlesource.com/c/infra/luci/luci-go/+/1541677/7/buildbucket/cmd/bb/base_command.go")
			So(err, ShouldBeNil)
			So(actual, ShouldResembleProto, &buildbucketpb.GerritChange{
				Host:     "chromium-review.googlesource.com",
				Project:  "infra/luci/luci-go",
				Change:   1541677,
				Patchset: 7,
			})
		})

		Convey("No project", func() {
			actual, err := parseCL("https://chromium-review.googlesource.com/c/1541677/7")
			So(err, ShouldBeNil)
			So(actual, ShouldResembleProto, &buildbucketpb.GerritChange{
				Host:     "chromium-review.googlesource.com",
				Change:   1541677,
				Patchset: 7,
			})
		})

		Convey("No patchset", func() {
			actual, err := parseCL("https://chromium-review.googlesource.com/c/infra/luci/luci-go/+/1541677")
			So(err, ShouldBeNil)
			So(actual, ShouldResembleProto, &buildbucketpb.GerritChange{
				Host:    "chromium-review.googlesource.com",
				Project: "infra/luci/luci-go",
				Change:  1541677,
			})
		})
	})
}

func TestParseCommit(t *testing.T) {
	Convey("ParseCommit", t, func() {

		Convey("7a63166bfab5de38ddb2cb8e29aca756bdc2a28d", func() {
			actual, confirm, err := parseCommit("https://chromium.googlesource.com/infra/luci/luci-go/+/7a63166bfab5de38ddb2cb8e29aca756bdc2a28d")
			So(err, ShouldBeNil)
			So(confirm, ShouldBeFalse)
			So(actual, ShouldResembleProto, &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "infra/luci/luci-go",
				Ref:     "",
				Id:      "7a63166bfab5de38ddb2cb8e29aca756bdc2a28d",
			})
		})

		Convey("refs/heads/x", func() {
			actual, confirm, err := parseCommit("https://chromium.googlesource.com/infra/luci/luci-go/+/refs/heads/x")
			So(err, ShouldBeNil)
			So(confirm, ShouldBeFalse)
			So(actual, ShouldResembleProto, &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "infra/luci/luci-go",
				Ref:     "refs/heads/x",
				Id:      "",
			})
		})

		Convey("refs/heads/x/y", func() {
			actual, confirm, err := parseCommit("https://chromium.googlesource.com/infra/luci/luci-go/+/refs/heads/x/y")
			So(err, ShouldBeNil)
			So(confirm, ShouldBeTrue)
			So(actual, ShouldResembleProto, &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "infra/luci/luci-go",
				Ref:     "refs/heads/x",
				Id:      "",
			})
		})

		Convey("refs/branch-heads/x", func() {
			actual, confirm, err := parseCommit("https://chromium.googlesource.com/infra/luci/luci-go/+/refs/branch-heads/x")
			So(err, ShouldBeNil)
			So(confirm, ShouldBeTrue)
			So(actual, ShouldResembleProto, &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "infra/luci/luci-go",
				Ref:     "refs/branch-heads/x",
				Id:      "",
			})
		})

		Convey("master", func() {
			actual, confirm, err := parseCommit("https://chromium.googlesource.com/infra/luci/luci-go/+/master")
			So(err, ShouldBeNil)
			So(confirm, ShouldBeTrue)
			So(actual, ShouldResembleProto, &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "infra/luci/luci-go",
				Ref:     "refs/heads/master",
				Id:      "",
			})
		})

		Convey("refs/x", func() {
			actual, confirm, err := parseCommit("https://chromium.googlesource.com/infra/luci/luci-go/+/refs/x")
			So(err, ShouldBeNil)
			So(confirm, ShouldBeTrue)
			So(actual, ShouldResembleProto, &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "infra/luci/luci-go",
				Ref:     "refs/x",
				Id:      "",
			})
		})
	})
}