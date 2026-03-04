package main

import "context"

// ProviderAdapter defines the minimal contract for provider adapters.
// Implementations should call the underlying provider API and return
// the raw string response or an error.
type ProviderAdapter interface {
	CallRaw(ctx context.Context, prompt string) (string, error)
}
