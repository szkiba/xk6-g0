// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewModuleInstance(t *testing.T) {
	t.Parallel()

	tc := newHelper(t) // nolint:varnamelen
	mod := tc.module

	exports := mod.Exports()

	assert.Empty(t, exports.Default)
	assert.Empty(t, exports.Named)
}
