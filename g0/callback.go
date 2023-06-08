// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

//
// "Abandon all hope, ye who enter here." - Dante
//
// code here is ugly, but works and maybe faster than a prettier solutions...

package g0

import (
	"context"
	"reflect"

	"github.com/szkiba/xk6-g0/internal/builtin/testify/assertions"
	"github.com/szkiba/xk6-g0/internal/builtin/testify/requirements"
)

type SetupFunc func(context.Context, *assertions.Assertions, *requirements.Assertions) (interface{}, error)

type TeardownFunc func(context.Context, *assertions.Assertions, *requirements.Assertions, interface{}) error

type DefaultFunc func(context.Context, *assertions.Assertions, *requirements.Assertions, interface{}) error

type HandleSummaryFunc func(map[string]interface{}) (map[string]interface{}, error)

type Options map[string]interface{}

func toSetupFunc(val reflect.Value, err error) SetupFunc { // nolint:funlen,cyclop
	if err != nil {
		return nil
	}

	iface := val.Interface()

	if fn, ok := iface.(func(context.Context, *assertions.Assertions) (interface{}, error)); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(ctx, a)
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions) (interface{}, error)); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			return fn(ctx, r)
		}
	}

	//

	if fn, ok := iface.(func(context.Context, *assertions.Assertions) interface{}); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(ctx, a), nil
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions) interface{}); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			return fn(ctx, r), nil
		}
	}

	//

	if fn, ok := iface.(func(context.Context, *assertions.Assertions) error); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			err := fn(ctx, a)

			return nil, err
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions) error); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			err := fn(ctx, r)

			return nil, err
		}
	}

	//

	if fn, ok := iface.(func(context.Context, *assertions.Assertions)); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			fn(ctx, a)

			return nil, nil
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions)); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			fn(ctx, r)

			return nil, nil
		}
	}

	//

	if fn, ok := iface.(func(context.Context) (interface{}, error)); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(ctx)
		}
	}

	if fn, ok := iface.(func(context.Context) interface{}); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(ctx), nil
		}
	}

	if fn, ok := iface.(func(context.Context) error); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			err := fn(ctx)

			return nil, err
		}
	}

	if fn, ok := iface.(func(context.Context)); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			fn(ctx)

			return nil, nil
		}
	}

	//

	if fn, ok := iface.(func(*assertions.Assertions) (interface{}, error)); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(a)
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions) (interface{}, error)); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			return fn(r)
		}
	}

	if fn, ok := iface.(func(*assertions.Assertions) interface{}); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(a), nil
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions) interface{}); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			return fn(r), nil
		}
	}

	if fn, ok := iface.(func(*assertions.Assertions) error); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			err := fn(a)

			return nil, err
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions) error); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			err := fn(r)

			return nil, err
		}
	}

	if fn, ok := iface.(func(*assertions.Assertions)); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			fn(a)

			return nil, nil
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions)); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions) (interface{}, error) {
			fn(r)

			return nil, nil
		}
	}

	//

	if fn, ok := iface.(func() (interface{}, error)); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn()
		}
	}

	if fn, ok := iface.(func() interface{}); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			return fn(), nil
		}
	}

	if fn, ok := iface.(func() error); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			err := fn()

			return nil, err
		}
	}

	if fn, ok := iface.(func()); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions) (interface{}, error) {
			fn()

			return nil, nil
		}
	}

	return nil
}

func toDefaultFunc(val reflect.Value, err error) DefaultFunc {
	return toCallbackFunc(val, err)
}

func toTeardownFunc(val reflect.Value, err error) TeardownFunc {
	return toCallbackFunc(val, err)
}

