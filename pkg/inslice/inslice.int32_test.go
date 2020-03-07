package inslice

import (
	"reflect"
	"testing"
)

type inputInt32Params struct {
	a    int32
	list []int32
}

type insliceInt32Test struct {
	in  inputInt32Params
	out bool
}

var int32Tests = []insliceInt32Test{
	{
		in: inputInt32Params{
			a:    5,
			list: nil,
		},
		out: false,
	},
	{
		in: inputInt32Params{
			a:    5,
			list: []int32{3, 7, 9},
		},
		out: false,
	},
	{
		in: inputInt32Params{
			a:    5,
			list: []int32{3, 5, 7},
		},
		out: true,
	},
}

func TestInt32(t *testing.T) {
	for i := 0; i < len(int32Tests); i++ {
		test := &int32Tests[i]
		out := Int32(test.in.a, &test.in.list)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Int32(%v, %v), Expected %v, Actual %v",
				test.in.a, test.in.list, test.out, out,
			)
		}
	}
}
