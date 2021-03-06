package interceptor

import "regexp"

const (
	Word   = ":word"
	Number = ":number"
	Any    = ":any"
)

func Predefined() map[string]string {
	return map[string]string{
		Word:   "([\\w]+)",
		Number: "([\\d]+)",
		Any:    ".*",
	}
}

// Регулярные выражения для проверки роутов
type HTTPRegexp struct {
	List map[string]*regexp.Regexp
}

func (reg *HTTPRegexp) Init() {
	reg.List = make(map[string]*regexp.Regexp)
}

func (reg *HTTPRegexp) AddPredefined(predefined []string) {
	all := Predefined()

	for _, key := range predefined {
		rx, found := all[key]
		if found {
			reg.Add(key, rx)
		}
	}
}

func (reg *HTTPRegexp) Add(pattern, regExpPattern string) {
	reg.List[pattern] = regexp.MustCompile(regExpPattern)
}

func (reg *HTTPRegexp) MatchURIParts(routeURIParts, requestURLParts []string) bool {
	if !(len(routeURIParts) > 0 && len(routeURIParts) == len(requestURLParts)) {
		return false
	}

	matched := true

	for index, requestURLPart := range requestURLParts {
		routeURIPart := routeURIParts[index]

		if requestURLPart == routeURIPart {
			// Прямое совпадение куска роута
			matched = matched && true
		} else {
			// Совпадение по паттерну
			r, found := reg.List[routeURIPart]
			if found {
				matched = matched && r.MatchString(requestURLPart)
			} else {
				matched = matched && false
			}
		}
	}

	return matched
}
