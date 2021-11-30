package yank

type AuthBearerToken struct {
	BaseAuth

	Token string
}

func NewAuthBearerToken(token string) Auth {
	return &AuthBearerToken{
		BaseAuth: BaseAuth{
			AuthType: AuthTypeBearerToken,
		},

		Token: token,
	}
}

func (a *AuthBearerToken) AddHeaders(r *HTTPRequest) {
	r.Header.Add("Authorization", "Bearer "+a.Token)
}
