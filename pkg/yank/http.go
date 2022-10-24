package yank

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/adapt"

	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/nobuenhombre/suikat/pkg/tracktime"
)

type HTTPHeaders interface {
	AddHeaders(r *HTTPRequest)
}

type HTTPRequest struct {
	*http.Request
	FollowRedirects bool
	Transport       http.RoundTripper
}

func (r *HTTPRequest) AddHeaders(headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
}

type HTTPResponse struct {
	*http.Response
	Timer *tracktime.Tracker

	RawBody []byte
}

func (r *HTTPRequest) Execute() (httpResponse *HTTPResponse, err error) {
	timer := tracktime.Start("HTTPRequest.Execute()")

	client := &http.Client{}

	if !r.FollowRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if !adapt.IsNil(r.Transport) {
		client.Transport = r.Transport
	}

	//nolint:bodyclose
	response, err := client.Do(r.Request)
	if err != nil {
		return nil, ge.Pin(err)
	}

	timer.Stop()

	return &HTTPResponse{Response: response, Timer: timer}, nil
}

func (rs *HTTPResponse) ReadBody() error {
	// parse Response JSON body
	//-------------------------
	respBody, err := io.ReadAll(rs.Body)
	if err != nil {
		return ge.Pin(err)
	}

	rs.RawBody = respBody

	return nil
}

func (rs *HTTPResponse) Parse(data interface{}) error {
	switch data.(type) {
	case nil:
		// nothing need to be parsed

	case *string:
		rv := reflect.ValueOf(data).Elem()
		rv.SetString(string(rs.RawBody))

	case http.Header:
		// write Response Headers to data
		//-------------------------------
		rh := reflect.ValueOf(rs.Header)
		rv := reflect.ValueOf(data)

		if rv.Kind() == reflect.Map {
			for _, k := range rh.MapKeys() {
				rv.SetMapIndex(k, rh.MapIndex(k))
			}
		}

	default:
		err := json.Unmarshal(rs.RawBody, data)
		if err != nil {
			return ge.Pin(err)
		}
	}

	return nil
}
