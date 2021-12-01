package yank

type URLConstructor interface {
	SetURL(url string)
}

type QueryConstructor interface {
	AddQuery(key, value string)
}

type HeaderConstructor interface {
	AddHeader(key, value string)
}

type AuthConstructor interface {
	AuthNone()
	AuthBasic(username, password string)
	AuthToken(token string)
	AuthBearerToken(token string)
	AuthCookieWithCSRFToken(cookie, token string)
	AuthEDS(keyID, body string)
}

type BodyConstructor interface {
	SetBody(content Content)
}
