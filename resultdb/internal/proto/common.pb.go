// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/internal/proto/common.proto

package internalpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/resultdb/proto/rpc/v1"
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

// A message for storing all the information attached to a page token.
type PageToken struct {
	// Position marks the cursor's start (exclusive). Its interpretation is
	// implementation-specific. For instance, for a Spanner cursor, this is a
	// string slice representation of the Spanner key corresponding to the entry
	// prior to the one at which to start reading, or empty if the cursor is to
	// start at the beginning.
	Position             []string `protobuf:"bytes,1,rep,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageToken) Reset()         { *m = PageToken{} }
func (m *PageToken) String() string { return proto.CompactTextString(m) }
func (*PageToken) ProtoMessage()    {}
func (*PageToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccc4743443d652e, []int{0}
}

func (m *PageToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageToken.Unmarshal(m, b)
}
func (m *PageToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageToken.Marshal(b, m, deterministic)
}
func (m *PageToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageToken.Merge(m, src)
}
func (m *PageToken) XXX_Size() int {
	return xxx_messageInfo_PageToken.Size(m)
}
func (m *PageToken) XXX_DiscardUnknown() {
	xxx_messageInfo_PageToken.DiscardUnknown(m)
}

var xxx_messageInfo_PageToken proto.InternalMessageInfo

func (m *PageToken) GetPosition() []string {
	if m != nil {
		return m.Position
	}
	return nil
}

// A container of artifacts.
// Used to store artifacts in Spanner.
type Artifacts struct {
	ArtifactsV1          []*v1.Artifact `protobuf:"bytes,1,rep,name=artifacts_v1,json=artifactsV1,proto3" json:"artifacts_v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Artifacts) Reset()         { *m = Artifacts{} }
func (m *Artifacts) String() string { return proto.CompactTextString(m) }
func (*Artifacts) ProtoMessage()    {}
func (*Artifacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccc4743443d652e, []int{1}
}

func (m *Artifacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifacts.Unmarshal(m, b)
}
func (m *Artifacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifacts.Marshal(b, m, deterministic)
}
func (m *Artifacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifacts.Merge(m, src)
}
func (m *Artifacts) XXX_Size() int {
	return xxx_messageInfo_Artifacts.Size(m)
}
func (m *Artifacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifacts.DiscardUnknown(m)
}

var xxx_messageInfo_Artifacts proto.InternalMessageInfo

func (m *Artifacts) GetArtifactsV1() []*v1.Artifact {
	if m != nil {
		return m.ArtifactsV1
	}
	return nil
}

func init() {
	proto.RegisterType((*PageToken)(nil), "internal.proto.PageToken")
	proto.RegisterType((*Artifacts)(nil), "internal.proto.Artifacts")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/internal/proto/common.proto", fileDescriptor_4ccc4743443d652e)
}

var fileDescriptor_4ccc4743443d652e = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x4b, 0xc5, 0x30,
	0x14, 0x85, 0x11, 0x41, 0x6c, 0xde, 0xc3, 0xa1, 0x93, 0x74, 0x10, 0xe9, 0xa2, 0xd3, 0x0d, 0xd1,
	0x49, 0x70, 0xa9, 0x3f, 0x40, 0xa4, 0x88, 0x83, 0x4b, 0x49, 0x63, 0xac, 0xc1, 0x26, 0x37, 0xdc,
	0xdc, 0xf6, 0xf7, 0x8b, 0x0d, 0xe9, 0xf8, 0xc6, 0x8f, 0x93, 0xef, 0x9c, 0x5c, 0xf1, 0x34, 0x21,
	0x98, 0x1f, 0x42, 0xef, 0x16, 0x0f, 0x48, 0x93, 0x9c, 0x17, 0xe3, 0x24, 0xd9, 0xb4, 0xcc, 0xfc,
	0x35, 0x4a, 0x17, 0xd8, 0x52, 0xd0, 0xb3, 0x8c, 0x84, 0x8c, 0xd2, 0xa0, 0xf7, 0x18, 0x60, 0x83,
	0xfa, 0xaa, 0x84, 0x99, 0x9b, 0xe7, 0xd3, 0x55, 0xb9, 0x81, 0xa2, 0x91, 0xab, 0x92, 0x6c, 0x13,
	0x0f, 0x39, 0xca, 0x76, 0x7b, 0x27, 0xaa, 0x37, 0x3d, 0xd9, 0x77, 0xfc, 0xb5, 0xa1, 0x6e, 0xc4,
	0x65, 0xc4, 0xe4, 0xd8, 0x61, 0xb8, 0x3e, 0xbb, 0x3d, 0xbf, 0xaf, 0xfa, 0x9d, 0xdb, 0x57, 0x51,
	0x75, 0xc4, 0xee, 0x5b, 0x1b, 0x4e, 0x75, 0x27, 0x8e, 0xba, 0xc0, 0xb0, 0xaa, 0xed, 0xf1, 0xe1,
	0xe1, 0x06, 0xfe, 0xa7, 0xa1, 0x4c, 0x03, 0x45, 0x03, 0xab, 0x82, 0xa2, 0xf5, 0x87, 0xdd, 0xf9,
	0x50, 0x2f, 0xc7, 0x4f, 0x51, 0x0e, 0x89, 0xe3, 0x78, 0xb1, 0xfd, 0xe6, 0xf1, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x63, 0xd0, 0xc8, 0x21, 0x18, 0x01, 0x00, 0x00,
}