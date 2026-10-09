package integrations_test

import (
	"errors"
	"net/http"
	"testing"

	integrations "github.com/seatling/integrations-shared"
)

func TestAPIError_Error(t *testing.T) {
	err1 := integrations.NewAPIError(integrations.ProviderGumroad, 401, "invalid access token", `{"error": "unauthorized"}`)
	expected1 := "gumroad api error (status 401): invalid access token"
	if err1.Error() != expected1 {
		t.Errorf("expected %q, got %q", expected1, err1.Error())
	}

	err2 := integrations.NewAPIError("", 500, "internal server error", "")
	expected2 := "provider api error (status 500): internal server error"
	if err2.Error() != expected2 {
		t.Errorf("expected %q, got %q", expected2, err2.Error())
	}
}

func TestAPIError_Is(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		target     error
		matches    bool
	}{
		{"401 is unauthorized", http.StatusUnauthorized, integrations.ErrUnauthorized, true},
		{"403 is unauthorized", http.StatusForbidden, integrations.ErrUnauthorized, true},
		{"404 is not found", http.StatusNotFound, integrations.ErrNotFound, true},
		{"429 is rate limited", http.StatusTooManyRequests, integrations.ErrRateLimited, true},
		{"409 is conflict", http.StatusConflict, integrations.ErrResourceConflict, true},
		{"500 is not unauthorized", http.StatusInternalServerError, integrations.ErrUnauthorized, false},
		{"200 is not not found", http.StatusOK, integrations.ErrNotFound, false},
		{"unrelated error", http.StatusBadRequest, errors.New("other"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := integrations.NewAPIError(integrations.ProviderLemonSqueezy, tt.statusCode, "test message", "")
			if got := errors.Is(apiErr, tt.target); got != tt.matches {
				t.Errorf("errors.Is(apiErr, %v) = %v, want %v", tt.target, got, tt.matches)
			}
		})
	}
}

func TestAPIError_Retryable(t *testing.T) {
	tests := []struct {
		statusCode int
		retryable  bool
	}{
		{http.StatusOK, false},
		{http.StatusBadRequest, false},
		{http.StatusUnauthorized, false},
		{http.StatusNotFound, false},
		{http.StatusTooManyRequests, true},
		{http.StatusInternalServerError, true},
		{http.StatusBadGateway, true},
		{http.StatusServiceUnavailable, true},
		{http.StatusGatewayTimeout, true},
	}

	for _, tt := range tests {
		apiErr := integrations.NewAPIError(integrations.ProviderStripe, tt.statusCode, "msg", "")
		if got := apiErr.Retryable(); got != tt.retryable {
			t.Errorf("APIError{StatusCode: %d}.Retryable() = %v, want %v", tt.statusCode, got, tt.retryable)
		}
	}
}
