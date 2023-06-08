package main

import (
	"net/http"

	"github.com/stretchr/testify/assert"
)

type M map[string]any

type A []any

var Options = M{
	"thresholds": M{
		"checks": A{"rate == 1"},
	},
}

func Default(assert *assert.Assertions) {
	res, err := http.Get("https://httpbin.test.k6.io/get")

	assert.NoError(err, "got response without error")
	assert.Equal(http.StatusOK, res.StatusCode, "status code was 200")
	assert.Equal("application/json", res.Header.Get("Content-Type"), "content type was application/json")
}
