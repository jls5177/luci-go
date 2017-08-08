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

package module

import (
	"net/http"

	// Importing pprof implicitly installs "/debug/*" profiling handlers.
	_ "net/http/pprof"

	"go.chromium.org/luci/appengine/gaemiddleware"
	"go.chromium.org/luci/logdog/appengine/coordinator"
	"go.chromium.org/luci/logdog/appengine/coordinator/tasks"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/tumble"

	// Include mutations package so its Mutations will register with tumble via
	// init().
	_ "go.chromium.org/luci/logdog/appengine/coordinator/mutations"
)

func init() {
	tmb := tumble.Service{}

	r := router.New()
	base := gaemiddleware.BaseProd().Extend(coordinator.ProdCoordinatorService)
	tmb.InstallHandlers(r, base)
	gaemiddleware.InstallHandlersWithMiddleware(r, base)
	tasks.InstallHandlers(r, base)

	http.Handle("/", r)
}
