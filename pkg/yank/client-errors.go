package yank

import (
	"fmt"
	"net/http"
)

// ApplyDefaultsError - Ошибка применения настроек по умолчанию (1)
//-----------------------------------------------------------------
type ApplyDefaultsError struct {
	Parent error
}

func (e *ApplyDefaultsError) Error() string {
	return fmt.Sprintf("cant apply defaults,\n"+
		" - error[%v]",
		e.Parent,
	)
}

// RawContentError - Ошибка формирования тела запроса из структуры данных (2)
//---------------------------------------------------------------------------
type RawContentError struct {
	Method   string
	URL      string
	MimeType string
	Parent   error
}

func (e *RawContentError) Error() string {
	return fmt.Sprintf("raw content error,\n"+
		" - method[%v],\n"+
		" - url[%v],\n"+
		" - mime[%v],\n"+
		" - error[%v]",
		e.Method,
		e.URL,
		e.MimeType,
		e.Parent,
	)
}

// CreateHTTPRequestError - Ошибка создания http запроса (3)
//----------------------------------------------------------
type CreateHTTPRequestError struct {
	Method         string
	URL            string
	MimeType       string
	RequestRawBody string
	Parent         error
}

func (e *CreateHTTPRequestError) Error() string {
	return fmt.Sprintf("create http request error,\n"+
		" - method[%v],\n"+
		" - url[%v],\n"+
		" - mime[%v],\n"+
		" - request[%v]\n"+
		" - error[%v]",
		e.Method,
		e.URL,
		e.MimeType,
		e.RequestRawBody,
		e.Parent,
	)
}

// ExecuteHTTPRequestError - Ошибка исполнения http запроса (4)
//-------------------------------------------------------------
type ExecuteHTTPRequestError struct {
	Method            string
	URL               string
	MimeType          string
	RequestRawBody    string
	RequestRawHeaders http.Header
	Parent            error
}

func (e *ExecuteHTTPRequestError) Error() string {
	return fmt.Sprintf("execute http request error,\n"+
		" - method[%v],\n"+
		" - url[%v],\n"+
		" - mime[%v],\n"+
		" - request[%v]\n"+
		" - request.headers[%v]\n"+
		" - error[%v]",
		e.Method,
		e.URL,
		e.MimeType,
		e.RequestRawBody,
		e.RequestRawHeaders,
		e.Parent,
	)
}

// ReadBodyHTTPRequestError - Ошибка чтения тела http ответа (5)
//--------------------------------------------------------------
type ReadBodyHTTPRequestError struct {
	Method             string
	URL                string
	MimeType           string
	RequestRawBody     string
	RequestRawHeaders  http.Header
	ResponseRawHeaders http.Header
	Parent             error
}

func (e *ReadBodyHTTPRequestError) Error() string {
	return fmt.Sprintf("execute http request error,\n"+
		" - method[%v],\n"+
		" - url[%v],\n"+
		" - mime[%v],\n"+
		" - request[%v]\n"+
		" - request.headers[%v]\n"+
		" - response.headers[%v]\n"+
		" - error[%v]",
		e.Method,
		e.URL,
		e.MimeType,
		e.RequestRawBody,
		e.RequestRawHeaders,
		e.ResponseRawHeaders,
		e.Parent,
	)
}

// WrongHTTPCodeError - Ошибка неверный http код (6)
//--------------------------------------------------
type WrongHTTPCodeError struct {
	Method             string
	URL                string
	MimeType           string
	RequestRawBody     string
	RequestRawHeaders  http.Header
	ResponseRawBody    string
	ResponseRawHeaders http.Header
	Expected           int
	Actual             int
}

func (e *WrongHTTPCodeError) Error() string {
	return fmt.Sprintf(
		"wrong status code,\n"+
			" - method[%v],\n"+
			" - url[%v],\n"+
			" - mime[%v],\n"+
			" - request[%v],\n"+
			" - request.headers[%v]\n"+
			" - response[%v],\n"+
			" - response.headers[%v]\n"+
			" - expected code[%v],\n"+
			" - actual code[%v]",
		e.Method,
		e.URL,
		e.MimeType,
		e.RequestRawBody,
		e.RequestRawHeaders,
		e.ResponseRawBody,
		e.ResponseRawHeaders,
		e.Expected,
		e.Actual,
	)
}

// ParseResponseError - Ошибка разбора тела ответа в структуру (7)
//----------------------------------------------------------------
type ParseResponseError struct {
	Method             string
	URL                string
	MimeType           string
	RequestRawBody     string
	RequestRawHeaders  http.Header
	ResponseRawBody    string
	ResponseRawHeaders http.Header
	Parent             error
}

func (e *ParseResponseError) Error() string {
	return fmt.Sprintf(
		"parse response body,\n"+
			" - method[%v],\n"+
			" - url[%v],\n"+
			" - mime[%v],\n"+
			" - request[%v],\n"+
			" - request.headers[%v]\n"+
			" - response[%v],\n"+
			" - response.headers[%v]\n"+
			" - error[%v]",
		e.Method,
		e.URL,
		e.MimeType,
		e.RequestRawBody,
		e.RequestRawHeaders,
		e.ResponseRawBody,
		e.ResponseRawHeaders,
		e.Parent,
	)
}
