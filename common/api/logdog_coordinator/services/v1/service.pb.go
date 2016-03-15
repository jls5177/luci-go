// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package logdog is a generated protocol buffer package.

It is generated from these files:
	service.proto
	state.proto
	tasks.proto

It has these top-level messages:
	GetConfigResponse
	RegisterStreamRequest
	RegisterStreamResponse
	LoadStreamRequest
	LoadStreamResponse
	TerminateStreamRequest
	ArchiveStreamRequest
	LogStreamState
	ArchiveTask
*/
package logdog

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import logpb "github.com/luci/luci-go/common/proto/logdog/logpb"
import google_protobuf2 "github.com/luci/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GetConfigResponse is the response structure for the user
// "GetConfig" endpoint.
type GetConfigResponse struct {
	// The API URL of the base "luci-config" service. If empty, the default
	// service URL will be used.
	ConfigServiceUrl string `protobuf:"bytes,1,opt,name=config_service_url" json:"config_service_url,omitempty"`
	// The name of the configuration set to load from.
	ConfigSet string `protobuf:"bytes,2,opt,name=config_set" json:"config_set,omitempty"`
	// The path of the text-serialized configuration protobuf.
	ConfigPath string `protobuf:"bytes,3,opt,name=config_path" json:"config_path,omitempty"`
}

