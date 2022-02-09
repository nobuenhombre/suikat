package inwords

import (
	"reflect"
	"testing"
)

type convertTest struct {
	in  float64
	out string
	err error
}

var convertTests = []convertTest{
	{
		in:  0.0,
		out: "ноль рублей 00 копеек",
		err: nil,
	},
	{
		in:  25230.33,
		out: "двадцать пять тысяч двести тридцать рублей 33 копейки",
		err: nil,
	},
	{
		in:  230.45,
		out: "двести тридцать рублей 45 копеек",
		err: nil,
	},
	{
		in:  3333.01,
		out: "три тысячи триста тридцать три рубля 01 копейка",
		err: nil,
	},
}

func TestRound(t *testing.T) {
	for i := 0; i < len(convertTests); i++ {
		test := &convertTests[i]
		out, err := Format(test.in)

		if !(reflect.DeepEqual(out, test.out) && reflect.DeepEqual(err, test.err)) {
			t.Errorf(
				"num2str(%#v), \nExpected (%#v, %#v) \nActual__ (%#v, %#v)",
				test.in,
				test.out, test.err,
				out, err,
			)
		}
	}
}
