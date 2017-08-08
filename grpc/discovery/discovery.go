// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package discovery implements RPC service introspection.
package discovery

import (
	"crypto/sha1"
	"fmt"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"go.chromium.org/luci/grpc/prpc"
)

// New creates a discovery server for all the given services.
// The service descriptions must be registered using
// RegisterDescriptorSetCompressed which is called by init() function
// generated by go.chromium.org/luci/grpc/cmd/cproto.
func New(serviceNames ...string) (DiscoveryServer, error) {
	desc, err := combineDescriptors(serviceNames)
	if err != nil {
		return nil, err
	}
	desc = proto.Clone(desc).(*descriptor.FileDescriptorSet)

	cpy := make([]string, len(serviceNames))
	copy(cpy, serviceNames)
	return &service{cpy, desc}, nil
}

// Enable registers a discovery service on the server.
// It makes all currently registered services and the discovery service
// discoverable.
func Enable(server *prpc.Server) {
	serviceNames := append(server.ServiceNames(), "discovery.Discovery")
	service, err := New(serviceNames...)
	if err != nil {
		panic(err)
	}
	RegisterDiscoveryServer(server, service)
}

type service struct {
	serviceNames []string
	description  *descriptor.FileDescriptorSet
}

func (s *service) Describe(c context.Context, _ *Void) (*DescribeResponse, error) {
	return &DescribeResponse{
		Description: s.description,
		Services:    s.serviceNames,
	}, nil
}

// combineDescriptors creates one FileDescriptorSet that covers all services
// and their dependencies.
func combineDescriptors(serviceNames []string) (*descriptor.FileDescriptorSet, error) {
	result := &descriptor.FileDescriptorSet{}
	// seenFiles is a set of descriptor files keyed by SHA1 of their contents..
	seenFiles := map[[sha1.Size]byte]bool{}

	for _, s := range serviceNames {
		desc, err := GetDescriptorSet(s)
		if err != nil {
			return nil, fmt.Errorf("service %s: %s", s, err)
		}
		if desc == nil {
			return nil, fmt.Errorf(
				"descriptor for service %q is not found. "+
					"Did you compile it with go.chromium.org/luci/grpc/cmd/cproto?",
				s)
		}
		for _, f := range desc.GetFile() {
			binary, err := proto.Marshal(f)
			if err != nil {
				return nil, fmt.Errorf("could not marshal description of %s", f.GetName())
			}
			hash := sha1.Sum(binary)
			if !seenFiles[hash] {
				result.File = append(result.File, f)
				seenFiles[hash] = true
			}

		}
	}
	return result, nil
}
