package main

import (
	"context"
	"errors"
	"testing"
)

// A minimal ProviderAdapter contract for testing
type mockAdapter struct {
	resp string
	err  error
}

func (m *mockAdapter) CallRaw(ctx context.Context, prompt string) (string, error) {
	return m.resp, m.err
}

func TestAdapterSuccess(t *testing.T) {
	ctx := context.Background()
	m := &mockAdapter{resp: "ok", err: nil}
	got, err := m.CallRaw(ctx, "hello")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != "ok" {
		t.Fatalf("unexpected resp: %s", got)
	}
}

func TestAdapterError(t *testing.T) {
	ctx := context.Background()
	ex := errors.New("fail")
	m := &mockAdapter{resp: "", err: ex}
	_, err := m.CallRaw(ctx, "hello")
	if err == nil {
		t.Fatalf("expected error")
	}
}

// Test fallback orchestration: try primary, on error try fallback
func TestFallbackOrchestration(t *testing.T) {
	ctx := context.Background()
	primary := &mockAdapter{resp: "", err: errors.New("primary fail")}
	fallback := &mockAdapter{resp: "fallback ok", err: nil}

	// minimal orchestrator function inline for test
	callWithFallback := func(ctx context.Context, primary, fallback ProviderAdapter, prompt string) (string, error) {
		res, err := primary.CallRaw(ctx, prompt)
		if err == nil && res != "" {
			return res, nil
		}
		return fallback.CallRaw(ctx, prompt)
	}

	res, err := callWithFallback(ctx, primary, fallback, "p")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if res != "fallback ok" {
		t.Fatalf("unexpected fallback res: %s", res)
	}
}
