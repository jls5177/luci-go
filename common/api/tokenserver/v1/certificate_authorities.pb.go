// Code generated by protoc-gen-go.
// source: certificate_authorities.proto
// DO NOT EDIT!

/*
Package tokenserver is a generated protocol buffer package.

It is generated from these files:
	certificate_authorities.proto
	config.proto
	service_accounts.proto

It has these top-level messages:
	ImportConfigResponse
	FetchCRLRequest
	FetchCRLResponse
	ListCAsResponse
	GetCAStatusRequest
	GetCAStatusResponse
	IsRevokedCertRequest
	IsRevokedCertResponse
	CheckCertificateRequest
	CheckCertificateResponse
	CRLStatus
	TokenServerConfig
	CertificateAuthorityConfig
	DomainConfig
	CreateServiceAccountRequest
	CreateServiceAccountResponse
	ServiceAccount
*/
package tokenserver

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf1 "github.com/luci/luci-go/common/proto/google"

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
const _ = proto.ProtoPackageIsVersion1

// ImportConfigResponse is returned by ImportConfig on success.
type ImportConfigResponse struct {
	Revision string `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
}

func (m *ImportConfigResponse) Reset()                    { *m = ImportConfigResponse{} }
func (m *ImportConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportConfigResponse) ProtoMessage()               {}
func (*ImportConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// FetchCRLRequest identifies a name of CA to fetch CRL for.
type FetchCRLRequest struct {
	Cn    string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
	Force bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *FetchCRLRequest) Reset()                    { *m = FetchCRLRequest{} }
func (m *FetchCRLRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLRequest) ProtoMessage()               {}
func (*FetchCRLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// FetchCRLResponse is returned by FetchCRL.
type FetchCRLResponse struct {
	CrlStatus *CRLStatus `protobuf:"bytes,1,opt,name=crl_status,json=crlStatus" json:"crl_status,omitempty"`
}

func (m *FetchCRLResponse) Reset()                    { *m = FetchCRLResponse{} }
func (m *FetchCRLResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLResponse) ProtoMessage()               {}
func (*FetchCRLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchCRLResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// ListCAsResponse is returned by ListCAs.
type ListCAsResponse struct {
	Cn []string `protobuf:"bytes,1,rep,name=cn" json:"cn,omitempty"`
}

func (m *ListCAsResponse) Reset()                    { *m = ListCAsResponse{} }
func (m *ListCAsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCAsResponse) ProtoMessage()               {}
func (*ListCAsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetCAStatusRequest identifies a name of CA to fetch.
type GetCAStatusRequest struct {
	Cn string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
}

func (m *GetCAStatusRequest) Reset()                    { *m = GetCAStatusRequest{} }
func (m *GetCAStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusRequest) ProtoMessage()               {}
func (*GetCAStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GetCAStatusResponse is returned by GetCAStatus method.
//
// If requested CA doesn't exist, all fields are empty.
type GetCAStatusResponse struct {
	Config     *CertificateAuthorityConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Cert       string                      `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
	Removed    bool                        `protobuf:"varint,3,opt,name=removed" json:"removed,omitempty"`
	Ready      bool                        `protobuf:"varint,4,opt,name=ready" json:"ready,omitempty"`
	AddedRev   string                      `protobuf:"bytes,5,opt,name=added_rev,json=addedRev" json:"added_rev,omitempty"`
	UpdatedRev string                      `protobuf:"bytes,6,opt,name=updated_rev,json=updatedRev" json:"updated_rev,omitempty"`
	RemovedRev string                      `protobuf:"bytes,7,opt,name=removed_rev,json=removedRev" json:"removed_rev,omitempty"`
	CrlStatus  *CRLStatus                  `protobuf:"bytes,8,opt,name=crl_status,json=crlStatus" json:"crl_status,omitempty"`
}

