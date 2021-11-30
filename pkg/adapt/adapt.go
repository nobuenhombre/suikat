package adapt

import (
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Config struct {
}

type Service interface {
	Check(val reflect.Value, expectType string) error
	Bool(v interface{}) (bool, error)
	Int(v interface{}) (int, error)
	String(v interface{}) (string, error)
}

func New() Service {
	return &Config{}
}

func (c *Config) Check(val reflect.Value, expectType string) error {
	if val.Type().String() != expectType {
		return ge.Pin(&ge.MismatchError{
			ComparedItems: "val.Type().String() vs expectType",
			Expected:      expectType,
			Actual:        val.Type().String(),
		})
	}

	return nil
}

func (c *Config) Bool(v interface{}) (bool, error) {
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

func (c *Config) Int(v interface{}) (int, error) {
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

func (c *Config) String(v interface{}) (string, error) {
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
