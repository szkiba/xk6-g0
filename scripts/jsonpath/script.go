package main

import (
	"net/http"

	"github.com/PaesslerAG/jsonpath"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

func Default(require *require.Assertions) {
	body := make(map[string]interface{})
	res, err := resty.New().R().SetResult(&body).Get("https://httpbin.test.k6.io/get")

	require.NoError(err, "request success")
	require.Equal(http.StatusOK, res.StatusCode(), "status code 200")

	val, err := jsonpath.Get("$.headers.Host", body)

	require.NoError(err, "$.headers.Host no error")
	require.Equal("httpbin.test.k6.io", val, "$.headers.Host value OK")
}
