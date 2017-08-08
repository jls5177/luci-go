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

package main

import (
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"go.chromium.org/luci/common/errors"
	log "go.chromium.org/luci/common/logging"
	"golang.org/x/net/context"
)

func unmarshalTextProtobuf(path string, msg proto.Message) error {
	data, err := ioutil.ReadFile(path)
	switch {
	case err == nil:
		if err := proto.UnmarshalText(string(data), msg); err != nil {
			return errors.Annotate(err, "failed to unmarshal %T from [%s]", msg, path).Err()
		}
		return nil

	case isNotExist(err):
		// Forward this so it can be tested.
		return err

	default:
		return errors.Annotate(err, "failed to read data from [%s]", path).Err()
	}
}

func unmarshalTextProtobufDir(base string, fis []os.FileInfo, msg proto.Message, cb func(name string) error) error {
	for _, fi := range fis {
		name := fi.Name()
		if isHidden(name) {
			continue
		}

		if err := unmarshalTextProtobuf(filepath.Join(base, name), msg); err != nil {
			return errors.Annotate(err, "failed to unmarshal file [%s]", name).Err()
		}
		if err := cb(name); err != nil {
			return errors.Annotate(err, "failed to process file [%s]", name).Err()
		}
	}
	return nil
}

func logError(c context.Context, err error, f string, args ...interface{}) {
	log.WithError(err).Errorf(c, f, args...)
	if log.IsLogging(c, log.Debug) {
		log.Debugf(c, "Error stack:\n%s", strings.Join(errors.RenderStack(err), "\n"))
	}
}

func isNotExist(err error) bool {
	return os.IsNotExist(errors.Unwrap(err))
}

func splitTitlePath(s string) (title, string) {
	switch v := strings.SplitN(s, "/", 2); len(v) {
	case 1:
		return title(v[0]), ""
	default:
		return title(v[0]), v[1]
	}
}

func joinPath(titles ...title) string {
	comps := make([]string, len(titles))
	for i, t := range titles {
		comps[i] = string(t)
	}
	return strings.Join(comps, "/")
}

func splitSourcePath(v string) (group, source title) {
	var tail string
	group, tail = splitTitlePath(v)
	source = title(tail)
	return
}

func splitComponentPath(v string) (deployment, component title) {
	var tail string
	deployment, tail = splitTitlePath(v)
	if tail != "" {
		component = title(tail)
	}
	return
}

func splitGoPackage(pkg string) []string {
	// Check intermediate paths to make sure there isn't a deployment
	// conflict.
	var (
		parts   []string
		lastIdx = 0
	)
	for {
		idx := strings.IndexRune(pkg[lastIdx:], '/')
		if idx < 0 {
			// Last component, don't check/register.
			return append(parts, pkg)
		}
		parts = append(parts, pkg[:lastIdx+idx])
		lastIdx += idx + len("/")
	}
}

func findGoPackage(pkg string, goPath []string) string {
	bctx := build.Default
	bctx.GOPATH = strings.Join(goPath, string(os.PathListSeparator))
	p, err := bctx.Import(pkg, "", build.FindOnly)
	if err != nil {
		return ""
	}
	return p.Dir
}
