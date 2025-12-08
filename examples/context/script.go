package main

import (
	"context"
	"net/http"
)

func Default(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)

	return err
}
