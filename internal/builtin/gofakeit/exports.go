// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package gofakeit

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name gofakeit github.com/brianvoe/gofakeit/v6

func Exports(vu modules.VU) interp.Exports {
	return Symbols
}
