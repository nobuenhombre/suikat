package yank

import (
	"net/http"

	"github.com/nobuenhombre/suikat/pkg/adapt"
	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/nobuenhombre/suikat/pkg/inslice"
	"github.com/nobuenhombre/suikat/pkg/mimes"
)

type LightClient struct {
	*Client
}

type LightService interface {
	Request(
		method,
		route string,
		send interface{},
		receiver interface{},
		expectedStatusCode int,
		mimeType string,
	) (statusCode int, rawBody []byte, err error)

	GET(
		route string,
		receiver interface{},
		expectedStatusCode int,
	) (statusCode int, rawBody []byte, err error)

	POST(
		route string,
		send interface{},
		receiver interface{},
		expectedStatusCode int,
		mimeType string,
	) (statusCode int, rawBody []byte, err error)

	PUT(
		route string,
		send interface{},
		receiver interface{},
		expectedStatusCode int,
		mimeType string,
	) (statusCode int, rawBody []byte, err error)

	PATCH(
		route string,
		send interface{},
		receiver interface{},
		expectedStatusCode int,
		mimeType string,
	) (statusCode int, rawBody []byte, err error)

	DELETE(
		route string,
		send interface{},
		receiver interface{},
		expectedStatusCode int,
		mimeType string,
	) (statusCode int, rawBody []byte, err error)
}

func (c *Client) Light() LightService {
	return &LightClient{
		Client: c,
	}
}

func (lc *LightClient) Request(
	method,
	route string,
	send interface{},
	receiver interface{},
	expectedStatusCode int,
	mimeType string,
) (statusCode int, rawBody []byte, err error) {
	request := NewRequest(route)
	request.Method = method

	methods := []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}
	if !adapt.IsNil(send) && inslice.String(method, &methods) {
		switch mimeType {
		case mimes.JSON:
			request.BodyConstructor().SetBody(NewJSON(send))
		case mimes.FormUrlencoded:
			request.BodyConstructor().SetBody(NewFormURLEncoded(send))
		case mimes.FormMultipartData:
			request.BodyConstructor().SetBody(NewFormMultipart(send))
		default:
			return 0, []byte{}, ge.Pin(&ge.UndefinedSwitchCaseError{
				Var: mimeType,
			})
		}
	}

	response := NewResponse(receiver, expectedStatusCode)

	err = lc.Client.Request(request, response, false)
	statusCode = response.HTTPCode
	rawBody = response.Raw

	return
}

func (lc *LightClient) GET(
	route string,
	receiver interface{},
	expectedStatusCode int,
) (statusCode int, rawBody []byte, err error) {
	return lc.Request(http.MethodGet, route, nil, receiver, expectedStatusCode, "")
}

func (lc *LightClient) POST(
	route string,
	send interface{},
	receiver interface{},
	expectedStatusCode int,
	mimeType string,
) (statusCode int, rawBody []byte, err error) {
	return lc.Request(http.MethodPost, route, send, receiver, expectedStatusCode, mimeType)
}

func (lc *LightClient) PUT(
	route string,
	send interface{},
	receiver interface{},
	expectedStatusCode int,
	mimeType string,
) (statusCode int, rawBody []byte, err error) {
	return lc.Request(http.MethodPut, route, send, receiver, expectedStatusCode, mimeType)
}

func (lc *LightClient) PATCH(
	route string,
	send interface{},
	receiver interface{},
	expectedStatusCode int,
	mimeType string,
) (statusCode int, rawBody []byte, err error) {
	return lc.Request(http.MethodPatch, route, send, receiver, expectedStatusCode, mimeType)
}

func (lc *LightClient) DELETE(
	route string,
	send interface{},
	receiver interface{},
	expectedStatusCode int,
	mimeType string,
) (statusCode int, rawBody []byte, err error) {
	return lc.Request(http.MethodDelete, route, send, receiver, expectedStatusCode, mimeType)
}
