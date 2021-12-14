package ge

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Way
// en: the structure contains the exact coordinates of the error, package, function, file name, line number in the file
// ru: структура содержит точные координаты ошибки, пакет, функция, имя файла, номер строки в файле
type Way struct {
	Package string
	Caller  string
	File    string
	Line    int
}

// View
// en: returns a formatted string with error coordinates
// ru: возвращает форматированную строку с координатами ошибки
func (w *Way) View() string {
	return fmt.Sprintf("%v : %v, file %v line %v", w.Package, w.Caller, w.File, w.Line)
}

const (
	skipCallers        = 2
	funcNameWithoutObj = 2
	funcNameWithObj    = 3
)

func getPathPlace(funcName string) (path, place string) {
	lastSlash := strings.LastIndexByte(funcName, '/')

	path = ""
	place = funcName

	if lastSlash > -1 {
		path = funcName[:lastSlash]
		place = funcName[lastSlash+1:]
	}

	return
}

func getPkgObjCaller(place string) (pkg, obj, caller string) {
	list := strings.Split(place, ".")

	switch len(list) {
	case funcNameWithoutObj:
		pkg = list[0]
		obj = ""
		caller = list[1]

	case funcNameWithObj:
		pkg = list[0]
		obj = list[1] + "."
		caller = list[2]

	default:
		pkg = list[0]
		obj = ""
		caller = ""
	}

	return
}

func getWay() *Way {
	pc, file, line, ok := runtime.Caller(skipCallers)
	if !ok {
		return nil
	}

	funcName := runtime.FuncForPC(pc).Name()
	path, place := getPathPlace(funcName)
	pkg, obj, caller := getPkgObjCaller(place)

	return &Way{
		Package: fmt.Sprintf("%v/%v", path, pkg),
		Caller:  fmt.Sprintf("%v%v()", obj, caller),
		File:    filepath.Base(file),
		Line:    line,
	}
}
