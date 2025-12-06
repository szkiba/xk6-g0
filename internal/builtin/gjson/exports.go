// Package gjson provides bindings for the gjson library.
package gjson

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the gjson package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name gjson github.com/tidwall/gjson

// Exports returns the exports for the gjson package.
func Exports(_ modules.VU) interp.Exports {
	return Symbols
}
