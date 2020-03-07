package inslice

import (
	"reflect"
	"testing"
)

type inputFloat64Params struct {
	a    float64
	list []float64
}

type insliceFloat64Test struct {
	in  inputFloat64Params
	out bool
}

var float64Tests = []insliceFloat64Test{
	{
		in: inputFloat64Params{
			a:    5.73,
			list: nil,
		},
		out: false,
	},
	{
		in: inputFloat64Params{
			a:    5.73,
			list: []float64{3.12, 7.45, 9.88},
		},
		out: false,
	},
	{
		in: inputFloat64Params{
			a:    5.73,
			list: []float64{3.12, 5.73, 9.88},
		},
		out: true,
	},
}

func TestFloat64(t *testing.T) {
	for i := 0; i < len(float64Tests); i++ {
		test := &float64Tests[i]
		out := Float64(test.in.a, &test.in.list)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Float64(%v, %v), Expected %v, Actual %v",
				test.in.a, test.in.list, test.out, out,
			)
		}
	}
}
