package interceptor

import (
	"net/http"
	"strings"
)

// Один Роут и методы его сравнения
type HTTPRoute struct {
	Method string
	URI    string
	F      func(http.ResponseWriter, *http.Request)
}

func (route *HTTPRoute) MatchMethod(r *http.Request) bool {
	return route.Method == r.Method
}

func (route *HTTPRoute) MatchURI(r *http.Request, regexp HTTPRegexp) bool {
	if route.URI == r.URL.Path {
		// Полное совпадение
		return true
	}

	URI := strings.Trim(route.URI, "/")
	Path := strings.Trim(r.URL.Path, "/")
	routeURIParts := strings.Split(URI, "/")
	requestURLParts := strings.Split(Path, "/")

	return regexp.MatchURIParts(routeURIParts, requestURLParts)
}
