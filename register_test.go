// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	assert.Panics(t, register) // already registered
}
