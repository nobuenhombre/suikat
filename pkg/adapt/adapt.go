package adapt

import (
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

func Check(val reflect.Value, expectType string) error {
	if val.Type().String() != expectType {
		return ge.Pin(&ge.MismatchError{
			ComparedItems: "val.Type().String() vs expectType",
			Expected:      expectType,
			Actual:        val.Type().String(),
		})
	}

	return nil
}

func Bool(v interface{}) (bool, error) {
	val := reflect.ValueOf(v)

	result, ok := val.Interface().(bool)
	if !ok {
		return false, ge.Pin(&ge.MismatchError{
			ComparedItems: "val.Type().String() vs bool",
			Expected:      "bool",
			Actual:        val.Type().String(),
		})
	}

	return result, nil
}

func Int(v interface{}) (int, error) {
	val := reflect.ValueOf(v)

	result, ok := val.Interface().(int)
	if !ok {
		return 0, ge.Pin(&ge.MismatchError{
			ComparedItems: "val.Type().String() vs int",
			Expected:      "int",
			Actual:        val.Type().String(),
		})
	}

	return result, nil
}

func String(v interface{}) (string, error) {
	val := reflect.ValueOf(v)

	result, ok := val.Interface().(string)
	if !ok {
		return "", ge.Pin(&ge.MismatchError{
			ComparedItems: "val.Type().String() vs string",
			Expected:      "string",
			Actual:        val.Type().String(),
		})
	}

	return result, nil
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	// nolint: exhaustive
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}
