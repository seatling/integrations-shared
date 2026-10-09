package integrations

// Provider identifies an external checkout, payment, or merchant-of-record platform.
type Provider string

const (
	// ProviderGumroad represents the Gumroad creator platform.
	ProviderGumroad Provider = "gumroad"

	// ProviderLemonSqueezy represents the Lemon Squeezy merchant-of-record platform.
	ProviderLemonSqueezy Provider = "lemon_squeezy"

	// ProviderPolar represents the Polar.sh developer monetization platform.
	ProviderPolar Provider = "polar"

	// ProviderStripe represents Stripe Checkout and billing.
	ProviderStripe Provider = "stripe"
)

// IsValid reports whether the provider is one of the supported providers.
func (p Provider) IsValid() bool {
	switch p {
	case ProviderGumroad, ProviderLemonSqueezy, ProviderPolar, ProviderStripe:
		return true
	default:
		return false
	}
}

// String returns the string representation of the provider.
func (p Provider) String() string {
	return string(p)
}

// EventType defines the normalized event type dispatched by providers.
type EventType string

const (
	// EventTypeOrderCreated represents a successful purchase, order, or sale.
	EventTypeOrderCreated EventType = "order.created"

	// EventTypeOrderRefunded represents a refund issued for an order.
	EventTypeOrderRefunded EventType = "order.refunded"

	// EventTypeOrderDisputed represents a dispute, chargeback, or inquiry opened on an order.
	EventTypeOrderDisputed EventType = "order.disputed"

	// EventTypeOrderCancelled represents an order cancellation.
	EventTypeOrderCancelled EventType = "order.cancelled"

	// EventTypeSubscriptionCreated represents a new recurring subscription activated.
	EventTypeSubscriptionCreated EventType = "subscription.created"

	// EventTypeSubscriptionUpdated represents changes to an active subscription.
	EventTypeSubscriptionUpdated EventType = "subscription.updated"

	// EventTypeSubscriptionCancelled represents a cancelled subscription.
	EventTypeSubscriptionCancelled EventType = "subscription.cancelled"

	// EventTypeSubscriptionPaused represents a paused subscription.
	EventTypeSubscriptionPaused EventType = "subscription.paused"

	// EventTypeSubscriptionResumed represents a resumed subscription.
	EventTypeSubscriptionResumed EventType = "subscription.resumed"

	// EventTypeSubscriptionExpired represents an expired subscription.
	EventTypeSubscriptionExpired EventType = "subscription.expired"

	// EventTypeLicenseKeyCreated represents a license key generated or provisioned.
	EventTypeLicenseKeyCreated EventType = "license_key.created"

	// EventTypeLicenseKeyUpdated represents a license key updated (e.g., seat count or disabled).
	EventTypeLicenseKeyUpdated EventType = "license_key.updated"
)

// SubscriptionStatus represents the lifecycle state of a recurring customer subscription.
type SubscriptionStatus string

const (
	// SubscriptionStatusActive indicates the subscription is in good standing and paid.
	SubscriptionStatusActive SubscriptionStatus = "active"

	// SubscriptionStatusPastDue indicates payment failure with grace retry pending.
	SubscriptionStatusPastDue SubscriptionStatus = "past_due"

	// SubscriptionStatusUnpaid indicates payment failed definitively and service should be suspended.
	SubscriptionStatusUnpaid SubscriptionStatus = "unpaid"

	// SubscriptionStatusCancelled indicates the buyer cancelled their subscription.
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled"

	// SubscriptionStatusExpired indicates the subscription period ended without renewal.
	SubscriptionStatusExpired SubscriptionStatus = "expired"

	// SubscriptionStatusPaused indicates billing is temporarily halted.
	SubscriptionStatusPaused SubscriptionStatus = "paused"

	// SubscriptionStatusTrialing indicates free or discounted trial period.
	SubscriptionStatusTrialing SubscriptionStatus = "trialing"
)

// IsValid reports whether the subscription status is known.
func (s SubscriptionStatus) IsValid() bool {
	switch s {
	case SubscriptionStatusActive,
		SubscriptionStatusPastDue,
		SubscriptionStatusUnpaid,
		SubscriptionStatusCancelled,
		SubscriptionStatusExpired,
		SubscriptionStatusPaused,
		SubscriptionStatusTrialing:
		return true
	default:
		return false
	}
}
