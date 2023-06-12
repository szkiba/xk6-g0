package main

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func Default(require *require.Assertions) {
	res, err := resty.New().R().Get("https://httpbin.test.k6.io/get")

	require.NoError(err, "request success")
	require.Equal(http.StatusOK, res.StatusCode(), "status code 200")

	body := res.Body()

	val := gjson.GetBytes(body, "headers.Host").Str

	require.Equal("httpbin.test.k6.io", val, "headers.Host value OK")
}
