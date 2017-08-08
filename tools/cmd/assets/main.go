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

// Package main hosts the utility that converts binary assets into assets.gen.go
// file, so that they can be baked directly into the executable. Intended to
// be used only for small files, like HTML templates.
//
// This utility is used via `go generate`. Corresponding incantation:
//   //go:generate go install go.chromium.org/luci/tools/cmd/assets
//   //go:generate assets
package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

// recognizedAssets lists glob patterns for files to put into generated
// *.go file.
var recognizedAssets = []string{
	"*.css",
	"*.html",
	"*.js",
}

// funcMap contains functions used when rendering assets.gen.go template.
var funcMap = template.FuncMap{
	"asByteArray": asByteArray,
}

// assetsGenGoTmpl is template for generated assets.gen.go file. Result of
// the execution will also be passed through gofmt.
var assetsGenGoTmpl = template.Must(template.New("tmpl").Funcs(funcMap).Parse(strings.TrimSpace(`
// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// Package {{.PackageName}} is generated by go.chromium.org/luci/tools/cmd/assets.
//
// It contains all {{.Patterns}} files found in the package as byte arrays.
package {{.PackageName}}

// GetAsset returns an asset by its name. Returns nil if no such asset.
func GetAsset(name string) []byte {
	return []byte(files[name])
}

// GetAssetString is version of GetAsset that returns string instead of byte
// slice. Returns empty string if no such asset.
func GetAssetString(name string) string {
	return files[name]
}

// Assets returns a map of all assets.
func Assets() map[string]string {
	cpy := make(map[string]string, len(files))
	for k, v := range files {
		cpy[k] = v
	}
	return cpy
}

var files = map[string]string{
{{range .Assets}}{{.Path | printf "%q"}}: string({{.Body | asByteArray }}),
{{end}}
}
`)))

// assetsTestTmpl is template to assets_test.go file.
var assetsTestTmpl = template.Must(template.New("tmpl").Funcs(funcMap).Parse(strings.TrimSpace(`
// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// This file is generated by go.chromium.org/luci/tools/cmd/assets.
//
// It contains tests that ensure that assets embedded into the binary are
// identical to files on disk.

package {{.PackageName}}

import (
	"go/build"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestAssets(t *testing.T) {
	t.Parallel()

	pkg, err := build.ImportDir(".", build.FindOnly)
	if err != nil {
		t.Fatalf("can't load package: %s", err)
	}

	fail := false
	for name := range Assets() {
		GetAsset(name) // for code coverage
		path := filepath.Join(pkg.Dir, filepath.FromSlash(name))
		blob, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("can't read file with assets %q (%s) - %s", name, path, err)
			fail = true
		} else if string(blob) != GetAssetString(name) {
			t.Errorf("embedded asset %q is out of date", name)
			fail = true
		}
	}

	if fail {
		t.Fatalf("run 'go generate' to update assets.gen.go")
	}
}
`)))

// templateData is passed to tmpl when rendering it.
type templateData struct {
	Patterns    []string
	PackageName string
	Assets      []asset
}

// asset is single file to be embedded into assets.gen.go.
type asset struct {
	Path string // path relative to package directory
	Body []byte // body of the file
}

type assetMap map[string]asset

func main() {
	var dir string
	switch len(os.Args) {
	case 1:
		dir = "."
	case 2:
		dir = os.Args[1]
	default:
		fmt.Fprintf(os.Stderr, "usage: assets [dir]\n")
		os.Exit(2)
	}

	if err := run(dir); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run generates assets.gen.go file with all assets discovered in the directory.
func run(dir string) error {
	pkg, err := build.ImportDir(dir, build.ImportComment)
	if err != nil {
		return fmt.Errorf("can't import dir %q - %s", dir, err)
	}

	assets, err := findAssets(pkg.Dir)
	if err != nil {
		return fmt.Errorf("can't find assets in %s - %s", pkg.Dir, err)
	}

	err = generate(assetsGenGoTmpl, pkg, assets, filepath.Join(pkg.Dir, "assets.gen.go"))
	if err != nil {
		return fmt.Errorf("can't generate assets.gen.go - %s", err)
	}

	err = generate(assetsTestTmpl, pkg, assets, filepath.Join(pkg.Dir, "assets_test.go"))
	if err != nil {
		return fmt.Errorf("can't generate assets_test.go - %s", err)
	}

	return nil
}

// findAssets recursively scans pkgDir for asset files.
func findAssets(pkgDir string) (assetMap, error) {
	assets := assetMap{}

	err := filepath.Walk(pkgDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !isAssetFile(path) {
			return err
		}
		rel, err := filepath.Rel(pkgDir, path)
		if err != nil {
			return err
		}
		blob, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		assets[filepath.ToSlash(rel)] = asset{
			Path: filepath.ToSlash(rel),
			Body: blob,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return assets, nil
}

// isAssetFile returns true if `path` base name matches some of
// `recognizedAssets` glob.
func isAssetFile(path string) bool {
	base := filepath.Base(path)
	for _, pattern := range recognizedAssets {
		if match, _ := filepath.Match(pattern, base); match {
			return true
		}
	}
	return false
}

// generate executes the template, runs output through gofmt and dumps it to disk.
func generate(t *template.Template, pkg *build.Package, assets assetMap, path string) error {
	keys := []string{}
	for k := range assets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	data := templateData{
		Patterns:    recognizedAssets,
		PackageName: pkg.Name,
	}
	for _, key := range keys {
		data.Assets = append(data.Assets, assets[key])
	}

	out := bytes.Buffer{}
	if err := t.Execute(&out, data); err != nil {
		return err
	}

	formatted, err := gofmt(out.Bytes())
	if err != nil {
		return fmt.Errorf("can't gofmt %s - %s", path, err)
	}

	return ioutil.WriteFile(path, formatted, 0666)
}

// gofmt applies "gofmt -s" to the content of the buffer.
func gofmt(blob []byte) ([]byte, error) {
	out := bytes.Buffer{}
	cmd := exec.Command("gofmt", "-s")
	cmd.Stdin = bytes.NewReader(blob)
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func asByteArray(blob []byte) string {
	out := &bytes.Buffer{}
	fmt.Fprintf(out, "[]byte{")
	for i := 0; i < len(blob); i++ {
		fmt.Fprintf(out, "%d, ", blob[i])
		if i%14 == 1 {
			fmt.Fprintln(out)
		}
	}
	fmt.Fprintf(out, "}")
	return out.String()
}
