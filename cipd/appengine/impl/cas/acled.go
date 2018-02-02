// Copyright 2017 The LUCI Authors.
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

package cas

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server/auth"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
)

// Public returns publicly exposed implementation of cipd.Storage service.
func Public() api.StorageServer {
	return &api.DecoratedStorage{
		Service: Internal(),

		Prelude: func(c context.Context, methodName string, req proto.Message) (context.Context, error) {
			acl, ok := perMethodACL[methodName]
			if !ok {
				panic(fmt.Sprintf("method %q is not defined in perMethodACL", methodName))
			}
			if acl.group != "*" {
				switch yep, err := auth.IsMember(c, acl.group); {
				case err != nil:
					logging.WithError(err).Errorf(c, "IsMember(%q) failed", acl.group)
					return nil, status.Errorf(codes.Internal, "failed to check ACL")
				case !yep:
					return nil, status.Errorf(codes.PermissionDenied, "not allowed")
				}
			}
			if acl.check != nil {
				if err := acl.check(c, req); err != nil {
					return nil, err
				}
			}
			return c, nil
		},
	}
}

// perMethodACL defines a group to check when authorizing an RPC call plus a
// callback for more detailed check.
//
// Group "*" means "allow anyone to call the method".
var perMethodACL = map[string]struct {
	group string
	check func(c context.Context, req proto.Message) error
}{
// TODO
}