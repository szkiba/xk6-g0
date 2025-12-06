// Package requirements provides wrappers around testify requirements.
package requirements

import "github.com/stretchr/testify/require"

// Assertions wraps testify assertions to provide better error messages in k6.
type Assertions struct {
	t require.TestingT
}

// New creates a new Assertions instance.
func New(t require.TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

type tHelper interface {
	Helper()
}

type tChecker interface {
	Check(name string, succ bool)
}

//go:generate go run github.com/stretchr/testify/_codegen -output-package=requirements -template=requirements_intercept.go.tmpl -include-format-funcs
//go:generate sh -c "cat requirements_intercept.go | sed -e 's;import (;import (\\n  \"fmt\";g' -e 's;time \"time\";time \"time\";g' > tmp.go; mv tmp.go requirements_intercept.go"
