package yank

import (
	"net/http"
	"testing"
)

func TestLightClient(t *testing.T) {
	var client Service

	// Create client service
	func() {
		defaults := NewDefaults("https://www.siarion.net")
		defaults.AuthConstructor().AuthBasic("user", "password")

		client = New(defaults)
	}()

	// Do request
	func() {
		result := new(string)

		status, rawBody, err := client.Light().GET("/rus/", result, http.StatusOK)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		t.Logf("status: %v, raw: %v", status, rawBody)
	}()
}
