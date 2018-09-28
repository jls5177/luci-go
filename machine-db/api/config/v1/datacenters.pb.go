// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/config/v1/datacenters.proto

package config

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/machine-db/api/common/v1"
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

// A switch to store in the database.
type Switch struct {
	// The name of this switch. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this switch.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The number of ports this switch has.
	Ports int32 `protobuf:"varint,3,opt,name=ports,proto3" json:"ports,omitempty"`
	// The state of this switch.
	State                v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Switch) Reset()         { *m = Switch{} }
func (m *Switch) String() string { return proto.CompactTextString(m) }
func (*Switch) ProtoMessage()    {}
func (*Switch) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd85d9fa4d73ee0, []int{0}
}

func (m *Switch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Switch.Unmarshal(m, b)
}
func (m *Switch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Switch.Marshal(b, m, deterministic)
}
func (m *Switch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Switch.Merge(m, src)
}
func (m *Switch) XXX_Size() int {
	return xxx_messageInfo_Switch.Size(m)
}
func (m *Switch) XXX_DiscardUnknown() {
	xxx_messageInfo_Switch.DiscardUnknown(m)
}

var xxx_messageInfo_Switch proto.InternalMessageInfo

func (m *Switch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Switch) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Switch) GetPorts() int32 {
	if m != nil {
		return m.Ports
	}
	return 0
}

func (m *Switch) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

// A KVM to store in the database.
type KVM struct {
	// The name of this KVM on the network. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this KVM.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The type of platform this KVM is.
	Platform string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	// The rack this KVM belongs to.
	// Must be the name of a rack in the same datacenter.
	Rack string `protobuf:"bytes,4,opt,name=rack,proto3" json:"rack,omitempty"`
	// The MAC address associated with this KVM.
	MacAddress string `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// The IPv4 address associated with this KVM.
	Ipv4 string `protobuf:"bytes,6,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// The state of this KVM.
	State                v1.State `protobuf:"varint,7,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVM) Reset()         { *m = KVM{} }
func (m *KVM) String() string { return proto.CompactTextString(m) }
func (*KVM) ProtoMessage()    {}
func (*KVM) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd85d9fa4d73ee0, []int{1}
}

func (m *KVM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVM.Unmarshal(m, b)
}
func (m *KVM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVM.Marshal(b, m, deterministic)
}
func (m *KVM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVM.Merge(m, src)
}
func (m *KVM) XXX_Size() int {
	return xxx_messageInfo_KVM.Size(m)
}
func (m *KVM) XXX_DiscardUnknown() {
	xxx_messageInfo_KVM.DiscardUnknown(m)
}

var xxx_messageInfo_KVM proto.InternalMessageInfo

func (m *KVM) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KVM) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *KVM) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *KVM) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *KVM) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *KVM) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *KVM) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

// A rack to store in the database.
type Rack struct {
	// The name of this rack. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this rack.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The switches belonging to this rack.
	Switch []*Switch `protobuf:"bytes,3,rep,name=switch,proto3" json:"switch,omitempty"`
	// The state of this rack.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The KVM serving this rack.
	Kvm                  string   `protobuf:"bytes,5,opt,name=kvm,proto3" json:"kvm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rack) Reset()         { *m = Rack{} }
func (m *Rack) String() string { return proto.CompactTextString(m) }
func (*Rack) ProtoMessage()    {}
func (*Rack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd85d9fa4d73ee0, []int{2}
}

func (m *Rack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rack.Unmarshal(m, b)
}
func (m *Rack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rack.Marshal(b, m, deterministic)
}
func (m *Rack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rack.Merge(m, src)
}
func (m *Rack) XXX_Size() int {
	return xxx_messageInfo_Rack.Size(m)
}
func (m *Rack) XXX_DiscardUnknown() {
	xxx_messageInfo_Rack.DiscardUnknown(m)
}

var xxx_messageInfo_Rack proto.InternalMessageInfo

func (m *Rack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rack) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Rack) GetSwitch() []*Switch {
	if m != nil {
		return m.Switch
	}
	return nil
}

func (m *Rack) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (m *Rack) GetKvm() string {
	if m != nil {
		return m.Kvm
	}
	return ""
}

