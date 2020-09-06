package interceptor

import (
	"net/http"
)

type HTTPRouter struct {
	Routes                []HTTPRoute
	DefaultNotFoundAnswer HTTPAnswer
	HTTPRegexp
}

func (router *HTTPRouter) Init() {
	router.Routes = make([]HTTPRoute, 0)
}

func (router *HTTPRouter) HandleFunc(method, uri string, f func(http.ResponseWriter, *http.Request)) {
	route := &HTTPRoute{
		Method: method,
		URI:    uri,
		F:      f,
	}
	router.Routes = append(router.Routes, *route)
}

func (router *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.Routes {
		// Method Matched
		if route.MatchMethod(r) {
			// URI Matched
			if route.MatchURI(r, router.HTTPRegexp) {
				route.F(w, r)

				return
			}
		}
	}

	router.DefaultNotFoundAnswer.Send(w)
}
