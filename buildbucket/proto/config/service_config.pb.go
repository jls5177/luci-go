// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/config/service_config.proto

package configpb

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

// Schema of settings.cfg file, a service config.
type SettingsCfg struct {
	// Swarmbucket settings.
	Swarming             *SwarmingSettings `protobuf:"bytes,1,opt,name=swarming,proto3" json:"swarming,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SettingsCfg) Reset()         { *m = SettingsCfg{} }
func (m *SettingsCfg) String() string { return proto.CompactTextString(m) }
func (*SettingsCfg) ProtoMessage()    {}
func (*SettingsCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d1f4664ae625b9b, []int{0}
}

func (m *SettingsCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsCfg.Unmarshal(m, b)
}
func (m *SettingsCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsCfg.Marshal(b, m, deterministic)
}
func (m *SettingsCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsCfg.Merge(m, src)
}
func (m *SettingsCfg) XXX_Size() int {
	return xxx_messageInfo_SettingsCfg.Size(m)
}
func (m *SettingsCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsCfg.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsCfg proto.InternalMessageInfo

func (m *SettingsCfg) GetSwarming() *SwarmingSettings {
	if m != nil {
		return m.Swarming
	}
	return nil
}

// Swarmbucket settings.
type SwarmingSettings struct {
	// Swarmbucket build URLs will point to this Milo instance.
	MiloHostname string `protobuf:"bytes,2,opt,name=milo_hostname,json=miloHostname,proto3" json:"milo_hostname,omitempty"`
	// Default value of swarming.task_template_canary_percentage field in
	// cr-buildbucket.cfg files.
	DefaultTaskTemplateCanaryPercentage int32    `protobuf:"varint,3,opt,name=default_task_template_canary_percentage,json=defaultTaskTemplateCanaryPercentage,proto3" json:"default_task_template_canary_percentage,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *SwarmingSettings) Reset()         { *m = SwarmingSettings{} }
func (m *SwarmingSettings) String() string { return proto.CompactTextString(m) }
func (*SwarmingSettings) ProtoMessage()    {}
func (*SwarmingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d1f4664ae625b9b, []int{1}
}

func (m *SwarmingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwarmingSettings.Unmarshal(m, b)
}
func (m *SwarmingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwarmingSettings.Marshal(b, m, deterministic)
}
func (m *SwarmingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwarmingSettings.Merge(m, src)
}
func (m *SwarmingSettings) XXX_Size() int {
	return xxx_messageInfo_SwarmingSettings.Size(m)
}
func (m *SwarmingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_SwarmingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_SwarmingSettings proto.InternalMessageInfo

func (m *SwarmingSettings) GetMiloHostname() string {
	if m != nil {
		return m.MiloHostname
	}
	return ""
}

func (m *SwarmingSettings) GetDefaultTaskTemplateCanaryPercentage() int32 {
	if m != nil {
		return m.DefaultTaskTemplateCanaryPercentage
	}
	return 0
}

func init() {
	proto.RegisterType((*SettingsCfg)(nil), "buildbucket.SettingsCfg")
	proto.RegisterType((*SwarmingSettings)(nil), "buildbucket.SwarmingSettings")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/config/service_config.proto", fileDescriptor_5d1f4664ae625b9b)
}

var fileDescriptor_5d1f4664ae625b9b = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x70, 0x94, 0xfe, 0x21, 0x95, 0x5b, 0x08, 0x9e, 0xbc, 0x14, 0x4c, 0x32, 0xd4, 0x93,
	0x0d, 0x2d, 0x94, 0x96, 0x6e, 0xc9, 0x12, 0x3a, 0x15, 0xc7, 0x53, 0x17, 0x21, 0x2b, 0x67, 0x45,
	0xd8, 0x92, 0x8c, 0x74, 0x6a, 0xe9, 0x73, 0xf4, 0x85, 0x4b, 0x1d, 0x37, 0x84, 0x6e, 0x99, 0x8e,
	0xfb, 0xbe, 0xdf, 0x2d, 0x47, 0x97, 0xd2, 0xe6, 0x62, 0xe7, 0xac, 0x56, 0x41, 0xe7, 0xd6, 0xc9,
	0xa2, 0x0b, 0x42, 0x15, 0x75, 0x50, 0xdd, 0xb6, 0x0e, 0xa2, 0x05, 0x2c, 0x7a, 0x67, 0xd1, 0x16,
	0xc2, 0x9a, 0x46, 0xc9, 0xc2, 0x83, 0xfb, 0x50, 0x02, 0xd8, 0x7e, 0xcd, 0x87, 0x2e, 0x8e, 0x8e,
	0xf8, 0x7c, 0x4d, 0xa3, 0x0d, 0x20, 0x2a, 0x23, 0xfd, 0xaa, 0x91, 0xf1, 0x33, 0x9d, 0xfa, 0x4f,
	0xee, 0xb4, 0x32, 0x32, 0x21, 0x29, 0xc9, 0xa2, 0xfb, 0xdb, 0xfc, 0x88, 0xe7, 0x9b, 0xb1, 0xfc,
	0xbb, 0x29, 0x0f, 0x7c, 0xfe, 0x4d, 0xe8, 0xec, 0x7f, 0x1d, 0x2f, 0xe8, 0x8d, 0x56, 0x9d, 0x65,
	0x3b, 0xeb, 0xd1, 0x70, 0x0d, 0xc9, 0x24, 0x25, 0xd9, 0x55, 0x79, 0xfd, 0x1b, 0xae, 0xc7, 0x2c,
	0xae, 0xe8, 0xdd, 0x16, 0x1a, 0x1e, 0x3a, 0x64, 0xc8, 0x7d, 0xcb, 0x10, 0x74, 0xdf, 0x71, 0x04,
	0x26, 0xb8, 0xe1, 0xee, 0x8b, 0xf5, 0xe0, 0x04, 0x18, 0xe4, 0x12, 0x92, 0xb3, 0x94, 0x64, 0x17,
	0xe5, 0x62, 0xe4, 0x15, 0xf7, 0x6d, 0x35, 0xe2, 0xd5, 0x60, 0xdf, 0x0e, 0xf4, 0xf5, 0x7c, 0x4a,
	0x66, 0x93, 0xe5, 0xd3, 0xfb, 0xe3, 0x49, 0x2f, 0x7b, 0xd9, 0x8f, 0xbe, 0xae, 0x2f, 0x87, 0xf8,
	0xe1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x98, 0x74, 0x44, 0x73, 0x01, 0x00, 0x00,
}
