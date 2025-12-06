// Package resty provides bindings for the resty library.
package resty

import (
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/imdario/mergo"
	"github.com/szkiba/xk6-g0/g0/addon"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the resty package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name resty github.com/go-resty/resty/v2

// Exports returns the exports for the resty package.
func Exports(vu modules.VU) interp.Exports {
	transport := addon.NewTransport(vu)
	newClient := func() *resty.Client {
		client := resty.New()

		client.SetTransport(transport)

		return client
	}
	exports := interp.Exports{
		"github.com/go-resty/resty/v2/resty": {
			"New": reflect.ValueOf(newClient),
		},
	}

	if err := mergo.Merge(&exports, Symbols); err != nil {
		panic(err)
	}

	return exports
}
