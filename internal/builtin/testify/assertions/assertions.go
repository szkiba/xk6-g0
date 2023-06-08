// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package assertions

import "github.com/stretchr/testify/assert"

type Assertions struct {
	t assert.TestingT
}

func New(t assert.TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

type tHelper interface {
	Helper()
}

type tChecker interface {
	Check(string, bool)
}

//go:generate go run github.com/stretchr/testify/_codegen -output-package=assertions -template=assertions_intercept.go.tmpl -include-format-funcs
//go:generate sh -c "cat assertions_intercept.go | sed -e 's;import (;import (\\n  \"fmt\";g' > tmp.go; mv tmp.go assertions_intercept.go"
