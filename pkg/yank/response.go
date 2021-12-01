package yank

import (
	"net/http"

	"github.com/nobuenhombre/suikat/pkg/tracktime"
)

type Response struct {
	ExpectedHTTPCode int

	HTTPCode int
	Headers  http.Header
	Data     interface{}
	Timer    *tracktime.Tracker
}

func NewResponse(data interface{}, expectedHTTPCode int) *Response {
	return &Response{
		Data:             data,
		ExpectedHTTPCode: expectedHTTPCode,
	}
}
