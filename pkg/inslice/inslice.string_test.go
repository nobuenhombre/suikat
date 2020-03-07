package inslice

import (
	"reflect"
	"testing"
)

type inputStringParams struct {
	a    string
	list *[]string
}

type insliceStringTest struct {
	in  inputStringParams
	out bool
}

var stringTests = []insliceStringTest{
	{
		in: inputStringParams{
			a:    "a",
			list: nil,
		},
		out: false,
	},
	{
		in: inputStringParams{
			a:    "a",
			list: &[]string{"b", "c", "d"},
		},
		out: false,
	},
	{
		in: inputStringParams{
			a:    "a",
			list: &[]string{"b", "a", "d"},
		},
		out: true,
	},
}

func TestString(t *testing.T) {
	for i := 0; i < len(stringTests); i++ {
		test := &stringTests[i]
		out := String(test.in.a, test.in.list)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"String(%v, %v), Expected %v, Actual %v",
				test.in.a, test.in.list, test.out, out,
			)
		}
	}
}
