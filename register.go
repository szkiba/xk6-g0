// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"github.com/szkiba/xk6-g0/g0"
	"go.k6.io/k6/js/modules"
)

func init() { //nolint:gochecknoinits
	register()
	g0.RedirectStdin()
}

func register() {
	modules.Register("k6/x/g0", g0.New())
}
