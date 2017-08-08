// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/config/project.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/milo/api/config/project.proto
	go.chromium.org/luci/milo/api/config/settings.proto

It has these top-level messages:
	Project
	Console
	Builder
	Settings
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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Project is a project definition for Milo.
type Project struct {
	// ID is the identifier for the project, if different from its repository name.
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	// Consoles is a list of consoles to define under /console/
	Consoles []*Console `protobuf:"bytes,2,rep,name=Consoles" json:"Consoles,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Project) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Project) GetConsoles() []*Console {
	if m != nil {
		return m.Consoles
	}
	return nil
}

// Console is a waterfall definition consisting of one or more builders.
type Console struct {
	// ID is the reference to the console, and will be the address to make the
	// console reachable from /console/<Project>/<ID>.
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	// Name is the longform name of the waterfall, and will be used to be
	// displayed in the title.
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// RepoURL is the URL of the git repository to display as the rows of the console.
	RepoURL string `protobuf:"bytes,3,opt,name=RepoURL" json:"RepoURL,omitempty"`
	// Ref is the ref to pull commits from when displaying the console.
	// Eg. refs/heads/master
	Ref string `protobuf:"bytes,4,opt,name=Ref" json:"Ref,omitempty"`
	// ManifestName is the name of the manifest the waterfall looks at.
	// By convention, Manifest Names can be:
	// * UNPATCHED - For non patched builds, such as continuous builds
	// * PATCHED - For patched builds, such as those on try jobs
	ManifestName string `protobuf:"bytes,5,opt,name=ManifestName" json:"ManifestName,omitempty"`
	// Builders is a list of builder configurations to display as the columns of the console.
	Builders []*Builder `protobuf:"bytes,6,rep,name=Builders" json:"Builders,omitempty"`
}

func (m *Console) Reset()                    { *m = Console{} }
func (m *Console) String() string            { return proto.CompactTextString(m) }
func (*Console) ProtoMessage()               {}
func (*Console) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Console) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Console) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Console) GetRepoURL() string {
	if m != nil {
		return m.RepoURL
	}
	return ""
}

func (m *Console) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *Console) GetManifestName() string {
	if m != nil {
		return m.ManifestName
	}
	return ""
}

func (m *Console) GetBuilders() []*Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

// A builder is a reference to a Milo builder.
type Builder struct {
	// Name is the identifier to find the builder, which includes the module.
	// Buildbot builds would be like "buildbot/chromium.linux/Linux Tests"
	// Buildbucket builds would be like "buildbucket/luci.chromium.try/linux_chromium_rel_ng"
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// Category describes the hierarchy of the builder on the header of the
	// console as a "|" delimited list.  Neighboring builders with common ancestors
	// will be have their headers merged.
	Category string `protobuf:"bytes,2,opt,name=Category" json:"Category,omitempty"`
	// ShortName is the 1-3 character abbreviation of the builder.
	ShortName string `protobuf:"bytes,3,opt,name=ShortName" json:"ShortName,omitempty"`
}

func (m *Builder) Reset()                    { *m = Builder{} }
func (m *Builder) String() string            { return proto.CompactTextString(m) }
func (*Builder) ProtoMessage()               {}
func (*Builder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Builder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Builder) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Builder) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "config.Project")
	proto.RegisterType((*Console)(nil), "config.Console")
	proto.RegisterType((*Builder)(nil), "config.Builder")
}

func init() { proto.RegisterFile("go.chromium.org/luci/milo/api/config/project.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0x3b, 0xdb, 0xed, 0x28, 0x2a, 0xb9, 0x0a, 0xe2, 0xc5, 0xe8, 0xd5, 0x40, 0x48,
	0x60, 0xbe, 0x81, 0x1b, 0xc2, 0x40, 0x45, 0x22, 0xe2, 0x75, 0xad, 0xa7, 0x5d, 0xa4, 0xed, 0x29,
	0x69, 0x7b, 0xe1, 0x2b, 0xf9, 0x94, 0xb2, 0x24, 0x56, 0xcb, 0xee, 0xfe, 0xff, 0x3b, 0x27, 0x87,
	0x8f, 0xc0, 0xba, 0x24, 0x91, 0xef, 0x0d, 0xd5, 0x7a, 0xa8, 0x05, 0x99, 0x52, 0x56, 0x43, 0xae,
	0x65, 0xad, 0x2b, 0x92, 0x59, 0xab, 0x65, 0x4e, 0x4d, 0xa1, 0x4b, 0xd9, 0x1a, 0xfa, 0xc4, 0xbc,
	0x17, 0xad, 0xa1, 0x9e, 0x58, 0xec, 0x68, 0x7a, 0x0f, 0xc9, 0xb3, 0x1b, 0xb0, 0x73, 0x08, 0x77,
	0x5b, 0x1e, 0x2c, 0x83, 0xd5, 0x42, 0x85, 0xbb, 0x2d, 0xbb, 0x81, 0xf9, 0x86, 0x9a, 0x8e, 0x2a,
	0xec, 0x78, 0xb8, 0x8c, 0x56, 0xa7, 0xeb, 0x0b, 0xe1, 0x5e, 0x09, 0xcf, 0xd5, 0xb8, 0x90, 0x7e,
	0x07, 0x90, 0xf8, 0x72, 0x74, 0x88, 0xc1, 0xec, 0x29, 0xab, 0x91, 0x87, 0x96, 0xd8, 0xcc, 0x38,
	0x24, 0x0a, 0x5b, 0x7a, 0x55, 0x0f, 0x3c, 0xb2, 0xf8, 0xb7, 0xb2, 0x4b, 0x88, 0x14, 0x16, 0x7c,
	0x66, 0xe9, 0x21, 0xb2, 0x14, 0xce, 0x1e, 0xb3, 0x46, 0x17, 0xd8, 0xf5, 0xf6, 0xce, 0x89, 0x1d,
	0x4d, 0xd8, 0x41, 0xf6, 0x6e, 0xd0, 0xd5, 0x07, 0x9a, 0x8e, 0xc7, 0x53, 0x59, 0xcf, 0xd5, 0xb8,
	0x90, 0xbe, 0x41, 0xe2, 0xf3, 0xe8, 0x16, 0xfc, 0x73, 0xbb, 0x82, 0xf9, 0x26, 0xeb, 0xb1, 0x24,
	0xf3, 0xe5, 0x9d, 0xc7, 0xce, 0xae, 0x61, 0xf1, 0xb2, 0x27, 0xe3, 0x44, 0x9c, 0xf9, 0x1f, 0x78,
	0x8f, 0xed, 0xe7, 0xde, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xb4, 0xea, 0xfc, 0x92, 0x01,
	0x00, 0x00,
}
