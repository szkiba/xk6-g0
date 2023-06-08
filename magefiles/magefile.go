// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

//go:build mage
// +build mage

package main

import (
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
	"github.com/princjef/mageutil/bintool"
	"github.com/princjef/mageutil/shellcmd"
)

var Default = All

func init() {
	//os.Setenv("XK6_BUILD_FLAGS", "-tags go1.19 -ldflags='-w -s'")
}

var linter = bintool.Must(bintool.New(
	"golangci-lint{{.BinExt}}",
	"1.51.1",
	"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
))

func Lint() error {
	if err := linter.Ensure(); err != nil {
		return err
	}

	return linter.Command(`run`).Run()
}

func Test() error {
	return shellcmd.Command(`go test -count 1 -coverprofile=coverage.txt . ./internal/... ./g0/...`).Run()
}

func Build() error {
	return shellcmd.Command(`xk6 build --with github.com/szkiba/xk6-g0=.`).Run()
}

func Run() error {
	return shellcmd.Command(`xk6 run --quiet --no-usage-report scripts/simple/script.go`).Run()
}

func It() error {
	all, err := filepath.Glob("scripts/**/*.go")
	if err != nil {
		return err
	}

	for _, script := range all {
		if strings.HasSuffix(script, "_test.go") {
			continue
		}

		err := shellcmd.Command("./k6 run --quiet --no-usage-report " + script).Run()
		if err != nil {
			return err
		}
	}

	return shellcmd.Command("go test ./scripts/...").Run()
}

func Coverage() error {
	return shellcmd.Command(`go tool cover -html=coverage.txt`).Run()
}

func glob(patterns ...string) (string, error) {
	buff := new(strings.Builder)

	for _, p := range patterns {
		m, err := filepath.Glob(p)
		if err != nil {
			return "", err
		}

		_, err = buff.WriteString(strings.Join(m, " ") + " ")
		if err != nil {
			return "", err
		}
	}

	return buff.String(), nil
}

func License() error {
	all, err := glob("*.go", "*/*.go", "internal/*/*/*.go", "internal/*/*/*/*.go", "internal/*/*.go", ".*.yml", ".gitignore", "*/.gitignore", ".github/workflows/*")
	if err != nil {
		return err
	}

	return shellcmd.Command(
		`reuse annotate --copyright "Iván Szkiba" --merge-copyrights --license MIT --skip-unrecognised ` + all,
	).Run()
}

func Clean() error {
	sh.Rm("magefiles/dist")
	sh.Rm("magefiles/bin")
	sh.Rm("coverage.txt")
	sh.Rm("bin")

	return nil
}

func All() error {
	if err := Lint(); err != nil {
		return err
	}

	if err := Test(); err != nil {
		return err
	}

	if err := Build(); err != nil {
		return err
	}

	return It()
}

func Prepare() error {
	return shellcmd.Command("pipx install reuse").Run()
}
