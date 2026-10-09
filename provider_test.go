package integrations_test

import (
	"testing"

	integrations "github.com/seatling/integrations-shared"
)

func TestProvider_IsValid(t *testing.T) {
	tests := []struct {
		provider integrations.Provider
		valid    bool
	}{
		{integrations.ProviderGumroad, true},
		{integrations.ProviderLemonSqueezy, true},
		{integrations.ProviderPolar, true},
		{integrations.ProviderStripe, true},
		{integrations.Provider("unknown"), false},
		{integrations.Provider(""), false},
	}

	for _, tt := range tests {
		if got := tt.provider.IsValid(); got != tt.valid {
			t.Errorf("Provider(%q).IsValid() = %v, want %v", tt.provider, got, tt.valid)
		}
	}
}

func TestProvider_String(t *testing.T) {
	p := integrations.ProviderGumroad
	if p.String() != "gumroad" {
		t.Errorf("expected 'gumroad', got %q", p.String())
	}
}

func TestSubscriptionStatus_IsValid(t *testing.T) {
	validStatuses := []integrations.SubscriptionStatus{
		integrations.SubscriptionStatusActive,
		integrations.SubscriptionStatusPastDue,
		integrations.SubscriptionStatusUnpaid,
		integrations.SubscriptionStatusCancelled,
		integrations.SubscriptionStatusExpired,
		integrations.SubscriptionStatusPaused,
		integrations.SubscriptionStatusTrialing,
	}

	for _, status := range validStatuses {
		if !status.IsValid() {
			t.Errorf("SubscriptionStatus(%q).IsValid() = false, want true", status)
		}
	}

	invalidStatuses := []integrations.SubscriptionStatus{
		integrations.SubscriptionStatus("unknown"),
		integrations.SubscriptionStatus(""),
	}

	for _, status := range invalidStatuses {
		if status.IsValid() {
			t.Errorf("SubscriptionStatus(%q).IsValid() = true, want false", status)
		}
	}
}
