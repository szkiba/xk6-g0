// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package requirements

import "github.com/stretchr/testify/require"

type Assertions struct {
	t require.TestingT
}

func New(t require.TestingT) *Assertions {
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

//go:generate go run github.com/stretchr/testify/_codegen -output-package=requirements -template=requirements_intercept.go.tmpl -include-format-funcs
//go:generate sh -c "cat requirements_intercept.go | sed -e 's;import (;import (\\n  \"fmt\";g' -e 's;time \"time\";time \"time\";g' > tmp.go; mv tmp.go requirements_intercept.go"
