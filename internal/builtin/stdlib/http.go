// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package stdlib

import (
	"net/http"
	"reflect"

	"github.com/szkiba/xk6-g0/g0/addon"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

func httpExports(vu modules.VU) interp.Exports {
	transport := addon.NewTransport(vu)
	client := &http.Client{Transport: transport}

	return interp.Exports{
		"net/http/http": {
			"DefaultTransport": reflect.ValueOf(&transport).Elem(),
			"DefaultClient":    reflect.ValueOf(&client).Elem(),
			"Get":              reflect.ValueOf(client.Get),
			"Head":             reflect.ValueOf(client.Head),
			"Post":             reflect.ValueOf(client.Post),
			"PostForm":         reflect.ValueOf(client.PostForm),
		},
	}
}
