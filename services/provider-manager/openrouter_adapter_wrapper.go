package main

import (
	"context"
)

// wrapper to adapt existing callOpenRouterRaw to ProviderAdapter interface
type openRouterAdapter struct {
	Model string
}

func (o *openRouterAdapter) CallRaw(ctx context.Context, prompt string) (string, error) {
	// callOpenRouterRaw now returns the parsed text when possible and
	// falls back to returning the raw body. We simply delegate to it.
	return callOpenRouterRaw(ctx, prompt, o.Model)
}
