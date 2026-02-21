// Package yank provides HTTP Client functions
package yank

import (
	"net/http"
	"time"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Client struct {
	Defaults *Defaults
}

type Service interface {
	Request(request *Request, response *Response, ignoreDefaults bool) error
	GET(request *Request, response *Response, ignoreDefaults bool) error
	POST(request *Request, response *Response, ignoreDefaults bool) error
	PUT(request *Request, response *Response, ignoreDefaults bool) error
	PATCH(request *Request, response *Response, ignoreDefaults bool) error
	DELETE(request *Request, response *Response, ignoreDefaults bool) error
	Light() LightService
	GetDefaults() *Defaults
}

func New(d *Defaults) Service {
	return &Client{
		Defaults: d,
	}
}

func (c *Client) GetDefaults() *Defaults {
	return c.Defaults
}

func (c *Client) ApplyDefaultsOnRequest(request *Request, ignoreDefaults bool) error {
	if len(request.Method) == 0 {
		if ignoreDefaults {
			return ge.New("request method undefined", ge.Params{"IgnoreDefaults": ignoreDefaults})
		}

		request.Method = http.MethodGet
	}

	if len(request.URL) == 0 {
		if ignoreDefaults {
			return ge.New("request URL undefined", ge.Params{"IgnoreDefaults": ignoreDefaults})
		}

		request.URL = c.Defaults.URL
	}

	if len(request.Route) == 0 {
		if ignoreDefaults {
			return ge.New("request Route undefined", ge.Params{"IgnoreDefaults": ignoreDefaults})
		}

		request.Route = "/"
	}

	if !request.FollowRedirects {
		if !ignoreDefaults && c.Defaults.FollowRedirects {
			request.SetFollowRedirects(c.Defaults.FollowRedirects)
		}
	}

	if request.Transport == nil {
		if !ignoreDefaults && c.Defaults.Transport != nil {
			request.SetTransport(c.Defaults.Transport)
		}
	}

	return nil
}

// Request - make http request
func (c *Client) Request(request *Request, response *Response, ignoreDefaults bool) (err error) {
	totalStart := time.Now()

	prepareStart := time.Now()
	err = c.ApplyDefaultsOnRequest(request, ignoreDefaults)
	if err != nil {
		return ge.Pin(&ApplyDefaultsError{
			Parent: err,
		})
	}

	uri := request.GetURI()

	raw, err := request.GetRawContent()
	if err != nil {
		return ge.Pin(&RawContentError{
			URL:    uri,
			Method: request.Method,
			Parent: err,
		})
	}

	httpRequest, err := request.NewHTTPRequest(uri, raw, c.Defaults, ignoreDefaults)
	if err != nil {
		return ge.Pin(&CreateHTTPRequestError{
			URL:            uri,
			Method:         request.Method,
			RequestRawBody: raw.Body,
			Parent:         err,
		})
	}

	prepareDuration := time.Since(prepareStart)

	httpResponse, err := httpRequest.Execute()
	if err != nil {
		return ge.Pin(&ExecuteHTTPRequestError{
			URL:               uri,
			Method:            request.Method,
			RequestRawBody:    raw.Body,
			RequestRawHeaders: httpRequest.Header,
			Parent:            err,
		})
	}

	defer func() {
		bodyCloseErr := httpResponse.Body.Close()
		if bodyCloseErr != nil {
			err = ge.Pin(bodyCloseErr, ge.Params{ge.BaseError: err})
		}
	}()

	response.Headers = httpResponse.Header

	downloadStart := time.Now()
	err = httpResponse.ReadBody()
	if err != nil {
		return ge.Pin(&ReadBodyHTTPRequestError{
			URL:                uri,
			Method:             request.Method,
			RequestRawBody:     raw.Body,
			RequestRawHeaders:  httpRequest.Header,
			ResponseRawHeaders: httpResponse.Header,
			Parent:             err,
		})
	}
	downloadDuration := time.Since(downloadStart)

	response.Raw = httpResponse.RawBody
	response.HTTPCode = httpResponse.StatusCode

	parseStart := time.Now()
	err = httpResponse.Parse(response.Data)
	if err != nil {
		return ge.Pin(&ParseResponseError{
			URL:                uri,
			Method:             request.Method,
			RequestRawBody:     raw.Body,
			RequestRawHeaders:  httpRequest.Header,
			ResponseRawBody:    string(httpResponse.RawBody),
			ResponseRawHeaders: httpResponse.Header,
			Parent:             err,
		})
	}
	parseDuration := time.Since(parseStart)

	response.Timing = &Timing{
		PrepareRequest:  prepareDuration,
		Connect:         httpResponse.Timing.Connect,
		SendRequest:     httpResponse.Timing.SendRequest,
		TimeToFirstByte: httpResponse.Timing.TimeToFirstByte,
		DownloadContent: downloadDuration,
		ParseResponse:   parseDuration,
		Total:           time.Since(totalStart),
	}

	if httpResponse.StatusCode != response.ExpectedHTTPCode {
		return ge.Pin(&WrongHTTPCodeError{
			URL:                uri,
			Method:             request.Method,
			RequestRawBody:     raw.Body,
			RequestRawHeaders:  httpRequest.Header,
			ResponseRawBody:    string(httpResponse.RawBody),
			ResponseRawHeaders: httpResponse.Header,
			Expected:           response.ExpectedHTTPCode,
			Actual:             httpResponse.StatusCode,
		})
	}

	return nil
}

func (c *Client) POST(request *Request, response *Response, ignoreDefaults bool) error {
	request.Method = http.MethodPost

	return c.Request(request, response, ignoreDefaults)
}

func (c *Client) PUT(request *Request, response *Response, ignoreDefaults bool) error {
	request.Method = http.MethodPut

	return c.Request(request, response, ignoreDefaults)
}

func (c *Client) PATCH(request *Request, response *Response, ignoreDefaults bool) error {
	request.Method = http.MethodPatch

	return c.Request(request, response, ignoreDefaults)
}

func (c *Client) DELETE(request *Request, response *Response, ignoreDefaults bool) error {
	request.Method = http.MethodDelete

	return c.Request(request, response, ignoreDefaults)
}

func (c *Client) GET(request *Request, response *Response, ignoreDefaults bool) error {
	request.Method = http.MethodGet

	return c.Request(request, response, ignoreDefaults)
}
