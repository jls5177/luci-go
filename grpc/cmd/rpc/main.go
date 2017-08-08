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

package main

import (
	"fmt"
	"os"

	"github.com/maruel/subcommands"
	"golang.org/x/net/context"

	"go.chromium.org/luci/client/authcli"
	"go.chromium.org/luci/common/auth"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/data/rand/mathrand"
	"go.chromium.org/luci/common/lhttp"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/grpc/prpc"

	"go.chromium.org/luci/hardcoded/chromeinfra"
)

const (
	userAgent = "luci-rpc"
)

var logCfg = gologger.LoggerConfig{
	Format: `%{message}`,
	Out:    os.Stderr,
}

// exit codes:
const (
	ecInvalidCommandLine = -iota
	ecAuthenticatedClientError
	ecOtherError
)

type exitCode struct {
	err  error
	code int
}

func (e *exitCode) Error() string { return e.err.Error() }

// cmdRun is a base of all rpc subcommands.
// It defines some common flags, such as logging and auth, and useful methods.
type cmdRun struct {
	subcommands.CommandRunBase
	verbose bool
	auth    authcli.Flags
}

// ModifyContext implements cli.ContextModificator.
func (r *cmdRun) ModifyContext(ctx context.Context) context.Context {
	if r.verbose {
		ctx = logging.SetLevel(ctx, logging.Debug)
	}
	return ctx
}

// registerBaseFlags registers common flags used by all subcommands.
func (r *cmdRun) registerBaseFlags(defaultAuthOpts auth.Options) {
	r.Flags.BoolVar(&r.verbose, "verbose", false, "Enable more logging.")
	r.auth.Register(&r.Flags, defaultAuthOpts)
}

func (r *cmdRun) authenticatedClient(ctx context.Context, host string) (*prpc.Client, error) {
	authOpts, err := r.auth.Options()
	if err != nil {
		return nil, err
	}
	a := auth.NewAuthenticator(ctx, auth.OptionalLogin, authOpts)
	httpClient, err := a.Client()
	if err != nil {
		return nil, err
	}

	client := prpc.Client{
		C:       httpClient,
		Host:    host,
		Options: prpc.DefaultOptions(),
	}
	client.Options.Insecure = lhttp.IsLocalHost(host)
	return &client, nil
}

// argErr prints an err and usage to stderr and returns an exit code.
func (r *cmdRun) argErr(shortDesc, usageLine, format string, a ...interface{}) int {
	if format != "" {
		fmt.Fprintf(os.Stderr, format+"\n", a...)
	}
	fmt.Fprintln(os.Stderr, shortDesc)
	fmt.Fprintln(os.Stderr, usageLine)
	fmt.Fprintln(os.Stderr, "\nFlags:")
	r.Flags.PrintDefaults()
	return ecInvalidCommandLine
}

// done prints err to stderr if it is not nil and returns an exit code.
func (r *cmdRun) done(err error) int {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err, ok := err.(*exitCode); ok {
			return err.code
		}
		return ecOtherError
	}
	return 0
}

func GetApplication(defaultAuthOpts auth.Options) *cli.Application {
	return &cli.Application{
		Name:  "rpc",
		Title: "Remote Procedure Call CLI",
		Context: func(ctx context.Context) context.Context {
			return logCfg.Use(ctx)
		},
		Commands: []*subcommands.Command{
			cmdCall(defaultAuthOpts),
			cmdShow(defaultAuthOpts),
			cmdFmt(defaultAuthOpts),
			authcli.SubcommandLogin(defaultAuthOpts, "login", false),
			authcli.SubcommandLogout(defaultAuthOpts, "logout", false),
			subcommands.CmdHelp,
		},
	}
}

func main() {
	mathrand.SeedRandomly()
	app := GetApplication(chromeinfra.DefaultAuthOptions())
	os.Exit(subcommands.Run(app, os.Args[1:]))
}
