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

package resp

import (
	"go.chromium.org/luci/milo/common/model"
)

// This file contains the structures for defining a Console view.
// Console: The main entry point and the overall struct for a console page.
// BuilderRef: Used both as an input to request a builder and headers for the console.
// CommitBuild: A row in the console.  References a commit with a list of build summaries.
// ConsoleBuild: A cell in the console. Contains all information required to render the cell.

// Console represents a console view.  Commit contains the full matrix of
// Commits x Builder, and BuilderRef contains information on how to render
// the header.  The two structs are expected to be consistent.  IE len(Console.[]BuilderRef)
// Should equal len(commit.Build) for all commit in Console.Commit.
type Console struct {
	Name string

	Commit []CommitBuild

	BuilderRef []BuilderRef
}

// BuilderRef is an unambiguous reference to a builder, along with metadata on how
// to lay it out for rendering.
type BuilderRef struct {
	// Name is the canonical reference to a specific builder.
	Name string
	// Category is a pipe "|" deliminated list of short strings used to catagorize
	// and organize builders.  Adjacent builders with common categories will be
	// merged on the header.
	Category []string
	// ShortName is a string of length 1-3 used to label the builder.
	ShortName string
}

// CommitBuild is a row in the console.  References a commit with a list of build summaries.
type CommitBuild struct {
	Commit
	Build []*model.BuildSummary
}
