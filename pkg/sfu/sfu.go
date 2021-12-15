package sfu

//-----------------------------------
//
// Convert struct to form url encoded
// to use with less-rest-client
//
//-----------------------------------

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/ge"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

func Convert(structData interface{}, parent string, form *url.Values) (err error) {
	tagProcessor := NewTagProcessor()

	structError := refavour.CheckStructure(structData)
	if structError != nil {
		return structError
	}

	canBeChangedError := refavour.CheckCanBeChanged(structData)
	if canBeChangedError != nil {
		return canBeChangedError
	}

	structureFields, getStructErr := refavour.GetStructureFieldsTypes(structData, tagProcessor)
	if getStructErr != nil {
		return getStructErr
	}

	for _, fieldInfo := range structureFields {
		t := fieldInfo.(*FieldInfo).Type.String()

		n := fieldInfo.(*FieldInfo).Name
		if len(parent) != 0 {
			n = fmt.Sprintf("%v[%v]", parent, fieldInfo.(*FieldInfo).Name)
		}

		value := fieldInfo.(*FieldInfo).Value

		if !value.CanInterface() {
			structValue := refavour.GetReflectValue(structData)

			return &ge.PrivateStructFieldError{
				Name: structValue.Type().String(),
			}
		}

		err := convertSimpleTypes(t, n, value, form)
		if err != nil {
			return ge.Pin(err)
		}
	}

	return nil
}

func convertSimpleTypes(itemType string, n string, value reflect.Value, form *url.Values) (err error) {
	switch itemType {
	case "string":
		v := value.String()
		form.Add(n, v)

	case "int64":
		vi := value.Int()
		v := fmt.Sprintf("%v", vi)
		form.Add(n, v)

	case "float64":
		vf := value.Float()
		v := fmt.Sprintf("%v", vf)
		form.Add(n, v)

	case "bool":
		vb := value.Bool()
		v := fmt.Sprintf("%v", vb)
		form.Add(n, v)

	default:
		err := convertComplexTypes(itemType, n, value, form)
		if err != nil {
			return ge.Pin(err)
		}
	}

	return nil
}

func convertComplexTypes(itemType string, n string, value reflect.Value, form *url.Values) (err error) {
	kind := value.Kind()

	switch kind {
	case reflect.Struct:
		data := value.Addr().Interface()

		err := Convert(data, n, form)
		if err != nil {
			return ge.Pin(err)
		}

	case reflect.Slice:
		err := convertSlice(n, value, form)
		if err != nil {
			return ge.Pin(err)
		}

	default:
		return &ge.UnknownTypeError{
			Type: itemType,
		}
	}

	return nil
}

func convertSlice(n string, value reflect.Value, form *url.Values) (err error) {
	for i := 0; i < value.Len(); i++ {
		sliceItem := value.Index(i)
		name := fmt.Sprintf("%v[%v]", n, i)
		ts := sliceItem.Type().String()

		err := convertSliceItem(ts, name, sliceItem, form)
		if err != nil {
			return ge.Pin(err)
		}
	}

	return nil
}

func convertSliceItem(itemType string, name string, sliceItem reflect.Value, form *url.Values) (err error) {
	switch itemType {
	case "string":
		v := sliceItem.String()
		form.Add(name, v)

	case "int64":
		vi := sliceItem.Int()
		v := fmt.Sprintf("%v", vi)
		form.Add(name, v)

	case "float64":
		vf := sliceItem.Float()
		v := fmt.Sprintf("%v", vf)
		form.Add(name, v)

	case "bool":
		vb := sliceItem.Bool()
		v := fmt.Sprintf("%v", vb)
		form.Add(name, v)

	default:
		err := convertSliceItemStruct(itemType, name, sliceItem, form)
		if err != nil {
			return ge.Pin(err)
		}
	}

	return nil
}

func convertSliceItemStruct(itemType string, name string, sliceItem reflect.Value, form *url.Values) (err error) {
	slKind := sliceItem.Kind()

	switch slKind {
	case reflect.Struct:
		slData := sliceItem.Addr().Interface()

		err := Convert(slData, name, form)
		if err != nil {
			return ge.Pin(err)
		}

	default:
		return &ge.UnknownTypeError{
			Type: itemType,
		}
	}

	return nil
}
