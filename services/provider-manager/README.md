Provider Manager

This service exposes a single endpoint `/v1/generate` that accepts generation requests and attempts them against a prioritized list of providers.

Adapters:
- openrouter
- copilot (stub)

Provider Adapter Contract:
- Implement the `ProviderAdapter` interface: `CallRaw(ctx context.Context, prompt string) (string, error)`.
- Adapters should return the provider's raw text response or an error.
- Keep error handling explicit; do not ignore errors.

Fallback behavior:
- The service attempts providers in priority order (e.g. `openrouter` then `copilot`).
- On error or empty response, the next provider is attempted.

Metrics:
- prometheus: provider_calls_total, provider_errors_total, provider_cost_total
