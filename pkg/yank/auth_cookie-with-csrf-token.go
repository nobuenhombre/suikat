package yank

type AuthCookieWithCSRFToken struct {
	BaseAuth

	Cookie string
	Token  string
}

func NewAuthCookieWithCSRFToken(cookie, token string) Auth {
	return &AuthCookieWithCSRFToken{
		BaseAuth: BaseAuth{
			AuthType: AuthTypeCookieWithCSRFToken,
		},

		Cookie: cookie,
		Token:  token,
	}
}

func (a *AuthCookieWithCSRFToken) AddHeaders(r *HTTPRequest) {
	r.Header.Add("Cookie", a.Cookie)
	r.Header.Add("X-CSRF-Token", a.Token)
}
