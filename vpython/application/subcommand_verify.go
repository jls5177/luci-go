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

package application

import (
	"strings"

	"github.com/maruel/subcommands"
	"golang.org/x/net/context"

	"go.chromium.org/luci/vpython/api/vpython"
	"go.chromium.org/luci/vpython/spec"

	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
)

var subcommandVerify = &subcommands.Command{
	UsageLine: "verify",
	ShortDesc: "verifies that a spec and its wheels are valid",
	LongDesc: "verifies that a spec file is valid, and that all of the wheels listed resolve for " +
		"all of the configured verification architectures.",
	Advanced: false,
	CommandRun: func() subcommands.CommandRun {
		return &verifyCommandRun{}
	},
}

type verifyCommandRun struct {
	subcommands.CommandRunBase
}

func (cr *verifyCommandRun) Run(app subcommands.Application, args []string, env subcommands.Env) int {
	c := cli.GetContext(app, cr, env)
	a := getApplication(c, args)

	return run(c, func(c context.Context) error {
		// Make sure that we can resolve the referenced specifiction.
		if err := a.opts.ResolveSpec(c); err != nil {
			return errors.Annotate(err, "failed to resolve specification").Err()
		}

		s := a.opts.EnvConfig.Spec
		if s == nil {
			s = &vpython.Spec{}
		}
		if s.Virtualenv == nil {
			s.Virtualenv = &a.opts.EnvConfig.Package
		}

		// Verify that the spec can be normalized. This may modify it, so we will
		// normalize a clone.
		if err := spec.NormalizeSpec(s.Clone(), nil); err != nil {
			return errors.Annotate(err, "failed to normalize specification").Err()
		}

		renderedSpec := spec.Render(s)
		logging.Infof(c, "Successfully verified specification:\n%s", renderedSpec)

		// Run our Verification generator and verify each generated environment.
		if a.WithVerificationConfig != nil {
			err := a.WithVerificationConfig(c, func(cfg Config, verificationScenarios []*vpython.PEP425Tag) error {
				if len(s.VerifyPep425Tag) > 0 {
					verificationScenarios = s.VerifyPep425Tag
				}
				if len(verificationScenarios) == 0 {
					return nil
				}

				var failures []string
				for _, vs := range verificationScenarios {
					// Create a verification environment to pass to our package loader.
					e := vpython.Environment{
						Spec:      s.Clone(),
						Pep425Tag: []*vpython.PEP425Tag{vs},
					}
					if err := spec.NormalizeEnvironment(&e); err != nil {
						logging.Errorf(c, "Failed to normalize environment against %q: %s", vs.TagString(), err)
						failures = append(failures, vs.TagString())
						continue
					}

					if err := cfg.PackageLoader.Resolve(c, &e); err != nil {
						logging.Errorf(c, "Failed to resolve against %q: %s", vs.TagString(), err)
						failures = append(failures, vs.TagString())
					}
				}

				if len(failures) > 0 {
					logging.Errorf(c, "Verification failed for %d scenario(s): %s\n%s",
						len(failures), strings.Join(failures, ", "), renderedSpec)
					return errors.New("verification failed")
				}

				logging.Infof(c, "Verified %d scenario(s).", len(verificationScenarios))
				return nil
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}
