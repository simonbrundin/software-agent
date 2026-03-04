package main

import (
	"context"
	"testing"
)

func TestCopilotFallback(t *testing.T) {
	// Simulate callOpenRouter returning error and copilot being attempted
	// For simplicity, call the copilot stub directly
	ctx := context.Background()
	_, err := callCopilotRaw(ctx, "hello", "copilot")
	if err == nil {
		t.Fatalf("expected error from copilot stub")
	}
}
