package interceptor

import (
	"net/http"
)

type HttpRouter struct {
	Routes                []HttpRoute
	DefaultNotFoundAnswer HttpAnswer
	HttpRegexp
}

func (router *HttpRouter) Init() {
	router.Routes = make([]HttpRoute, 0)
}

func (router *HttpRouter) HandleFunc(method, uri string, f func(http.ResponseWriter, *http.Request)) {
	route := &HttpRoute{
		Method: method,
		URI:    uri,
		F:      f,
	}
	router.Routes = append(router.Routes, *route)
}

func (router *HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.Routes {
		// Method Matched
		if route.MatchMethod(r) {
			// URI Matched
			if route.MatchURI(r, router.HttpRegexp) {
				route.F(w, r)
				return
			}
		}
	}

	router.DefaultNotFoundAnswer.Send(w)
}
