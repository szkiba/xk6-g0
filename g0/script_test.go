// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

// placeholder test suite, to be continued...

package g0

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type scriptSuite struct {
	suiteBase
}

func TestScript(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(scriptSuite))
}

func (suite *scriptSuite) TestSimple() {
	suite.g0(`
	func Default() {}
	`)

	exports := suite.module.Exports()

	suite.Len(exports.Named, 1) // workaround for missing handleSummary
	suite.NotEmpty(exports.Default)
}
