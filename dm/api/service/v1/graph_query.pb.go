// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/service/v1/graph_query.proto

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GraphQuery_Search_Domain int32

const (
	GraphQuery_Search_QUEST   GraphQuery_Search_Domain = 0
	GraphQuery_Search_ATTEMPT GraphQuery_Search_Domain = 1
)

var GraphQuery_Search_Domain_name = map[int32]string{
	0: "QUEST",
	1: "ATTEMPT",
}
var GraphQuery_Search_Domain_value = map[string]int32{
	"QUEST":   0,
	"ATTEMPT": 1,
}

func (x GraphQuery_Search_Domain) String() string {
	return proto.EnumName(GraphQuery_Search_Domain_name, int32(x))
}
func (GraphQuery_Search_Domain) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 1, 0}
}

// GraphQuery represents a single query into the state of DM's dependency graph.
// It's a required parameter for WalkGraphReq.
type GraphQuery struct {
	// AttemptList allows you to list one or more specific attempts as the result
	// of the query. If a quest contains the attempt number 0, or is empty, it
	// means 'all attempts for this quest'.
	AttemptList *AttemptList `protobuf:"bytes,1,opt,name=attempt_list,json=attemptList" json:"attempt_list,omitempty"`
	// attempt_range allows you to list a range of attempts in a single quest.
	// low must be > 0, and high must be > low. The range is [low, high). High may
	// be higher than the highest attempt, and low may be lower than the lowest
	// attempt (but not 0).
	AttemptRange []*GraphQuery_AttemptRange `protobuf:"bytes,2,rep,name=attempt_range,json=attemptRange" json:"attempt_range,omitempty"`
	Search       []*GraphQuery_Search       `protobuf:"bytes,3,rep,name=search" json:"search,omitempty"`
}

func (m *GraphQuery) Reset()                    { *m = GraphQuery{} }
func (m *GraphQuery) String() string            { return proto.CompactTextString(m) }
func (*GraphQuery) ProtoMessage()               {}
func (*GraphQuery) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *GraphQuery) GetAttemptList() *AttemptList {
	if m != nil {
		return m.AttemptList
	}
	return nil
}

func (m *GraphQuery) GetAttemptRange() []*GraphQuery_AttemptRange {
	if m != nil {
		return m.AttemptRange
	}
	return nil
}

func (m *GraphQuery) GetSearch() []*GraphQuery_Search {
	if m != nil {
		return m.Search
	}
	return nil
}

type GraphQuery_AttemptRange struct {
	Quest string `protobuf:"bytes,1,opt,name=quest" json:"quest,omitempty"`
	Low   uint32 `protobuf:"varint,2,opt,name=low" json:"low,omitempty"`
	High  uint32 `protobuf:"varint,3,opt,name=high" json:"high,omitempty"`
}

func (m *GraphQuery_AttemptRange) Reset()                    { *m = GraphQuery_AttemptRange{} }
func (m *GraphQuery_AttemptRange) String() string            { return proto.CompactTextString(m) }
func (*GraphQuery_AttemptRange) ProtoMessage()               {}
func (*GraphQuery_AttemptRange) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func (m *GraphQuery_AttemptRange) GetQuest() string {
	if m != nil {
		return m.Quest
	}
	return ""
}

func (m *GraphQuery_AttemptRange) GetLow() uint32 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *GraphQuery_AttemptRange) GetHigh() uint32 {
	if m != nil {
		return m.High
	}
	return 0
}

