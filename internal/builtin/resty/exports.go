// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package resty

import (
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/imdario/mergo"
	"github.com/szkiba/xk6-g0/internal/builtin/stdlib"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate yaegi extract -name resty github.com/go-resty/resty/v2

func Exports(vu modules.VU) interp.Exports {
	transport := stdlib.NewTransport(vu)
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
