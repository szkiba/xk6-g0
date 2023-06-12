// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"github.com/szkiba/xk6-g0/internal/builtin/gjson"
	"github.com/szkiba/xk6-g0/internal/builtin/gofakeit"
	"github.com/szkiba/xk6-g0/internal/builtin/goquery"
	"github.com/szkiba/xk6-g0/internal/builtin/jsonpath"
	"github.com/szkiba/xk6-g0/internal/builtin/jsonschema"
	"github.com/szkiba/xk6-g0/internal/builtin/logrus"
	"github.com/szkiba/xk6-g0/internal/builtin/resty"
	"github.com/szkiba/xk6-g0/internal/builtin/stdlib"
	"github.com/szkiba/xk6-g0/internal/builtin/testify"
	"go.k6.io/k6/js/modules"
)

func registerBuiltins() {
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

func registerExtension() {
	modules.Register("k6/x/g0", New())
}

func Bootstrap() {
	redirectStdin()
	registerBuiltins()
	registerExtension()
}
