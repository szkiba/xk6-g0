// Package websocket provides the exports for the websocket module.
package websocket

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the websocket package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name websocket github.com/gorilla/websocket
//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name websocket github.com/coder/websocket

// Exports returns the exports for the websocket package.
func Exports(_ modules.VU) interp.Exports {
	return Symbols
}
