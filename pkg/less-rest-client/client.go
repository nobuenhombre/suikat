package lessrestclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	AuthTypeEDS
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
	SignKeyID   string
	SignBody    string
}

func New(c *Client) LRC {
	return c
}

func (c *Client) GetBodyFormMultipartData(
	urlAddr, method string,
	inData interface{},
	addHeader *map[string]string,
) (reqBody io.Reader, reqRawBody string, err error) {
	reqBody = &bytes.Buffer{}

	w := multipart.NewWriter(reqBody.(io.Writer))

	for k, v := range inData.(url.Values) {
		for _, iv := range v {
			if strings.HasPrefix(k, "@") { // file
				err = addFile(w, k[1:], iv)
				if err != nil {
					return
				}
			} else { // form value
				err = w.WriteField(k, iv)
				if err != nil {
					return
				}
			}
		}
	}

	// Set the Boundary in the Content-Type
	(*addHeader)["Content-Type"] = w.FormDataContentType()

	// Set Content-Length
	(*addHeader)["Content-Length"] = fmt.Sprintf("%d", reqBody.(*bytes.Buffer).Len())

	err = w.Close()
	if err != nil {
		return
	}

	reqRawBody = reqBody.(*bytes.Buffer).String()

	return
}

func (c *Client) GetBodyFormUrlencoded(
	urlAddr, method string,
	inData interface{},
	addHeader *map[string]string,
) (reqBody io.Reader, reqRawBody string, err error) {
	var (
		q  url.Values
		ok bool
	)

	q, ok = inData.(url.Values)
	if !ok {
		var encoder = schema.NewEncoder()

		q = url.Values{}

		err = encoder.Encode(inData, q)
		if err != nil {
			err = &EncoderError{
				URL:    urlAddr,
				Method: method,
				Data:   inData,
				Parent: err,
			}

			return
		}
	}

	reqRawBody = q.Encode()
	reqBody = strings.NewReader(reqRawBody)

	return
}

func (c *Client) GetBodyJSON(
	urlAddr, method string,
	inData interface{},
	addHeader *map[string]string,
) (reqBody io.Reader, reqRawBody string, err error) {
	var tmp []byte

	tmp, err = json.Marshal(inData)
	if err != nil {
		err = &MarshalingError{
			URL:    urlAddr,
			Method: method,
			Data:   inData,
			Parent: err,
		}

		return
	}

	reqBody = bytes.NewBuffer(tmp)
	reqRawBody = string(tmp)

	return
}

func (c *Client) ErrorRawBody(contentType, rawBody string) string {
	switch c.ContentType {
	case mimes.FormMultipartData:
		return fmt.Sprintf("FormMultipartData(len=%v)", len(rawBody))
	case mimes.FormUrlencoded:
		return fmt.Sprintf("FormUrlencoded(%v)", rawBody)
	case mimes.JSON:
		return fmt.Sprintf("JSON(%v)", rawBody)
	default:
		return fmt.Sprintf("Unknown(len=%v)", len(rawBody))
	}
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

	addHeader := make(map[string]string)

	URL := c.URL + route

	switch inData.(type) {
	case nil:
		reqBody = nil
		reqRawBody = "nil"
	default:
		switch c.ContentType {
		case mimes.FormMultipartData:
			reqBody, reqRawBody, err = c.GetBodyFormMultipartData(URL, method, inData, &addHeader)
			if err != nil {
				err = &BodyFormMultipartDataError{
					FormDataError{
						URL:    URL,
						Method: method,
						Parent: err,
					},
				}

				return
			}

		case mimes.FormUrlencoded:
			reqBody, reqRawBody, err = c.GetBodyFormUrlencoded(URL, method, inData, &addHeader)
			if err != nil {
				err = &BodyFormUrlencodedError{
					FormDataError{
						URL:    URL,
						Method: method,
						Parent: err,
					},
				}

				return
			}

		case mimes.JSON:
			reqBody, reqRawBody, err = c.GetBodyJSON(URL, method, inData, &addHeader)
			if err != nil {
				err = &BodyJSONError{
					FormDataError{
						URL:    URL,
						Method: method,
						Parent: err,
					},
				}

				return
			}

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
			RawBody: c.ErrorRawBody(c.ContentType, reqRawBody),
			Parent:  err,
		}

		return
	}

	contentType, found := addHeader["Content-Type"]
	if found {
		req.Header.Add("Content-Type", contentType)
		delete(addHeader, "Content-Type")
	} else {
		req.Header.Add("Content-Type", c.ContentType)
	}

	for key, value := range addHeader {
		req.Header.Add(key, value)
	}

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
	case AuthTypeEDS:
		req.Header.Add("Sign-Key-Id", c.SignKeyID)
		req.Header.Add("Sign-Body", c.SignBody)
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
			RawBody: c.ErrorRawBody(c.ContentType, reqRawBody),
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
			RawBody: c.ErrorRawBody(c.ContentType, reqRawBody),
			Parent:  err,
		}

		return
	}

	respRawBody := string(respBody)

	if statusCode != expectedStatusCode {
		err = &WrongStatusCodeError{
			URL:             URL,
			Method:          method,
			RequestRawBody:  c.ErrorRawBody(c.ContentType, reqRawBody),
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

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)

	return err
}
