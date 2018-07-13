// Copyright 2018 The LUCI Authors.
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

package cipd

import (
	"encoding/json"
	"fmt"
	"time"

	"go.chromium.org/luci/common/proto/google"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
	"go.chromium.org/luci/cipd/common"
)

// Helper structs and functions for working with JSON representation of CIPD
// domain objects.
//
// See also acl.go for ACL-related structs and action_plan.go structs related
// to EnsurePackages call.
//
// These structs largely define public API of 'cipd ... -json-output ...'.

// UnixTime is time.Time that serializes to integer unix timestamp in JSON
// (represented as a number of seconds since January 1, 1970 UTC).
type UnixTime time.Time

// String is needed to be able to print UnixTime.
func (t UnixTime) String() string {
	return time.Time(t).String()
}

// Before is used to compare UnixTime objects.
func (t UnixTime) Before(t2 UnixTime) bool {
	return time.Time(t).Before(time.Time(t2))
}

// IsZero reports whether t represents the zero time instant.
func (t UnixTime) IsZero() bool {
	return time.Time(t).IsZero()
}

// MarshalJSON is used by JSON encoder.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("0"), nil
	}
	return []byte(fmt.Sprintf("%d", time.Time(t).Unix())), nil
}

// JSONError is wrapper around Error that serializes it as string.
type JSONError struct {
	error
}

// MarshalJSON is used by JSON encoder.
func (e JSONError) MarshalJSON() ([]byte, error) {
	if e.error == nil {
		return []byte("null"), nil
	}
	return json.Marshal(e.Error())
}

// InstanceInfo is information about single package instance.
type InstanceInfo struct {
	// Pin identifies package instance.
	Pin common.Pin `json:"pin"`
	// RegisteredBy is identity of whoever uploaded this instance.
	RegisteredBy string `json:"registered_by"`
	// RegisteredTs is when the instance was registered.
	RegisteredTs UnixTime `json:"registered_ts"`
}

// TagInfo is returned by DescribeInstance.
type TagInfo struct {
	// Tag is actual tag name ("key:value" pair).
	Tag string `json:"tag"`
	// RegisteredBy is identity of whoever attached this tag.
	RegisteredBy string `json:"registered_by"`
	// RegisteredTs is when the tag was registered.
	RegisteredTs UnixTime `json:"registered_ts"`
}

// RefInfo is returned by DescribeInstance and FetchPackageRefs.
type RefInfo struct {
	// Ref is the ref name.
	Ref string `json:"ref"`
	// InstanceID is ID of a package instance the ref points to.
	InstanceID string `json:"instance_id"`
	// ModifiedBy is identity of whoever modified this ref last time.
	ModifiedBy string `json:"modified_by"`
	// ModifiedTs is when the ref was modified last time.
	ModifiedTs UnixTime `json:"modified_ts"`
}

// InstanceDescription contains extended information about an instance as
// returned by DescribeInstance.
type InstanceDescription struct {
	InstanceInfo

	// Refs is a list of refs pointing to the instance, sorted by modification
	// timestamp (newest first)
	//
	// Present only if DescribeRefs in DescribeInstanceOpts is true.
	Refs []RefInfo `json:"refs,omitempty"`

	// Tags is a list of tags attached to the instance, sorted by tag key and
	// creation timestamp (newest first).
	//
	// Present only if DescribeTags in DescribeInstanceOpts is true.
	Tags []TagInfo `json:"tags,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// Converters from proto API to JSON output structs.

func apiInstanceToInfo(inst *api.Instance) InstanceInfo {
	return InstanceInfo{
		Pin: common.Pin{
			PackageName: inst.Package,
			InstanceID:  common.ObjectRefToInstanceID(inst.Instance),
		},
		RegisteredBy: inst.RegisteredBy,
		RegisteredTs: UnixTime(google.TimeFromProto(inst.RegisteredTs)),
	}
}

func apiRefToInfo(r *api.Ref) RefInfo {
	return RefInfo{
		Ref:        r.Name,
		InstanceID: common.ObjectRefToInstanceID(r.Instance),
		ModifiedBy: r.ModifiedBy,
		ModifiedTs: UnixTime(google.TimeFromProto(r.ModifiedTs)),
	}
}

func apiTagToInfo(t *api.Tag) TagInfo {
	return TagInfo{
		Tag:          common.JoinInstanceTag(t),
		RegisteredBy: t.AttachedBy,
		RegisteredTs: UnixTime(google.TimeFromProto(t.AttachedTs)),
	}
}

func apiDescToInfo(d *api.DescribeInstanceResponse) *InstanceDescription {
	desc := &InstanceDescription{
		InstanceInfo: apiInstanceToInfo(d.Instance),
	}
	if len(d.Refs) != 0 {
		desc.Refs = make([]RefInfo, len(d.Refs))
		for i, r := range d.Refs {
			desc.Refs[i] = apiRefToInfo(r)
		}
	}
	if len(d.Tags) != 0 {
		desc.Tags = make([]TagInfo, len(d.Tags))
		for i, t := range d.Tags {
			desc.Tags[i] = apiTagToInfo(t)
		}
	}
	return desc
}