func (m *GetCAStatusResponse) Reset()                    { *m = GetCAStatusResponse{} }
func (m *GetCAStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusResponse) ProtoMessage()               {}
func (*GetCAStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetCAStatusResponse) GetConfig() *CertificateAuthorityConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GetCAStatusResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// IsRevokedCertRequest contains a name of the CA and a cert serial number.
type IsRevokedCertRequest struct {
	Ca string `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Sn string `protobuf:"bytes,2,opt,name=sn" json:"sn,omitempty"`
}

func (m *IsRevokedCertRequest) Reset()                    { *m = IsRevokedCertRequest{} }
func (m *IsRevokedCertRequest) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertRequest) ProtoMessage()               {}
func (*IsRevokedCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// IsRevokedCertResponse is returned by IsRevokedCert
type IsRevokedCertResponse struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked" json:"revoked,omitempty"`
}

func (m *IsRevokedCertResponse) Reset()                    { *m = IsRevokedCertResponse{} }
func (m *IsRevokedCertResponse) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertResponse) ProtoMessage()               {}
func (*IsRevokedCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// CheckCertificateRequest contains a pem encoded certificate to check.
type CheckCertificateRequest struct {
	CertPem string `protobuf:"bytes,1,opt,name=cert_pem,json=certPem" json:"cert_pem,omitempty"`
}

func (m *CheckCertificateRequest) Reset()                    { *m = CheckCertificateRequest{} }
func (m *CheckCertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateRequest) ProtoMessage()               {}
func (*CheckCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// CheckCertificateResponse is returned by CheckCertificate.
type CheckCertificateResponse struct {
	IsValid       bool   `protobuf:"varint,1,opt,name=is_valid,json=isValid" json:"is_valid,omitempty"`
	InvalidReason string `protobuf:"bytes,2,opt,name=invalid_reason,json=invalidReason" json:"invalid_reason,omitempty"`
}

func (m *CheckCertificateResponse) Reset()                    { *m = CheckCertificateResponse{} }
func (m *CheckCertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateResponse) ProtoMessage()               {}
func (*CheckCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// CRLStatus describes the latest known state of imported CRL.
type CRLStatus struct {
	LastUpdateTime    *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	LastFetchTime     *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=last_fetch_time,json=lastFetchTime" json:"last_fetch_time,omitempty"`
	LastFetchEtag     string                      `protobuf:"bytes,3,opt,name=last_fetch_etag,json=lastFetchEtag" json:"last_fetch_etag,omitempty"`
	RevokedCertsCount int64                       `protobuf:"varint,4,opt,name=revoked_certs_count,json=revokedCertsCount" json:"revoked_certs_count,omitempty"`
}

func (m *CRLStatus) Reset()                    { *m = CRLStatus{} }
func (m *CRLStatus) String() string            { return proto.CompactTextString(m) }
func (*CRLStatus) ProtoMessage()               {}
func (*CRLStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CRLStatus) GetLastUpdateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *CRLStatus) GetLastFetchTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastFetchTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportConfigResponse)(nil), "tokenserver.ImportConfigResponse")
	proto.RegisterType((*FetchCRLRequest)(nil), "tokenserver.FetchCRLRequest")
	proto.RegisterType((*FetchCRLResponse)(nil), "tokenserver.FetchCRLResponse")
	proto.RegisterType((*ListCAsResponse)(nil), "tokenserver.ListCAsResponse")
	proto.RegisterType((*GetCAStatusRequest)(nil), "tokenserver.GetCAStatusRequest")
	proto.RegisterType((*GetCAStatusResponse)(nil), "tokenserver.GetCAStatusResponse")
	proto.RegisterType((*IsRevokedCertRequest)(nil), "tokenserver.IsRevokedCertRequest")
	proto.RegisterType((*IsRevokedCertResponse)(nil), "tokenserver.IsRevokedCertResponse")
	proto.RegisterType((*CheckCertificateRequest)(nil), "tokenserver.CheckCertificateRequest")
	proto.RegisterType((*CheckCertificateResponse)(nil), "tokenserver.CheckCertificateResponse")
	proto.RegisterType((*CRLStatus)(nil), "tokenserver.CRLStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for CertificateAuthorities service

type CertificateAuthoritiesClient interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error)
}
type certificateAuthoritiesPRPCClient struct {
	client *prpccommon.Client
}

