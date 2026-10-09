package integrations_test

import (
	"encoding/json"
	"testing"
	"time"

	integrations "github.com/seatling/integrations-shared"
)

func TestModels_JSONSerialization(t *testing.T) {
	now := time.Date(2026, 10, 9, 12, 0, 0, 0, time.UTC)

	t.Run("User JSON roundtrip", func(t *testing.T) {
		u := integrations.User{
			ID:                "usr_123",
			Provider:          integrations.ProviderGumroad,
			Name:              "Jane Doe",
			Email:             "jane@example.com",
			ProfileURL:        "https://gumroad.com/janedoe",
			Currency:          "USD",
			CustomDomain:      "https://example.com",
			SubscribedToEmail: true,
			RawData:           map[string]any{"tier": "pro"},
		}

		data, err := json.Marshal(u)
		if err != nil {
			t.Fatalf("marshal user: %v", err)
		}

		var decoded integrations.User
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal user: %v", err)
		}

		if decoded.ID != u.ID || decoded.Email != u.Email || decoded.Provider != u.Provider {
			t.Errorf("decoded user mismatch: %+v", decoded)
		}
	})

	t.Run("Product JSON roundtrip", func(t *testing.T) {
		p := integrations.Product{
			ID:             "prod_456",
			Provider:       integrations.ProviderLemonSqueezy,
			Name:           "CodeCraft Pro",
			Description:    "Developer editor",
			URL:            "https://store.lemonsqueezy.com/buy/123",
			PriceCents:     4900,
			Currency:       "USD",
			FormattedPrice: "$49.00",
			Published:      true,
			Variants: []integrations.Variant{
				{
					ID:             "var_789",
					ProductID:      "prod_456",
					Name:           "Studio Edition",
					PriceCents:     9900,
					Currency:       "USD",
					FormattedPrice: "$99.00",
					IsDefault:      false,
				},
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		}

		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("marshal product: %v", err)
		}

		var decoded integrations.Product
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal product: %v", err)
		}

		if decoded.ID != p.ID || len(decoded.Variants) != 1 || decoded.Variants[0].PriceCents != 9900 {
			t.Errorf("decoded product mismatch: %+v", decoded)
		}
	})

	t.Run("Order JSON roundtrip", func(t *testing.T) {
		ord := integrations.Order{
			ID:          "ord_111",
			Provider:    integrations.ProviderPolar,
			ProductID:   "prod_456",
			ProductName: "CodeCraft Pro",
			Customer: integrations.Customer{
				ID:      "cus_999",
				Email:   "buyer@company.org",
				Name:    "Buyer Name",
				Country: "US",
			},
			LicenseKey:          "ABCD-EFGH-IJKL-MNOP",
			PriceCents:          4900,
			Currency:            "USD",
			Refunded:            false,
			Disputed:            false,
			OrderDate:           now,
			CustomFields:        map[string]string{"seat_count": "5"},
			RawData:             map[string]any{"source": "polar"},
		}

		data, err := json.Marshal(ord)
		if err != nil {
			t.Fatalf("marshal order: %v", err)
		}

		var decoded integrations.Order
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal order: %v", err)
		}

		if decoded.ID != ord.ID || decoded.LicenseKey != ord.LicenseKey || decoded.Customer.Email != ord.Customer.Email {
			t.Errorf("decoded order mismatch: %+v", decoded)
		}
	})

	t.Run("Subscription JSON roundtrip", func(t *testing.T) {
		sub := integrations.Subscription{
			ID:                 "sub_333",
			Provider:           integrations.ProviderStripe,
			ProductID:          "prod_456",
			CustomerID:         "cus_999",
			CustomerEmail:      "buyer@company.org",
			Status:             integrations.SubscriptionStatusActive,
			LicenseKey:         "ABCD-EFGH-IJKL-MNOP",
			CurrentPeriodStart: &now,
			CurrentPeriodEnd:   &now,
			CreatedAt:          now,
		}

		data, err := json.Marshal(sub)
		if err != nil {
			t.Fatalf("marshal subscription: %v", err)
		}

		var decoded integrations.Subscription
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal subscription: %v", err)
		}

		if decoded.ID != sub.ID || decoded.Status != integrations.SubscriptionStatusActive {
			t.Errorf("decoded subscription mismatch: %+v", decoded)
		}
	})

	t.Run("NormalizedWebhookEvent JSON roundtrip", func(t *testing.T) {
		evt := integrations.NormalizedWebhookEvent{
			EventID:       "evt_0123456789",
			Provider:      integrations.ProviderGumroad,
			EventType:     integrations.EventTypeOrderCreated,
			OccurredAt:    now,
			SellerID:      "seller_abc",
			ProductID:     "prod_456",
			CustomerEmail: "buyer@example.com",
			LicenseKey:    "LIC-KEY-999",
			RawPayload:    map[string]any{"seller_id": "seller_abc"},
		}

		data, err := json.Marshal(evt)
		if err != nil {
			t.Fatalf("marshal event: %v", err)
		}

		var decoded integrations.NormalizedWebhookEvent
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal event: %v", err)
		}

		if decoded.EventID != evt.EventID || decoded.EventType != evt.EventType {
			t.Errorf("decoded event mismatch: %+v", decoded)
		}
	})

	t.Run("LicenseVerificationResult JSON roundtrip", func(t *testing.T) {
		res := integrations.LicenseVerificationResult{
			Valid:         true,
			Provider:      integrations.ProviderGumroad,
			LicenseKey:    "LIC-KEY-999",
			ProductID:     "prod_456",
			UsesCount:     2,
			MaxUses:       5,
			CustomerEmail: "buyer@example.com",
		}

		data, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("marshal verification: %v", err)
		}

		var decoded integrations.LicenseVerificationResult
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal verification: %v", err)
		}

		if !decoded.Valid || decoded.UsesCount != 2 || decoded.MaxUses != 5 {
			t.Errorf("decoded verification mismatch: %+v", decoded)
		}
	})

	t.Run("WebhookSubscription JSON roundtrip", func(t *testing.T) {
		sub := integrations.WebhookSubscription{
			ID:        "sub_123",
			Provider:  integrations.ProviderLemonSqueezy,
			TargetURL: "https://example.com/webhook",
			Events:    []string{"order_created"},
			Secret:    "secret_key_123",
		}

		data, err := json.Marshal(sub)
		if err != nil {
			t.Fatalf("marshal webhook subscription: %v", err)
		}

		var decoded integrations.WebhookSubscription
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal webhook subscription: %v", err)
		}

		if decoded.ID != sub.ID || decoded.Secret != sub.Secret {
			t.Errorf("decoded webhook subscription mismatch: %+v", decoded)
		}
	})
}
