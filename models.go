package integrations

import (
	"time"
)

// User represents an authenticated seller, merchant, or creator profile returned by a provider.
type User struct {
	ID                string         `bson:"id"                  json:"id"`
	Provider          Provider       `bson:"provider"            json:"provider"`
	Name              string         `bson:"name"                json:"name"`
	Email             string         `bson:"email"               json:"email"`
	ProfileURL        string         `bson:"profile_url,omitempty" json:"profile_url,omitempty"`
	Currency          string         `bson:"currency,omitempty"   json:"currency,omitempty"`
	CustomDomain      string         `bson:"custom_domain,omitempty" json:"custom_domain,omitempty"`
	SubscribedToEmail bool           `bson:"subscribed_to_email,omitempty" json:"subscribed_to_email,omitempty"`
	RawData           map[string]any `bson:"raw_data,omitempty"  json:"raw_data,omitempty"`
}

// Product represents a software application, digital good, or product synced from a provider.
type Product struct {
	ID             string         `bson:"id"                       json:"id"`
	Provider       Provider       `bson:"provider"                 json:"provider"`
	Name           string         `bson:"name"                     json:"name"`
	Description    string         `bson:"description,omitempty"     json:"description,omitempty"`
	URL            string         `bson:"url,omitempty"             json:"url,omitempty"` // Storefront or checkout permalink
	PriceCents     int            `bson:"price_cents"              json:"price_cents"`
	Currency       string         `bson:"currency"                 json:"currency"`
	FormattedPrice string         `bson:"formatted_price,omitempty" json:"formatted_price,omitempty"`
	Published      bool           `bson:"published"                json:"published"`
	Variants       []Variant      `bson:"variants,omitempty"       json:"variants,omitempty"`
	CreatedAt      *time.Time     `bson:"created_at,omitempty"      json:"created_at,omitempty"`
	UpdatedAt      *time.Time     `bson:"updated_at,omitempty"      json:"updated_at,omitempty"`
	RawData        map[string]any `bson:"raw_data,omitempty"       json:"raw_data,omitempty"`
}

// Variant represents a SKU, pricing tier, or variant of a product.
type Variant struct {
	ID             string         `bson:"id"                       json:"id"`
	ProductID      string         `bson:"product_id"               json:"product_id"`
	Name           string         `bson:"name"                     json:"name"`
	PriceCents     int            `bson:"price_cents"              json:"price_cents"`
	Currency       string         `bson:"currency"                 json:"currency"`
	FormattedPrice string         `bson:"formatted_price,omitempty" json:"formatted_price,omitempty"`
	IsDefault      bool           `bson:"is_default,omitempty"      json:"is_default,omitempty"`
	RawData        map[string]any `bson:"raw_data,omitempty"       json:"raw_data,omitempty"`
}

// Customer represents the buyer or licensee of an order.
type Customer struct {
	ID        string         `bson:"id,omitempty"         json:"id,omitempty"`
	Email     string         `bson:"email"                json:"email"`
	Name      string         `bson:"name,omitempty"       json:"name,omitempty"`
	Country   string         `bson:"country,omitempty"    json:"country,omitempty"`
	IPAddress string         `bson:"ip_address,omitempty" json:"ip_address,omitempty"`
	RawData   map[string]any `bson:"raw_data,omitempty"   json:"raw_data,omitempty"`
}

// Order represents a completed purchase, sale, or checkout invoice from a provider.
type Order struct {
	ID                  string            `bson:"id"                            json:"id"`
	Provider            Provider          `bson:"provider"                      json:"provider"`
	ProductID           string            `bson:"product_id"                    json:"product_id"`
	ProductName         string            `bson:"product_name,omitempty"        json:"product_name,omitempty"`
	VariantID           string            `bson:"variant_id,omitempty"          json:"variant_id,omitempty"`
	VariantName         string            `bson:"variant_name,omitempty"        json:"variant_name,omitempty"`
	Customer            Customer          `bson:"customer"                      json:"customer"`
	LicenseKey          string            `bson:"license_key,omitempty"         json:"license_key,omitempty"`
	PriceCents          int               `bson:"price_cents"                   json:"price_cents"`
	Currency            string            `bson:"currency"                      json:"currency"`
	TaxCents            int               `bson:"tax_cents,omitempty"           json:"tax_cents,omitempty"`
	DiscountCents       int               `bson:"discount_cents,omitempty"      json:"discount_cents,omitempty"`
	Refunded            bool              `bson:"refunded"                      json:"refunded"`
	RefundedAmountCents int               `bson:"refunded_amount_cents,omitempty" json:"refunded_amount_cents,omitempty"`
	Disputed            bool              `bson:"disputed"                      json:"disputed"`
	SubscriptionID      string            `bson:"subscription_id,omitempty"     json:"subscription_id,omitempty"`
	OrderDate           time.Time         `bson:"order_date"                    json:"order_date"`
	CustomFields        map[string]string `bson:"custom_fields,omitempty"       json:"custom_fields,omitempty"`
	RawData             map[string]any    `bson:"raw_data,omitempty"            json:"raw_data,omitempty"`
}

