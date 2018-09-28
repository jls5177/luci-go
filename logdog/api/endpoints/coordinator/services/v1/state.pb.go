// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/endpoints/coordinator/services/v1/state.proto

package logdog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// LogStreamState is the log stream state value communicated to services.
type LogStreamState struct {
	// ProtoVersion is the protobuf version for this stream.
	ProtoVersion string `protobuf:"bytes,1,opt,name=proto_version,json=protoVersion,proto3" json:"proto_version,omitempty"`
	// The log stream's secret.
	//
	// Note that the secret is returned! This is okay, since this endpoint is only
	// accessible to trusted services. The secret can be cached by services to
	// validate stream information without needing to ping the Coordinator in
	// between each update.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// The stream index of the log stream's terminal message. If the value is -1,
	// the log is still streaming.
	TerminalIndex int64 `protobuf:"varint,3,opt,name=terminal_index,json=terminalIndex,proto3" json:"terminal_index,omitempty"`
	// If the log stream has been archived.
	Archived bool `protobuf:"varint,4,opt,name=archived,proto3" json:"archived,omitempty"`
	// If the log stream has been purged.
	Purged               bool     `protobuf:"varint,5,opt,name=purged,proto3" json:"purged,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStreamState) Reset()         { *m = LogStreamState{} }
func (m *LogStreamState) String() string { return proto.CompactTextString(m) }
func (*LogStreamState) ProtoMessage()    {}
func (*LogStreamState) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa585312e55e4e16, []int{0}
}

func (m *LogStreamState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStreamState.Unmarshal(m, b)
}
func (m *LogStreamState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStreamState.Marshal(b, m, deterministic)
}
func (m *LogStreamState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStreamState.Merge(m, src)
}
func (m *LogStreamState) XXX_Size() int {
	return xxx_messageInfo_LogStreamState.Size(m)
}
func (m *LogStreamState) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStreamState.DiscardUnknown(m)
}

var xxx_messageInfo_LogStreamState proto.InternalMessageInfo

func (m *LogStreamState) GetProtoVersion() string {
	if m != nil {
		return m.ProtoVersion
	}
	return ""
}

func (m *LogStreamState) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *LogStreamState) GetTerminalIndex() int64 {
	if m != nil {
		return m.TerminalIndex
	}
	return 0
}

func (m *LogStreamState) GetArchived() bool {
	if m != nil {
		return m.Archived
	}
	return false
}

func (m *LogStreamState) GetPurged() bool {
	if m != nil {
		return m.Purged
	}
	return false
}

func init() {
	proto.RegisterType((*LogStreamState)(nil), "logdog.LogStreamState")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/logdog/api/endpoints/coordinator/services/v1/state.proto", fileDescriptor_aa585312e55e4e16)
}

var fileDescriptor_aa585312e55e4e16 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x40, 0x89, 0xab, 0x65, 0x0d, 0xbb, 0x7b, 0xc8, 0x41, 0x82, 0xa7, 0xa2, 0x08, 0x3d, 0x35,
	0x88, 0x5f, 0x21, 0xe8, 0xa5, 0x0b, 0x5e, 0x97, 0x98, 0x0c, 0xd9, 0x81, 0x36, 0x53, 0x26, 0x69,
	0xf1, 0x93, 0xfc, 0x4c, 0x69, 0xb6, 0xee, 0xf1, 0xbd, 0x79, 0x33, 0x8c, 0xfc, 0x0c, 0xd4, 0xba,
	0x33, 0xd3, 0x80, 0xd3, 0xd0, 0x12, 0x07, 0xd3, 0x4f, 0x0e, 0x4d, 0x4f, 0xc1, 0x53, 0x30, 0x76,
	0x44, 0x03, 0xd1, 0x8f, 0x84, 0x31, 0x27, 0xe3, 0x88, 0xd8, 0x63, 0xb4, 0x99, 0xd8, 0x24, 0xe0,
	0x19, 0x1d, 0x24, 0x33, 0xbf, 0x9a, 0x94, 0x6d, 0x86, 0x76, 0x64, 0xca, 0xa4, 0xaa, 0xcb, 0xe6,
	0xd3, 0xaf, 0x90, 0x87, 0x0f, 0x0a, 0xc7, 0xcc, 0x60, 0x87, 0xe3, 0x12, 0xa8, 0x67, 0xb9, 0x2f,
	0xcd, 0x69, 0x06, 0x4e, 0x48, 0x51, 0x8b, 0x5a, 0x34, 0xf7, 0xdd, 0xae, 0xc8, 0xaf, 0x8b, 0x53,
	0x0f, 0xb2, 0x4a, 0xe0, 0x18, 0xb2, 0xbe, 0xa9, 0x45, 0xb3, 0xeb, 0x56, 0x52, 0x2f, 0xf2, 0x90,
	0x81, 0x07, 0x8c, 0xb6, 0x3f, 0x61, 0xf4, 0xf0, 0xa3, 0x37, 0xb5, 0x68, 0x36, 0xdd, 0xfe, 0xdf,
	0xbe, 0x2f, 0x52, 0x3d, 0xca, 0xad, 0x65, 0x77, 0xc6, 0x19, 0xbc, 0xbe, 0xad, 0x45, 0xb3, 0xed,
	0xae, 0xbc, 0x9c, 0x1e, 0x27, 0x0e, 0xe0, 0xf5, 0x5d, 0x99, 0xac, 0xf4, 0x5d, 0x95, 0x07, 0xde,
	0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xda, 0x1d, 0x68, 0x2d, 0x0a, 0x01, 0x00, 0x00,
}
