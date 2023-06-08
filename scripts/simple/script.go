package main

import "net/http"

func Default() {
	http.Get("https://httpbin.test.k6.io/get")
}