func NewCertificateAuthoritiesPRPCClient(client *prpccommon.Client) CertificateAuthoritiesClient {
	return &certificateAuthoritiesPRPCClient{client}
}

func (c *certificateAuthoritiesPRPCClient) ImportConfig(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "ImportConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "FetchCRL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "ListCAs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "GetCAStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "IsRevokedCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "CheckCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type certificateAuthoritiesClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthoritiesClient(cc *grpc.ClientConn) CertificateAuthoritiesClient {
	return &certificateAuthoritiesClient{cc}
}

func (c *certificateAuthoritiesClient) ImportConfig(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/ImportConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/FetchCRL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/ListCAs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/GetCAStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/IsRevokedCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/CheckCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthorities service

type CertificateAuthoritiesServer interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(context.Context, *google_protobuf.Empty) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(context.Context, *FetchCRLRequest) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(context.Context, *google_protobuf.Empty) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(context.Context, *GetCAStatusRequest) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(context.Context, *IsRevokedCertRequest) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(context.Context, *CheckCertificateRequest) (*CheckCertificateResponse, error)
}

func RegisterCertificateAuthoritiesServer(s prpc.Registrar, srv CertificateAuthoritiesServer) {
	s.RegisterService(&_CertificateAuthorities_serviceDesc, srv)
}

func _CertificateAuthorities_ImportConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).ImportConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_FetchCRL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FetchCRLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).FetchCRL(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_ListCAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).ListCAs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_GetCAStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetCAStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).GetCAStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_IsRevokedCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IsRevokedCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).IsRevokedCert(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_CheckCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CheckCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).CheckCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CertificateAuthorities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.CertificateAuthorities",
	HandlerType: (*CertificateAuthoritiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportConfig",
			Handler:    _CertificateAuthorities_ImportConfig_Handler,
		},
		{
			MethodName: "FetchCRL",
			Handler:    _CertificateAuthorities_FetchCRL_Handler,
		},
		{
			MethodName: "ListCAs",
			Handler:    _CertificateAuthorities_ListCAs_Handler,
		},
		{
			MethodName: "GetCAStatus",
			Handler:    _CertificateAuthorities_GetCAStatus_Handler,
		},
		{
			MethodName: "IsRevokedCert",
			Handler:    _CertificateAuthorities_IsRevokedCert_Handler,
		},
		{
			MethodName: "CheckCertificate",
			Handler:    _CertificateAuthorities_CheckCertificate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0x16, 0x09, 0x90, 0x64, 0xc2, 0xd7, 0xbb, 0xf0, 0x52, 0xd7, 0x14, 0x01, 0x16, 0xb4, 0x3d,
	0x19, 0x95, 0x7e, 0x1d, 0x2b, 0x9a, 0x52, 0x84, 0xc4, 0x01, 0xb9, 0x2d, 0xa7, 0x4a, 0x96, 0xb1,
	0x37, 0x61, 0x45, 0x9c, 0x75, 0x77, 0x37, 0x91, 0xf8, 0x57, 0xfd, 0x63, 0xbd, 0xf7, 0xd8, 0xdd,
	0xd9, 0x75, 0x88, 0x43, 0x28, 0xbd, 0x79, 0x66, 0x9e, 0x79, 0x66, 0xfc, 0xec, 0xcc, 0xc0, 0x76,
	0x4a, 0x85, 0x62, 0x5d, 0x96, 0x26, 0x8a, 0xc6, 0xc9, 0x50, 0x5d, 0x73, 0xc1, 0x14, 0xa3, 0x32,
	0x2c, 0x04, 0x57, 0x9c, 0xb4, 0x15, 0xbf, 0xa1, 0x03, 0x49, 0xc5, 0x88, 0x0a, 0x7f, 0x29, 0xe5,
	0x83, 0x2e, 0xeb, 0xd9, 0x90, 0xbf, 0xd5, 0xe3, 0xbc, 0xd7, 0xa7, 0x87, 0x68, 0x5d, 0x0d, 0xbb,
	0x87, 0x34, 0x2f, 0xd4, 0xad, 0x0b, 0xee, 0x4c, 0x07, 0x15, 0xcb, 0xa9, 0x54, 0x49, 0x5e, 0x58,
	0x40, 0x70, 0x04, 0x1b, 0x67, 0x79, 0xc1, 0x85, 0xea, 0x20, 0x67, 0x44, 0x65, 0xc1, 0x75, 0x1d,
	0xe2, 0x43, 0x53, 0xd0, 0x11, 0x93, 0x8c, 0x0f, 0xbc, 0xb9, 0xdd, 0xb9, 0x97, 0xad, 0x68, 0x6c,
	0x07, 0xef, 0x61, 0xf5, 0x33, 0x55, 0xe9, 0x75, 0x27, 0x3a, 0x8f, 0xe8, 0x8f, 0xa1, 0xe6, 0x23,
	0x2b, 0x50, 0x4b, 0x4b, 0xa0, 0xfe, 0x22, 0x1b, 0xb0, 0xd0, 0xe5, 0x22, 0xa5, 0x5e, 0x4d, 0xbb,
	0x9a, 0x91, 0x35, 0x82, 0x33, 0x58, 0xbb, 0x4b, 0x74, 0x85, 0xde, 0x02, 0xa4, 0xa2, 0x1f, 0xeb,
	0x9e, 0xd4, 0x50, 0x22, 0x43, 0xfb, 0x68, 0x33, 0x9c, 0xf8, 0xdd, 0x50, 0xa3, 0xbf, 0x60, 0x34,
	0x6a, 0x69, 0xa4, 0xfd, 0x0c, 0xf6, 0x60, 0xf5, 0x9c, 0x49, 0xd5, 0x39, 0x96, 0x63, 0xa6, 0xb2,
	0x87, 0xba, 0xed, 0x21, 0xd8, 0x07, 0x72, 0x4a, 0x35, 0xc2, 0x25, 0xcf, 0xee, 0x34, 0xf8, 0x59,
	0x83, 0xf5, 0x0a, 0xcc, 0xb1, 0x7d, 0x80, 0x45, 0x2b, 0xb3, 0xeb, 0xe9, 0x45, 0xb5, 0xa7, 0xbb,
	0xd7, 0x3a, 0x76, 0x8f, 0x75, 0xeb, 0x14, 0x74, 0x69, 0x84, 0xc0, 0xbc, 0x79, 0x53, 0x54, 0xa0,
	0x15, 0xe1, 0x37, 0xf1, 0xa0, 0x21, 0x68, 0xce, 0x47, 0x34, 0xf3, 0xea, 0x28, 0x4c, 0x69, 0x1a,
	0xc1, 0x04, 0x4d, 0xb2, 0x5b, 0x6f, 0xde, 0x0a, 0x86, 0x06, 0xd9, 0x82, 0x56, 0x92, 0x65, 0x34,
	0x8b, 0xb5, 0xf6, 0xde, 0x82, 0x7d, 0x06, 0x74, 0x44, 0x74, 0x44, 0x76, 0xa0, 0x3d, 0x2c, 0x32,
	0xdd, 0x81, 0x0d, 0x2f, 0x62, 0x18, 0x9c, 0xcb, 0x01, 0x1c, 0x3d, 0x02, 0x1a, 0x16, 0xe0, 0x5c,
	0x06, 0x50, 0xd5, 0xbe, 0xf9, 0xaf, 0xda, 0xbf, 0xd3, 0x33, 0xa3, 0x85, 0x1a, 0x69, 0x5c, 0x66,
	0x84, 0x98, 0x94, 0x36, 0x19, 0x4b, 0x9b, 0x18, 0x5b, 0x0e, 0xdc, 0xff, 0xeb, 0xaf, 0xe0, 0x15,
	0xfc, 0x3f, 0x95, 0xe7, 0xb4, 0x46, 0x59, 0xd0, 0x8d, 0xd9, 0x28, 0x0b, 0x9a, 0xc1, 0x1b, 0x78,
	0xd2, 0xb9, 0xa6, 0xe9, 0xcd, 0x84, 0xde, 0x65, 0xb5, 0xa7, 0xd0, 0x34, 0x9a, 0xc6, 0x05, 0xcd,
	0x5d, 0xcd, 0x86, 0xb1, 0x2f, 0x68, 0x1e, 0x7c, 0x07, 0xef, 0x7e, 0x96, 0xab, 0xa5, 0xd3, 0x98,
	0x8c, 0x47, 0x49, 0x9f, 0x8d, 0x8b, 0x31, 0x79, 0x69, 0x4c, 0x72, 0x00, 0x2b, 0x6c, 0x80, 0x11,
	0xad, 0x57, 0x22, 0x79, 0xd9, 0xfb, 0xb2, 0xf3, 0x46, 0xe8, 0x0c, 0x7e, 0xcd, 0x41, 0x6b, 0xac,
	0x0b, 0xf9, 0x04, 0x6b, 0xfd, 0x44, 0xaa, 0xd8, 0xea, 0x1e, 0x9b, 0xfd, 0x72, 0x13, 0xe3, 0x87,
	0x76, 0xf9, 0xc2, 0x72, 0xf9, 0xc2, 0xaf, 0xe5, 0xf2, 0x45, 0x2b, 0x26, 0xe7, 0x1b, 0xa6, 0x18,
	0x27, 0xf9, 0x08, 0xab, 0xc8, 0xd2, 0x35, 0xeb, 0x61, 0x49, 0x6a, 0x8f, 0x92, 0x2c, 0x9b, 0x14,
	0x5c, 0x28, 0xe4, 0x78, 0x5e, 0xe1, 0xa0, 0x2a, 0xe9, 0xe1, 0x90, 0xb5, 0x26, 0x70, 0x27, 0xda,
	0x49, 0x42, 0x58, 0x77, 0xf2, 0xc6, 0x46, 0x30, 0x19, 0xa7, 0x7c, 0x38, 0x50, 0x38, 0x78, 0xf5,
	0xe8, 0x3f, 0x71, 0xf7, 0x3e, 0xb2, 0x63, 0x02, 0x47, 0xbf, 0xeb, 0xb0, 0x39, 0x63, 0xde, 0xf5,
	0x71, 0x22, 0x67, 0xb0, 0x34, 0x79, 0x3d, 0xc8, 0xe6, 0xbd, 0x6e, 0x4f, 0xcc, 0x31, 0xf2, 0xf7,
	0x2a, 0x43, 0x35, 0xf3, 0xe0, 0x9c, 0x42, 0xb3, 0xbc, 0x0d, 0xe4, 0x59, 0x05, 0x3e, 0x75, 0x6b,
	0xfc, 0xed, 0x07, 0xa2, 0xe3, 0xc5, 0x6d, 0xb8, 0xcb, 0xf0, 0x60, 0x3b, 0x55, 0xfe, 0xe9, 0x3b,
	0x72, 0x01, 0xed, 0x89, 0x83, 0x40, 0x76, 0x2a, 0xe0, 0xfb, 0x17, 0xc5, 0xdf, 0x7d, 0x18, 0xe0,
	0x18, 0x2f, 0x61, 0xb9, 0x32, 0xf8, 0x64, 0x4a, 0x8f, 0x19, 0xcb, 0xe4, 0x07, 0x7f, 0x83, 0x38,
	0xde, 0x18, 0xd6, 0xa6, 0xe7, 0x9c, 0xec, 0x57, 0xf7, 0x77, 0xf6, 0xf2, 0xf8, 0x07, 0x8f, 0xa0,
	0x6c, 0x81, 0xab, 0x45, 0x14, 0xee, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x21, 0x67,
	0x83, 0x9e, 0x06, 0x00, 0x00,
}
