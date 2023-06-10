// nolint:forcetypeassert
package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

func Default(ctx context.Context) {
	vu := ctx.Value("__VU").(int64)
	env := ctx.Value("__ENV").(map[string]string)
	iter := ctx.Value("__ITER").(int64)

	logrus.Info(vu)
	logrus.Info(iter)
	logrus.Info(env["PATH"])
	logrus.Info(ctx.Value("execution.scenario.name"))
}
