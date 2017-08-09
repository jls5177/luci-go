// Copyright 2017 The LUCI Authors.
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

package ensure

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/cipd/client/cipd/common"
)

func f(lines ...string) string {
	return strings.Join(lines, "\n")
}

func p(pkg, ver string) common.Pin {
	return common.Pin{PackageName: pkg, InstanceID: ver}
}

var goodEnsureFiles = []struct {
	name   string
	file   string
	expect *ResolvedFile
}{
	{
		"old_style",
		f(
			"# comment",
			"",
			"path/to/package deadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
			"path/to/other_package some_tag:version",
			"path/to/yet_another a_ref",
		),
		&ResolvedFile{"", common.PinSliceBySubdir{
			"": {
				p("path/to/package", "deadbeefdeadbeefdeadbeefdeadbeefdeadbeef"),
				p("path/to/other_package", "some_tag:version"),
				p("path/to/yet_another", "a_ref"),
			},
		}},
	},

	{
		"templates",
		f(
			"path/to/package/${os}-${arch} latest",
			"path/to/other/${platform} latest",
		),
		&ResolvedFile{"", common.PinSliceBySubdir{
			"": {
				p("path/to/package/test_os-test_arch", "latest"),
				p("path/to/other/test_os-test_arch", "latest"),
			},
		}},
	},

	{
		"optional_templates",
		f(
			"path/to/package/${os}-${arch=neep,test_arch} latest",
			"path/to/other/${platform=test_os-test_arch} latest",
		),
		&ResolvedFile{"", common.PinSliceBySubdir{
			"": {
				p("path/to/package/test_os-test_arch", "latest"),
				p("path/to/other/test_os-test_arch", "latest"),
			},
		}},
	},

	{
		"optional_templates_no_match",
		f(
			"path/to/package/${os=spaz}-${arch=neep,test_arch} latest",
			"path/to/package/${platform=neep-foo} latest",
		),
		&ResolvedFile{"", common.PinSliceBySubdir{}},
	},

	{
		"Subdir directives",
		f(
			"some/package latest",
			"",
			"@Subdir a/subdir with spaces",
			"some/package canary",
			"some/other/package tag:value",
			"",
			"@Subdir", // reset back to empty
			"cool/package beef",
			"@Subdir ${platform}", // template expansion
			"some/package latest",
			"@Subdir ${platform}/subdir",
			"some/package canary",
		),
		&ResolvedFile{"", common.PinSliceBySubdir{
			"": {
				p("some/package", "latest"),
				p("cool/package", "beef"),
			},
			"a/subdir with spaces": {
				p("some/package", "canary"),
				p("some/other/package", "tag:value"),
			},
			common.CurrentOS() + "-" + common.CurrentArchitecture(): {
				p("some/package", "latest"),
			},
			common.CurrentOS() + "-" + common.CurrentArchitecture() + "/subdir": {
				p("some/package", "canary"),
			},
		}},
	},

	{
		"ServiceURL setting",
		f(
			"$ServiceURL https://cipd.example.com/path/to/thing",
			"",
			"some/package version",
		),
		&ResolvedFile{"https://cipd.example.com/path/to/thing", common.PinSliceBySubdir{
			"": {
				p("some/package", "version"),
			},
		}},
	},

	{
		"empty",
		"",
		&ResolvedFile{},
	},

	{
		"wacky spaces",
		f(
			"path/to/package           latest",
			"tabs/to/package\t\t\t\tlatest",
			"\ttabs/and/spaces  \t  \t  \tlatest   \t",
		),
		&ResolvedFile{"", common.PinSliceBySubdir{
			"": {
				p("path/to/package", "latest"),
				p("tabs/to/package", "latest"),
				p("tabs/and/spaces", "latest"),
			},
		}},
	},
}

func testResolver(pkg, vers string) (common.Pin, error) {
	if strings.Contains(vers, "error") {
		return p("", ""), errors.New("testResolver returned error")
	}
	return p(pkg, vers), nil
}

func TestGoodEnsureFiles(t *testing.T) {
	t.Parallel()

	Convey("good ensure files", t, func() {
		for _, tc := range goodEnsureFiles {
			Convey(tc.name, func() {
				buf := bytes.NewBufferString(tc.file)
				f, err := ParseFile(buf)
				So(err, ShouldBeNil)
				rf, err := f.ResolveWith(testResolver, map[string]string{
					"os":       "test_os",
					"arch":     "test_arch",
					"platform": "test_os-test_arch",
				})
				So(err, ShouldBeNil)
				So(rf, ShouldResemble, tc.expect)
			})
		}
	})
}
