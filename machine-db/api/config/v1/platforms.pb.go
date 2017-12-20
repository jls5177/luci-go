// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/config/v1/platforms.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Platform describes a platform.
type Platform struct {
	// The name of this platform.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this platform.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Platform) Reset()                    { *m = Platform{} }
func (m *Platform) String() string            { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()               {}
func (*Platform) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Platform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Platforms enumerates platforms.
type Platforms struct {
	// A list of platforms.
	Platform []*Platform `protobuf:"bytes,1,rep,name=platform" json:"platform,omitempty"`
}

func (m *Platforms) Reset()                    { *m = Platforms{} }
func (m *Platforms) String() string            { return proto.CompactTextString(m) }
func (*Platforms) ProtoMessage()               {}
func (*Platforms) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Platforms) GetPlatform() []*Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func init() {
	proto.RegisterType((*Platform)(nil), "config.Platform")
	proto.RegisterType((*Platforms)(nil), "config.Platforms")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/config/v1/platforms.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xcd, 0x0a, 0x82, 0x40,
	0x10, 0x80, 0xb1, 0x42, 0x74, 0xbc, 0xc4, 0x9e, 0x3c, 0x8a, 0x27, 0x0f, 0xb5, 0x43, 0x75, 0xea,
	0x10, 0xf4, 0x08, 0xe1, 0x1b, 0xac, 0xeb, 0xdf, 0x80, 0xbb, 0xb3, 0xac, 0xda, 0xf3, 0x07, 0x9a,
	0xd1, 0x6d, 0x98, 0xef, 0x63, 0xbe, 0x81, 0x47, 0xc7, 0x52, 0xf7, 0x9e, 0x0d, 0xcd, 0x46, 0xb2,
	0xef, 0x70, 0x98, 0x35, 0xa1, 0x51, 0xba, 0x27, 0xdb, 0x9c, 0xeb, 0x0a, 0x95, 0x23, 0xd4, 0x6c,
	0x5b, 0xea, 0xf0, 0x7d, 0x41, 0x37, 0xa8, 0xa9, 0x65, 0x6f, 0x46, 0xe9, 0x3c, 0x4f, 0x2c, 0xc2,
	0x15, 0xe5, 0x4f, 0x88, 0x5e, 0x5f, 0x24, 0x04, 0x1c, 0xac, 0x32, 0x4d, 0x1a, 0x64, 0x41, 0x11,
	0x97, 0xcb, 0x2c, 0x32, 0x48, 0xea, 0x66, 0xd4, 0x9e, 0xdc, 0x44, 0x6c, 0xd3, 0xdd, 0x82, 0xfe,
	0x57, 0xf9, 0x1d, 0xe2, 0xed, 0xc2, 0x28, 0x4e, 0x10, 0x6d, 0xa5, 0x34, 0xc8, 0xf6, 0x45, 0x72,
	0x3d, 0xca, 0xb5, 0x24, 0x37, 0xa9, 0xfc, 0x19, 0x55, 0xb8, 0xfc, 0x72, 0xfb, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x96, 0x69, 0x5f, 0x7b, 0xcc, 0x00, 0x00, 0x00,
}
