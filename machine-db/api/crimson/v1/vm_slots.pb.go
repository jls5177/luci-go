// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/vm_slots.proto

package crimson

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

// A request to find available VM slots in the database.
type FindVMSlotsRequest struct {
	// The number of available VM slots to find.
	// Values < 1 return all available VM slots.
	Slots int32 `protobuf:"varint,1,opt,name=slots,proto3" json:"slots,omitempty"`
	// The platform manufacturers to filter found VM slots on.
	Manufacturers []string `protobuf:"bytes,2,rep,name=manufacturers,proto3" json:"manufacturers,omitempty"`
	// The virtual datacenter to filter found VM slots on.
	VirtualDatacenters []string `protobuf:"bytes,3,rep,name=virtual_datacenters,json=virtualDatacenters,proto3" json:"virtual_datacenters,omitempty"`
	// The states to filter found VM slots on.
	States               []v1.State `protobuf:"varint,6,rep,packed,name=states,proto3,enum=common.State" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindVMSlotsRequest) Reset()         { *m = FindVMSlotsRequest{} }
func (m *FindVMSlotsRequest) String() string { return proto.CompactTextString(m) }
func (*FindVMSlotsRequest) ProtoMessage()    {}
func (*FindVMSlotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_620049fabde60637, []int{0}
}

func (m *FindVMSlotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindVMSlotsRequest.Unmarshal(m, b)
}
func (m *FindVMSlotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindVMSlotsRequest.Marshal(b, m, deterministic)
}
func (m *FindVMSlotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindVMSlotsRequest.Merge(m, src)
}
func (m *FindVMSlotsRequest) XXX_Size() int {
	return xxx_messageInfo_FindVMSlotsRequest.Size(m)
}
func (m *FindVMSlotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindVMSlotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindVMSlotsRequest proto.InternalMessageInfo

func (m *FindVMSlotsRequest) GetSlots() int32 {
	if m != nil {
		return m.Slots
	}
	return 0
}

func (m *FindVMSlotsRequest) GetManufacturers() []string {
	if m != nil {
		return m.Manufacturers
	}
	return nil
}

func (m *FindVMSlotsRequest) GetVirtualDatacenters() []string {
	if m != nil {
		return m.VirtualDatacenters
	}
	return nil
}

func (m *FindVMSlotsRequest) GetStates() []v1.State {
	if m != nil {
		return m.States
	}
	return nil
}

// A response containing a list of available VM slots in the database.
type FindVMSlotsResponse struct {
	// The hosts with available VM slots.
	// Only includes name, vlan_id, and vm_slots.
	// vm_slots in this context means the number of available VM slots.
	Hosts                []*PhysicalHost `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FindVMSlotsResponse) Reset()         { *m = FindVMSlotsResponse{} }
func (m *FindVMSlotsResponse) String() string { return proto.CompactTextString(m) }
func (*FindVMSlotsResponse) ProtoMessage()    {}
func (*FindVMSlotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_620049fabde60637, []int{1}
}

func (m *FindVMSlotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindVMSlotsResponse.Unmarshal(m, b)
}
func (m *FindVMSlotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindVMSlotsResponse.Marshal(b, m, deterministic)
}
func (m *FindVMSlotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindVMSlotsResponse.Merge(m, src)
}
func (m *FindVMSlotsResponse) XXX_Size() int {
	return xxx_messageInfo_FindVMSlotsResponse.Size(m)
}
func (m *FindVMSlotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindVMSlotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindVMSlotsResponse proto.InternalMessageInfo

func (m *FindVMSlotsResponse) GetHosts() []*PhysicalHost {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func init() {
	proto.RegisterType((*FindVMSlotsRequest)(nil), "crimson.FindVMSlotsRequest")
	proto.RegisterType((*FindVMSlotsResponse)(nil), "crimson.FindVMSlotsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/vm_slots.proto", fileDescriptor_620049fabde60637)
}

var fileDescriptor_620049fabde60637 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0xa9, 0x65, 0x13, 0x23, 0xf3, 0x90, 0x29, 0x94, 0x9d, 0xca, 0x50, 0x28, 0x88, 0x09,
	0xce, 0x9b, 0xe0, 0x45, 0x45, 0xbc, 0x08, 0xd2, 0x81, 0xd7, 0x91, 0xa5, 0x71, 0x0d, 0x34, 0xf9,
	0x6a, 0xbe, 0xa4, 0xe0, 0x7f, 0xe4, 0x9f, 0x29, 0x6d, 0x23, 0xb2, 0x9b, 0x1e, 0xf3, 0xde, 0xcb,
	0xe3, 0xf7, 0x3d, 0x72, 0xb7, 0x03, 0x26, 0x6b, 0x07, 0x46, 0x07, 0xc3, 0xc0, 0xed, 0x78, 0x13,
	0xa4, 0xe6, 0x46, 0xc8, 0x5a, 0x5b, 0x75, 0x55, 0x6d, 0xb9, 0x68, 0x35, 0x97, 0x4e, 0x1b, 0x04,
	0xcb, 0xbb, 0x6b, 0xde, 0x99, 0x0d, 0x36, 0xe0, 0x91, 0xb5, 0x0e, 0x3c, 0xd0, 0xc3, 0x68, 0x2d,
	0x1e, 0xfe, 0xd9, 0xd3, 0xd6, 0x9f, 0xa8, 0xa5, 0x68, 0x36, 0x35, 0xe0, 0x4f, 0xdb, 0xe2, 0xf6,
	0x4f, 0x25, 0x60, 0xcc, 0xd8, 0x81, 0x5e, 0x78, 0x15, 0xff, 0x2e, 0xbf, 0x12, 0x42, 0x9f, 0xb4,
	0xad, 0xde, 0x5e, 0xd6, 0x3d, 0x5f, 0xa9, 0x3e, 0x82, 0x42, 0x4f, 0x4f, 0xc9, 0x64, 0xe0, 0xcd,
	0x92, 0x3c, 0x29, 0x26, 0xe5, 0xf8, 0xa0, 0xe7, 0x64, 0x66, 0x84, 0x0d, 0xef, 0x42, 0xfa, 0xe0,
	0x94, 0xc3, 0xec, 0x20, 0x4f, 0x8b, 0xa3, 0x72, 0x5f, 0xa4, 0x9c, 0xcc, 0x3b, 0xed, 0x7c, 0x10,
	0xcd, 0xa6, 0x12, 0x5e, 0x48, 0x65, 0x7d, 0x9f, 0x4d, 0x87, 0x2c, 0x8d, 0xd6, 0xe3, 0xaf, 0x43,
	0x2f, 0xc8, 0x74, 0x64, 0xca, 0xa6, 0x79, 0x5a, 0x9c, 0xac, 0x66, 0x6c, 0x84, 0x65, 0xeb, 0x5e,
	0x2d, 0xa3, 0xb9, 0xbc, 0x27, 0xf3, 0x3d, 0x52, 0x6c, 0xc1, 0xa2, 0xa2, 0x97, 0x64, 0x32, 0x8c,
	0x91, 0x25, 0x79, 0x5a, 0x1c, 0xaf, 0xce, 0x58, 0x9c, 0x8b, 0xbd, 0xc6, 0xad, 0x9e, 0x01, 0x7d,
	0x39, 0x66, 0xb6, 0xd3, 0xe1, 0xea, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xb1, 0xcf,
	0x39, 0xc0, 0x01, 0x00, 0x00,
}
