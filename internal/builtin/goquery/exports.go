// Package goquery provides bindings for the goquery library.
package goquery

import (
	"net/http"
	"reflect"

	"dario.cat/mergo"
	"github.com/PuerkitoBio/goquery"
	"github.com/szkiba/xk6-g0/g0/addon"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// Symbols holds the exported symbols of the goquery package.
var Symbols = interp.Exports{} //nolint:gochecknoglobals

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name goquery github.com/PuerkitoBio/goquery

// Exports returns the exports for the goquery package.
func Exports(vu modules.VU) interp.Exports {
	newDocument := func(url string) (*goquery.Document, error) {
		client := &http.Client{Transport: addon.NewTransport(vu)}

		req, err := http.NewRequestWithContext(vu.Context(), http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = resp.Body.Close()
		}()

		return goquery.NewDocumentFromReader(resp.Body)
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
