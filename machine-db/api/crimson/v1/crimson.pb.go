// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto

/*
Package crimson is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/datacenters.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/hosts.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/racks.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/switches.proto
	go.chromium.org/luci/machine-db/api/crimson/v1/vlans.proto

It has these top-level messages:
	DatacentersRequest
	Datacenter
	DatacentersResponse
	Host
	AddHostRequest
	AddHostResponse
	DeleteHostRequest
	DeleteHostResponse
	EditHostRequest
	EditHostResponse
	GetHostsRequest
	GetHostsResponse
	Machine
	AddMachineRequest
	AddMachineResponse
	DeleteMachineRequest
	DeleteMachineResponse
	EditMachineRequest
	EditMachineResponse
	GetMachinesRequest
	GetMachinesResponse
	OperatingSystemsRequest
	OperatingSystem
	OperatingSystemsResponse
	PlatformsRequest
	Platform
	PlatformsResponse
	RacksRequest
	Rack
	RacksResponse
	SwitchesRequest
	Switch
	SwitchesResponse
	VLANsRequest
	VLAN
	VLANsResponse
*/
package crimson

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Crimson service

type CrimsonClient interface {
	// GetDatacenters retrieves datacenters.
	GetDatacenters(ctx context.Context, in *DatacentersRequest, opts ...grpc.CallOption) (*DatacentersResponse, error)
	// GetRacks retrieves racks.
	GetRacks(ctx context.Context, in *RacksRequest, opts ...grpc.CallOption) (*RacksResponse, error)
}
type crimsonPRPCClient struct {
	client *prpc.Client
}

func NewCrimsonPRPCClient(client *prpc.Client) CrimsonClient {
	return &crimsonPRPCClient{client}
}

func (c *crimsonPRPCClient) GetDatacenters(ctx context.Context, in *DatacentersRequest, opts ...grpc.CallOption) (*DatacentersResponse, error) {
	out := new(DatacentersResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "GetDatacenters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonPRPCClient) GetRacks(ctx context.Context, in *RacksRequest, opts ...grpc.CallOption) (*RacksResponse, error) {
	out := new(RacksResponse)
	err := c.client.Call(ctx, "crimson.Crimson", "GetRacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type crimsonClient struct {
	cc *grpc.ClientConn
}

func NewCrimsonClient(cc *grpc.ClientConn) CrimsonClient {
	return &crimsonClient{cc}
}

func (c *crimsonClient) GetDatacenters(ctx context.Context, in *DatacentersRequest, opts ...grpc.CallOption) (*DatacentersResponse, error) {
	out := new(DatacentersResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/GetDatacenters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crimsonClient) GetRacks(ctx context.Context, in *RacksRequest, opts ...grpc.CallOption) (*RacksResponse, error) {
	out := new(RacksResponse)
	err := grpc.Invoke(ctx, "/crimson.Crimson/GetRacks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Crimson service

type CrimsonServer interface {
	// GetDatacenters retrieves datacenters.
	GetDatacenters(context.Context, *DatacentersRequest) (*DatacentersResponse, error)
	// GetRacks retrieves racks.
	GetRacks(context.Context, *RacksRequest) (*RacksResponse, error)
}

func RegisterCrimsonServer(s prpc.Registrar, srv CrimsonServer) {
	s.RegisterService(&_Crimson_serviceDesc, srv)
}

func _Crimson_GetDatacenters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatacentersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).GetDatacenters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/GetDatacenters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).GetDatacenters(ctx, req.(*DatacentersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crimson_GetRacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrimsonServer).GetRacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crimson.Crimson/GetRacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrimsonServer).GetRacks(ctx, req.(*RacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crimson_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crimson.Crimson",
	HandlerType: (*CrimsonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatacenters",
			Handler:    _Crimson_GetDatacenters_Handler,
		},
		{
			MethodName: "GetRacks",
			Handler:    _Crimson_GetRacks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/crimson.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0xcf, 0x4d, 0x4c, 0xce, 0xc8, 0xcc, 0x4b, 0xd5, 0x4d, 0x49, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2e, 0xca, 0xcc, 0x2d, 0xce, 0xcf, 0xd3, 0x2f, 0x33, 0x84, 0x31, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xd8, 0xa1, 0x5c, 0x29, 0x07, 0x12, 0x8d, 0x49, 0x49, 0x2c, 0x49, 0x4c, 0x4e,
	0xcd, 0x2b, 0x49, 0x2d, 0x2a, 0x86, 0x18, 0x25, 0x65, 0x45, 0xa2, 0x09, 0x45, 0x89, 0xc9, 0xd9,
	0x50, 0xbd, 0x46, 0x13, 0x19, 0xb9, 0xd8, 0x9d, 0x21, 0x52, 0x42, 0xde, 0x5c, 0x7c, 0xee, 0xa9,
	0x25, 0x2e, 0x08, 0xf3, 0x85, 0xa4, 0xf5, 0x60, 0x8e, 0x46, 0x12, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x91, 0x92, 0xc1, 0x2e, 0x59, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xc9, 0xc5,
	0xe1, 0x9e, 0x5a, 0x12, 0x04, 0xb2, 0x4a, 0x48, 0x14, 0xae, 0x12, 0xcc, 0x87, 0x19, 0x20, 0x86,
	0x2e, 0x0c, 0xd1, 0x9a, 0xc4, 0x06, 0x76, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xae,
	0xbd, 0x90, 0x61, 0x01, 0x00, 0x00,
}
