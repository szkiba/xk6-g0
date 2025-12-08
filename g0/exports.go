// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"sync"

	"dario.cat/mergo"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

// ExportsFunc defines a function that returns the exports for a module.
type ExportsFunc func(modules.VU) interp.Exports

// RegisterExports registers one or more ExportsFunc to the g0 module.
func RegisterExports(fn ...ExportsFunc) {
	registry.register(fn...)
}

type exportsRegistry struct {
	exports []ExportsFunc
	mu      sync.RWMutex
}

var registry exportsRegistry //nolint:gochecknoglobals

func (r *exportsRegistry) register(fn ...ExportsFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.exports = append(r.exports, fn...)
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