// A datacenter to store in the database.
type Datacenter struct {
	// The name of this datacenter. Must be globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this datacenter.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The racks belonging to this datacenter.
	Rack []*Rack `protobuf:"bytes,3,rep,name=rack,proto3" json:"rack,omitempty"`
	// The state of this datacenter.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The KVMs belonging to this datacenter.
	Kvm                  []*KVM   `protobuf:"bytes,5,rep,name=kvm,proto3" json:"kvm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Datacenter) Reset()         { *m = Datacenter{} }
func (m *Datacenter) String() string { return proto.CompactTextString(m) }
func (*Datacenter) ProtoMessage()    {}
func (*Datacenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd85d9fa4d73ee0, []int{3}
}

func (m *Datacenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Datacenter.Unmarshal(m, b)
}
func (m *Datacenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Datacenter.Marshal(b, m, deterministic)
}
func (m *Datacenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Datacenter.Merge(m, src)
}
func (m *Datacenter) XXX_Size() int {
	return xxx_messageInfo_Datacenter.Size(m)
}
func (m *Datacenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Datacenter.DiscardUnknown(m)
}

var xxx_messageInfo_Datacenter proto.InternalMessageInfo

func (m *Datacenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Datacenter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Datacenter) GetRack() []*Rack {
	if m != nil {
		return m.Rack
	}
	return nil
}

func (m *Datacenter) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (m *Datacenter) GetKvm() []*KVM {
	if m != nil {
		return m.Kvm
	}
	return nil
}

// A list of datacenter config files.
type Datacenters struct {
	// A list of names of datacenter config files.
	Datacenter           []string `protobuf:"bytes,1,rep,name=datacenter,proto3" json:"datacenter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Datacenters) Reset()         { *m = Datacenters{} }
func (m *Datacenters) String() string { return proto.CompactTextString(m) }
func (*Datacenters) ProtoMessage()    {}
func (*Datacenters) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddd85d9fa4d73ee0, []int{4}
}

func (m *Datacenters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Datacenters.Unmarshal(m, b)
}
func (m *Datacenters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Datacenters.Marshal(b, m, deterministic)
}
func (m *Datacenters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Datacenters.Merge(m, src)
}
func (m *Datacenters) XXX_Size() int {
	return xxx_messageInfo_Datacenters.Size(m)
}
func (m *Datacenters) XXX_DiscardUnknown() {
	xxx_messageInfo_Datacenters.DiscardUnknown(m)
}

var xxx_messageInfo_Datacenters proto.InternalMessageInfo

func (m *Datacenters) GetDatacenter() []string {
	if m != nil {
		return m.Datacenter
	}
	return nil
}

func init() {
	proto.RegisterType((*Switch)(nil), "config.Switch")
	proto.RegisterType((*KVM)(nil), "config.KVM")
	proto.RegisterType((*Rack)(nil), "config.Rack")
	proto.RegisterType((*Datacenter)(nil), "config.Datacenter")
	proto.RegisterType((*Datacenters)(nil), "config.Datacenters")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/config/v1/datacenters.proto", fileDescriptor_ddd85d9fa4d73ee0)
}

