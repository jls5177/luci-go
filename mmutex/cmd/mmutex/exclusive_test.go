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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/danjacques/gofslock/fslock"
	. "github.com/smartystreets/goconvey/convey"

	. "go.chromium.org/luci/common/testing/assertions"
)

func TestExclusive(t *testing.T) {
	Convey("RunExclusive", t, func() {
		var lockFileDir string
		var err error
		if lockFileDir, err = ioutil.TempDir("", ""); err != nil {
			panic(err)
		}
		lockFilePath := filepath.Join(lockFileDir, "mmutex.lock")
		defer func() {
			os.Remove(lockFileDir)
		}()

		Convey("executes the command", func() {
			var tempDir string
			var err error

			if tempDir, err = ioutil.TempDir("", ""); err != nil {
				panic(err)
			}
			defer os.Remove(tempDir)

			testFilePath := filepath.Join(tempDir, "test")
			var command []string
			if runtime.GOOS == "windows" {
				command = createCommand([]string{"copy", "NUL", testFilePath})
			} else {
				command = createCommand([]string{"touch", testFilePath})
			}

			So(RunExclusive(command, lockFilePath, 0, 0), ShouldBeNil)

			_, err = os.Stat(testFilePath)
			So(err, ShouldBeNil)
		})

		Convey("returns error from the command", func() {
			So(RunExclusive([]string{"nonexistent_command"}, lockFilePath, 0, 0), ShouldErrLike, "executable file not found")
		})

		Convey("times out if exclusive lock isn't released", func() {
			var handle fslock.Handle
			var err error

			if handle, err = fslock.Lock(lockFilePath); err != nil {
				panic(err)
			}
			defer handle.Unlock()

			So(RunExclusive([]string{"echo", "should_fail"}, lockFilePath, 0, 0), ShouldErrLike, "fslock: lock is held")
		})

		Convey("times out if shared lock isn't released", func() {
			var handle fslock.Handle
			var err error

			if handle, err = fslock.LockShared(lockFilePath); err != nil {
				panic(err)
			}
			defer handle.Unlock()

			So(RunExclusive(createCommand([]string{"echo", "should_fail"}), lockFilePath, 0, 0), ShouldErrLike, "fslock: lock is held")
		})

		Convey("respects timeout", func() {
			var handle fslock.Handle
			var err error

			if handle, err = fslock.Lock(lockFilePath); err != nil {
				panic(err)
			}
			defer handle.Unlock()

			start := time.Now()
			RunExclusive(createCommand([]string{"echo", "should_succeed"}), lockFilePath, 5*time.Millisecond, 0)
			So(time.Now(), ShouldHappenOnOrAfter, start.Add(5*time.Millisecond))
		})

		// TODO(charliea): Add a test to ensure that a drain file is created when RunExclusive() is called.
	})

	Convey("RunExclusive acts as a passthrough if lockFilePath is empty", t, func() {
		So(RunExclusive(createCommand([]string{"echo", "should_succeed"}), "", 0, 0), ShouldBeNil)
	})
}
