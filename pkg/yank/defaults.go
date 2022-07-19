package yank

import "net/http"

type Defaults struct {
	URL             string
	Auth            Auth
	Headers         http.Header
	FollowRedirects bool
	Transport       http.RoundTripper
}

func NewDefaults(url string) *Defaults {
	return &Defaults{
		URL: url,
	}
}

type DefaultsConstructor interface {
	HeaderConstructor
	AuthConstructor
}

func (d *Defaults) SetFollowRedirects(value bool) {
	d.FollowRedirects = value
}

func (d *Defaults) SetTransport(value http.RoundTripper) {
	d.Transport = value
}

func (d *Defaults) AuthConstructor() AuthConstructor {
	return d
}

func (d *Defaults) HeaderConstructor() HeaderConstructor {
	return d
}

func (d *Defaults) AddHeader(key, value string) {
	if d.Headers == nil {
		d.Headers = http.Header{}
	}

	d.Headers.Add(key, value)
}

func (d *Defaults) AuthNone() {
	d.Auth = NewAuthNone()
}

func (d *Defaults) AuthBasic(username, password string) {
	d.Auth = NewAuthBasic(username, password)
}

func (d *Defaults) AuthToken(token string) {
	d.Auth = NewAuthToken(token)
}

func (d *Defaults) AuthBearerToken(token string) {
	d.Auth = NewAuthBearerToken(token)
}

func (d *Defaults) AuthCookieWithCSRFToken(cookie, token string) {
	d.Auth = NewAuthCookieWithCSRFToken(cookie, token)
}

func (d *Defaults) AuthEDS(keyID, body string) {
	d.Auth = NewAuthEDS(keyID, body)
}

func (d *Defaults) AddHeaders(r *HTTPRequest) {
	if d.Headers != nil {
		r.AddHeaders(d.Headers)
	}
}
