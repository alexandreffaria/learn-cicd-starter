package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := make(http.Header)

	_, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("Expected 'ErrNoAuthHeaderIncluded', got '%v'", err)
	}
}

func TestGetAPIKey_HeaderWithoutApiKeyPrefix(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "Notvalidkey")

	_, err := GetAPIKey(headers)

	if err == nil || err.Error() != "malformed authorization gheader" {
		t.Errorf("Expected 'malformed authorization header' error, got '%v'", err)
	}
}
