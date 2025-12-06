// Package jsonpath provides bindings for the jsonpath library.
package jsonpath

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the jsonpath package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name jsonpath github.com/PaesslerAG/jsonpath github.com/PaesslerAG/gval

// Exports returns the exports for the jsonpath package.
func Exports(_ modules.VU) interp.Exports {
	return Symbols
}
