package integrations

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrUnauthorized is returned when provider credentials or access tokens are invalid or expired.
	ErrUnauthorized = errors.New("integrations: unauthorized or invalid credentials")

	// ErrNotFound is returned when the requested upstream resource does not exist.
	ErrNotFound = errors.New("integrations: resource not found")

	// ErrRateLimited is returned when upstream provider returns HTTP 429 Too Many Requests.
	ErrRateLimited = errors.New("integrations: rate limit exceeded")

	// ErrMissingAccessToken is returned when an API call requiring an access token receives an empty token.
	ErrMissingAccessToken = errors.New("integrations: access token is required")

	// ErrInvalidParameter is returned when an input argument fails validation constraints.
	ErrInvalidParameter = errors.New("integrations: invalid parameter")

	// ErrWebhookVerificationFailed is returned when webhook signatures, tokens, or passcodes do not match.
	ErrWebhookVerificationFailed = errors.New("integrations: webhook verification failed")

	// ErrInvalidPayload is returned when a webhook body is empty, malformed, or cannot be parsed.
	ErrInvalidPayload = errors.New("integrations: invalid webhook payload")

	// ErrUnsupportedEventType is returned when a provider webhook event is unrecognized or unsupported.
	ErrUnsupportedEventType = errors.New("integrations: unsupported event type")

	// ErrResourceConflict is returned when an upstream entity already exists or conflicts (HTTP 409).
	ErrResourceConflict = errors.New("integrations: resource conflict")
)

// APIError represents an upstream error response returned by a provider REST API.
type APIError struct {
	Provider   Provider `json:"provider,omitempty"`
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	RawBody    string   `json:"raw_body,omitempty"`
}

// Error returns a formatted error string.
func (e *APIError) Error() string {
	if e.Provider != "" {
		return fmt.Sprintf("%s api error (status %d): %s", e.Provider, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("provider api error (status %d): %s", e.StatusCode, e.Message)
}

// Is reports whether the error matches sentinel errors based on its HTTP status code.
func (e *APIError) Is(target error) bool {
	switch {
	case errors.Is(target, ErrUnauthorized):
		return e.StatusCode == http.StatusUnauthorized || e.StatusCode == http.StatusForbidden
	case errors.Is(target, ErrNotFound):
		return e.StatusCode == http.StatusNotFound
	case errors.Is(target, ErrRateLimited):
		return e.StatusCode == http.StatusTooManyRequests
	case errors.Is(target, ErrResourceConflict):
		return e.StatusCode == http.StatusConflict
	default:
		return false
	}
}

// Retryable reports whether the failure is transient (429 Too Many Requests or 5xx Server Error).
func (e *APIError) Retryable() bool {
	return e.StatusCode == http.StatusTooManyRequests || (e.StatusCode >= 500 && e.StatusCode <= 599)
}

// NewAPIError constructs a new APIError.
func NewAPIError(provider Provider, statusCode int, message string, rawBody string) *APIError {
	return &APIError{
		Provider:   provider,
		StatusCode: statusCode,
		Message:    message,
		RawBody:    rawBody,
	}
}
