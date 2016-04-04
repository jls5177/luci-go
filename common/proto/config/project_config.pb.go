// Code generated by protoc-gen-go.
// source: project_config.proto
// DO NOT EDIT!

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	project_config.proto
	service_config.proto

It has these top-level messages:
	ProjectCfg
	RefsCfg
	ConfigSetLocation
	Project
	ProjectsCfg
	Service
	ServiceDynamicMetadata
	ServicesCfg
	AclCfg
	ImportCfg
	SchemasCfg
	ConfigPattern
	Validator
	ValidationRequestMessage
	ValidationResponseMessage
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Schema for project.cfg
type ProjectCfg struct {
	// Full name of the project.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A list of identities that have read-only access to the. An element is one
	// of:
	// * "group:<group>", where group is defined on auth server.
	// * "<email>"
	// * "<identity>"
	//
	// If not specified, only admins and trusted services have access.
	// Talk to admins to determine the group name appropriate for your project.
	Access           []string `protobuf:"bytes,2,rep,name=access" json:"access,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProjectCfg) Reset()                    { *m = ProjectCfg{} }
func (m *ProjectCfg) String() string            { return proto.CompactTextString(m) }
func (*ProjectCfg) ProtoMessage()               {}
func (*ProjectCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProjectCfg) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProjectCfg) GetAccess() []string {
	if m != nil {
		return m.Access
	}
	return nil
}

// Schema of refs.cfg.
type RefsCfg struct {
	// List of refs that have configuration files and need to be fetched into
	// luci-config. Refs are accessible through get_refs() API endpoint.
	// A CI service can read all refs of all projects and build them.
	Refs             []*RefsCfg_Ref `protobuf:"bytes,1,rep,name=refs" json:"refs,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RefsCfg) Reset()                    { *m = RefsCfg{} }
func (m *RefsCfg) String() string            { return proto.CompactTextString(m) }
func (*RefsCfg) ProtoMessage()               {}
func (*RefsCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RefsCfg) GetRefs() []*RefsCfg_Ref {
	if m != nil {
		return m.Refs
	}
	return nil
}

type RefsCfg_Ref struct {
	// Name of the ref. Must start with "refs/".
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Path to config directory for the ref. Defaults to "infra/config".
	ConfigPath       *string `protobuf:"bytes,3,opt,name=config_path,json=configPath" json:"config_path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RefsCfg_Ref) Reset()                    { *m = RefsCfg_Ref{} }
func (m *RefsCfg_Ref) String() string            { return proto.CompactTextString(m) }
func (*RefsCfg_Ref) ProtoMessage()               {}
func (*RefsCfg_Ref) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *RefsCfg_Ref) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RefsCfg_Ref) GetConfigPath() string {
	if m != nil && m.ConfigPath != nil {
		return *m.ConfigPath
	}
	return ""
}

func init() {
	proto.RegisterType((*ProjectCfg)(nil), "config.ProjectCfg")
	proto.RegisterType((*RefsCfg)(nil), "config.RefsCfg")
	proto.RegisterType((*RefsCfg_Ref)(nil), "config.RefsCfg.Ref")
}

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0x89, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x03, 0x72, 0x4b, 0xf2, 0x85,
	0xd8, 0x20, 0x3c, 0x25, 0x0b, 0x2e, 0xae, 0x00, 0x88, 0xbc, 0x73, 0x5a, 0xba, 0x90, 0x10, 0x17,
	0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc6,
	0xc5, 0x96, 0x98, 0x9c, 0x9c, 0x5a, 0x5c, 0x2c, 0xc1, 0xa4, 0xc0, 0x0c, 0x14, 0x85, 0xf2, 0x94,
	0xf2, 0xb8, 0xd8, 0x83, 0x52, 0xd3, 0x8a, 0x41, 0xda, 0xd4, 0xb9, 0x58, 0x8a, 0x80, 0x4c, 0xa0,
	0x36, 0x66, 0x0d, 0x6e, 0x23, 0x61, 0x3d, 0xa8, 0x4d, 0x50, 0x69, 0x10, 0x1d, 0x04, 0x56, 0x20,
	0x65, 0xc5, 0xc5, 0x0c, 0xe4, 0x60, 0xb5, 0x46, 0x9e, 0x8b, 0x1b, 0xa2, 0x2d, 0xbe, 0x20, 0xb1,
	0x24, 0x43, 0x82, 0x19, 0x2c, 0xc5, 0x05, 0x11, 0x0a, 0x00, 0x8a, 0x00, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0xef, 0x48, 0x4c, 0xc8, 0x00, 0x00, 0x00,
}
