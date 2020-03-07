package inslice

import (
	"reflect"
	"testing"
)

type inputFloat32Params struct {
	a    float32
	list *[]float32
}

type insliceFloat32Test struct {
	in  inputFloat32Params
	out bool
}

var float32Tests = []insliceFloat32Test{
	{
		in: inputFloat32Params{
			a:    5.73,
			list: nil,
		},
		out: false,
	},
	{
		in: inputFloat32Params{
			a:    5.73,
			list: &[]float32{3.12, 7.45, 9.88},
		},
		out: false,
	},
	{
		in: inputFloat32Params{
			a:    5.73,
			list: &[]float32{3.12, 5.73, 9.88},
		},
		out: true,
	},
}

func TestFloat32(t *testing.T) {
	for i := 0; i < len(float32Tests); i++ {
		test := &float32Tests[i]
		out := Float32(test.in.a, test.in.list)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Float32(%v, %v), Expected %v, Actual %v",
				test.in.a, test.in.list, test.out, out,
			)
		}
	}
}
