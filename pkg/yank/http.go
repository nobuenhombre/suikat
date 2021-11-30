package yank

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type HTTPHeaders interface {
	AddHeaders(r *HTTPRequest)
}

type HTTPRequest struct {
	*http.Request
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

	RawBody []byte
}

func (r *HTTPRequest) Execute() (httpResponse *HTTPResponse, err error) {
	client := &http.Client{}

	//nolint:bodyclose
	response, err := client.Do(r.Request)
	if err != nil {
		return nil, ge.Pin(err)
	}

	return &HTTPResponse{Response: response}, nil
}

func (rs *HTTPResponse) ReadBody() error {
	// parse Response JSON body
	//-------------------------
	respBody, err := ioutil.ReadAll(rs.Body)
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
