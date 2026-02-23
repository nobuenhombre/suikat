package yank

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptrace"
	"reflect"
	"time"

	"github.com/nobuenhombre/suikat/pkg/adapt"

	"github.com/nobuenhombre/suikat/pkg/ge"
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

	RawBody []byte
	Timing  *Timing
}

func (r *HTTPRequest) Execute() (httpResponse *HTTPResponse, err error) {
	var (
		startTime    = time.Now()
		connectStart time.Time
		connectDone  time.Time
		wroteRequest time.Time
		gotFirstByte time.Time
	)

	trace := &httptrace.ClientTrace{
		ConnectStart: func(network, addr string) {
			connectStart = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			connectDone = time.Now()
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			wroteRequest = time.Now()
		},
		GotFirstResponseByte: func() {
			gotFirstByte = time.Now()
		},
	}

	ctx := httptrace.WithClientTrace(context.Background(), trace)
	req := r.Request.WithContext(ctx)

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
	response, err := client.Do(req)
	if err != nil {
		return nil, ge.Pin(err)
	}

	var connect, sendRequest time.Duration
	if !connectStart.IsZero() && !connectDone.IsZero() {
		connect = connectDone.Sub(connectStart)
		sendRequest = wroteRequest.Sub(connectDone)
	} else {
		// Если соединение переиспользовано, считаем от начала запроса
		connect = 0
		if !wroteRequest.IsZero() {
			sendRequest = wroteRequest.Sub(startTime)
		}
	}

	// Считаем времена
	timing := &Timing{
		Connect:         connect,
		SendRequest:     sendRequest,
		TimeToFirstByte: gotFirstByte.Sub(wroteRequest),
	}

	return &HTTPResponse{Response: response, Timing: timing}, nil
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