// Subscription represents a recurring subscription agreement managed by a provider.
type Subscription struct {
	ID                 string             `bson:"id"                            json:"id"`
	Provider           Provider           `bson:"provider"                      json:"provider"`
	ProductID          string             `bson:"product_id"                    json:"product_id"`
	VariantID          string             `bson:"variant_id,omitempty"          json:"variant_id,omitempty"`
	CustomerID         string             `bson:"customer_id,omitempty"         json:"customer_id,omitempty"`
	CustomerEmail      string             `bson:"customer_email"                json:"customer_email"`
	Status             SubscriptionStatus `bson:"status"                        json:"status"`
	LicenseKey         string             `bson:"license_key,omitempty"         json:"license_key,omitempty"`
	CurrentPeriodStart *time.Time         `bson:"current_period_start,omitempty" json:"current_period_start,omitempty"`
	CurrentPeriodEnd   *time.Time         `bson:"current_period_end,omitempty"   json:"current_period_end,omitempty"`
	CancelledAt        *time.Time         `bson:"cancelled_at,omitempty"        json:"cancelled_at,omitempty"`
	RenewsAt           *time.Time         `bson:"renews_at,omitempty"           json:"renews_at,omitempty"`
	EndsAt             *time.Time         `bson:"ends_at,omitempty"             json:"ends_at,omitempty"`
	CreatedAt          time.Time          `bson:"created_at"                    json:"created_at"`
	UpdatedAt          *time.Time         `bson:"updated_at,omitempty"          json:"updated_at,omitempty"`
	RawData            map[string]any     `bson:"raw_data,omitempty"            json:"raw_data,omitempty"`
}

// Refund represents refund information for an order or transaction.
type Refund struct {
	ID          string         `bson:"id,omitempty"   json:"id,omitempty"`
	OrderID     string         `bson:"order_id"       json:"order_id"`
	Provider    Provider       `bson:"provider"       json:"provider"`
	AmountCents int            `bson:"amount_cents"   json:"amount_cents"`
	Currency    string         `bson:"currency"       json:"currency"`
	Reason      string         `bson:"reason,omitempty" json:"reason,omitempty"`
	RefundedAt  time.Time      `bson:"refunded_at"    json:"refunded_at"`
	RawData     map[string]any `bson:"raw_data,omitempty" json:"raw_data,omitempty"`
}

// LicenseVerificationRequest contains parameters for validating a license directly with a provider API.
type LicenseVerificationRequest struct {
	LicenseKey         string `bson:"license_key"                   json:"license_key"`
	ProductID          string `bson:"product_id,omitempty"           json:"product_id,omitempty"`
	ProductPermalink   string `bson:"product_permalink,omitempty"   json:"product_permalink,omitempty"`
	InstanceID         string `bson:"instance_id,omitempty"         json:"instance_id,omitempty"`
	IncrementUsesCount *bool  `bson:"increment_uses_count,omitempty" json:"increment_uses_count,omitempty"`
}

// LicenseVerificationResult represents the normalized result of verifying a license key with a provider.
type LicenseVerificationResult struct {
	Valid          bool           `bson:"valid"                     json:"valid"`
	Provider       Provider       `bson:"provider"                  json:"provider"`
	LicenseKey     string         `bson:"license_key"               json:"license_key"`
	ProductID      string         `bson:"product_id,omitempty"       json:"product_id,omitempty"`
	UsesCount      int            `bson:"uses_count"                json:"uses_count"`
	MaxUses        int            `bson:"max_uses,omitempty"        json:"max_uses,omitempty"`
	Status         string         `bson:"status,omitempty"          json:"status,omitempty"`
	CustomerEmail  string         `bson:"customer_email,omitempty"  json:"customer_email,omitempty"`
	CustomerName   string         `bson:"customer_name,omitempty"   json:"customer_name,omitempty"`
	OrderID        string         `bson:"order_id,omitempty"        json:"order_id,omitempty"`
	SubscriptionID string         `bson:"subscription_id,omitempty" json:"subscription_id,omitempty"`
	ExpiresAt      *time.Time     `bson:"expires_at,omitempty"      json:"expires_at,omitempty"`
	Message        string         `bson:"message,omitempty"         json:"message,omitempty"`
	RawData        map[string]any `bson:"raw_data,omitempty"        json:"raw_data,omitempty"`
}

// WebhookSubscription represents a webhook listener endpoint registered with a provider.
type WebhookSubscription struct {
	ID        string         `bson:"id"                  json:"id"`
	Provider  Provider       `bson:"provider"            json:"provider"`
	TargetURL string         `bson:"target_url"          json:"target_url"`
	Events    []string       `bson:"events,omitempty"    json:"events,omitempty"`
	Secret    string         `bson:"secret,omitempty"    json:"secret,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	RawData   map[string]any `bson:"raw_data,omitempty"  json:"raw_data,omitempty"`
}

// CreateWebhookSubscriptionRequest contains parameters for registering a webhook listener with a provider.
type CreateWebhookSubscriptionRequest struct {
	TargetURL string   `bson:"target_url"       json:"target_url"`
	Events    []string `bson:"events,omitempty" json:"events,omitempty"`
	Secret    string   `bson:"secret,omitempty" json:"secret,omitempty"`
}

// NormalizedWebhookEvent represents a parsed, normalized webhook event payload from any provider.
type NormalizedWebhookEvent struct {
	EventID       string         `bson:"event_id"                 json:"event_id"`
	Provider      Provider       `bson:"provider"                 json:"provider"`
	EventType     EventType      `bson:"event_type"               json:"event_type"`
	OccurredAt    time.Time      `bson:"occurred_at"              json:"occurred_at"`
	SellerID      string         `bson:"seller_id,omitempty"      json:"seller_id,omitempty"`
	ProductID     string         `bson:"product_id,omitempty"     json:"product_id,omitempty"`
	CustomerEmail string         `bson:"customer_email,omitempty" json:"customer_email,omitempty"`
	LicenseKey    string         `bson:"license_key,omitempty"    json:"license_key,omitempty"`
	Order         *Order         `bson:"order,omitempty"          json:"order,omitempty"`
	Subscription  *Subscription  `bson:"subscription,omitempty"   json:"subscription,omitempty"`
	RawPayload    map[string]any `bson:"raw_payload"              json:"raw_payload"`
}
