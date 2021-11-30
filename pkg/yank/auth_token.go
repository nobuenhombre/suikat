package yank

type AuthToken struct {
	BaseAuth

	Token string
}

func NewAuthToken(token string) Auth {
	return &AuthToken{
		BaseAuth: BaseAuth{
			AuthType: AuthTypeToken,
		},

		Token: token,
	}
}

func (a *AuthToken) AddHeaders(r *HTTPRequest) {
	r.Header.Add("Authorization", a.Token)
}
