package main

import (
	"encoding/json"
	"net/http"
)

func Default() {
	_, _ = http.Get("https://httpbin.org/get")
}

type M map[string]interface{}

func HandleSummary(data M) (M, error) {
	bin, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}

	return M{
		"stdout": bin,
	}, nil
}
