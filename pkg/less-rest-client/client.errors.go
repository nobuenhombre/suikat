package lessrestclient

import (
	"fmt"
)

type UnknownAuthTypeError struct {
	URL    string
	Method string
	Data   interface{}
}

func (e *UnknownAuthTypeError) Error() string {
	return fmt.Sprintf(
		"unknown auth type error,\n - method[%v],\n - url[%v],\n - data[%#v]",
		e.Method,
		e.URL,
		e.Data,
	)
}

type UnknownContentTypeError struct {
	URL    string
	Method string
	Data   interface{}
}

func (e *UnknownContentTypeError) Error() string {
	return fmt.Sprintf(
		"unknown content type error,\n - method[%v],\n - url[%v],\n - data[%#v]",
		e.Method,
		e.URL,
		e.Data,
	)
}

type EncoderError struct {
	URL    string
	Method string
	Data   interface{}
	Parent error
}

func (e *EncoderError) Error() string {
	return fmt.Sprintf(
		"api encoder error,\n - method[%v],\n - url[%v],\n - data[%#v],\n - error[%v]",
		e.Method,
		e.URL,
		e.Data,
		e.Parent,
	)
}

type MarshalingError struct {
	URL    string
	Method string
	Data   interface{}
	Parent error
}

func (e *MarshalingError) Error() string {
	return fmt.Sprintf(
		"api marshaling error,\n - method[%v],\n - url[%v],\n - data[%#v],\n - error[%v]",
		e.Method,
		e.URL,
		e.Data,
		e.Parent,
	)
}

type UnMarshalingError struct {
	URL     string
	Method  string
	RawBody string
	Parent  error
}

func (e *UnMarshalingError) Error() string {
	return fmt.Sprintf(
		"api unmarshaling error,\n - method[%v],\n - url[%v],\n - rawbody[%v],\n - error[%v]",
		e.Method,
		e.URL,
		e.RawBody,
		e.Parent,
	)
}

type CreateRequestError struct {
	URL     string
	Method  string
	RawBody string
	Parent  error
}

func (e *CreateRequestError) Error() string {
	return fmt.Sprintf(
		"api create request error,\n - method[%v],\n - url[%v],\n - rawbody[%v],\n - error[%v]",
		e.Method,
		e.URL,
		e.RawBody,
		e.Parent,
	)
}

type ClientError struct {
	URL     string
	Method  string
	RawBody string
	Parent  error
}

func (e *ClientError) Error() string {
	return fmt.Sprintf(
		"api client error,\n - method[%v],\n - url[%v],\n - rawbody[%v],\n - error[%v]",
		e.Method,
		e.URL,
		e.RawBody,
		e.Parent,
	)
}

type ReadResponseBodyError struct {
	URL     string
	Method  string
	RawBody string
	Parent  error
}

func (e *ReadResponseBodyError) Error() string {
	return fmt.Sprintf(
		"api read response body error,\n - method[%v],\n - url[%v],\n - rawbody[%v],\n - error[%v]",
		e.Method,
		e.URL,
		e.RawBody,
		e.Parent,
	)
}

type WrongStatusCodeError struct {
	URL             string
	Method          string
	RequestRawBody  string
	ResponseRawBody string
	Expected        int
	Actual          int
}

func (e *WrongStatusCodeError) Error() string {
	return fmt.Sprintf(
		"api return wrong status code,\n"+
			" - method[%v],\n - url[%v],\n - request[%v],\n - response[%v],\n"+
			" - expected code[%v],\n - actual code[%v]",
		e.Method,
		e.URL,
		e.RequestRawBody,
		e.ResponseRawBody,
		e.Expected,
		e.Actual,
	)
}

type FormDataError struct {
	URL    string
	Method string
	Parent error
}

func (e *FormDataError) Error() string {
	return fmt.Sprintf("form multipart data error,\n"+
		" - method[%v],\n - url[%v],\n"+
		" - error[%v]",
		e.Method,
		e.URL,
		e.Parent,
	)
}

type BodyFormMultipartDataError struct {
	FormDataError
}

type BodyFormUrlencodedError struct {
	FormDataError
}

type BodyJSONError struct {
	FormDataError
}

type IsNotPointerError struct {
	Name string
}

func (e *IsNotPointerError) Error() string {
	return fmt.Sprintf(
		"var [%#v] is not pointer",
		e.Name,
	)
}
