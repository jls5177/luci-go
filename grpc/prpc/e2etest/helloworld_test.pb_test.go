// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/grpc/prpc/e2etest/helloworld_test.proto

/*
Package e2etest is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/grpc/prpc/e2etest/helloworld_test.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package e2etest

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "e2etest.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "e2etest.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}
type helloPRPCClient struct {
	client *prpc.Client
}

func NewHelloPRPCClient(client *prpc.Client) HelloClient {
	return &helloPRPCClient{client}
}

func (c *helloPRPCClient) Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.client.Call(ctx, "e2etest.Hello", "Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/e2etest.Hello/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	Greet(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterHelloServer(s prpc.Registrar, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e2etest.Hello/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Greet(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "e2etest.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Hello_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/prpc/e2etest/helloworld_test.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/grpc/prpc/e2etest/helloworld_test.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x00, 0x11, 0xa9, 0x46, 0xa9, 0x25, 0xa9, 0xc5, 0x25,
	0xfa, 0x19, 0xa9, 0x39, 0x39, 0xf9, 0xe5, 0xf9, 0x45, 0x39, 0x29, 0xf1, 0x20, 0xbe, 0x5e, 0x41,
	0x51, 0x7e, 0x49, 0xbe, 0x10, 0x3b, 0x54, 0x5a, 0x49, 0x89, 0x8b, 0xc7, 0x03, 0xa4, 0x22, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x52, 0xe3, 0xe2, 0x82, 0xaa, 0x29, 0xc8, 0xa9, 0x14,
	0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0x29, 0x82, 0x71, 0x8d, 0x6c, 0xb8,
	0x58, 0xc1, 0xea, 0x84, 0x8c, 0xb9, 0x58, 0xdd, 0x8b, 0x52, 0x53, 0x4b, 0x84, 0x44, 0xf5, 0xa0,
	0xf6, 0xe8, 0x21, 0x5b, 0x22, 0x25, 0x8c, 0x2e, 0x5c, 0x90, 0x53, 0x99, 0xc4, 0x06, 0x76, 0x99,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xab, 0x37, 0xd6, 0x0f, 0xd9, 0x00, 0x00, 0x00,
}
