// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package stdlib

import (
	"github.com/imdario/mergo"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"go.k6.io/k6/js/modules"
)

func Exports(vu modules.VU) interp.Exports {
	exports := interp.Exports{}

	if err := mergo.Merge(&exports, httpExports(vu)); err != nil {
		panic(err)
	}

	if err := mergo.Merge(&exports, interp.Exports(stdlib.Symbols)); err != nil {
		panic(err)
	}

	return exports
}
