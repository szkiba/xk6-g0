// nolint:forcetypeassert,gomnd
package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

func Default(ctx context.Context) {
	add := ctx.Value("add").(func(...any) any)

	logrus.Info(add(2, 3))
	logrus.Info(ctx.Value("welcome"))
}

const JS = `//js

global.add = function (a, b) {
  return a + b
}

global.welcome = 'Hello, World!'
`
