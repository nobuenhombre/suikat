package yank

import (
	"net/http"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	var client Service

	// Create client service
	func() {
		defaults := NewDefaults("https://www.siarion.net")
		defaults.AuthConstructor().AuthBasic("user", "password")

		client = New(defaults)
	}()

	// Do request
	func() {
		request := NewRequest("/rus/")
		request.AuthConstructor().AuthNone()
		request.QueryConstructor().AddQuery("hello", "world")
		request.HeaderConstructor().AddHeader("wake-up", "neo")
		request.BodyConstructor().SetBody(nil)

		result := new(string)

		response := NewResponse(result, http.StatusOK)

		err := client.GET(request, response, false)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}()
}
