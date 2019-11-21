// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/invocation.proto

package rpcpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_type "go.chromium.org/luci/resultdb/proto/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Invocation_State int32

const (
	// The default value. This value is used if the state is omitted.
	Invocation_STATE_UNSPECIFIED Invocation_State = 0
	// The invocation was created and accepts new results.
	Invocation_ACTIVE Invocation_State = 1
	// The invocation is finalized and contains all the results that the
	// associated computation was expected to compute; unlike INTERRUPTED state.
	//
	// The invocation is immutable and no longer accepts new results.
	Invocation_COMPLETED Invocation_State = 2
	// The invocation is finalized and does NOT contain all the results that the
	// associated computation was expected to compute.
	// The computation was interrupted prematurely.
	//
	// Such invocation should be discarded.
	// Often the associated computation is retried.
	//
	// The invocation is immutable and no longer accepts new results.
	Invocation_INTERRUPTED Invocation_State = 3
)

var Invocation_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "COMPLETED",
	3: "INTERRUPTED",
}

var Invocation_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"COMPLETED":         2,
	"INTERRUPTED":       3,
}

func (x Invocation_State) String() string {
	return proto.EnumName(Invocation_State_name, int32(x))
}

func (Invocation_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{0, 0}
}

// A conceptual container of results. Immutable once finalized.
// It represents all results of some computation; examples: swarming task,
// buildbucket build, CQ attempt.
// Composable: can include other invocations, see inclusion.proto.
type Invocation struct {
	// Can be used to refer to this invocation, e.g. in ResultDB.GetInvocation
	// RPC.
	// Format: invocations/{INVOCATION_ID}
	// See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Current state of the invocation.
	State Invocation_State `protobuf:"varint,2,opt,name=state,proto3,enum=luci.resultdb.rpc.v1.Invocation_State" json:"state,omitempty"`
	// When the invocation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*_type.StringPair `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// When the invocation was finalized, i.e. transitioned to COMPLETED or
	// INTERRUPTED state.
	// If this field is set, implies that the invocation is finalized.
	FinalizeTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=finalize_time,json=finalizeTime,proto3" json:"finalize_time,omitempty"`
	// Timestamp when the invocation will be forcefully finalized.
	// Can be extended with UpdateInvocation until finalized.
	Deadline *timestamp.Timestamp `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Names of invocations included into this one. Overall results of this
	// invocation is a UNION of results directly included into this invocation
	// and results from the included invocations, recursively.
	// For example, a Buildbucket build invocation may include invocations of its
	// child swarming tasks and represent overall result of the build,
	// encapsulating the internal structure of the build.
	//
	// The graph is directed.
	// There can be at most one edge between a given pair of invocations.
	// The shape of the graph does not matter. What matters is only the set of
	// reachable invocations. Thus cycles are allowed and are noop.
	//
	// QueryTestResults returns test results from the transitive closure of
	// invocations.
	//
	// Use Recorder.Include RPC to modify this field.
	IncludedInvocations  []string `protobuf:"bytes,7,rep,name=included_invocations,json=includedInvocations,proto3" json:"included_invocations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{0}
}

func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (m *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(m, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Invocation) GetState() Invocation_State {
	if m != nil {
		return m.State
	}
	return Invocation_STATE_UNSPECIFIED
}

func (m *Invocation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Invocation) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Invocation) GetFinalizeTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinalizeTime
	}
	return nil
}

func (m *Invocation) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Invocation) GetIncludedInvocations() []string {
	if m != nil {
		return m.IncludedInvocations
	}
	return nil
}

func init() {
	proto.RegisterEnum("luci.resultdb.rpc.v1.Invocation_State", Invocation_State_name, Invocation_State_value)
	proto.RegisterType((*Invocation)(nil), "luci.resultdb.rpc.v1.Invocation")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/invocation.proto", fileDescriptor_4005c8951497aaef)
}

