// Package refavour provides interface to work with tags of struct and several functions for working with reflect
package refavour

import (
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

// TagProcessor interface to provide parsing tags of struct
type TagProcessor interface {
	GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error)
}

// GetReflectValue return reflect.Value of interface{}
func GetReflectValue(s interface{}) reflect.Value {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	return val
}

// CheckKind check is that interface{} equal expected Kind
// if equal then return nil
// else return ge.MismatchError
func CheckKind(value interface{}, expectedKind reflect.Kind) error {
	reflectValue := GetReflectValue(value)
	reflectKind := reflectValue.Kind()

	if reflectKind != expectedKind {
		return &ge.MismatchError{
			Expected: expectedKind.String(),
			Actual:   reflectKind.String(),
		}
	}

	return nil
}

// CheckStructure check is that interface{} struct
// if struct then return nil
// else return ge.MismatchError
func CheckStructure(data interface{}) error {
	return CheckKind(data, reflect.Struct)
}

// CheckMap check is that interface{} map
// if map then return nil
// else return ge.MismatchError
func CheckMap(data interface{}) error {
	return CheckKind(data, reflect.Map)
}

// CheckSlice check is that interface{} slice
// if slice then return nil
// else return ge.MismatchError
func CheckSlice(data interface{}) error {
	return CheckKind(data, reflect.Slice)
}

// CheckCanBeChanged Check whether the data receiver can be changed
func CheckCanBeChanged(value interface{}) error {
	reflectValue := GetReflectValue(value)
	if !reflectValue.CanSet() {
		return &ge.CantBeSetError{}
	}

	return nil
}

// FieldsInfo Types of structure fields
type FieldsInfo map[string]interface{}

// GetStructureFieldsTypes reads the structure from the interface and generates a list of field names and their types
func GetStructureFieldsTypes(value interface{}, tagProcessor TagProcessor) (FieldsInfo, error) {
	reflectValue := GetReflectValue(value)
	reflectedValueType := reflectValue.Type()

	result := make(FieldsInfo)

	for i := 0; i < reflectValue.NumField(); i++ {
		valueField := reflectValue.Field(i)
		typeField := reflectedValueType.Field(i)

		fieldInfo, err := tagProcessor.GetFieldInfo(typeField, valueField)
		if err != nil {
			return nil, err
		}

		result[typeField.Name] = fieldInfo
	}

	return result, nil
}
