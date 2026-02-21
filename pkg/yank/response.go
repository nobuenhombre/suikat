package yank

import (
	"net/http"
)

type Response struct {
	ExpectedHTTPCode int

	HTTPCode int
	Headers  http.Header
	Raw      []byte
	Data     interface{}
	Timing   *Timing
}

func NewResponse(data interface{}, expectedHTTPCode int) *Response {
	return &Response{
		Data:             data,
		ExpectedHTTPCode: expectedHTTPCode,
	}
}
