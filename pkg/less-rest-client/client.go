package lessrestclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/mimes"

	"github.com/gorilla/schema"
)

const (
	AuthTypeNone = iota
	AuthTypeLoginPass
	AuthTypeToken
	AuthTypeBearerToken
	AuthTypeCookieWithCSRFToken
)

type Client struct {
	URL         string
	ContentType string
	AuthType    int
	Username    string
	Password    string
	Token       string
	BearerToken string
	Cookie      string
	XCSRFToken  string
}

func New(c *Client) LRC {
	return c
}

func (c *Client) Request(
	method, route string,
	inData interface{},
	outData interface{},
	expectedStatusCode int,
) (statusCode int, respBody []byte, err error) {
	var (
		reqBody    io.Reader
		reqRawBody string
	)

	URL := c.URL + route

	switch inData.(type) {
	case nil:
		reqBody = nil
		reqRawBody = "nil"
	default:
		switch c.ContentType {
		case mimes.FormUrlencoded:
			var encoder = schema.NewEncoder()

			q := url.Values{}

			err = encoder.Encode(inData, q)
			if err != nil {
				err = &EncoderError{
					URL:    URL,
					Method: method,
					Data:   inData,
					Parent: err,
				}

				return
			}

			reqRawBody = q.Encode()
			reqBody = strings.NewReader(reqRawBody)
		case mimes.JSON:
			var tmp []byte

			tmp, err = json.Marshal(inData)
			if err != nil {
				err = &MarshalingError{
					URL:    URL,
					Method: method,
					Data:   inData,
					Parent: err,
				}

				return
			}

			reqBody = bytes.NewBuffer(tmp)
			reqRawBody = string(tmp)
		default:
			err = &UnknownContentTypeError{
				URL:    URL,
				Method: method,
				Data:   inData,
			}

			return
		}
	}

	req, err := http.NewRequest(method, URL, reqBody)
	if err != nil {
		err = &CreateRequestError{
			URL:     URL,
			Method:  method,
			RawBody: reqRawBody,
			Parent:  err,
		}

		return
	}

	req.Header.Add("Content-Type", c.ContentType)

	switch c.AuthType {
	case AuthTypeNone:
		// AuthTypeNone
	case AuthTypeLoginPass:
		req.SetBasicAuth(c.Username, c.Password)
	case AuthTypeToken:
		req.Header.Add("Authorization", c.Token)
	case AuthTypeBearerToken:
		req.Header.Add("Authorization", "Bearer "+c.BearerToken)
	case AuthTypeCookieWithCSRFToken:
		req.Header.Add("Cookie", c.Cookie)
		req.Header.Add("X-CSRF-Token", c.XCSRFToken)
	default:
		err = &UnknownAuthTypeError{
			URL:    URL,
			Method: method,
			Data:   inData,
		}

		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		err = &ClientError{
			URL:     URL,
			Method:  method,
			RawBody: reqRawBody,
			Parent:  err,
		}

		return
	}

	defer resp.Body.Close()

	statusCode = resp.StatusCode

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = &ReadResponseBodyError{
			URL:     URL,
			Method:  method,
			RawBody: reqRawBody,
			Parent:  err,
		}

		return
	}

	respRawBody := string(respBody)

	if statusCode != expectedStatusCode {
		err = &WrongStatusCodeError{
			URL:             URL,
			Method:          method,
			RequestRawBody:  reqRawBody,
			ResponseRawBody: respRawBody,
			Expected:        expectedStatusCode,
			Actual:          statusCode,
		}

		return
	}

	switch outData.(type) {
	case nil:
	default:
		err = json.Unmarshal(respBody, outData)
		if err != nil {
			err = &UnMarshalingError{
				URL:     URL,
				Method:  method,
				RawBody: respRawBody,
				Parent:  err,
			}

			return
		}
	}

	return
}

func (c *Client) POST(
	route string,
	inData interface{},
	outData interface{},
	expectedStatusCode int,
) (statusCode int, respBody []byte, err error) {
	return c.Request(http.MethodPost, route, inData, outData, expectedStatusCode)
}

func (c *Client) PUT(
	route string,
	inData interface{},
	outData interface{},
	expectedStatusCode int,
) (statusCode int, respBody []byte, err error) {
	return c.Request(http.MethodPut, route, inData, outData, expectedStatusCode)
}

func (c *Client) PATCH(
	route string,
	inData interface{},
	outData interface{},
	expectedStatusCode int,
) (statusCode int, respBody []byte, err error) {
	return c.Request(http.MethodPatch, route, inData, outData, expectedStatusCode)
}

func (c *Client) DELETE(
	route string,
	inData interface{},
	outData interface{},
	expectedStatusCode int,
) (statusCode int, respBody []byte, err error) {
	return c.Request(http.MethodDelete, route, inData, outData, expectedStatusCode)
}

func (c *Client) GET(
	route string,
	outData interface{},
	expectedStatusCode int,
) (statusCode int, respBody []byte, err error) {
	return c.Request(http.MethodGet, route, nil, outData, expectedStatusCode)
}
