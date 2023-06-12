// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package jsonschema

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate yaegi extract -name jsonschema github.com/santhosh-tekuri/jsonschema/v5

func Exports(vu modules.VU) interp.Exports {
	return Symbols
}