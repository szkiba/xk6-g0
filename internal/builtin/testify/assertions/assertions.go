// Package assertions provides wrappers around testify assertions.
package assertions

import "github.com/stretchr/testify/assert"

// Assertions wraps testify assertions to provide better error messages in k6.
type Assertions struct {
	t assert.TestingT
}

// New creates a new Assertions instance.
func New(t assert.TestingT) *Assertions {
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

//go:generate go run github.com/stretchr/testify/_codegen -output-package=assertions -template=assertions_intercept.go.tmpl -include-format-funcs
//go:generate sh -c "cat assertions_intercept.go | sed -e 's;import (;import (\\n  \"fmt\";g' > tmp.go; mv tmp.go assertions_intercept.go"
