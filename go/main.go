package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type RequestParams struct {
	Method  string
	Path    string
	Body    any
	Headers map[string]string
}

type ResponseResult struct {
	Method     string
	StatusCode int
	Data       []byte
	Error      error
}

func NewAPIClient(baseURL string, timeout time.Duration) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *APIClient) Do(ctx context.Context, params RequestParams) ResponseResult {
	var reqBody io.Reader
	if params.Body != nil {
		bodyBytes, _ := json.Marshal(params.Body)
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, params.Method, c.BaseURL+params.Path, reqBody)
	if err != nil {
		return ResponseResult{Method: params.Method, Error: err}
	}

	for k, v := range params.Headers {
		req.Header.Set(k, v)
	}
	if params.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return ResponseResult{Method: params.Method, Error: err}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	return ResponseResult{
		Method:     params.Method,
		StatusCode: resp.StatusCode,
		Data:       respBody,
		Error:      err,
	}
}

func (c *APIClient) DoConcurrent(ctx context.Context, requests []RequestParams) []ResponseResult {
	var wg sync.WaitGroup
	results := make([]ResponseResult, len(requests))

	for i, req := range requests {
		wg.Add(1)
		c.executeAsync(ctx, req, i, &wg, results)
	}

	wg.Wait()
	return results
}

func (c *APIClient) executeAsync(ctx context.Context, p RequestParams, idx int, wg *sync.WaitGroup, res []ResponseResult) {
	go func() {
		defer wg.Done()
		res[idx] = c.Do(ctx, p)
	}()
}

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

	results := client.DoConcurrent(ctx, requests)

	for _, res := range results {
		fmt.Printf("Method: %-7s | Status: %3d | Error: %v\n", res.Method, res.StatusCode, res.Error)
	}
}
