// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package jsonpath

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name jsonpath github.com/PaesslerAG/jsonpath github.com/PaesslerAG/gval

func Exports(vu modules.VU) interp.Exports {
	return Symbols
}
