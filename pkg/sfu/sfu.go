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
			value := fieldInfo.(*FieldInfo).Value
			kind := value.Kind()
			data := value.Addr().Interface()

			switch kind {
			case reflect.Struct:
				err := Convert(data, n, form)
				if err != nil {
					return err
				}

			case reflect.Slice:
				for i := 0; i < value.Len(); i++ {
					sliceItem := value.Index(i)
					name := fmt.Sprintf("%v[%v]", n, i)
					ts := sliceItem.Type().String()

					switch ts {
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
						slKind := sliceItem.Kind()
						slData := sliceItem.Addr().Interface()

						switch slKind {
						case reflect.Struct:
							err := Convert(slData, name, form)
							if err != nil {
								return err
							}

						default:
							return &UnknownTypeError{
								Type: ts,
							}
						}
					}
				}

			default:
				return &UnknownTypeError{
					Type: t,
				}
			}
		}
	}

	return nil
}
