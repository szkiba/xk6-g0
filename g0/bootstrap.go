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
	registerExports(stdlib.Exports)
	registerExports(logrus.Exports)
	registerExports(resty.Exports)
	registerExports(testify.Exports)
	registerExports(goquery.Exports)
	registerExports(gjson.Exports)
	registerExports(jsonpath.Exports)
	registerExports(jsonschema.Exports)
	registerExports(gofakeit.Exports)
}

func registerExtension() {
	modules.Register("k6/x/g0", New())
}

func Bootstrap(addons ...ExportsFunc) {
	registerBuiltins()
	registerExports(addons...)
	redirectStdin()
	registerExtension()
}
