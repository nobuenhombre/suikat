package yank

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Request struct {
	Method          string
	URL             string
	Route           string
	Query           url.Values
	Auth            Auth
	Headers         http.Header
	Content         Content
	FollowRedirects bool
}

type RequestConstructor interface {
	URLConstructor
	QueryConstructor
	HeaderConstructor
	AuthConstructor
	BodyConstructor
}

func NewRequest(route string) *Request {
	return &Request{
		Route: route,
	}
}

func (r *Request) URLConstructor() URLConstructor {
	return r
}

func (r *Request) QueryConstructor() QueryConstructor {
	return r
}

func (r *Request) HeaderConstructor() HeaderConstructor {
	return r
}

func (r *Request) AuthConstructor() AuthConstructor {
	return r
}

func (r *Request) RequestConstructor() RequestConstructor {
	return r
}

func (r *Request) BodyConstructor() BodyConstructor {
	return r
}

func (r *Request) SetURL(url string) {
	r.URL = url
}

func (r *Request) SetFollowRedirects(followRedirects bool) {
	r.FollowRedirects = followRedirects
}

func (r *Request) AddQuery(key, value string) {
	if r.Query == nil {
		r.Query = url.Values{}
	}

	r.Query.Add(key, value)
}

func (r *Request) AddHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = http.Header{}
	}

	r.Headers.Add(key, value)
}

func (r *Request) AuthNone() {
	r.Auth = NewAuthNone()
}

func (r *Request) AuthBasic(username, password string) {
	r.Auth = NewAuthBasic(username, password)
}

func (r *Request) AuthToken(token string) {
	r.Auth = NewAuthToken(token)
}

func (r *Request) AuthBearerToken(token string) {
	r.Auth = NewAuthBearerToken(token)
}

func (r *Request) AuthCookieWithCSRFToken(cookie, token string) {
	r.Auth = NewAuthCookieWithCSRFToken(cookie, token)
}

func (r *Request) AuthEDS(keyID, body string) {
	r.Auth = NewAuthEDS(keyID, body)
}

func (r *Request) SetBody(content Content) {
	r.Content = content
}

func (r *Request) GetURI() string {
	if r.Query != nil {
		return fmt.Sprintf("%v%v?%v", r.URL, r.Route, r.Query.Encode())
	}

	return fmt.Sprintf("%v%v", r.URL, r.Route)
}

func (r *Request) GetRawContent() (raw *RawContent, err error) {
	if r.Content == nil {
		raw = NewRawContent(nil, "", "")
	} else {
		raw, err = r.Content.GetRawContent()
		if err != nil {
			return nil, ge.Pin(err, ge.Params{"request": r})
		}
	}

	return
}

func (r *Request) NewHTTPRequest(
	uri string,
	raw *RawContent,
	defaults *Defaults,
	ignoreDefaults bool,
) (httpRequest *HTTPRequest, err error) {
	body := raw.GetBuffer()

	req, err := http.NewRequest(r.Method, uri, body)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"request": r})
	}

	httpRequest = &HTTPRequest{
		Request:         req,
		FollowRedirects: r.FollowRedirects,
	}

	raw.AddHeaders(httpRequest)

	// Установим авторизацию из запроса если она задана!
	if r.Auth != nil {
		r.Auth.AddHeaders(httpRequest)
	}

	if !ignoreDefaults {
		defaults.AddHeaders(httpRequest)

		// Установим авторизацию из настроек по умолчанию, если авторизация не указана в запросе!
		// Если авторизация указана в запросе - она имеет приоритет
		// и авторизация из настроек не применяется
		if r.Auth == nil && defaults.Auth != nil {
			defaults.Auth.AddHeaders(httpRequest)
		}
	}

	return httpRequest, nil
}
