package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
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

	results := client.DoConcurrent(context.Background(), requests, 2)

	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results) )
	}

	for _, res := range results {
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", res.StatusCode)
		}
	}
}

func TestAPIClient_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Set a very short timeout
	client := NewAPIClient(server.URL, 20*time.Millisecond)
	res := client.Do(context.Background(), RequestParams{
		Method: http.MethodGet,
		Path:   "/timeout",
	})

	if res.Error == nil {
		t.Error("Expected timeout error, got nil")
	}
}

func TestAPIClient_HeadersAndBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test") != "value" {
			t.Errorf("Expected header X-Test: value, got %s", r.Header.Get("X-Test"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL, 5*time.Second)
	client.Do(context.Background(), RequestParams{
		Method:  http.MethodPost,
		Path:    "/post",
		Headers: map[string]string{"X-Test": "value"},
		Body:    map[string]string{"foo": "bar"},
	})
}

func TestAPIClient_DoConcurrent_Throttling(t *testing.T) {
	var mu sync.Mutex
	active, maxActive := 0, 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		active++
		if active > maxActive {
			maxActive = active
		}
		mu.Unlock()
		time.Sleep(20 * time.Millisecond)
		mu.Lock()
		active--
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	client := NewAPIClient(server.URL, 5*time.Second)
	requests := []RequestParams{
		{Method: http.MethodGet, Path: "/1"},
		{Method: http.MethodGet, Path: "/2"},
		{Method: http.MethodGet, Path: "/3"},
		{Method: http.MethodGet, Path: "/4"},
	}
	client.DoConcurrent(context.Background(), requests, 2)
	if maxActive > 2 {
		t.Errorf("Expected max concurrent requests to be at most 2, got %d", maxActive)
	}
	if maxActive == 0 {
		t.Error("Expected some active requests, got 0")
	}
}
