// Package gofakeit provides bindings for the gofakeit library.
package gofakeit

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the gofakeit package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name gofakeit github.com/brianvoe/gofakeit/v6

// Exports returns the exports for the gofakeit package.
func Exports(_ modules.VU) interp.Exports {
	return Symbols
}
