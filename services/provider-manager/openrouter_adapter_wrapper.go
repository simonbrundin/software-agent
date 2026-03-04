package main

import (
	"context"
)

// wrapper to adapt existing callOpenRouterRaw to ProviderAdapter interface
type openRouterAdapter struct {
	Model string
}

func (o *openRouterAdapter) CallRaw(ctx context.Context, prompt string) (string, error) {
	raw, err := callOpenRouterRaw(ctx, prompt, o.Model)
	if err != nil {
		return "", err
	}
	// If callOpenRouterRaw already returns parsed text in future, this
	// double-parsing is safe because extractOpenRouterText will handle
	// a JSON payload or plain text gracefully (current implementation
	// expects JSON). For now return raw if parsing fails.
	parsed, perr := extractOpenRouterText([]byte(raw))
	if perr != nil {
		return raw, nil
	}
	return parsed, nil
}
