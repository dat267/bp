package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAPIClient_Do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := NewAPIClient(server.URL, 5*time.Second)
	res := client.Do(context.Background(), RequestParams{
		Method: http.MethodGet,
		Path:   "/test",
	})

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
}

func TestAPIClient_DoConcurrent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL, 5*time.Second)
	requests := []RequestParams{
		{Method: http.MethodGet, Path: "/1"},
		{Method: http.MethodGet, Path: "/2"},
		{Method: http.MethodGet, Path: "/3"},
	}

	results := client.DoConcurrent(context.Background(), requests)

	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}

	for _, res := range results {
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", res.StatusCode)
		}
	}
}
