package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCallOpenRouterRawReturnsParsedText(t *testing.T) {
	// Start a test server that returns a valid OpenRouter JSON response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":"1","choices":[{"message":{"content":"Hello test"}}]}`))
	}))
	defer ts.Close()

	// Set env vars expected by callOpenRouterRaw
	os.Setenv("OPENROUTER_API_URL", ts.URL)
	os.Setenv("OPENROUTER_API_KEY", "dummy")
	defer os.Unsetenv("OPENROUTER_API_URL")
	defer os.Unsetenv("OPENROUTER_API_KEY")

	ctx := context.Background()
	text, err := callOpenRouterRaw(ctx, "hi", "model-X")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if text != "Hello test" {
		t.Fatalf("expected parsed text, got: %s", text)
	}
}
