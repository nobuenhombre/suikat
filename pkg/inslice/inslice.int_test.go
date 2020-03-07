package inslice

import (
	"reflect"
	"testing"
)

type inputIntParams struct {
	a    int
	list []int
}

type insliceIntTest struct {
	in  inputIntParams
	out bool
}

var intTests = []insliceIntTest{
	{
		in: inputIntParams{
			a:    5,
			list: nil,
		},
		out: false,
	},
	{
		in: inputIntParams{
			a:    5,
			list: []int{3, 7, 9},
		},
		out: false,
	},
	{
		in: inputIntParams{
			a:    5,
			list: []int{3, 5, 7},
		},
		out: true,
	},
}

func TestInt(t *testing.T) {
	for i := 0; i < len(intTests); i++ {
		test := &intTests[i]
		out := Int(test.in.a, &test.in.list)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Int(%v, %v), Expected %v, Actual %v",
				test.in.a, test.in.list, test.out, out,
			)
		}
	}
}
