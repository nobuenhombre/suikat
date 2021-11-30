package yank

type AuthNone struct {
	BaseAuth
}

func NewAuthNone() Auth {
	return &AuthNone{
		BaseAuth: BaseAuth{
			AuthType: AuthTypeNone,
		},
	}
}

func (a *AuthNone) AddHeaders(r *HTTPRequest) {

}
