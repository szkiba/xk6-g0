// Package jsonschema provides bindings for the jsonschema library.
package jsonschema

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the jsonschema package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name jsonschema github.com/santhosh-tekuri/jsonschema/v5

// Exports returns the exports for the jsonschema package.
func Exports(_ modules.VU) interp.Exports {
	return Symbols
}
