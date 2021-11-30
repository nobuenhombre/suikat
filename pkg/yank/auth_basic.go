package yank

type AuthBasic struct {
	BaseAuth

	Username string
	Password string
}

func NewAuthBasic(username, password string) Auth {
	return &AuthBasic{
		BaseAuth: BaseAuth{
			AuthType: AuthTypeBasic,
		},

		Username: username,
		Password: password,
	}
}

func (a *AuthBasic) AddHeaders(r *HTTPRequest) {
	r.SetBasicAuth(a.Username, a.Password)
}
