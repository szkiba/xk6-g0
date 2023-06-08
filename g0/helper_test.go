// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.k6.io/k6/js/modulestest"
)

type testHelper struct {
	runtime *modulestest.Runtime
	vu      *modulestest.VU
	module  *Module
}

func newHelper(t *testing.T) *testHelper {
	t.Helper()

	runtime := modulestest.NewRuntime(t)
	vu := runtime.VU // nolint:varnamelen

	assert.NoError(t, vu.Runtime().Set("__VU", 1))

	root := New()

	var module *Module

	assert.NotPanics(t, func() { module = root.NewModuleInstance(vu).(*Module) }) // nolint:forcetypeassert
	assert.NotNil(t, module)

	return &testHelper{
		runtime: runtime,
		vu:      vu,
		module:  module,
	}
}

type suiteBase struct {
	suite.Suite
	*testHelper
}

func (suite *suiteBase) SetupSuite() {
	suite.testHelper = newHelper(suite.T())
	suite.NoError(suite.module.initYaegi())
}

func (suite *suiteBase) run(script string) error {
	_, err := suite.module.yaegi.Eval(script)
	if err != nil {
		return err
	}

	suite.module.initCallbacks()

	return nil
}

func (suite *suiteBase) g0(script string) {
	suite.NoError(suite.run(script))
}
