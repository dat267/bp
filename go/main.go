package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := NewAPIClient("https://httpbin.org", 10*time.Second)
	ctx := context.Background()

	requests := []RequestParams{
		{Method: http.MethodGet, Path: "/get"},
		{Method: http.MethodPost, Path: "/post", Body: map[string]string{"msg": "hello"}},
		{Method: http.MethodPut, Path: "/put", Body: map[string]string{"update": "true"}},
		{Method: http.MethodDelete, Path: "/delete"},
		{Method: http.MethodPatch, Path: "/patch", Body: map[string]string{"patch": "true"}},
		{Method: http.MethodOptions, Path: "/options"},
		{Method: http.MethodHead, Path: "/get"},
		{Method: http.MethodTrace, Path: "/trace"},
	}

	results := client.DoConcurrent(ctx, requests, 3)

	for _, res := range results {
		fmt.Printf("Method: %-7s | Status: %3d | Error: %v\n", res.Method, res.StatusCode, res.Error)
	}
}
