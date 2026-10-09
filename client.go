package integrations

import (
	"context"
	"net/http"
)

// HTTPClient represents the minimal interface required for making outbound HTTP calls.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ProviderClient defines the standard REST operations supported by checkout/licensing providers.
type ProviderClient interface {
	// Provider returns the provider identifier this client communicates with.
	Provider() Provider

	// GetUser retrieves the seller/creator account details from the provider.
	GetUser(ctx context.Context, accessToken string) (*User, error)

	// ListProducts retrieves the creator's product catalog from the provider.
	ListProducts(ctx context.Context, accessToken string) ([]Product, error)

	// GetProduct retrieves a single product by its ID or permalink.
	GetProduct(ctx context.Context, accessToken string, productID string) (*Product, error)

	// VerifyLicense verifies a license key directly with the provider.
	VerifyLicense(ctx context.Context, req LicenseVerificationRequest) (*LicenseVerificationResult, error)

	// CreateWebhookSubscription registers a webhook listener with the provider.
	CreateWebhookSubscription(ctx context.Context, accessToken string, req CreateWebhookSubscriptionRequest) (*WebhookSubscription, error)

	// ListWebhookSubscriptions lists registered webhooks with the provider.
	ListWebhookSubscriptions(ctx context.Context, accessToken string) ([]WebhookSubscription, error)

	// DeleteWebhookSubscription unregisters a webhook listener with the provider.
	DeleteWebhookSubscription(ctx context.Context, accessToken string, subscriptionID string) error
}

// WebhookParser defines the standard contract for validating and normalizing incoming webhooks.
type WebhookParser interface {
	// Provider returns the provider this parser handles.
	Provider() Provider

	// ParseWebhook parses and validates an incoming HTTP webhook request into a NormalizedWebhookEvent.
	ParseWebhook(r *http.Request, secret string) (*NormalizedWebhookEvent, error)

	// ParseWebhookPayload parses and validates raw payload bytes and headers into a NormalizedWebhookEvent.
	ParseWebhookPayload(payload []byte, headers http.Header, secret string) (*NormalizedWebhookEvent, error)
}

// Integration represents a complete provider integration implementing both API client operations and webhook parsing.
type Integration interface {
	ProviderClient
	WebhookParser
}
