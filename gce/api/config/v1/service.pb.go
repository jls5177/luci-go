// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/gce/api/config/v1/service.proto

package config

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// A request to create a VMs block.
type CreateVMsRequest struct {
	Vm                   *Block   `protobuf:"bytes,1,opt,name=vm,proto3" json:"vm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMsRequest) Reset()         { *m = CreateVMsRequest{} }
func (m *CreateVMsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVMsRequest) ProtoMessage()    {}
func (*CreateVMsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4344ecc87758d8d4, []int{0}
}

func (m *CreateVMsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMsRequest.Unmarshal(m, b)
}
func (m *CreateVMsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMsRequest.Marshal(b, m, deterministic)
}
func (m *CreateVMsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMsRequest.Merge(m, src)
}
func (m *CreateVMsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVMsRequest.Size(m)
}
func (m *CreateVMsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMsRequest proto.InternalMessageInfo

func (m *CreateVMsRequest) GetVm() *Block {
	if m != nil {
		return m.Vm
	}
	return nil
}

// A request to delete a VMs block.
type DeleteVMsRequest struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteVMsRequest) Reset()         { *m = DeleteVMsRequest{} }
func (m *DeleteVMsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVMsRequest) ProtoMessage()    {}
func (*DeleteVMsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4344ecc87758d8d4, []int{1}
}

func (m *DeleteVMsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVMsRequest.Unmarshal(m, b)
}
func (m *DeleteVMsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVMsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteVMsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVMsRequest.Merge(m, src)
}
func (m *DeleteVMsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteVMsRequest.Size(m)
}
func (m *DeleteVMsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVMsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVMsRequest proto.InternalMessageInfo

func (m *DeleteVMsRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// A request to get a VMs block.
type GetVMsRequest struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMsRequest) Reset()         { *m = GetVMsRequest{} }
func (m *GetVMsRequest) String() string { return proto.CompactTextString(m) }
func (*GetVMsRequest) ProtoMessage()    {}
func (*GetVMsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4344ecc87758d8d4, []int{2}
}

func (m *GetVMsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMsRequest.Unmarshal(m, b)
}
func (m *GetVMsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMsRequest.Marshal(b, m, deterministic)
}
func (m *GetVMsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMsRequest.Merge(m, src)
}
func (m *GetVMsRequest) XXX_Size() int {
	return xxx_messageInfo_GetVMsRequest.Size(m)
}
func (m *GetVMsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMsRequest proto.InternalMessageInfo

func (m *GetVMsRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVMsRequest)(nil), "config.CreateVMsRequest")
	proto.RegisterType((*DeleteVMsRequest)(nil), "config.DeleteVMsRequest")
	proto.RegisterType((*GetVMsRequest)(nil), "config.GetVMsRequest")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/gce/api/config/v1/service.proto", fileDescriptor_4344ecc87758d8d4)
}

var fileDescriptor_4344ecc87758d8d4 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0x0e, 0x81, 0xbe, 0x32, 0x18, 0x01, 0xc7, 0xa8, 0x08, 0xd2, 0x8b, 0xe2, 0x21,
	0x61, 0x7f, 0xae, 0x5e, 0x9c, 0xe2, 0xc9, 0x4b, 0x0f, 0xde, 0xb7, 0xf0, 0x36, 0x06, 0x1b, 0xdf,
	0x9a, 0xa6, 0x45, 0x3f, 0x94, 0xdf, 0x51, 0x4c, 0xd6, 0x49, 0x7b, 0x71, 0xb7, 0x36, 0x79, 0x7e,
	0x79, 0x7e, 0x3c, 0xb0, 0xd1, 0x24, 0xd4, 0xab, 0x23, 0x6b, 0x5a, 0x2b, 0xc8, 0x69, 0x59, 0xb5,
	0xca, 0x48, 0xad, 0x50, 0xee, 0x6a, 0x23, 0x15, 0xbd, 0x97, 0x46, 0xcb, 0x6e, 0x29, 0x1b, 0x74,
	0x9d, 0x51, 0x28, 0x6a, 0x47, 0x9e, 0x38, 0x8b, 0x17, 0xd9, 0x85, 0x26, 0xd2, 0x15, 0xca, 0x70,
	0xba, 0x6f, 0x4b, 0x89, 0xb6, 0xf6, 0x5f, 0x31, 0x94, 0xad, 0x4f, 0x7c, 0x3a, 0x7e, 0x45, 0x28,
	0x5f, 0xc2, 0x6c, 0xeb, 0x70, 0xe7, 0xf1, 0xe5, 0xb9, 0x29, 0xf0, 0xa3, 0xc5, 0xc6, 0xf3, 0x4b,
	0x98, 0x74, 0x76, 0x91, 0x5c, 0x25, 0x37, 0x67, 0xab, 0xa9, 0x38, 0xc4, 0xef, 0x2b, 0x52, 0x6f,
	0xc5, 0xa4, 0xb3, 0xf9, 0x2d, 0xcc, 0x1e, 0xb0, 0xc2, 0x01, 0x32, 0x07, 0x56, 0x3b, 0x2c, 0xcd,
	0x67, 0xc0, 0xd2, 0xe2, 0xf0, 0x97, 0x5f, 0xc3, 0xf4, 0x09, 0xfd, 0xff, 0xc1, 0xd5, 0x77, 0x02,
	0x6c, 0x1b, 0x9a, 0xf8, 0x06, 0xd2, 0xa3, 0x12, 0x5f, 0xf4, 0xfd, 0x63, 0xcb, 0x6c, 0x68, 0xc6,
	0xef, 0x20, 0x3d, 0x5a, 0xfd, 0x51, 0x63, 0xd1, 0x6c, 0x2e, 0xe2, 0x84, 0xa2, 0x9f, 0x50, 0x3c,
	0xfe, 0x4e, 0xc8, 0x05, 0xb0, 0x28, 0xca, 0xcf, 0x7b, 0x76, 0x20, 0x3e, 0xaa, 0xdb, 0xb3, 0xc0,
	0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x27, 0xa6, 0x23, 0xd0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigClient interface {
	// CreateVMs creates a new VMs block.
	// Internal API.
	CreateVMs(ctx context.Context, in *CreateVMsRequest, opts ...grpc.CallOption) (*Block, error)
	// DeleteVMs deletes an existing VMs block.
	// Internal API.
	DeleteVMs(ctx context.Context, in *DeleteVMsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetVMs returns a configured VMs block.
	GetVMs(ctx context.Context, in *GetVMsRequest, opts ...grpc.CallOption) (*Block, error)
}
type configPRPCClient struct {
	client *prpc.Client
}

func NewConfigPRPCClient(client *prpc.Client) ConfigClient {
	return &configPRPCClient{client}
}

func (c *configPRPCClient) CreateVMs(ctx context.Context, in *CreateVMsRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.client.Call(ctx, "config.Config", "CreateVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configPRPCClient) DeleteVMs(ctx context.Context, in *DeleteVMsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "config.Config", "DeleteVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configPRPCClient) GetVMs(ctx context.Context, in *GetVMsRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.client.Call(ctx, "config.Config", "GetVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type configClient struct {
	cc *grpc.ClientConn
}

func NewConfigClient(cc *grpc.ClientConn) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) CreateVMs(ctx context.Context, in *CreateVMsRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/config.Config/CreateVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) DeleteVMs(ctx context.Context, in *DeleteVMsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/config.Config/DeleteVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetVMs(ctx context.Context, in *GetVMsRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/config.Config/GetVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
type ConfigServer interface {
	// CreateVMs creates a new VMs block.
	// Internal API.
	CreateVMs(context.Context, *CreateVMsRequest) (*Block, error)
	// DeleteVMs deletes an existing VMs block.
	// Internal API.
	DeleteVMs(context.Context, *DeleteVMsRequest) (*empty.Empty, error)
	// GetVMs returns a configured VMs block.
	GetVMs(context.Context, *GetVMsRequest) (*Block, error)
}

func RegisterConfigServer(s prpc.Registrar, srv ConfigServer) {
	s.RegisterService(&_Config_serviceDesc, srv)
}

func _Config_CreateVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).CreateVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/CreateVMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).CreateVMs(ctx, req.(*CreateVMsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_DeleteVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVMsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).DeleteVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/DeleteVMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).DeleteVMs(ctx, req.(*DeleteVMsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVMsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Config/GetVMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetVMs(ctx, req.(*GetVMsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Config_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVMs",
			Handler:    _Config_CreateVMs_Handler,
		},
		{
			MethodName: "DeleteVMs",
			Handler:    _Config_DeleteVMs_Handler,
		},
		{
			MethodName: "GetVMs",
			Handler:    _Config_GetVMs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/gce/api/config/v1/service.proto",
}