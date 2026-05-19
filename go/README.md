# Go API Client Boilerplate

A thread-safe, concurrent API client implementation in Go.

## Features
- **Concurrent Requests:** Execute multiple HTTP requests in parallel using goroutines.
- **Context Support:** Integrated with `context` for cancellation and timeouts.
- **Generic Requests:** Handles JSON bodies and custom headers.

## Usage

```go
client := NewAPIClient("https://api.example.com", 10*time.Second)
ctx := context.Background()

requests := []RequestParams{
    {Method: http.MethodGet, Path: "/resource"},
}

results := client.DoConcurrent(ctx, requests)
```

## Running Tests

```bash
go test ./...
```
