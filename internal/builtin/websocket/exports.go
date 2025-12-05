// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package websocket

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name websocket github.com/gorilla/websocket
//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name websocket github.com/coder/websocket

func Exports(vu modules.VU) interp.Exports {
	return Symbols
}
