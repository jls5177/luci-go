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

package field

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"

	pb "go.chromium.org/luci/common/tsmon/ts_mon_proto"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSerializeDescriptor(t *testing.T) {
	data := []struct {
		fields []Field
		want   []*pb.MetricsDataSet_MetricFieldDescriptor
	}{
		{
			fields: []Field{String("foo")},
			want: []*pb.MetricsDataSet_MetricFieldDescriptor{
				{
					Name:      proto.String("foo"),
					FieldType: pb.MetricsDataSet_MetricFieldDescriptor_STRING.Enum(),
				},
			},
		},
		{
			fields: []Field{Int("foo")},
			want: []*pb.MetricsDataSet_MetricFieldDescriptor{
				{
					Name:      proto.String("foo"),
					FieldType: pb.MetricsDataSet_MetricFieldDescriptor_INT64.Enum(),
				},
			},
		},
		{
			fields: []Field{Bool("foo")},
			want: []*pb.MetricsDataSet_MetricFieldDescriptor{
				{
					Name:      proto.String("foo"),
					FieldType: pb.MetricsDataSet_MetricFieldDescriptor_BOOL.Enum(),
				},
			},
		},
	}

	for i, d := range data {
		Convey(fmt.Sprintf("%d. SerializeDescriptor(%v)", i, d.fields), t, func() {
			got := SerializeDescriptor(d.fields)
			So(got, ShouldResemble, d.want)
		})
	}
}

func TestSerialize(t *testing.T) {
	data := []struct {
		fields []Field
		values []interface{}
		want   []*pb.MetricsData_MetricField
	}{
		{
			fields: []Field{String("foo")},
			values: makeInterfaceSlice("v"),
			want: []*pb.MetricsData_MetricField{
				{
					Name:  proto.String("foo"),
					Value: &pb.MetricsData_MetricField_StringValue{"v"},
				},
			},
		},
		{
			fields: []Field{Int("foo")},
			values: makeInterfaceSlice(int64(123)),
			want: []*pb.MetricsData_MetricField{
				{
					Name:  proto.String("foo"),
					Value: &pb.MetricsData_MetricField_Int64Value{123},
				},
			},
		},
		{
			fields: []Field{Bool("foo")},
			values: makeInterfaceSlice(true),
			want: []*pb.MetricsData_MetricField{
				{
					Name:  proto.String("foo"),
					Value: &pb.MetricsData_MetricField_BoolValue{true},
				},
			},
		},
	}

	for i, d := range data {
		Convey(fmt.Sprintf("%d. Serialize(%v, %v)", i, d.fields, d.values), t, func() {
			got := Serialize(d.fields, d.values)
			So(got, ShouldResemble, d.want)
		})
	}
}
