package interceptor

import "regexp"

// Регулярные выражения для проверки роутов
type HttpRegexp struct {
	List map[string]*regexp.Regexp
}

func (reg *HttpRegexp) Init() {
	reg.List = make(map[string]*regexp.Regexp, 0)
}

func (reg *HttpRegexp) Add(pattern, regExpPattern string) {
	reg.List[pattern] = regexp.MustCompile(regExpPattern)
}
