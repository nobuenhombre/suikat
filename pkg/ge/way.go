package ge

import (
	"fmt"
	"runtime"
	"strings"
)

type Way struct {
	Package string
	Caller  string
	Line    int
}

func (w *Way) View() string {
	return fmt.Sprintf("%v : %v, line %v", w.Package, w.Caller, w.Line)
}

// main.Father.Hello
// main.(*Father).Hello
// main.Hello
//
// repco/internal/pkg/daddy.Father.Hello
// repco/internal/pkg/daddy.(*Father).Hello
// repco/internal/pkg/daddy.Hello
//
// github.com/nobuenhombre/suikat/pkg/daddy.Father.Hello
// github.com/nobuenhombre/suikat/pkg/daddy.(*Father).Hello
// github.com/nobuenhombre/suikat/pkg/daddy.Hello

const (
	skipCallers     = 2
	funcNameWithObj = 3
)

func getWay() *Way {
	pc, _, line, ok := runtime.Caller(skipCallers)
	if !ok {
		return nil
	}

	funcName := runtime.FuncForPC(pc).Name()

	lastSlash := strings.LastIndexByte(funcName, '/')

	path := ""
	place := funcName

	if lastSlash > -1 {
		path = funcName[:lastSlash]
		place = funcName[lastSlash+1:]
	}

	list := strings.Split(place, ".")

	pkg := list[0]
	obj := ""
	caller := list[1]

	if len(list) == funcNameWithObj {
		pkg = list[0]
		obj = list[1] + "."
		caller = list[2]
	}

	return &Way{
		Package: fmt.Sprintf("%v/%v", path, pkg),
		Caller:  fmt.Sprintf("%v%v()", obj, caller),
		Line:    line,
	}
}
