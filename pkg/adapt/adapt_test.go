package adapt

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type checkParams struct {
	val        reflect.Value
	expectType string
}

type checkTest struct {
	in  checkParams
	err error
}

var checkTests = []checkTest{
	{
		in: checkParams{
			val:        reflect.ValueOf(int(5)),
			expectType: "int",
		},
		err: nil,
	},
	{
		in: checkParams{
			val:        reflect.ValueOf(string("Hello")),
			expectType: "string",
		},
		err: nil,
	},
	{
		in: checkParams{
			val:        reflect.ValueOf(bool(true)),
			expectType: "bool",
		},
		err: nil,
	},
	{
		in: checkParams{
			val:        reflect.ValueOf(int(5)),
			expectType: "bool",
		},
		err: &ge.MismatchError{
			ComparedItems: "val.Type().String() vs expectType",
			Expected:      "bool",
			Actual:        "int",
		},
	},
}

func TestCheck(t *testing.T) {
	for i := 0; i < len(checkTests); i++ {
		test := &checkTests[i]

		err := Check(test.in.val, test.in.expectType)
		if err != nil {
			var eI *ge.IdentityError
			if errors.As(err, &eI) {
				if !reflect.DeepEqual(eI.Parent, test.err) {
					t.Errorf(
						"Check(%v, %v), Expected %v, Actual %v",
						test.in.val, test.in.expectType, test.err, eI.Parent,
					)
				}
			} else {
				t.Errorf(
					"Check(%v, %v), Expected %v, Actual %v",
					test.in.val, test.in.expectType, test.err, err,
				)
			}
		}
	}
}

type boolTest struct {
	in  interface{}
	out bool
	err error
}

var boolTests = []boolTest{
	{
		in:  true,
		out: true,
		err: nil,
	},
	{
		in:  5,
		out: false,
		err: &ge.MismatchError{
			ComparedItems: "val.Type().String() vs bool",
			Expected:      "bool",
			Actual:        "int",
		},
	},
}

func TestBool(t *testing.T) {
	for i := 0; i < len(boolTests); i++ {
		test := &boolTests[i]

		out, err := Bool(test.in)
		if err != nil {
			var eI *ge.IdentityError
			if errors.As(err, &eI) {
				if !reflect.DeepEqual(eI.Parent, test.err) {
					t.Errorf(
						"Bool(%v), Expected (%v, %v), Actual (%v, %v)",
						test.in, test.out, test.err, out, eI.Parent,
					)
				}
			} else {
				t.Errorf(
					"Bool(%v), Expected (%v, %v), Actual (%v, %v)",
					test.in, test.out, test.err, out, err,
				)
			}
		}

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Bool(%v), Expected (%v, %v), Actual (%v, %v)",
				test.in, test.out, test.err, out, err,
			)
		}
	}
}

type intTest struct {
	in  interface{}
	out int
	err error
}

var intTests = []intTest{
	{
		in:  5,
		out: 5,
		err: nil,
	},
	{
		in:  false,
		out: 0,
		err: &ge.MismatchError{
			ComparedItems: "val.Type().String() vs int",
			Expected:      "int",
			Actual:        "bool",
		},
	},
}

func TestInt(t *testing.T) {
	for i := 0; i < len(intTests); i++ {
		test := &intTests[i]

		out, err := Int(test.in)
		if err != nil {
			var eI *ge.IdentityError
			if errors.As(err, &eI) {
				if !reflect.DeepEqual(eI.Parent, test.err) {
					t.Errorf(
						"Int(%v), Expected (%v, %v), Actual (%v, %v)",
						test.in, test.out, test.err, out, eI.Parent,
					)
				}
			} else {
				t.Errorf(
					"Int(%v), Expected (%v, %v), Actual (%v, %v)",
					test.in, test.out, test.err, out, err,
				)
			}
		}

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Int(%v), Expected (%v, %v), Actual (%v, %v)",
				test.in, test.out, test.err, out, err,
			)
		}
	}
}

type stringTest struct {
	in  interface{}
	out string
	err error
}

var stringTests = []stringTest{
	{
		in:  "Hello",
		out: "Hello",
		err: nil,
	},
	{
		in:  false,
		out: "",
		err: &ge.MismatchError{
			ComparedItems: "val.Type().String() vs string",
			Expected:      "string",
			Actual:        "bool",
		},
	},
}

func TestString(t *testing.T) {
	for i := 0; i < len(stringTests); i++ {
		test := &stringTests[i]

		out, err := String(test.in)
		if err != nil {
			var eI *ge.IdentityError
			if errors.As(err, &eI) {
				if !reflect.DeepEqual(eI.Parent, test.err) {
					t.Errorf(
						"String(%v), Expected (%v, %v), Actual (%v, %v)",
						test.in, test.out, test.err, out, eI.Parent,
					)
				}
			} else {
				t.Errorf(
					"String(%v), Expected (%v, %v), Actual (%v, %v)",
					test.in, test.out, test.err, out, err,
				)
			}
		}

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"String(%v), Expected (%v, %v), Actual (%v, %v)",
				test.in, test.out, test.err, out, err,
			)
		}
	}
}

type isNilTest struct {
	in            interface{}
	out           bool
	simpleCompare bool
}

var (
	test          *bytes.Buffer
	testInterface io.Reader = test
	nilInterface  io.Reader
)

var isNilTests = []isNilTest{
	{
		in:            test,
		out:           true,
		simpleCompare: false,
	},
	{
		in:            testInterface,
		out:           true,
		simpleCompare: false,
	},
	{
		in:            nilInterface,
		out:           true,
		simpleCompare: true,
	},
}

func TestIsNil(t *testing.T) {
	for i := 0; i < len(isNilTests); i++ {
		test := &isNilTests[i]

		out := IsNil(test.in)
		simpleCompare := test.in == nil

		if !reflect.DeepEqual(out, test.out) && !reflect.DeepEqual(simpleCompare, test.simpleCompare) {
			t.Errorf(
				"IsNil(%#v), Expected %v, Actual %v,\n"+
					" - simpleCompare Expected %v, Actual %v",
				test.in, test.out, out,
				test.simpleCompare, simpleCompare,
			)
		}
	}
}
