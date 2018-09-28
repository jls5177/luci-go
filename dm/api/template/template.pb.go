// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/template/template.proto

package dmTemplate

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	templateproto "go.chromium.org/luci/common/data/text/templateproto"
	v1 "go.chromium.org/luci/dm/api/service/v1"
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

// File represents a file full of DM template definitions.
type File struct {
	Template             map[string]*File_Template `protobuf:"bytes,1,rep,name=template,proto3" json:"template,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_374a6f94bcadcedb, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetTemplate() map[string]*File_Template {
	if m != nil {
		return m.Template
	}
	return nil
}

// Template defines a single template.
type File_Template struct {
	DistributorConfigName string                       `protobuf:"bytes,1,opt,name=distributor_config_name,json=distributorConfigName,proto3" json:"distributor_config_name,omitempty"`
	Parameters            *templateproto.File_Template `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	DistributorParameters *templateproto.File_Template `protobuf:"bytes,3,opt,name=distributor_parameters,json=distributorParameters,proto3" json:"distributor_parameters,omitempty"`
	Meta                  *v1.Quest_Desc_Meta          `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                     `json:"-"`
	XXX_unrecognized      []byte                       `json:"-"`
	XXX_sizecache         int32                        `json:"-"`
}

func (m *File_Template) Reset()         { *m = File_Template{} }
func (m *File_Template) String() string { return proto.CompactTextString(m) }
func (*File_Template) ProtoMessage()    {}
func (*File_Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_374a6f94bcadcedb, []int{0, 0}
}

func (m *File_Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File_Template.Unmarshal(m, b)
}
func (m *File_Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File_Template.Marshal(b, m, deterministic)
}
func (m *File_Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File_Template.Merge(m, src)
}
func (m *File_Template) XXX_Size() int {
	return xxx_messageInfo_File_Template.Size(m)
}
func (m *File_Template) XXX_DiscardUnknown() {
	xxx_messageInfo_File_Template.DiscardUnknown(m)
}

var xxx_messageInfo_File_Template proto.InternalMessageInfo

func (m *File_Template) GetDistributorConfigName() string {
	if m != nil {
		return m.DistributorConfigName
	}
	return ""
}

func (m *File_Template) GetParameters() *templateproto.File_Template {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *File_Template) GetDistributorParameters() *templateproto.File_Template {
	if m != nil {
		return m.DistributorParameters
	}
	return nil
}

func (m *File_Template) GetMeta() *v1.Quest_Desc_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "dmTemplate.File")
	proto.RegisterMapType((map[string]*File_Template)(nil), "dmTemplate.File.TemplateEntry")
	proto.RegisterType((*File_Template)(nil), "dmTemplate.File.Template")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/dm/api/template/template.proto", fileDescriptor_374a6f94bcadcedb)
}

var fileDescriptor_374a6f94bcadcedb = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xe9, 0x3a, 0x65, 0xde, 0x21, 0x48, 0x44, 0x9d, 0x45, 0x64, 0xf8, 0xe2, 0x9e, 0x12,
	0xdc, 0x40, 0x65, 0xf8, 0xe4, 0xbf, 0x37, 0x45, 0xab, 0xf8, 0x3a, 0xb2, 0xf6, 0xda, 0x05, 0x9b,
	0xa6, 0xa4, 0xb7, 0xc3, 0x7d, 0x16, 0xbf, 0xab, 0x48, 0xbb, 0x7f, 0x9d, 0xa8, 0xf8, 0x76, 0x93,
	0x73, 0xf2, 0xcb, 0x39, 0x17, 0x7a, 0x91, 0xe1, 0xc1, 0xc8, 0x1a, 0xad, 0x72, 0xcd, 0x8d, 0x8d,
	0x44, 0x9c, 0x07, 0x4a, 0x84, 0x5a, 0xc8, 0x54, 0x09, 0x42, 0x9d, 0xc6, 0x92, 0x70, 0x31, 0xf0,
	0xd4, 0x1a, 0x32, 0x0c, 0x42, 0xfd, 0x3c, 0xbb, 0xf1, 0x2e, 0x7f, 0x04, 0x04, 0x46, 0x6b, 0x93,
	0x88, 0x50, 0x92, 0x14, 0x84, 0xef, 0xb4, 0x20, 0x94, 0x80, 0x6f, 0x3c, 0xef, 0xec, 0xaf, 0x10,
	0x19, 0xda, 0xb1, 0x0a, 0x50, 0x8c, 0x4f, 0x44, 0x64, 0x65, 0x3a, 0x1a, 0x14, 0xd0, 0xe9, 0xc3,
	0xa3, 0x0f, 0x17, 0xea, 0xb7, 0x2a, 0x46, 0xd6, 0x87, 0xc6, 0x9c, 0xd9, 0x72, 0xda, 0x6e, 0xa7,
	0xd9, 0x3d, 0xe4, 0xcb, 0x90, 0xbc, 0xf0, 0xf0, 0xf9, 0xe9, 0x26, 0x21, 0x3b, 0xf1, 0x17, 0x7e,
	0xef, 0xd3, 0x81, 0xc6, 0x5c, 0x63, 0xa7, 0xb0, 0x17, 0xaa, 0x8c, 0xac, 0x1a, 0xe6, 0x64, 0xec,
	0x20, 0x30, 0xc9, 0xab, 0x8a, 0x06, 0x89, 0xd4, 0x05, 0xd7, 0xe9, 0x6c, 0xf8, 0x3b, 0x15, 0xf9,
	0xaa, 0x54, 0xef, 0xa5, 0x46, 0x76, 0x01, 0x90, 0x4a, 0x2b, 0x35, 0x12, 0xda, 0xac, 0x55, 0x6b,
	0x3b, 0x9d, 0x66, 0xf7, 0x80, 0xaf, 0xb4, 0x5e, 0x4d, 0xe1, 0x57, 0xfc, 0xec, 0x09, 0x76, 0xab,
	0xbf, 0x56, 0x48, 0xee, 0x3f, 0x48, 0xd5, 0x48, 0x0f, 0x4b, 0xe8, 0x31, 0xd4, 0x35, 0x92, 0x6c,
	0xd5, 0x4b, 0xc4, 0x36, 0x0f, 0x35, 0x7f, 0xcc, 0x31, 0x23, 0x7e, 0x8d, 0x59, 0xc0, 0xef, 0x90,
	0xa4, 0x5f, 0x1a, 0xbc, 0x17, 0xd8, 0x5c, 0xd9, 0x0d, 0xdb, 0x02, 0xf7, 0x0d, 0x27, 0xb3, 0xc2,
	0xc5, 0xc8, 0x04, 0xac, 0x8d, 0x65, 0x9c, 0xe3, 0xac, 0xd9, 0xfe, 0xaf, 0xcb, 0xf5, 0xa7, 0xbe,
	0x7e, 0xed, 0xdc, 0x19, 0xae, 0x97, 0x61, 0x7b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xcf,
	0xa9, 0x67, 0x64, 0x02, 0x00, 0x00,
}
