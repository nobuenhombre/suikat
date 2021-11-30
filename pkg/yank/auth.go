package yank

const (
	AuthTypeNone = iota
	AuthTypeBasic
	AuthTypeToken
	AuthTypeBearerToken
	AuthTypeCookieWithCSRFToken
	AuthTypeEDS
)

type BaseAuth struct {
	AuthType int
}

type Auth interface {
	HTTPHeaders
}