func toCallbackFunc(val reflect.Value, err error) func(context.Context, *assertions.Assertions, *requirements.Assertions, interface{}) error { // nolint:funlen,cyclop
	if err != nil {
		return nil
	}

	iface := val.Interface()

	if fn, ok := iface.(func(context.Context, *assertions.Assertions, interface{}) error); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			return fn(ctx, a, data)
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions, interface{}) error); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions, data interface{}) error {
			return fn(ctx, r, data)
		}
	}

	//

	if fn, ok := iface.(func(context.Context, *assertions.Assertions, interface{})); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			fn(ctx, a, data)

			return nil
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions, interface{})); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions, data interface{}) error {
			fn(ctx, r, data)

			return nil
		}
	}

	//

	if fn, ok := iface.(func(context.Context, *assertions.Assertions) error); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			return fn(ctx, a)
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions) error); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions, data interface{}) error {
			return fn(ctx, r)
		}
	}

	//

	if fn, ok := iface.(func(context.Context, *assertions.Assertions)); ok {
		return func(ctx context.Context, a *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			fn(ctx, a)

			return nil
		}
	}

	if fn, ok := iface.(func(context.Context, *requirements.Assertions)); ok {
		return func(ctx context.Context, _ *assertions.Assertions, r *requirements.Assertions, data interface{}) error {
			fn(ctx, r)

			return nil
		}
	}

	//

	if fn, ok := iface.(func(context.Context, interface{}) error); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			return fn(ctx, data)
		}
	}

	if fn, ok := iface.(func(context.Context, interface{})); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			fn(ctx, data)

			return nil
		}
	}

	if fn, ok := iface.(func(context.Context) error); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			return fn(ctx)
		}
	}

	if fn, ok := iface.(func(context.Context)); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			fn(ctx)

			return nil
		}
	}

	//

	if fn, ok := iface.(func(*assertions.Assertions, interface{}) error); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			return fn(a, data)
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions, interface{}) error); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions, data interface{}) error {
			return fn(r, data)
		}
	}

	if fn, ok := iface.(func(*assertions.Assertions, interface{})); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			fn(a, data)

			return nil
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions, interface{})); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions, data interface{}) error {
			fn(r, data)

			return nil
		}
	}

	if fn, ok := iface.(func(*assertions.Assertions) error); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions, _ interface{}) error {
			return fn(a)
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions) error); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions, _ interface{}) error {
			return fn(r)
		}
	}

	if fn, ok := iface.(func(*assertions.Assertions)); ok {
		return func(_ context.Context, a *assertions.Assertions, _ *requirements.Assertions, _ interface{}) error {
			fn(a)

			return nil
		}
	}

	if fn, ok := iface.(func(*requirements.Assertions)); ok {
		return func(_ context.Context, _ *assertions.Assertions, r *requirements.Assertions, _ interface{}) error {
			fn(r)

			return nil
		}
	}

	//

	if fn, ok := iface.(func(interface{}) error); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			return fn(data)
		}
	}

	if fn, ok := iface.(func(interface{})); ok {
		return func(ctx context.Context, _ *assertions.Assertions, _ *requirements.Assertions, data interface{}) error {
			fn(data)

			return nil
		}
	}

	if fn, ok := iface.(func() error); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions, _ interface{}) error {
			return fn()
		}
	}

	if fn, ok := iface.(func()); ok {
		return func(_ context.Context, _ *assertions.Assertions, _ *requirements.Assertions, _ interface{}) error {
			fn()

			return nil
		}
	}

	return nil
}

func toOptions(val reflect.Value, err error) Options {
	if err != nil {
		return nil
	}

	iface := val.Interface()

	if m, ok := iface.(map[string]interface{}); ok {
		return m
	}

	return nil
}

func toHandleSummaryFunc(val reflect.Value, err error) HandleSummaryFunc {
	if err != nil {
		return nil
	}

	iface := val.Interface()

	if fn, ok := iface.(func(map[string]interface{}) (map[string]interface{}, error)); ok {
		return fn
	}

	if fn, ok := iface.(func(map[string]interface{}) map[string]interface{}); ok {
		return func(data map[string]interface{}) (map[string]interface{}, error) {
			return fn(data), nil
		}
	}

	return nil
}
