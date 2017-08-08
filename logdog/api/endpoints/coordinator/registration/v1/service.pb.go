// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto

/*
Package logdog is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto

It has these top-level messages:
	RegisterPrefixRequest
	RegisterPrefixResponse
*/
package logdog

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

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

// RegisterPrefixRequest registers a new Prefix with the Coordinator.
type RegisterPrefixRequest struct {
	// The log stream's project.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The log stream prefix to register.
	Prefix string `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	// Optional information about the registering agent.
	SourceInfo []string `protobuf:"bytes,3,rep,name=source_info,json=sourceInfo" json:"source_info,omitempty"`
	// The prefix expiration time. If <= 0, the project's default prefix
	// expiration period will be applied.
	//
	// The prefix will be closed by the Coordinator after its expiration period.
	// Once closed, new stream registration requests will no longer be accepted.
	//
	// If supplied, this value should exceed the timeout of the local task, else
	// some of the task's streams may be dropped due to failing registration.
	Expiration *google_protobuf.Duration `protobuf:"bytes,10,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *RegisterPrefixRequest) Reset()                    { *m = RegisterPrefixRequest{} }
func (m *RegisterPrefixRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterPrefixRequest) ProtoMessage()               {}
func (*RegisterPrefixRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterPrefixRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RegisterPrefixRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RegisterPrefixRequest) GetSourceInfo() []string {
	if m != nil {
		return m.SourceInfo
	}
	return nil
}

func (m *RegisterPrefixRequest) GetExpiration() *google_protobuf.Duration {
	if m != nil {
		return m.Expiration
	}
	return nil
}

// The response message for the RegisterPrefix RPC.
type RegisterPrefixResponse struct {
	// Secret is the prefix's secret. This must be included verbatim in Butler
	// bundles to assert ownership of this prefix.
	Secret []byte `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	// The name of the Pub/Sub topic to publish butlerproto-formatted Butler log
	// bundles to.
	LogBundleTopic string `protobuf:"bytes,2,opt,name=log_bundle_topic,json=logBundleTopic" json:"log_bundle_topic,omitempty"`
}

func (m *RegisterPrefixResponse) Reset()                    { *m = RegisterPrefixResponse{} }
func (m *RegisterPrefixResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterPrefixResponse) ProtoMessage()               {}
func (*RegisterPrefixResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterPrefixResponse) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *RegisterPrefixResponse) GetLogBundleTopic() string {
	if m != nil {
		return m.LogBundleTopic
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterPrefixRequest)(nil), "logdog.RegisterPrefixRequest")
	proto.RegisterType((*RegisterPrefixResponse)(nil), "logdog.RegisterPrefixResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Registration service

type RegistrationClient interface {
	// RegisterStream allows a Butler instance to register a log stream with the
	// Coordinator. Upon success, the Coordinator will return registration
	// information and streaming parameters to the Butler.
	//
	// This should be called by a Butler instance to gain the ability to publish
	// to a prefix space. The caller must have WRITE access to its project's
	// stream space. If WRITE access is not present, this will fail with the
	// "PermissionDenied" gRPC code.
	//
	// A stream prefix may be registered at most once. Additional registration
	// requests will fail with the "AlreadyExists" gRPC code.
	RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error)
}
type registrationPRPCClient struct {
	client *prpc.Client
}

func NewRegistrationPRPCClient(client *prpc.Client) RegistrationClient {
	return &registrationPRPCClient{client}
}

func (c *registrationPRPCClient) RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error) {
	out := new(RegisterPrefixResponse)
	err := c.client.Call(ctx, "logdog.Registration", "RegisterPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error) {
	out := new(RegisterPrefixResponse)
	err := grpc.Invoke(ctx, "/logdog.Registration/RegisterPrefix", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registration service

type RegistrationServer interface {
	// RegisterStream allows a Butler instance to register a log stream with the
	// Coordinator. Upon success, the Coordinator will return registration
	// information and streaming parameters to the Butler.
	//
	// This should be called by a Butler instance to gain the ability to publish
	// to a prefix space. The caller must have WRITE access to its project's
	// stream space. If WRITE access is not present, this will fail with the
	// "PermissionDenied" gRPC code.
	//
	// A stream prefix may be registered at most once. Additional registration
	// requests will fail with the "AlreadyExists" gRPC code.
	RegisterPrefix(context.Context, *RegisterPrefixRequest) (*RegisterPrefixResponse, error)
}

func RegisterRegistrationServer(s prpc.Registrar, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_RegisterPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).RegisterPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Registration/RegisterPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).RegisterPrefix(ctx, req.(*RegisterPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logdog.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPrefix",
			Handler:    _Registration_RegisterPrefix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x99, 0xaf, 0xd0, 0x8f, 0xa6, 0xa5, 0x48, 0xc0, 0x32, 0x16, 0xac, 0xa5, 0xab, 0x59,
	0x25, 0x58, 0x57, 0x6e, 0xc5, 0x8d, 0x2b, 0x25, 0xba, 0x72, 0x33, 0x74, 0x32, 0x77, 0x62, 0x64,
	0x9a, 0x1b, 0xf3, 0xa7, 0xf4, 0x85, 0x7c, 0x4f, 0xe9, 0x64, 0x0a, 0x55, 0x74, 0x79, 0xce, 0x3d,
	0x37, 0xf9, 0xdd, 0x43, 0x9e, 0x15, 0x32, 0xf9, 0xe6, 0x70, 0xab, 0xe3, 0x96, 0xa1, 0x53, 0xbc,
	0x8d, 0x52, 0xf3, 0x16, 0x55, 0x8d, 0x8a, 0x6f, 0xac, 0xe6, 0x60, 0x6a, 0x8b, 0xda, 0x04, 0xcf,
	0x25, 0xa2, 0xab, 0xb5, 0xd9, 0x04, 0x74, 0xdc, 0x81, 0xd2, 0x3e, 0xb8, 0x4d, 0xd0, 0x68, 0xf8,
	0xee, 0x9a, 0x7b, 0x70, 0x3b, 0x2d, 0x81, 0x59, 0x87, 0x01, 0xe9, 0x30, 0xed, 0xcf, 0x17, 0x0a,
	0x51, 0xb5, 0xc0, 0x3b, 0xb7, 0x8a, 0x0d, 0xaf, 0x63, 0x5a, 0x49, 0xb9, 0xd5, 0x67, 0x46, 0xce,
	0x45, 0xf7, 0x12, 0xb8, 0x27, 0x07, 0x8d, 0xde, 0x0b, 0xf8, 0x88, 0xe0, 0x03, 0xcd, 0xc9, 0x7f,
	0xeb, 0xf0, 0x1d, 0x64, 0xc8, 0xb3, 0x65, 0x56, 0x8c, 0xc4, 0x51, 0xd2, 0x19, 0x19, 0xda, 0x2e,
	0x9a, 0xff, 0xeb, 0x06, 0xbd, 0xa2, 0x57, 0x64, 0xec, 0x31, 0x3a, 0x09, 0xa5, 0x36, 0x0d, 0xe6,
	0x83, 0xe5, 0xa0, 0x18, 0x09, 0x92, 0xac, 0x07, 0xd3, 0x20, 0xbd, 0x25, 0x04, 0xf6, 0x56, 0x27,
	0x80, 0x9c, 0x2c, 0xb3, 0x62, 0xbc, 0xbe, 0x60, 0x89, 0x90, 0x1d, 0x09, 0xd9, 0x7d, 0x4f, 0x28,
	0x4e, 0xc2, 0xab, 0x57, 0x32, 0xfb, 0x89, 0xe9, 0x2d, 0x1a, 0x0f, 0x07, 0x1a, 0x0f, 0xd2, 0x41,
	0xc2, 0x9c, 0x88, 0x5e, 0xd1, 0x82, 0x9c, 0xb5, 0xa8, 0xca, 0x2a, 0x9a, 0xba, 0x85, 0x32, 0xa0,
	0xd5, 0xb2, 0xe7, 0x9d, 0xb6, 0xa8, 0xee, 0x3a, 0xfb, 0xe5, 0xe0, 0xae, 0x4b, 0x32, 0x11, 0x27,
	0x65, 0xd2, 0x47, 0x32, 0xfd, 0xfe, 0x17, 0xbd, 0x64, 0xa9, 0x4e, 0xf6, 0x6b, 0x55, 0xf3, 0xc5,
	0x5f, 0xe3, 0x84, 0x58, 0x0d, 0xbb, 0xdb, 0x6e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x52,
	0x29, 0x53, 0xea, 0x01, 0x00, 0x00,
}
