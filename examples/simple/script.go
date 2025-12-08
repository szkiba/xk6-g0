package main

import "net/http"

func Default() {
	_, _ = http.Get("https://httpbin.org/get")
}
