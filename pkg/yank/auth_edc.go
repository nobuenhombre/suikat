package yank

type AuthEDS struct {
	BaseAuth

	SignKeyID string
	SignBody  string
}

func NewAuthEDS(keyID, body string) Auth {
	return &AuthEDS{
		BaseAuth: BaseAuth{
			AuthType: AuthTypeEDS,
		},

		SignKeyID: keyID,
		SignBody:  body,
	}
}

func (a *AuthEDS) AddHeaders(r *HTTPRequest) {
	r.Header.Add("Sign-Key-Id", a.SignKeyID)
	r.Header.Add("Sign-Body", a.SignBody)
}