func (m *GetConfigResponse) Reset()                    { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()               {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// RegisterStreamRequest is the set of caller-supplied data for the
// RegisterStream Coordinator service endpoint.
type RegisterStreamRequest struct {
	// The log stream's path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The log stream's secret.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// The protobuf version string for this stream.
	ProtoVersion string `protobuf:"bytes,3,opt,name=proto_version" json:"proto_version,omitempty"`
	// The serialized LogStreamDescriptor protobuf for this stream.
	Desc *logpb.LogStreamDescriptor `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
}

func (m *RegisterStreamRequest) Reset()                    { *m = RegisterStreamRequest{} }
func (m *RegisterStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterStreamRequest) ProtoMessage()               {}
func (*RegisterStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterStreamRequest) GetDesc() *logpb.LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

// The response message for the RegisterStream RPC.
type RegisterStreamResponse struct {
	// The state of the requested log stream.
	State *LogStreamState `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	// The log stream's secret.
	//
	// Note that the secret is returned! This is okay, since this endpoint is only
	// accessible to trusted services. The secret can be cached by services to
	// validate stream information without needing to ping the Coordinator in
	// between each update.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *RegisterStreamResponse) Reset()                    { *m = RegisterStreamResponse{} }
func (m *RegisterStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterStreamResponse) ProtoMessage()               {}
func (*RegisterStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterStreamResponse) GetState() *LogStreamState {
	if m != nil {
		return m.State
	}
	return nil
}

// LoadStreamRequest loads the current state of a log stream.
type LoadStreamRequest struct {
	// The log stream's path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// If true, include the log stream descriptor.
	Desc bool `protobuf:"varint,2,opt,name=desc" json:"desc,omitempty"`
}

func (m *LoadStreamRequest) Reset()                    { *m = LoadStreamRequest{} }
func (m *LoadStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadStreamRequest) ProtoMessage()               {}
func (*LoadStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// The response message for the LoadStream RPC.
type LoadStreamResponse struct {
	// The state of the requested log stream.
	State *LogStreamState `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	// If requested, the serialized log stream descriptor. The protobuf version
	// of this descriptor will match the "proto_version" field in "state".
	Desc []byte `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (m *LoadStreamResponse) Reset()                    { *m = LoadStreamResponse{} }
func (m *LoadStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*LoadStreamResponse) ProtoMessage()               {}
func (*LoadStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoadStreamResponse) GetState() *LogStreamState {
	if m != nil {
		return m.State
	}
	return nil
}

// TerminateStreamRequest is the set of caller-supplied data for the
// TerminateStream service endpoint.
type TerminateStreamRequest struct {
	// The log stream's path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The log stream's secret.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// The terminal index of the stream.
	TerminalIndex int64 `protobuf:"varint,3,opt,name=terminal_index" json:"terminal_index,omitempty"`
}

func (m *TerminateStreamRequest) Reset()                    { *m = TerminateStreamRequest{} }
func (m *TerminateStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminateStreamRequest) ProtoMessage()               {}
func (*TerminateStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// ArchiveStreamRequest is the set of caller-supplied data for the ArchiveStream
// service endpoint.
type ArchiveStreamRequest struct {
	// The path of the log stream that was archived.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// If true, the archive includes the full set of log stream entries between
	// [0..terminal_index].
	Complete bool `protobuf:"varint,2,opt,name=complete" json:"complete,omitempty"`
	// The highest log stream index that was archived.
	TerminalIndex int64 `protobuf:"varint,3,opt,name=terminal_index" json:"terminal_index,omitempty"`
	// The archive URL of the log stream's stream data.
	StreamUrl string `protobuf:"bytes,10,opt,name=stream_url" json:"stream_url,omitempty"`
	// The size of the log stream's stream data.
	StreamSize int64 `protobuf:"varint,11,opt,name=stream_size" json:"stream_size,omitempty"`
	// The archive URL of the log stream's index data.
	IndexUrl string `protobuf:"bytes,20,opt,name=index_url" json:"index_url,omitempty"`
	// The size of the log stream's index data.
	IndexSize int64 `protobuf:"varint,21,opt,name=index_size" json:"index_size,omitempty"`
	// The archive URL of the log stream's binary data.
	DataUrl string `protobuf:"bytes,30,opt,name=data_url" json:"data_url,omitempty"`
	// The size of the log stream's binary data.
	DataSize int64 `protobuf:"varint,31,opt,name=data_size" json:"data_size,omitempty"`
}

func (m *ArchiveStreamRequest) Reset()                    { *m = ArchiveStreamRequest{} }
func (m *ArchiveStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*ArchiveStreamRequest) ProtoMessage()               {}
func (*ArchiveStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*GetConfigResponse)(nil), "logdog.GetConfigResponse")
	proto.RegisterType((*RegisterStreamRequest)(nil), "logdog.RegisterStreamRequest")
	proto.RegisterType((*RegisterStreamResponse)(nil), "logdog.RegisterStreamResponse")
	proto.RegisterType((*LoadStreamRequest)(nil), "logdog.LoadStreamRequest")
	proto.RegisterType((*LoadStreamResponse)(nil), "logdog.LoadStreamResponse")
	proto.RegisterType((*TerminateStreamRequest)(nil), "logdog.TerminateStreamRequest")
	proto.RegisterType((*ArchiveStreamRequest)(nil), "logdog.ArchiveStreamRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Services service

type ServicesClient interface {
	// GetConfig allows a service to retrieve the current service configuration
	// parameters.
	GetConfig(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// RegisterStream is an idempotent stream state register operation.
	RegisterStream(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (*RegisterStreamResponse, error)
	// LoadStream loads the current state of a log stream.
	LoadStream(ctx context.Context, in *LoadStreamRequest, opts ...grpc.CallOption) (*LoadStreamResponse, error)
	// TerminateStream is an idempotent operation to update the stream's terminal
	// index.
	TerminateStream(ctx context.Context, in *TerminateStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// ArchiveStream is an idempotent operation to record a log stream's archival
	// parameters. It is used by the Archivist service upon successful stream
	// archival.
	ArchiveStream(ctx context.Context, in *ArchiveStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}
type servicesPRPCClient struct {
	client *prpccommon.Client
}

func NewServicesPRPCClient(client *prpccommon.Client) ServicesClient {
	return &servicesPRPCClient{client}
}

func (c *servicesPRPCClient) GetConfig(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.client.Call(ctx, "logdog.Services", "GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesPRPCClient) RegisterStream(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (*RegisterStreamResponse, error) {
	out := new(RegisterStreamResponse)
	err := c.client.Call(ctx, "logdog.Services", "RegisterStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesPRPCClient) LoadStream(ctx context.Context, in *LoadStreamRequest, opts ...grpc.CallOption) (*LoadStreamResponse, error) {
	out := new(LoadStreamResponse)
	err := c.client.Call(ctx, "logdog.Services", "LoadStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesPRPCClient) TerminateStream(ctx context.Context, in *TerminateStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := c.client.Call(ctx, "logdog.Services", "TerminateStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesPRPCClient) ArchiveStream(ctx context.Context, in *ArchiveStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := c.client.Call(ctx, "logdog.Services", "ArchiveStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type servicesClient struct {
	cc *grpc.ClientConn
}

func NewServicesClient(cc *grpc.ClientConn) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) GetConfig(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := grpc.Invoke(ctx, "/logdog.Services/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) RegisterStream(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (*RegisterStreamResponse, error) {
	out := new(RegisterStreamResponse)
	err := grpc.Invoke(ctx, "/logdog.Services/RegisterStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) LoadStream(ctx context.Context, in *LoadStreamRequest, opts ...grpc.CallOption) (*LoadStreamResponse, error) {
	out := new(LoadStreamResponse)
	err := grpc.Invoke(ctx, "/logdog.Services/LoadStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) TerminateStream(ctx context.Context, in *TerminateStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/logdog.Services/TerminateStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) ArchiveStream(ctx context.Context, in *ArchiveStreamRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/logdog.Services/ArchiveStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Services service

type ServicesServer interface {
	// GetConfig allows a service to retrieve the current service configuration
	// parameters.
	GetConfig(context.Context, *google_protobuf2.Empty) (*GetConfigResponse, error)
	// RegisterStream is an idempotent stream state register operation.
	RegisterStream(context.Context, *RegisterStreamRequest) (*RegisterStreamResponse, error)
	// LoadStream loads the current state of a log stream.
	LoadStream(context.Context, *LoadStreamRequest) (*LoadStreamResponse, error)
	// TerminateStream is an idempotent operation to update the stream's terminal
	// index.
	TerminateStream(context.Context, *TerminateStreamRequest) (*google_protobuf2.Empty, error)
	// ArchiveStream is an idempotent operation to record a log stream's archival
	// parameters. It is used by the Archivist service upon successful stream
	// archival.
	ArchiveStream(context.Context, *ArchiveStreamRequest) (*google_protobuf2.Empty, error)
}

func RegisterServicesServer(s prpc.Registrar, srv ServicesServer) {
	s.RegisterService(&_Services_serviceDesc, srv)
}

func _Services_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).GetConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Services_RegisterStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RegisterStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).RegisterStream(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Services_LoadStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LoadStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).LoadStream(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Services_TerminateStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TerminateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).TerminateStream(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Services_ArchiveStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ArchiveStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ServicesServer).ArchiveStream(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Services_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logdog.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Services_GetConfig_Handler,
		},
		{
			MethodName: "RegisterStream",
			Handler:    _Services_RegisterStream_Handler,
		},
		{
			MethodName: "LoadStream",
			Handler:    _Services_LoadStream_Handler,
		},
		{
			MethodName: "TerminateStream",
			Handler:    _Services_TerminateStream_Handler,
		},
		{
			MethodName: "ArchiveStream",
			Handler:    _Services_ArchiveStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xa5, 0xbb, 0xb5, 0xb4, 0xb7, 0x1f, 0xda, 0x71, 0x5b, 0xba, 0x51, 0x57, 0x09, 0x08, 0xfb,
	0x62, 0x02, 0xf5, 0x51, 0x10, 0x64, 0x5d, 0x64, 0x61, 0x71, 0xa1, 0xf5, 0xc1, 0xb7, 0x92, 0x26,
	0x77, 0xd3, 0x81, 0x24, 0x13, 0x67, 0x26, 0x45, 0xfd, 0x7b, 0xfe, 0x23, 0x7f, 0x81, 0xc9, 0x9d,
	0x49, 0xbb, 0x5b, 0xdb, 0x85, 0x7d, 0x49, 0x98, 0x73, 0xcf, 0x3d, 0xf7, 0xeb, 0x40, 0x5f, 0xa1,
	0x5c, 0xf3, 0x10, 0xbd, 0x5c, 0x0a, 0x2d, 0x58, 0x2b, 0x11, 0x71, 0x24, 0x62, 0xa7, 0xab, 0x74,
	0xa0, 0x2d, 0xe8, 0x7c, 0x88, 0xb9, 0x5e, 0x15, 0x4b, 0x2f, 0x14, 0xa9, 0x9f, 0x14, 0x21, 0xa7,
	0xcf, 0xbb, 0x58, 0xf8, 0x25, 0x90, 0x8a, 0xcc, 0x27, 0x96, 0x6f, 0x32, 0xab, 0x5f, 0xbe, 0xac,
	0xbe, 0x36, 0xf9, 0x45, 0x2c, 0x44, 0x9c, 0xa0, 0x21, 0x2d, 0x8b, 0x5b, 0x1f, 0xd3, 0x5c, 0xff,
	0x32, 0x41, 0xf7, 0x3b, 0x0c, 0xbf, 0xa0, 0xbe, 0x10, 0xd9, 0x2d, 0x8f, 0x67, 0xa8, 0x72, 0x91,
	0x29, 0x64, 0x0e, 0xb0, 0x90, 0x90, 0x85, 0xed, 0x6d, 0x51, 0xc8, 0x64, 0xd2, 0x78, 0xd3, 0x38,
	0xef, 0x30, 0x06, 0xb0, 0x89, 0xe9, 0xc9, 0x11, 0x61, 0xcf, 0xa1, 0x6b, 0xb1, 0x3c, 0xd0, 0xab,
	0xc9, 0x71, 0x05, 0xba, 0x6b, 0x18, 0xcd, 0x30, 0xe6, 0x4a, 0xa3, 0x9c, 0x6b, 0x89, 0x41, 0x3a,
	0xc3, 0x1f, 0x05, 0x2a, 0xcd, 0x7a, 0xd0, 0x24, 0x9a, 0xd1, 0x1b, 0x40, 0x4b, 0x61, 0x28, 0xad,
	0x56, 0x8f, 0x8d, 0xa0, 0x4f, 0x9d, 0x2d, 0xd6, 0x28, 0x15, 0x17, 0x99, 0x51, 0x63, 0xe7, 0xd0,
	0x8c, 0x50, 0x85, 0x93, 0x66, 0xf9, 0xea, 0x4e, 0x1d, 0x8f, 0x86, 0xf4, 0xae, 0x45, 0x6c, 0xb4,
	0x3f, 0x97, 0x31, 0xc9, 0x73, 0x2d, 0xa4, 0x7b, 0x03, 0xe3, 0xdd, 0xba, 0x76, 0xac, 0xb7, 0xf0,
	0x84, 0x96, 0x4a, 0x95, 0xbb, 0xd3, 0xb1, 0x67, 0x16, 0xb6, 0x55, 0x99, 0x57, 0xd1, 0xdd, 0x8e,
	0x5c, 0x1f, 0x86, 0xd7, 0x22, 0x88, 0x1e, 0x1a, 0xa2, 0x67, 0xbb, 0xab, 0x12, 0xda, 0xee, 0x15,
	0xb0, 0xbb, 0x09, 0x8f, 0xab, 0x7e, 0x57, 0xaa, 0xe7, 0x7e, 0x85, 0xf1, 0x37, 0x94, 0x29, 0xcf,
	0xca, 0xd0, 0x63, 0xb6, 0x38, 0x86, 0x81, 0x36, 0x79, 0xc9, 0x82, 0x67, 0x11, 0xfe, 0xa4, 0x35,
	0x1e, 0xbb, 0x7f, 0x1a, 0x70, 0xf2, 0x49, 0x86, 0x2b, 0xbe, 0x7e, 0x50, 0xee, 0x19, 0xb4, 0x4b,
	0x67, 0xe5, 0x09, 0x96, 0xed, 0xd2, 0x4c, 0x87, 0x04, 0x2b, 0x3b, 0x28, 0x12, 0x22, 0x8b, 0x40,
	0x6d, 0x07, 0x8b, 0x29, 0xfe, 0x1b, 0x27, 0x5d, 0x22, 0x0e, 0xa1, 0x43, 0x79, 0xc4, 0x3b, 0xa9,
	0xad, 0x64, 0x20, 0xa2, 0x8d, 0x88, 0x56, 0x56, 0x8e, 0x02, 0x1d, 0x10, 0xeb, 0x8c, 0x58, 0x65,
	0x22, 0x21, 0x44, 0x7a, 0x5d, 0x91, 0xa6, 0x7f, 0x8f, 0xa0, 0x3d, 0x37, 0xce, 0x54, 0xec, 0x23,
	0x74, 0x36, 0x0e, 0x66, 0x63, 0xcf, 0x98, 0xdd, 0xab, 0xcd, 0xee, 0x5d, 0x56, 0x66, 0x77, 0x4e,
	0xeb, 0x6d, 0xff, 0x6f, 0xf6, 0x1b, 0x18, 0xdc, 0xf7, 0x0b, 0x7b, 0x55, 0x93, 0xf7, 0xfa, 0xd7,
	0x39, 0x3b, 0x14, 0xb6, 0x82, 0x17, 0x00, 0xdb, 0xf3, 0xb3, 0xd3, 0xed, 0x9d, 0x77, 0x3c, 0xe4,
	0x38, 0xfb, 0x42, 0x56, 0xe4, 0x0a, 0x9e, 0xee, 0x1c, 0x9e, 0x6d, 0xea, 0xee, 0x77, 0x84, 0x73,
	0x60, 0x76, 0x76, 0x09, 0xfd, 0x7b, 0x27, 0x67, 0x2f, 0x6b, 0xa1, 0x7d, 0x4e, 0x38, 0x24, 0xb3,
	0x6c, 0xd1, 0xfb, 0xfd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x13, 0x5d, 0xbb, 0xb0, 0x04,
	0x00, 0x00,
}
