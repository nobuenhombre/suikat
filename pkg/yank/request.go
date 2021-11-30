package yank

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Request struct {
	Method  string
	URL     string
	Route   string
	Query   url.Values
	Auth    Auth
	Headers http.Header
	Content Content
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
	raw *RawContent,
	defaults *Defaults,
	ignoreDefaults bool,
) (httpRequest *HTTPRequest, err error) {
	req, err := http.NewRequest(r.Method, r.GetURI(), raw.GetReader())
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"request": r})
	}

	httpRequest = &HTTPRequest{Request: req}

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
