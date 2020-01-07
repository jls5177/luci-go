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

package appstatus

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/errors"
)

var appStatusTagKey = errors.NewTagKey("application-specific response status")

// ToError converts an application-specific status to an error.
func ToError(s *status.Status) error {
	return Attach(s.Err(), s)
}

// Error returns an error with an application-specific status.
// The message will be shared with the RPC client as is.
func Error(code codes.Code, msg string) error {
	return ToError(status.New(code, msg))
}

// Errorf returns an error with an application-specific status.
// The message will be shared with the RPC client as is.
func Errorf(code codes.Code, format string, args ...interface{}) error {
	return ToError(status.Newf(code, format, args...))
}

// Attach attaches an application-specific status to the error.
// The status will be shared with the RPC client as is.
// If err already has an application-specific status attached, panics.
func Attach(err error, status *status.Status) error {
	if status.Code() == codes.OK {
		panic("cannot attach an OK status")
	}

	if _, ok := Get(err); ok {
		panic("err already has an application-specific status")
	}

	tag := errors.TagValue{
		Key:   appStatusTagKey,
		Value: status,
	}
	return errors.Annotate(err, "attaching a status").Tag(tag).Err()
}

// Attachf is a shortcut for Attach(err, status.Newf(...))
func Attachf(err error, code codes.Code, format string, args ...interface{}) error {
	return Attach(err, status.Newf(code, format, args...))
}

// Get returns an application-specific Status attached to err using this
// package. If not explicitly set or if err is nil, then ok is false.
func Get(err error) (st *status.Status, ok bool) {
	v, ok := errors.TagValueIn(appStatusTagKey, err)
	if !ok {
		return nil, false
	}

	return v.(*status.Status), true
}
