// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package goquery

import (
	"net/http"
	"reflect"

	"github.com/PuerkitoBio/goquery"
	"github.com/imdario/mergo"
	"github.com/szkiba/xk6-g0/g0/addon"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name goquery github.com/PuerkitoBio/goquery

func Exports(vu modules.VU) interp.Exports {
	newDocument := func(url string) (*goquery.Document, error) {
		client := &http.Client{Transport: addon.NewTransport(vu)}

		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}

		return goquery.NewDocumentFromResponse(resp)
	}
	exports := interp.Exports{
		"github.com/PuerkitoBio/goquery/goquery": {
			"NewDocument": reflect.ValueOf(newDocument),
		},
	}

	if err := mergo.Merge(&exports, Symbols); err != nil {
		panic(err)
	}

	return exports
}