// A Search allows you to query objects whose properties match all of the
// provided filters. Filters take the form of a dot-delimited path. For
// example, say that we had the following objects:
//
//   Quest(id=deadbeef):
//     created = <timestamp>  #sort
//     descriptor.distributor_config_name = "foo"
//     descriptor.json_payload = {
//       "key": "value",
//       "multi": ["some", 10, "values", true],
//       "sub": [{"msg": 11}, {"msg": 12}],
//     }
//
//   Attempt(id=deadbeef|1):
//     created = <timestamp>  #sort
//     attempt_type = Finished
//     finished.expiration = <timestamp>
//     finished.json_result = {
//       "rslt": "yes",
//       "ok": true,
//     }
//
// Then you could query (in pseudo-proto):
//   domain: Attempt
//   approx_filters: {
//     "attempt_type": ["Finished"],
//     "$quest.descriptor.json_payload.multi": [true, 10],
//     "$quest.descriptor.json_payload.sub.msg": [11, 10],
//     "finished.json_result.ok": [true],
//   }
//
// Or:
//
//   domain: Attempt
//   exact_filters: {
//     "$quest.descriptor.json_payload.multi[1]": [10],
//     "$quest.descriptor.json_payload.sub[0].msg": [11],
//   }
//
// Literal '.' and '[' characters may be escaped with a backslash.
type GraphQuery_Search struct {
	// Domain indicates which class of objects your query applies to. The fields
	// available to query are defined by the `data` field in the corresponding
	// GraphData message.
	//
	// Additionally `Attempt` has a special field $quest whose subfields are
	// queriable in the exact same way that a search in a Quest domain works.
	Domain GraphQuery_Search_Domain `protobuf:"varint,1,opt,name=domain,enum=dm.GraphQuery_Search_Domain" json:"domain,omitempty"`
	// Start and End are optional restrictions on the first sort property. For
	// now, these are just restrictions on the 'created' timestamp for either
	// the Quest or Attempt, depending on the SearchDomain.
	Start *PropertyValue `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	End   *PropertyValue `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
	// ApproxFilters allows you to filter on 'approximate' fields. Approximate
	// fields are the json path to the value, without any array subscripts. For
	// example, if your document looked like:
	//
	//   {
	//     "some": ["list", {"of": ["data", "and", "stuff"]}],
	//   }
	//
	// Then the following approximate filters would match:
	//   "some" = ["list"]
	//   "some.of" = ["data"]
	//   "some.of" = ["and"]
	//   "some.of" = ["stuff"]
	//   "some.of" = ["stuff", "and"]
	//   "some.of" = ["stuff", "and", "data"]
	//
	// This is useful for filtering documents where the order of parameters
	// in a list or sublist isn't known, or doesn't matter.
	ApproxFilters map[string]*MultiPropertyValue `protobuf:"bytes,5,rep,name=approx_filters,json=approxFilters" json:"approx_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ExactFilters allows you to filter on 'exact' fields. Exact fields are the
	// json path to the value, including array subscripts. For example if your
	// document looked like:
	//
	//   {
	//     "some": ["list", {"of": ["data", "and", "stuff"]}],
	//   }
	//
	// Then the following exact filters would match:
	//   "some[0]" = "list"
	//   "some[1].of[0]" = "data"
	//   "some[1].of[1]" = "and"
	//   "some[1].of[2]" = "stuff"
	//
	// This is useful for filtering documents where the order of parameters
	// in a list or sublist matters.
	ExactFilters map[string]*PropertyValue `protobuf:"bytes,6,rep,name=exact_filters,json=exactFilters" json:"exact_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GraphQuery_Search) Reset()                    { *m = GraphQuery_Search{} }
func (m *GraphQuery_Search) String() string            { return proto.CompactTextString(m) }
func (*GraphQuery_Search) ProtoMessage()               {}
func (*GraphQuery_Search) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 1} }

func (m *GraphQuery_Search) GetDomain() GraphQuery_Search_Domain {
	if m != nil {
		return m.Domain
	}
	return GraphQuery_Search_QUEST
}

func (m *GraphQuery_Search) GetStart() *PropertyValue {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *GraphQuery_Search) GetEnd() *PropertyValue {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *GraphQuery_Search) GetApproxFilters() map[string]*MultiPropertyValue {
	if m != nil {
		return m.ApproxFilters
	}
	return nil
}

