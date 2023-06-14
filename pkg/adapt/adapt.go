// Package adapt provides functions
// to check vars interface{} type
// and convert interface{} to basic types: bool, int, string.
package adapt

import (
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

// Check compares reflect.Value with a string representation of the expected type.
// If the type does not match the expected one, the ge.MismatchError{} error is returned
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

// Bool convert v interface{} into bool.
// If v can't be converted - return ge.MismatchError{}
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

// Int convert v interface{} into int.
// If v can't be converted - return ge.MismatchError{}
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

// String convert v interface{} into string.
// If v can't be converted - return ge.MismatchError{}
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

// IsNil check interface{} is nil.
// and return bool value
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	// nolint: exhaustive
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}
