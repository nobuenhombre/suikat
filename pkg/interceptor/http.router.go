package interceptor

import (
	"net/http"

	"go.uber.org/zap"
)

type HTTPRouter struct {
	Logger                *zap.Logger
	Routes                []HTTPRoute
	DefaultNotFoundAnswer HTTPAnswer
	HTTPRegexp
}

func (router *HTTPRouter) Init() {
	router.Routes = make([]HTTPRoute, 0)
}

func (router *HTTPRouter) HandleFunc(method, uri string, f HandlerFunc) {
	route := &HTTPRoute{
		Method: method,
		URI:    uri,
		F:      f,
	}
	router.Routes = append(router.Routes, *route)
}

func (router *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	for _, route := range router.Routes {
		// Method Matched
		if route.MatchMethod(r) {
			// URI Matched
			if route.MatchURI(r, router.HTTPRegexp) {
				err = route.F(w, r)
				if err != nil && router.Logger != nil {
					router.Logger.Error("route.F() error", zap.Error(err))
				}

				return
			}
		}
	}

	err = router.DefaultNotFoundAnswer.Send(w, r)
	if err != nil && router.Logger != nil {
		router.Logger.Error("router.DefaultNotFoundAnswer.Send(w, r) error", zap.Error(err))
	}
}
