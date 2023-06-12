// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"context"
	"time"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

type contextWrapper struct {
	vu modules.VU
}

var _ context.Context = (*contextWrapper)(nil)

func newContextWrapper(vu modules.VU) *contextWrapper {
	return &contextWrapper{vu: vu}
}

func (w *contextWrapper) Deadline() (time.Time, bool) {
	return w.vu.Context().Deadline()
}

func (w *contextWrapper) Done() <-chan struct{} {
	return w.vu.Context().Done()
}

func (w *contextWrapper) Err() error {
	return w.vu.Context().Err()
}

func (w *contextWrapper) Value(key any) any {
	switch expr := key.(type) {
	case string:
		if val := w.vu.Context().Value(key); val != nil {
			return val
		}

		return w.eval(expr)
	default:
		return w.vu.Context().Value(key)
	}
}

func (w *contextWrapper) eval(expr string) any {
	val, err := w.vu.Runtime().RunString(expr)
	if err != nil {
		common.Throw(w.vu.Runtime(), err)
	}

	return w.export(val)
}

func (w *contextWrapper) export(val goja.Value) any {
	if function, ok := goja.AssertFunction(val); ok {
		return func(goargs ...any) any {
			jsargs := make([]goja.Value, 0, len(goargs))

			for _, v := range goargs {
				jsargs = append(jsargs, w.vu.Runtime().ToValue(v))
			}

			val, err := function(goja.Undefined(), jsargs...)
			if err != nil {
				common.Throw(w.vu.Runtime(), err)
			}

			return val.Export()
		}
	}

	return val.Export()
}
