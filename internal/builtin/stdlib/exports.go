// Package stdlib provides the standard library exports for k6.
package stdlib

import (
	"github.com/imdario/mergo"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"go.k6.io/k6/js/modules"
)

// Exports returns the exports for the standard library.
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