var fileDescriptor_4005c8951497aaef = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8f, 0x93, 0x40,
	0x1c, 0xc5, 0x65, 0x69, 0xab, 0xfd, 0xd7, 0xd5, 0x3a, 0xae, 0x09, 0xe9, 0x41, 0xc9, 0x1e, 0x0c,
	0xa7, 0x19, 0x17, 0xe3, 0x5e, 0xf6, 0x44, 0x5b, 0x4c, 0x30, 0xba, 0x36, 0x94, 0xf5, 0xe0, 0xa5,
	0x19, 0x86, 0x29, 0x3b, 0x09, 0x30, 0x64, 0x18, 0x9a, 0xe8, 0xa7, 0xed, 0xc9, 0xcf, 0x61, 0x80,
	0xa5, 0x4d, 0x8c, 0xc9, 0xf6, 0xc8, 0xe3, 0xf7, 0xe6, 0xbd, 0x37, 0x19, 0xb8, 0x49, 0x25, 0x66,
	0xf7, 0x4a, 0xe6, 0xa2, 0xce, 0xb1, 0x54, 0x29, 0xc9, 0x6a, 0x26, 0x88, 0xe2, 0x55, 0x9d, 0xe9,
	0x24, 0x26, 0xa5, 0x92, 0x5a, 0x12, 0x55, 0x32, 0xb2, 0xbb, 0x22, 0xa2, 0xd8, 0x49, 0x46, 0xb5,
	0x90, 0x05, 0x6e, 0x75, 0x74, 0xd1, 0xc0, 0xb8, 0x87, 0xb1, 0x2a, 0x19, 0xde, 0x5d, 0xcd, 0xde,
	0xa5, 0x52, 0xa6, 0x19, 0x27, 0xb4, 0x14, 0x64, 0x2b, 0x78, 0x96, 0x6c, 0x62, 0x7e, 0x4f, 0x77,
	0x42, 0xaa, 0xce, 0x76, 0x00, 0xda, 0xaf, 0xb8, 0xde, 0x12, 0x2d, 0x72, 0x5e, 0x69, 0x9a, 0x97,
	0x0f, 0xc0, 0xa7, 0x53, 0x4a, 0xe9, 0x5f, 0x25, 0x27, 0x4c, 0xe6, 0x79, 0x5f, 0xe7, 0xf2, 0x8f,
	0x09, 0x10, 0x1c, 0x3a, 0xa2, 0x19, 0x0c, 0x0a, 0x9a, 0x73, 0xcb, 0xb0, 0x0d, 0x67, 0x3c, 0x1f,
	0xed, 0x3d, 0x73, 0xef, 0x0d, 0xc3, 0x56, 0x43, 0x1e, 0x0c, 0x2b, 0x4d, 0x35, 0xb7, 0xce, 0x6c,
	0xc3, 0x79, 0xe1, 0xbe, 0xc7, 0xff, 0x5b, 0x82, 0x8f, 0x87, 0xe1, 0x75, 0x43, 0xcf, 0xcd, 0xbd,
	0x67, 0x86, 0x9d, 0x13, 0x2d, 0x60, 0xc2, 0x14, 0xa7, 0x9a, 0x6f, 0x9a, 0xfa, 0x96, 0x69, 0x1b,
	0xce, 0xc4, 0x9d, 0xe1, 0x6e, 0x1b, 0xee, 0xb7, 0xe1, 0xa8, 0xdf, 0x76, 0x68, 0x00, 0x9d, 0xad,
	0xf9, 0x81, 0x5c, 0x18, 0x68, 0x9a, 0x56, 0xd6, 0xc0, 0x36, 0x9d, 0x89, 0xfb, 0xf6, 0x9f, 0x1a,
	0xcd, 0x44, 0xbc, 0xd6, 0x4a, 0x14, 0xe9, 0x8a, 0x0a, 0x15, 0xb6, 0x2c, 0x5a, 0xc2, 0xf9, 0x56,
	0x14, 0x34, 0x13, 0xbf, 0x1f, 0xa2, 0x87, 0x8f, 0x46, 0xb7, 0xbd, 0x9f, 0xf7, 0xae, 0x36, 0xf9,
	0x1a, 0x9e, 0x25, 0x9c, 0x26, 0x99, 0x28, 0xb8, 0x35, 0x7a, 0xec, 0x80, 0xf0, 0xc0, 0xa2, 0x6b,
	0xb8, 0x10, 0x05, 0xcb, 0xea, 0x84, 0x27, 0x9b, 0xe3, 0x83, 0xa8, 0xac, 0xa7, 0xb6, 0xe9, 0x8c,
	0xbb, 0xa0, 0xd7, 0x3d, 0x70, 0xbc, 0xbf, 0xea, 0xf2, 0x0b, 0x0c, 0xdb, 0x3b, 0x44, 0x6f, 0xe0,
	0xd5, 0x3a, 0xf2, 0x22, 0x7f, 0x73, 0x77, 0xbb, 0x5e, 0xf9, 0x8b, 0xe0, 0x73, 0xe0, 0x2f, 0xa7,
	0x4f, 0x10, 0xc0, 0xc8, 0x5b, 0x44, 0xc1, 0x0f, 0x7f, 0x6a, 0xa0, 0x73, 0x18, 0x2f, 0xbe, 0x7f,
	0x5b, 0x7d, 0xf5, 0x23, 0x7f, 0x39, 0x3d, 0x43, 0x2f, 0x61, 0x12, 0xdc, 0x46, 0x7e, 0x18, 0xde,
	0xad, 0x1a, 0xc1, 0x9c, 0xbb, 0x3f, 0x3f, 0x9c, 0xfe, 0x6c, 0x6f, 0x54, 0xc9, 0xca, 0x38, 0x1e,
	0xb5, 0xda, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0x1a, 0x76, 0x2f, 0xf1, 0x02, 0x00,
	0x00,
}