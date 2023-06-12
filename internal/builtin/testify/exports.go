// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package testify

import (
	"reflect"

	"github.com/imdario/mergo"
	"github.com/szkiba/xk6-g0/internal/builtin/testify/assertions"
	"github.com/szkiba/xk6-g0/internal/builtin/testify/requirements"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name testify github.com/stretchr/testify/assert github.com/stretchr/testify/require

func Exports(vu modules.VU) interp.Exports {
	exports := interp.Exports{
		"github.com/stretchr/testify/assert/assert": {
			"Assertions": reflect.ValueOf((*assertions.Assertions)(nil)),
			"New":        reflect.ValueOf(assertions.New),
		},
		"github.com/stretchr/testify/require/require": {
			"Assertions": reflect.ValueOf((*requirements.Assertions)(nil)),
			"New":        reflect.ValueOf(requirements.New),
		},
	}

	if err := mergo.Merge(&exports, Symbols); err != nil {
		panic(err)
	}

	return exports
}
