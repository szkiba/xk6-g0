// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"os"

	"github.com/szkiba/xk6-g0/g0/addon"
	"github.com/szkiba/xk6-g0/internal/builtin/testify/assertions"
	"github.com/szkiba/xk6-g0/internal/builtin/testify/requirements"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

const envScript = "XK6_G0_SCRIPT"

type RootModule struct{}

func New() modules.Module {
	return new(RootModule)
}

func (root *RootModule) NewModuleInstance(vu modules.VU) modules.Instance { // nolint:varnamelen
	yaegi := interp.New(interp.Options{}) //nolint:exhaustruct
	mod := &Module{                       //nolint:exhaustruct
		vu:      vu,
		yaegi:   yaegi,
		assert:  assertions.New(addon.NewTestingT(vu, false)),
		require: requirements.New(addon.NewTestingT(vu, true)),
	}

	if err := mod.init(os.Getenv(envScript)); err != nil {
		common.Throw(vu.Runtime(), err)
	}

	return mod
}

type Module struct {
	vu    modules.VU
	yaegi *interp.Interpreter

	setupFunc         SetupFunc
	teardownFunc      TeardownFunc
	defaultFunc       DefaultFunc
	options           map[string]interface{}
	handleSummaryFunc HandleSummaryFunc
	ctx               *contextWrapper

	assert  *assertions.Assertions
	require *requirements.Assertions
}

var (
	_ modules.Module   = (*RootModule)(nil)
	_ modules.Instance = (*Module)(nil)
)

func (mod *Module) initYaegi() error {
	symbols, err := registry.merge(mod.vu)
	if err != nil {
		return err
	}

	return mod.yaegi.Use(symbols)
}

func (mod *Module) init(filename string) error {
	if err := mod.initYaegi(); err != nil {
		return err
	}

	if len(filename) != 0 {
		return mod.compile(filename)
	}

	return nil
}

func (mod *Module) compile(filename string) error {
	prog, err := mod.yaegi.CompilePath(filename)
	if err != nil {
		return err
	}

	_, err = mod.yaegi.Execute(prog)
	if err != nil {
		return err
	}

	mod.initCallbacks()

	return nil
}

func (mod *Module) initCallbacks() {
	mod.setupFunc = toSetupFunc(mod.yaegi.Eval("Setup"))
	mod.teardownFunc = toTeardownFunc(mod.yaegi.Eval("Teardown"))
	mod.defaultFunc = toDefaultFunc(mod.yaegi.Eval("Default"))
	mod.options = toOptions(mod.yaegi.Eval("Options"))
	mod.handleSummaryFunc = toHandleSummaryFunc(mod.yaegi.Eval("HandleSummary"))
	mod.ctx = newContextWrapper(mod.vu)
}

func (mod *Module) Exports() modules.Exports {
	toValue := mod.vu.Runtime().ToValue

	exports := modules.Exports{
		Default: nil,
		Named:   make(map[string]interface{}),
	}

	if mod.defaultFunc != nil {
		exports.Default = toValue(mod.callDefault)
	}

	if mod.setupFunc != nil {
		exports.Named["setup"] = toValue(mod.callSetup)
	}

	if mod.teardownFunc != nil {
		exports.Named["teardown"] = toValue(mod.callTeardown)
	}

	if mod.handleSummaryFunc != nil {
		exports.Named["handleSummary"] = toValue(mod.callHandleSummary)
	}

	if mod.options != nil {
		exports.Named["options"] = toValue(mod.options)
	}

	return exports
}

func (mod *Module) callSetup() (interface{}, error) {
	mod.vu.State().Logger.Debug("Calling Setup")

	return mod.setupFunc(mod.ctx, mod.assert, mod.require)
}

func (mod *Module) callTeardown(data interface{}) error {
	mod.vu.State().Logger.Debug("Calling Teardown")

	return mod.teardownFunc(mod.ctx, mod.assert, mod.require, data)
}

func (mod *Module) callDefault(data interface{}) error {
	mod.vu.State().Logger.Debug("Calling Default")

	err := mod.defaultFunc(mod.ctx, mod.assert, mod.require, data)
	if err != nil {
		common.Throw(mod.vu.Runtime(), err)
	}

	return err
}

func (mod *Module) callHandleSummary(data map[string]interface{}) map[string]interface{} {
	mod.vu.State().Logger.Debug("Calling HandleSummary")

	res, err := mod.handleSummaryFunc(data)
	if err != nil {
		common.Throw(mod.vu.Runtime(), err)
	}

	return res
}
