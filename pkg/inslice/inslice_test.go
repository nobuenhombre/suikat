package inslice

import (
	"reflect"
	"testing"
)

type inputIsIndexExistsParams struct {
	index int
	list  []interface{}
}

type insliceIsIndexExistsTest struct {
	in  inputIsIndexExistsParams
	out bool
}

var isIndexExistsTests = []insliceIsIndexExistsTest{
	{
		in: inputIsIndexExistsParams{
			index: 5,
			list:  []interface{}{1, 2, 3, 4, 5, 6},
		},
		out: true,
	},
	{
		in: inputIsIndexExistsParams{
			index: 7,
			list:  []interface{}{1, 2, 3, 4, 5, 6},
		},
		out: false,
	},
	{
		in: inputIsIndexExistsParams{
			index: -1,
			list:  []interface{}{1, 2, 3, 4, 5, 6},
		},
		out: false,
	},
}

func TestIIndexExists(t *testing.T) {
	for i := 0; i < len(isIndexExistsTests); i++ {
		test := &isIndexExistsTests[i]
		out := IsIndexExists(test.in.index, test.in.list)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"IsIndexExists(%v, %v), Expected %v, Actual %v",
				test.in.index, test.in.list, test.out, out,
			)
		}
	}
}
