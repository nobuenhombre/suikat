package yank

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLightClient(t *testing.T) {
	var client Service

	// Create client service
	func() {
		defaults := NewDefaults("https://nextjs.org")
		// defaults.AuthConstructor().AuthBasic("user", "password")

		client = New(defaults)
	}()

	// Do request
	func() {
		result := new(string)

		status, rawBody, err := client.Light().GET("/", result, http.StatusOK)
		require.NoError(t, err)

		t.Logf("status: %v, raw: %v", status, rawBody)
	}()
}
