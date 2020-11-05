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

		switch t {
		case "string":
			v := fieldInfo.(*FieldInfo).Value.String()
			form.Add(n, v)

		case "int64":
			vi := fieldInfo.(*FieldInfo).Value.Int()
			v := fmt.Sprintf("%v", vi)
			form.Add(n, v)

		case "float64":
			vf := fieldInfo.(*FieldInfo).Value.Float()
			v := fmt.Sprintf("%v", vf)
			form.Add(n, v)

		case "bool":
			vb := fieldInfo.(*FieldInfo).Value.Bool()
			v := fmt.Sprintf("%v", vb)
			form.Add(n, v)

		default:
			if fieldInfo.(*FieldInfo).Value.Kind() == reflect.Struct {
				err := Convert(fieldInfo.(*FieldInfo).Value.Addr().Interface(), n, form)
				if err != nil {
					return err
				}
			} else {
				return &UnknownTypeError{
					Type: t,
				}
			}
		}
	}

	return nil
}
