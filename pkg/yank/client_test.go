package yank

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPRequest(t *testing.T) {
	var client Service

	// Create client service
	func() {
		defaults := NewDefaults("https://nextjs.org")
		// defaults.AuthConstructor().AuthBasic("user", "password")

		client = New(defaults)
	}()

	// Do request
	func() {
		request := NewRequest("/")
		request.AuthConstructor().AuthNone()
		request.QueryConstructor().AddQuery("hello", "world")
		request.HeaderConstructor().AddHeader("wake-up", "neo")
		request.BodyConstructor().SetBody(nil)

		result := new(string)

		response := NewResponse(result, http.StatusOK)

		err := client.GET(request, response, false)
		require.NoError(t, err)

		t.Logf("%#v", response.Timing)
	}()
}