func (m *GraphQuery_Search) GetExactFilters() map[string]*PropertyValue {
	if m != nil {
		return m.ExactFilters
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphQuery)(nil), "dm.GraphQuery")
	proto.RegisterType((*GraphQuery_AttemptRange)(nil), "dm.GraphQuery.AttemptRange")
	proto.RegisterType((*GraphQuery_Search)(nil), "dm.GraphQuery.Search")
	proto.RegisterEnum("dm.GraphQuery_Search_Domain", GraphQuery_Search_Domain_name, GraphQuery_Search_Domain_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/dm/api/service/v1/graph_query.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xb1, 0x1d, 0xbb, 0x74, 0x9c, 0x94, 0x74, 0x05, 0xc8, 0x32, 0x1c, 0xa2, 0x72, 0x48,
	0x0e, 0x60, 0x0b, 0xc3, 0xa1, 0xe2, 0x44, 0x24, 0x02, 0x52, 0xd5, 0x8a, 0x76, 0x1b, 0x10, 0xb7,
	0x68, 0x6b, 0x2f, 0xce, 0x0a, 0x3b, 0xeb, 0xae, 0xd7, 0xa1, 0xfe, 0x0b, 0xfd, 0xd5, 0x68, 0xd7,
	0x4e, 0x9b, 0xb6, 0x46, 0xea, 0x6d, 0x67, 0xe6, 0xbd, 0x6f, 0x26, 0x2f, 0x32, 0x1c, 0xa6, 0x3c,
	0x88, 0x97, 0x82, 0xe7, 0xac, 0xca, 0x03, 0x2e, 0xd2, 0x30, 0xab, 0x62, 0x16, 0x26, 0x79, 0x48,
	0x0a, 0x16, 0x96, 0x54, 0xac, 0x59, 0x4c, 0xc3, 0xf5, 0xfb, 0x30, 0x15, 0xa4, 0x58, 0x2e, 0x2e,
	0x2b, 0x2a, 0xea, 0xa0, 0x10, 0x5c, 0x72, 0x64, 0x26, 0xb9, 0x1f, 0x3d, 0xd2, 0x2d, 0xeb, 0x82,
	0x96, 0x8d, 0xef, 0xe0, 0xda, 0x01, 0xf8, 0xa6, 0x68, 0x67, 0x0a, 0x86, 0x22, 0xe8, 0x13, 0x29,
	0x69, 0x5e, 0xc8, 0x45, 0xc6, 0x4a, 0xe9, 0x19, 0x23, 0x63, 0xe2, 0x46, 0xcf, 0x82, 0x24, 0x0f,
	0xa6, 0x4d, 0xff, 0x98, 0x95, 0x12, 0xbb, 0xe4, 0xb6, 0x40, 0x9f, 0x61, 0xb0, 0xf1, 0x08, 0xb2,
	0x4a, 0xa9, 0x67, 0x8e, 0xac, 0x89, 0x1b, 0xbd, 0x52, 0xa6, 0x5b, 0xf4, 0xc6, 0x8f, 0x95, 0x04,
	0x6f, 0xb6, 0xe8, 0x0a, 0xbd, 0x03, 0xa7, 0xa4, 0x44, 0xc4, 0x4b, 0xcf, 0xd2, 0xd6, 0x17, 0xf7,
	0xac, 0xe7, 0x7a, 0x88, 0x5b, 0x91, 0x7f, 0x04, 0xfd, 0x6d, 0x18, 0x7a, 0x0e, 0xf6, 0x65, 0x45,
	0xdb, 0x6b, 0x77, 0x71, 0x53, 0xa0, 0x21, 0x58, 0x19, 0xff, 0xeb, 0x99, 0x23, 0x63, 0x32, 0xc0,
	0xea, 0x89, 0x10, 0xf4, 0x96, 0x2c, 0x55, 0x4b, 0x54, 0x4b, 0xbf, 0xfd, 0xeb, 0x1e, 0x38, 0x0d,
	0x1e, 0x7d, 0x04, 0x27, 0xe1, 0x39, 0x61, 0x2b, 0xcd, 0xd9, 0x8b, 0x5e, 0x77, 0x5e, 0x11, 0x7c,
	0xd1, 0x1a, 0xdc, 0x6a, 0xd1, 0x18, 0xec, 0x52, 0x12, 0x21, 0x35, 0xd5, 0x8d, 0xf6, 0x95, 0xe9,
	0x54, 0xf0, 0x82, 0x0a, 0x59, 0xff, 0x24, 0x59, 0x45, 0x71, 0x33, 0x47, 0x6f, 0xc0, 0xa2, 0xab,
	0xc4, 0xeb, 0xfd, 0x4f, 0xa6, 0xa6, 0xe8, 0x3b, 0xec, 0x91, 0xa2, 0x10, 0xfc, 0x6a, 0xf1, 0x9b,
	0x65, 0x92, 0x8a, 0xd2, 0xb3, 0x75, 0x22, 0x93, 0xee, 0x5b, 0xa6, 0x5a, 0xfb, 0xb5, 0x91, 0xce,
	0x56, 0x52, 0xd4, 0x78, 0x40, 0xb6, 0x7b, 0xe8, 0x18, 0x06, 0xf4, 0x8a, 0xc4, 0xf2, 0x86, 0xe7,
	0x68, 0xde, 0xb8, 0x9b, 0x37, 0x53, 0xd2, 0x3b, 0xb8, 0x3e, 0xdd, 0x6a, 0xf9, 0xbf, 0x00, 0x3d,
	0x5c, 0xa9, 0x92, 0xfe, 0x43, 0xeb, 0x36, 0x7d, 0xf5, 0x44, 0x6f, 0xc1, 0x5e, 0xab, 0x1f, 0xa5,
	0xd3, 0x77, 0xa3, 0x97, 0x6a, 0xdb, 0x49, 0x95, 0x49, 0x76, 0x2f, 0x19, 0x2d, 0xfa, 0x64, 0x1e,
	0x1a, 0x3e, 0x86, 0xfd, 0x07, 0xcb, 0x3b, 0xc0, 0xe3, 0xbb, 0xe0, 0xae, 0xb4, 0x6f, 0x98, 0x07,
	0x23, 0x70, 0x9a, 0x3f, 0x0b, 0xed, 0x82, 0x7d, 0xf6, 0x63, 0x76, 0x3e, 0x1f, 0x3e, 0x41, 0x2e,
	0xec, 0x4c, 0xe7, 0xf3, 0xd9, 0xc9, 0xe9, 0x7c, 0x68, 0x1c, 0xf5, 0x9e, 0x9a, 0x43, 0x0b, 0xef,
	0x94, 0x5c, 0xc8, 0xc5, 0x45, 0x7d, 0xe1, 0xe8, 0x6f, 0xe2, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xed, 0x5b, 0xb4, 0x87, 0x03, 0x00, 0x00,
}
