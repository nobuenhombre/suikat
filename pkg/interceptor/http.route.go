package interceptor

import (
	"net/http"
	"strings"
)

// Один Роут и методы его сравнения
type HttpRoute struct {
	Method string
	URI    string
	F      func(http.ResponseWriter, *http.Request)
}

func (route *HttpRoute) MatchMethod(r *http.Request) bool {
	return route.Method == r.Method
}

func (route *HttpRoute) MatchURI(r *http.Request, regexp HttpRegexp) bool {
	if route.URI == r.URL.Path {
		// Полное совпадение
		return true
	} else {
		URI := strings.Trim(route.URI, "/")
		Path := strings.Trim(r.URL.Path, "/")
		URIParts := strings.Split(URI, "/")
		PathParts := strings.Split(Path, "/")

		var pattern string

		if len(URIParts) > 0 && len(URIParts) == len(PathParts) {
			matched := true
			for index, pathPart := range PathParts {
				if len(pathPart) > 0 && len(URIParts[index]) > 0 {
					if pathPart == URIParts[index] {
						// Прямое совпадение куска роута
						matched = matched && true
					} else {
						// Совпадение по паттерну
						pattern = URIParts[index]
						r, found := regexp.List[pattern]
						if found {
							matched = matched && r.MatchString(pathPart)
						} else {
							matched = matched && false
						}
					}
				} else {
					return false
				}
			}
			return matched
		}

		return false
	}
}
