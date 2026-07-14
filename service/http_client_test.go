package service

import (
	"net/http"
	"testing"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetTextHttpClient exercises GetTextHttpClient's self-healing cache directly against the
// package-level httpClient/common.TextRelayTimeout state (same pattern as
// TestGetSSRFProtectedHTTPClientFallsBackToDefaultClientWhenProtectionDisabled in
// protected_fetch_client_test.go). Not run in parallel: it mutates shared package state.
func TestGetTextHttpClient(t *testing.T) {
	originalHTTPClient := httpClient
	originalTimeout := common.TextRelayTimeout
	t.Cleanup(func() {
		httpClient = originalHTTPClient
		common.TextRelayTimeout = originalTimeout
	})

	httpClient = &http.Client{Transport: &http.Transport{}}

	common.TextRelayTimeout = 0
	require.Same(t, httpClient, GetTextHttpClient(), "disabled (<=0) must fall back to the plain client")

	common.TextRelayTimeout = 30
	client := GetTextHttpClient()
	assert.Same(t, httpClient.Transport, client.Transport, "must reuse the shared transport/connection pool")
	assert.Equal(t, 30*time.Second, client.Timeout)

	same := GetTextHttpClient()
	assert.Same(t, client, same, "must cache when the configured value hasn't changed")

	common.TextRelayTimeout = 60
	rebuilt := GetTextHttpClient()
	assert.NotSame(t, client, rebuilt, "must rebuild when the configured value changes")
	assert.Equal(t, 60*time.Second, rebuilt.Timeout)
}
