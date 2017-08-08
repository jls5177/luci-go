// Copyright 2015 The LUCI Authors.
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

package gaetesting

import (
	"golang.org/x/net/context"

	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/luci/common/data/caching/proccache"
	"go.chromium.org/luci/server/secrets/testsecrets"
)

// TestingContext returns context with base services installed:
//   * go.chromium.org/gae/impl/memory (in-memory appengine services)
//   * go.chromium.org/luci/server/secrets/testsecrets (access to fake secret keys)
func TestingContext() context.Context {
	c := context.Background()
	c = memory.Use(c)
	c = testsecrets.Use(c)
	c = proccache.Use(c, &proccache.Cache{})
	return c
}
