package yank

import (
	"net/http"
)

type Response struct {
	ExpectedHTTPCode int

	HTTPCode int
	Headers  http.Header
	Data     interface{}
}
