// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"sync"

	"github.com/imdario/mergo"
	"github.com/szkiba/xk6-g0/internal/builtin/gjson"
	"github.com/szkiba/xk6-g0/internal/builtin/gofakeit"
	"github.com/szkiba/xk6-g0/internal/builtin/goquery"
	"github.com/szkiba/xk6-g0/internal/builtin/jsonpath"
	jsonschema "github.com/szkiba/xk6-g0/internal/builtin/jsonscema"
	"github.com/szkiba/xk6-g0/internal/builtin/logrus"
	"github.com/szkiba/xk6-g0/internal/builtin/resty"
	"github.com/szkiba/xk6-g0/internal/builtin/stdlib"
	"github.com/szkiba/xk6-g0/internal/builtin/testify"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

type ExportsFunc func(modules.VU) interp.Exports

func RegisterExports(fn ExportsFunc) {
	registry.register(fn)
}

type exportsRegistry struct {
	exports []ExportsFunc
	mu      sync.RWMutex
}

var registry exportsRegistry

func (r *exportsRegistry) register(fn ExportsFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.exports = append(r.exports, fn)
}

func (r *exportsRegistry) merge(vu modules.VU) (interp.Exports, error) { //nolint:varnamelen
	r.mu.RLock()
	defer r.mu.RUnlock()

	symbols := interp.Exports{}

	for _, fn := range r.exports {
		if err := mergo.Merge(&symbols, fn(vu)); err != nil {
			return nil, err
		}
	}

	return symbols, nil
}

func init() {
	RegisterExports(stdlib.Exports)
	RegisterExports(logrus.Exports)
	RegisterExports(resty.Exports)
	RegisterExports(testify.Exports)
	RegisterExports(goquery.Exports)
	RegisterExports(gjson.Exports)
	RegisterExports(jsonpath.Exports)
	RegisterExports(jsonschema.Exports)
	RegisterExports(gofakeit.Exports)
}
