package main

import "github.com/sirupsen/logrus"

func Setup() interface{} {
	logrus.Info("Setup")

	return map[string]interface{}{
		"foo": "bar",
	}
}

func Default(data interface{}) {
	logrus.Info("Default", data)
}

func Teardown(data interface{}) {
	logrus.Info("Teardown", data)
}

func init() {
	logrus.Info("init")
}
