package inslice

import (
	"reflect"
	"testing"
)

type inputIsIndexExistsParams struct {
	index int
	list  interface{}
}

type insliceIsIndexExistsTest struct {
	in  inputIsIndexExistsParams
	out bool
	err error
}

var isIndexExistsTests = []insliceIsIndexExistsTest{
	{
		in: inputIsIndexExistsParams{
			index: 0,
			list:  []bool{true, false},
		},
		out: true,
		err: nil,
	},
	{
		in: inputIsIndexExistsParams{
			index: 5,
			list:  []int{1, 2, 3, 4, 5, 6},
		},
		out: true,
		err: nil,
	},
	{
		in: inputIsIndexExistsParams{
			index: 7,
			list:  []string{"1", "2", "3", "4", "5", "6"},
		},
		out: false,
		err: &IndexNotExistsError{
			Index: 7,
		},
	},
	{
		in: inputIsIndexExistsParams{
			index: -1,
			list:  []float64{1.23, 2.34, 3.45, 4.56, 5.67, 6.78},
		},
		out: false,
		err: &IndexNotExistsError{
			Index: -1,
		},
	},
	{
		in: inputIsIndexExistsParams{
			index: 1,
			list:  99,
		},
		out: false,
		err: &IndexNotExistsError{
			Index: 1,
		},
	},
	{
		in: inputIsIndexExistsParams{
			index: 1,
			list:  nil,
		},
		out: false,
		err: &IndexNotExistsError{
			Index: 1,
		},
	},
}

func TestIsIndexExists(t *testing.T) {
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

func TestCheckIndex(t *testing.T) {
	for i := 0; i < len(isIndexExistsTests); i++ {
		test := &isIndexExistsTests[i]
		err := CheckIndex(test.in.index, test.in.list)

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf(
				"CheckIndex(%v, %v), Expected %v, Actual %v",
				test.in.index, test.in.list, test.err, err,
			)
		}
	}
}

type indexNotExistsErrorTest struct {
	err     *IndexNotExistsError
	message string
}

var testIndexNotExistsErrorTest = indexNotExistsErrorTest{
	err: &IndexNotExistsError{
		Index: 5,
	},
	message: "index [5] not exists",
}

func TestIndexNotExistsError(t *testing.T) {
	test := &testIndexNotExistsErrorTest

	if !reflect.DeepEqual(test.err.Error(), test.message) {
		t.Errorf(
			"index=%v, test.err.Error(), Expected %v, Actual %v",
			test.err.Index, test.message, test.err.Error(),
		)
	}
}
