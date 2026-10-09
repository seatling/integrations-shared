# Seatling Integrations Shared (`integrations-shared`)

`github.com/seatling/integrations-shared` defines the canonical **types, models, and interfaces** for Seatling's checkout and payment integrations (Gumroad, Lemon Squeezy, Polar, Stripe, etc.).

---

## 1. Architectural Role & Boundary

This repository is a **pure contract / interface module**:
- **Interface & Type Specification**: Defines standard Go interfaces (`ProviderClient`, `WebhookParser`, `Integration`) and normalized data models (`Order`, `Subscription`, `Product`, `User`, `NormalizedWebhookEvent`).
- **Persistence & API Interoperability**: All models include dual, column-aligned `bson:"..."` and `json:"..."` annotations for seamless serialization and storage across MongoDB pipelines and Go JSON encoders.
- **No Concrete Implementations**: Concrete HTTP requests, webhook signature verification algorithms, and provider-specific payload decoders are implemented in **specialized provider libraries** (e.g. `seatling-integrations-gumroad`, `seatling-integrations-lemonsqueezy`, `seatling-integrations-polar`, `seatling-integrations-stripe`).
- **Service Isolation**: Seatling backend services import the specialized provider libraries and interact with them via these shared contracts, ensuring all providers return identical, strongly-typed data structures without coupling services to provider quirks.
- **Zero Third-Party Dependencies**: 100% Go standard library.

```
+-----------------------------------------------------------------------------------------+
|                  github.com/seatling/integrations-shared (This Module)                  |
|  - Interfaces: ProviderClient, WebhookParser, Integration                                |
|  - Models (BSON & JSON): User, Product, Variant, Customer, Order, Subscription          |
|  - Enums: Provider, EventType, SubscriptionStatus                                       |
|  - Errors: APIError, Sentinels                                                          |
+-----------------------------------------------------------------------------------------+
                                             ^
                                             | implements interfaces & returns shared types
        +------------------------------------+------------------------------------+
        |                                    |                                    |
+--------------------------+  +-------------------------------+  +--------------------------+
|  integrations-gumroad    |  |  integrations-lemonsqueezy    |  |   integrations-polar     |
| (Specialized Go library) |  |   (Specialized Go library)    |  | (Specialized Go library) |
+--------------------------+  +-------------------------------+  +--------------------------+
        ^                                    ^                                    ^
        |                                    |                                    |
        +------------------------------------+------------------------------------+
                                             | imports specialized library & consumes contracts
+-----------------------------------------------------------------------------------------+
|                    Seatling Microservices (services/gumroad, workers)                   |
+-----------------------------------------------------------------------------------------+
```

---

## 2. Core Interfaces

```go
package integrations

// ProviderClient defines the standard REST operations supported by checkout/licensing providers.
type ProviderClient interface {
    Provider() Provider
    GetUser(ctx context.Context, accessToken string) (*User, error)
    ListProducts(ctx context.Context, accessToken string) ([]Product, error)
    GetProduct(ctx context.Context, accessToken string, productID string) (*Product, error)
    VerifyLicense(ctx context.Context, req LicenseVerificationRequest) (*LicenseVerificationResult, error)
    CreateWebhookSubscription(ctx context.Context, accessToken string, req CreateWebhookSubscriptionRequest) (*WebhookSubscription, error)
    ListWebhookSubscriptions(ctx context.Context, accessToken string) ([]WebhookSubscription, error)
    DeleteWebhookSubscription(ctx context.Context, accessToken string, subscriptionID string) error
}

// WebhookParser defines the standard contract for validating and normalizing incoming webhooks.
type WebhookParser interface {
    Provider() Provider
    ParseWebhook(r *http.Request, secret string) (*NormalizedWebhookEvent, error)
    ParseWebhookPayload(payload []byte, headers http.Header, secret string) (*NormalizedWebhookEvent, error)
}

// Integration represents a complete provider integration implementing both API client operations and webhook parsing.
type Integration interface {
    ProviderClient
    WebhookParser
}
```

---

## 3. Normalized Models & Types

All provider implementations translate provider-specific schemas into these normalized structs:

| Model | Purpose |
| :--- | :--- |
| `User` | Authenticated merchant/creator account profile from `GetUser`. |
| `Product` | Catalog product information, permalinks, pricing, and variants. |
| `Variant` | SKU / pricing tier variation of a product. |
| `Customer` | Buyer identity and contact info. |
| `Order` | Completed checkout purchase or sale record. |
| `Subscription` | Recurring billing agreement and lifecycle status (`active`, `past_due`, `cancelled`, etc.). |
| `Refund` | Refund transaction information. |
| `LicenseVerificationRequest` / `Result` | Standardized parameters and results for upstream license validation. |
| `WebhookSubscription` | Webhook listener endpoint registered with a provider. |
| `NormalizedWebhookEvent` | Canonical webhook event envelope emitted by any provider's parser. |

---

## 4. Provider Taxonomy & Enums

- **Providers**: `ProviderGumroad ("gumroad")`, `ProviderLemonSqueezy ("lemon_squeezy")`, `ProviderPolar ("polar")`, `ProviderStripe ("stripe")`.
- **Event Types**: `order.created`, `order.refunded`, `order.disputed`, `order.cancelled`, `subscription.created`, `subscription.updated`, `subscription.cancelled`, `subscription.paused`, `subscription.resumed`, `subscription.expired`, `license_key.created`, `license_key.updated`.
- **Subscription Statuses**: `active`, `past_due`, `unpaid`, `cancelled`, `expired`, `paused`, `trialing`.

---

## 5. Standard Error Taxonomy

Common sentinel errors for upstream error discrimination:
- `ErrUnauthorized`: HTTP 401/403 or invalid access token.
- `ErrNotFound`: HTTP 404 resource not found.
- `ErrRateLimited`: HTTP 429 rate limit exceeded.
- `ErrMissingAccessToken`: Required access token was omitted.
- `ErrInvalidParameter`: Input argument failed validation.
- `ErrWebhookVerificationFailed`: Signature or passcode mismatch.
- `ErrInvalidPayload`: Malformed request body.
- `ErrResourceConflict`: HTTP 409 resource conflict.

The `*APIError` type implements `error`, `errors.Is`, and provides a `Retryable() bool` method for automated retry backoff.

---

## 6. Development & Testing

```bash
# Run tests
go test -count=1 -v -cover ./...

# Run static analysis
go vet ./...
```
