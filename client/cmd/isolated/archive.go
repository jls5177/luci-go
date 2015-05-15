// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/luci/luci-go/client/archiver"
	"github.com/luci/luci-go/client/internal/common"
	"github.com/luci/luci-go/client/isolatedclient"
	"github.com/maruel/subcommands"
)

var cmdArchive = &subcommands.Command{
	UsageLine: "archive <options>...",
	ShortDesc: "creates a .isolated file and uploads the tree to an isolate server.",
	LongDesc:  "All the files listed in the .isolated file are put in the isolate server.",
	CommandRun: func() subcommands.CommandRun {
		c := archiveRun{}
		c.commonFlags.Init(&c.CommandRunBase)
		c.commonServerFlags.Init(&c.CommandRunBase)
		c.Flags.Var(&c.dirs, "dirs", "Directory(ies) to archive")
		c.Flags.Var(&c.files, "files", "Individual file(s) to archive")
		c.Flags.Var(&c.blacklist, "blacklist",
			"List of regexp to use as blacklist filter when uploading directories")
		return &c
	},
}

type archiveRun struct {
	subcommands.CommandRunBase
	commonFlags
	commonServerFlags
	dirs      common.Strings
	files     common.Strings
	blacklist common.Strings
}

func (c *archiveRun) Parse(a subcommands.Application, args []string) error {
	if err := c.commonFlags.Parse(); err != nil {
		return err
	}
	if err := c.commonServerFlags.Parse(); err != nil {
		return err
	}
	if len(args) != 0 {
		return errors.New("position arguments not expected")
	}
	return nil
}

func (c *archiveRun) main(a subcommands.Application, args []string) error {
	start := time.Now()
	out := os.Stdout
	prefix := "\n"
	if c.quiet {
		out = nil
		prefix = ""
	}
	arch := archiver.New(isolatedclient.New(c.serverURL, c.namespace), out)
	common.CancelOnCtrlC(arch)
	futures := []archiver.Future{}
	names := []string{}
	for _, file := range c.files {
		futures = append(futures, arch.PushFile(file, file))
		names = append(names, file)
	}

	for _, d := range c.dirs {
		futures = append(futures, archiver.PushDirectory(arch, d, "", nil))
		names = append(names, d)
	}

	for i, future := range futures {
		future.WaitForHashed()
		if err := future.Error(); err == nil {
			fmt.Printf("%s%s  %s\n", prefix, future.Digest(), names[i])
		} else {
			fmt.Printf("%s%s failed: %s\n", prefix, names[i], err)
		}
	}
	// This waits for all uploads.
	err := arch.Close()
	if !c.quiet {
		duration := time.Since(start)
		stats := arch.Stats()
		fmt.Fprintf(os.Stderr, "Hits    : %5d (%s)\n", stats.TotalHits(), stats.TotalBytesHits())
		fmt.Fprintf(os.Stderr, "Misses  : %5d (%s)\n", stats.TotalMisses(), stats.TotalBytesPushed())
		fmt.Fprintf(os.Stderr, "Duration: %s\n", common.Round(duration, time.Millisecond))
	}
	return err
}

func (c *archiveRun) Run(a subcommands.Application, args []string) int {
	defer c.Close()
	if err := c.Parse(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	if err := c.main(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
