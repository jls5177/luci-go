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

package frontend

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/templates"

	"go.chromium.org/luci/milo/api/resp"
	"go.chromium.org/luci/milo/buildsource"
	"go.chromium.org/luci/milo/common"
	"go.chromium.org/luci/milo/common/model"
	"go.chromium.org/luci/milo/git"
)

// getConsoleDef finds the console definition as defined by any project.
// If the user is not a reader of the project, this will return a 404.
// TODO(hinoka): If the user is not a reader of any of of the builders returned,
// that builder will be removed from list of results.
func getConsoleDef(c context.Context, project, name string) (*common.Console, error) {
	cs, err := common.GetConsole(c, project, name)
	if err != nil {
		return nil, err
	}
	// TODO(hinoka): Remove builders that the user does not have access to.
	return cs, nil
}

// shortname calculates a short name (3 char max long name) out of a full name
// by splitting on delimiters and taking the first letter of each "word".
// name is expected to be a builderID, which is module/<bucket or master>/buildername
func shortname(name string) string {
	builderNameComp := strings.SplitN(name, "/", 3)
	if len(builderNameComp) == 3 {
		name = builderNameComp[2]
	}
	tokens := strings.FieldsFunc(name, func(r rune) bool {
		switch r {
		case '_', '-', ' ':
			return true
		}
		return false
	})
	numLetters := len(tokens)
	if numLetters > 3 {
		numLetters = 3
	}
	short := ""
	for i := 0; i < numLetters; i++ {
		short += string(tokens[i][0])
	}
	return strings.ToLower(short)
}

func console(c context.Context, project, name string, limit int) (*resp.Console, error) {
	tStart := clock.Now(c)
	def, err := getConsoleDef(c, project, name)
	if err != nil {
		return nil, err
	}
	commitInfo, err := git.GetHistory(c, def.RepoURL, def.Ref, limit)
	if err != nil {
		return nil, err
	}
	tGitiles := clock.Now(c)
	logging.Debugf(c, "Loading commits took %s.", tGitiles.Sub(tStart))

	commitNames := make([]string, len(commitInfo.Commits))
	for i, commit := range commitInfo.Commits {
		commitNames[i] = hex.EncodeToString(commit.Hash)
	}

	rows, err := buildsource.GetConsoleRows(c, project, def, commitNames, def.Builders)
	tConsole := clock.Now(c)
	logging.Debugf(c, "Loading the console took a total of %s.", tConsole.Sub(tGitiles))
	if err != nil {
		return nil, err
	}

	// Build list of commits.
	commits := make([]resp.Commit, len(commitInfo.Commits))
	for row, commit := range commitInfo.Commits {
		commits[row] = resp.Commit{
			AuthorName:  commit.AuthorName,
			AuthorEmail: commit.AuthorEmail,
			CommitTime:  google.TimeFromProto(commit.CommitTime),
			Repo:        def.RepoURL,
			Branch:      def.Ref, // TODO(hinoka): Actually this doesn't match, change branch to ref.
			Description: commit.Msg,
			Revision:    resp.NewLink(commitNames[row], def.RepoURL+"/+/"+commitNames[row]),
		}
	}

	// Build console table tree from builders.
	categoryTree := resp.NewCategory("")
	depth := 0
	for col, b := range def.Builders {
		meta := def.BuilderMetas[col]
		short := meta.ShortName
		if short == "" {
			short = shortname(b)
		}
		name := buildsource.BuilderID(b)

		// Group together all builds for this builder.
		builds := make([]*model.BuildSummary, len(commits))
		for row := 0; row < len(commits); row++ {
			if summaries := rows[row].Builds[name]; len(summaries) > 0 {
				builds[row] = summaries[0]
			}
		}
		builderRef := &resp.BuilderRef{
			Name:      b,
			ShortName: short,
			Build:     builds,
		}
		categories := strings.Split(meta.Category, "|")
		if len(categories) > depth {
			depth = len(categories)
		}
		categoryTree.AddBuilder(categories, builderRef)
	}

	return &resp.Console{
		Name:     def.ID,
		Commit:   commits,
		Table:    *categoryTree,
		MaxDepth: depth + 1,
	}, nil
}

// consoleRenderer is a wrapper around Console to provide additional methods.
type consoleRenderer struct {
	*resp.Console
}

// ConsoleTable generates the main console table html.
//
// This cannot be generated with templates due to the 'recursive' nature of
// this layout.
func (c consoleRenderer) ConsoleTable() template.HTML {
	var buffer bytes.Buffer
	// The first node is a dummy node
	for _, column := range c.Table.Children {
		column.RenderHTML(&buffer, 1, c.MaxDepth)
	}
	return template.HTML(buffer.String())
}

func (c consoleRenderer) BuilderLink(bs *model.BuildSummary) (*resp.Link, error) {
	_, _, builderName, err := buildsource.BuilderID(bs.BuilderID).Split()
	if err != nil {
		return nil, err
	}
	return resp.NewLink(builderName, "/"+bs.BuilderID), nil
}

// ConsoleHandler renders the console page.
func ConsoleHandler(c *router.Context) {
	project := c.Params.ByName("project")
	if project == "" {
		ErrorHandler(c, errors.New("Missing Project", common.CodeParameterError))
		return
	}
	name := c.Params.ByName("name")
	const defaultLimit = 25
	const maxLimit = 1000
	limit := defaultLimit
	if tLimit := GetLimit(c.Request, -1); tLimit >= 0 {
		limit = tLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	result, err := console(c.Context, project, name, limit)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	templates.MustRender(c.Context, c.Writer, "pages/console.html", templates.Args{
		"Console": consoleRenderer{result},
	})
}

// ConsoleMainHandler is a redirect handler that redirects the user to the main
// console for a particular project.
func ConsoleMainHandler(ctx *router.Context) {
	w, r, p := ctx.Writer, ctx.Request, ctx.Params
	proj := p.ByName("project")
	http.Redirect(w, r, fmt.Sprintf("/console/%s/main", proj), http.StatusMovedPermanently)
	return
}
