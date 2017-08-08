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

package auth

import (
	"go.chromium.org/luci/server/router"
)

// Authenticate returns a middleware that performs authentication.
//
// This is simplest form of this middleware that uses only one authentication
// method. It is sufficient in most cases.
//
// This middleware either updates the context by injecting the authentication
// state into it (enabling functions like CurrentIdentity and IsMember), or
// aborts the request with an HTTP 401 or HTTP 500 error.
//
// Note that it passes through anonymous requests. CurrentIdentity returns
// identity.AnonymousIdentity in this case. Use separate authorization layer to
// further restrict the access, if necessary.
func Authenticate(m Method) router.Middleware {
	a := &Authenticator{Methods: []Method{m}}
	return a.GetMiddleware()
}
