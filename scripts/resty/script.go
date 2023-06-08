package main

import (
	"github.com/go-resty/resty/v2"
)

func Default() error {
	_, err := client.R().Get("https://httpbin.test.k6.io/get")

	return err
}

var client *resty.Client

func init() {
	client = resty.New()
}
