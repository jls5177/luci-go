// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/vpython/api/vpython/pep425.proto

package vpython

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

// Represents a Python PEP425 tag.
type PEP425Tag struct {
	// Python is the PEP425 python tag (e.g., "cp27").
	Python string `protobuf:"bytes,1,opt,name=python,proto3" json:"python,omitempty"`
	// ABI is the PEP425 "python ABI" tag (e.g., "cp27mu", "none").
	Abi string `protobuf:"bytes,2,opt,name=abi,proto3" json:"abi,omitempty"`
	// Platform is the PEP425 "python platform" tag (e.g., "linux_x86_64",
	// "armv7l", "any").
	Platform             string   `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PEP425Tag) Reset()         { *m = PEP425Tag{} }
func (m *PEP425Tag) String() string { return proto.CompactTextString(m) }
func (*PEP425Tag) ProtoMessage()    {}
func (*PEP425Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_06702d5b59b98781, []int{0}
}

func (m *PEP425Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PEP425Tag.Unmarshal(m, b)
}
func (m *PEP425Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PEP425Tag.Marshal(b, m, deterministic)
}
func (m *PEP425Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PEP425Tag.Merge(m, src)
}
func (m *PEP425Tag) XXX_Size() int {
	return xxx_messageInfo_PEP425Tag.Size(m)
}
func (m *PEP425Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_PEP425Tag.DiscardUnknown(m)
}

var xxx_messageInfo_PEP425Tag proto.InternalMessageInfo

func (m *PEP425Tag) GetPython() string {
	if m != nil {
		return m.Python
	}
	return ""
}

func (m *PEP425Tag) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *PEP425Tag) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*PEP425Tag)(nil), "vpython.PEP425Tag")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/vpython/api/vpython/pep425.proto", fileDescriptor_06702d5b59b98781)
}

var fileDescriptor_06702d5b59b98781 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0x2b, 0xa8, 0x2c, 0xc9, 0xc8, 0xcf, 0xd3, 0x4f, 0x2c, 0x40, 0xb0, 0x0b, 0x52, 0x0b,
	0x4c, 0x8c, 0x4c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0xa1, 0xa2, 0x4a, 0x81, 0x5c,
	0x9c, 0x01, 0xae, 0x01, 0x26, 0x46, 0xa6, 0x21, 0x89, 0xe9, 0x42, 0x62, 0x5c, 0x6c, 0x10, 0x61,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48, 0x80, 0x8b, 0x39, 0x31, 0x29, 0x53,
	0x82, 0x09, 0x2c, 0x08, 0x62, 0x0a, 0x49, 0x71, 0x71, 0x14, 0xe4, 0x24, 0x96, 0xa4, 0xe5, 0x17,
	0xe5, 0x4a, 0x30, 0x83, 0x85, 0xe1, 0xfc, 0x24, 0x36, 0xb0, 0x15, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x62, 0xd3, 0xb4, 0x6f, 0x9b, 0x00, 0x00, 0x00,
}
