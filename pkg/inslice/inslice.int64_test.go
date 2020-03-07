package inslice

import (
	"reflect"
	"testing"
)

type inputInt64Params struct {
	a    int64
	list []int64
}

type insliceInt64Test struct {
	in  inputInt64Params
	out bool
}

var int64Tests = []insliceInt64Test{
	{
		in: inputInt64Params{
			a:    5,
			list: nil,
		},
		out: false,
	},
	{
		in: inputInt64Params{
			a:    5,
			list: []int64{3, 7, 9},
		},
		out: false,
	},
	{
		in: inputInt64Params{
			a:    5,
			list: []int64{3, 5, 7},
		},
		out: true,
	},
}

func TestInt64(t *testing.T) {
	for i := 0; i < len(int64Tests); i++ {
		test := &int64Tests[i]
		out := Int64(test.in.a, &test.in.list)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Int64(%v, %v), Expected %v, Actual %v",
				test.in.a, test.in.list, test.out, out,
			)
		}
	}
}
