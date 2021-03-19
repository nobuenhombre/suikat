package ge

import (
	"errors"
	"reflect"
	"testing"
)

type getPathPlaceTest struct {
	funcName string
	path     string
	place    string
}

var getPathPlaceTests = []getPathPlaceTest{
	{
		funcName: "hmm",
		path:     "",
		place:    "hmm",
	},
	{
		funcName: "main.Hello",
		path:     "",
		place:    "main.Hello",
	},
	{
		funcName: "main.Father.Hello",
		path:     "",
		place:    "main.Father.Hello",
	},
	{
		funcName: "main.(*Father).Hello",
		path:     "",
		place:    "main.(*Father).Hello",
	},

	{
		funcName: "repco/internal/pkg/daddy.Hello",
		path:     "repco/internal/pkg",
		place:    "daddy.Hello",
	},
	{
		funcName: "repco/internal/pkg/daddy.Father.Hello",
		path:     "repco/internal/pkg",
		place:    "daddy.Father.Hello",
	},
	{
		funcName: "repco/internal/pkg/daddy.(*Father).Hello",
		path:     "repco/internal/pkg",
		place:    "daddy.(*Father).Hello",
	},

	{
		funcName: "github.com/nobuenhombre/suikat/pkg/daddy.Hello",
		path:     "github.com/nobuenhombre/suikat/pkg",
		place:    "daddy.Hello",
	},
	{
		funcName: "github.com/nobuenhombre/suikat/pkg/daddy.Father.Hello",
		path:     "github.com/nobuenhombre/suikat/pkg",
		place:    "daddy.Father.Hello",
	},
	{
		funcName: "github.com/nobuenhombre/suikat/pkg/daddy.(*Father).Hello",
		path:     "github.com/nobuenhombre/suikat/pkg",
		place:    "daddy.(*Father).Hello",
	},
}

func TestGetPathPlace(t *testing.T) {
	for i := 0; i < len(getPathPlaceTests); i++ {
		test := &getPathPlaceTests[i]
		path, place := getPathPlace(test.funcName)

		if !(reflect.DeepEqual(path, test.path) && reflect.DeepEqual(place, test.place)) {
			t.Errorf(
				"getPathPlace(%v),\n "+
					"Expected (%v, %v),\n "+
					"Actual (%v, %v)",
				test.funcName,
				test.path, test.place,
				path, place,
			)
		}
	}
}

type getPkgObjCallerTest struct {
	place  string
	pkg    string
	obj    string
	caller string
}

var getPkgObjCallerTests = []getPkgObjCallerTest{
	{
		place:  "main.Hello",
		pkg:    "main",
		obj:    "",
		caller: "Hello",
	},
	{
		place:  "main.Father.Hello",
		pkg:    "main",
		obj:    "Father.",
		caller: "Hello",
	},
	{
		place:  "main.(*Father).Hello",
		pkg:    "main",
		obj:    "(*Father).",
		caller: "Hello",
	},
}

func TestGetPkgObjCaller(t *testing.T) {
	for i := 0; i < len(getPkgObjCallerTests); i++ {
		test := &getPkgObjCallerTests[i]
		pkg, obj, caller := getPkgObjCaller(test.place)

		if !(reflect.DeepEqual(pkg, test.pkg) && reflect.DeepEqual(obj, test.obj) && reflect.DeepEqual(caller, test.caller)) {
			t.Errorf(
				"getPkgObjCaller(%v),\n "+
					"Expected (%v, %v, %v),\n "+
					"Actual (%v, %v, %v)",
				test.place,
				test.pkg, test.obj, test.caller,
				pkg, obj, caller,
			)
		}
	}
}

var wayTest = &Way{
	Package: "github.com/nobuenhombre/suikat/pkg/ge",
	Caller:  "TestGetWay()",
	File:    "way_test.go",
	Line:    142,
}

func TestGetWay(t *testing.T) {
	err := New("Hello")

	var e *IdentityError
	if errors.As(err, &e) {
		if !reflect.DeepEqual(e.Way, wayTest) {
			t.Errorf(
				"getWay(),\n "+
					"Expected (%v),\n "+
					"Actual (%v)",
				wayTest, e.Way,
			)
		}
	}
}