var fileDescriptor_ddd85d9fa4d73ee0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x6a, 0xe3, 0x30,
	0x10, 0x46, 0xeb, 0xd8, 0xbb, 0x19, 0xef, 0x86, 0x45, 0xec, 0x41, 0x04, 0x76, 0xd7, 0x64, 0x61,
	0xf1, 0x25, 0x36, 0x4d, 0x7b, 0xea, 0xa5, 0x14, 0x7a, 0x0b, 0xb9, 0x38, 0x90, 0x6b, 0x51, 0x64,
	0x27, 0x11, 0x89, 0x2c, 0x23, 0x29, 0x29, 0x7d, 0x96, 0xbe, 0x40, 0x5f, 0xa4, 0xef, 0x55, 0x24,
	0x39, 0x71, 0x4f, 0x25, 0xe4, 0x36, 0x9a, 0x0f, 0x7d, 0x3f, 0x33, 0x03, 0x77, 0x6b, 0x99, 0xb1,
	0x8d, 0x92, 0x82, 0xef, 0x45, 0x26, 0xd5, 0x3a, 0xdf, 0xed, 0x19, 0xcf, 0x05, 0x65, 0x1b, 0x5e,
	0x57, 0xe3, 0x72, 0x99, 0xd3, 0x86, 0xe7, 0x4c, 0xd6, 0x2b, 0xbe, 0xce, 0x0f, 0x57, 0x79, 0x49,
	0x0d, 0x65, 0x55, 0x6d, 0x2a, 0xa5, 0xb3, 0x46, 0x49, 0x23, 0x71, 0xe4, 0xc1, 0xe1, 0xed, 0x79,
	0x44, 0x42, 0xc8, 0xda, 0x12, 0x69, 0x43, 0x4d, 0xd5, 0x72, 0x8c, 0x9e, 0x21, 0x9a, 0x3f, 0x71,
	0xc3, 0x36, 0x18, 0x43, 0xaf, 0xa6, 0xa2, 0x22, 0x28, 0x41, 0x69, 0xbf, 0x70, 0x35, 0x4e, 0x20,
	0x2e, 0x2b, 0xcd, 0x14, 0x6f, 0x0c, 0x97, 0x35, 0xf9, 0xe2, 0xa0, 0x8f, 0x2d, 0xfc, 0x0b, 0xc2,
	0x46, 0x2a, 0xa3, 0x49, 0x90, 0xa0, 0x34, 0x2c, 0xfc, 0x03, 0xff, 0x83, 0xd0, 0xa9, 0x90, 0x5e,
	0x82, 0xd2, 0xc1, 0xe4, 0x47, 0xe6, 0xd5, 0xb3, 0xb9, 0x6d, 0x16, 0x1e, 0x1b, 0xbd, 0x21, 0x08,
	0xa6, 0x8b, 0xd9, 0x85, 0xc2, 0x43, 0xf8, 0xd6, 0xec, 0xa8, 0x59, 0x49, 0x25, 0x9c, 0x76, 0xbf,
	0x38, 0xbd, 0x2d, 0xa3, 0xa2, 0x6c, 0xeb, 0xd4, 0xfb, 0x85, 0xab, 0xf1, 0x5f, 0x88, 0x05, 0x65,
	0x8f, 0xb4, 0x2c, 0x55, 0xa5, 0x35, 0x09, 0x1d, 0x04, 0x82, 0xb2, 0x7b, 0xdf, 0xb1, 0x9f, 0x78,
	0x73, 0xb8, 0x21, 0x91, 0xff, 0x64, 0xeb, 0x2e, 0xc7, 0xd7, 0x4f, 0x72, 0xbc, 0x20, 0xe8, 0x15,
	0x56, 0xe2, 0xb2, 0x20, 0xff, 0x21, 0xd2, 0x6e, 0x03, 0x24, 0x48, 0x82, 0x34, 0x9e, 0x0c, 0x32,
	0xbf, 0xd6, 0xcc, 0xef, 0xa5, 0x68, 0xd1, 0xb3, 0x66, 0x8a, 0x7f, 0x42, 0xb0, 0x3d, 0x88, 0x36,
	0x9d, 0x2d, 0x47, 0xaf, 0x08, 0xe0, 0xe1, 0x74, 0x3a, 0x17, 0x7a, 0x4c, 0xda, 0x81, 0x7a, 0x87,
	0xdf, 0x8f, 0x0e, 0x6d, 0xea, 0x76, 0xbc, 0x67, 0xb9, 0xfb, 0x7d, 0x74, 0x67, 0x59, 0xe2, 0x23,
	0xcb, 0x74, 0x31, 0xf3, 0x56, 0xc7, 0x10, 0x77, 0x4e, 0x35, 0xfe, 0x03, 0xd0, 0xdd, 0x3c, 0x41,
	0x49, 0x60, 0x17, 0xd6, 0x75, 0x96, 0x91, 0xbb, 0xe0, 0xeb, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa9, 0x6e, 0x8a, 0xc9, 0x48, 0x03, 0x00, 0x00,
}
