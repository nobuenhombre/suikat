package envar

import (
	"os"
	"reflect"
	"strconv"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

type EnvVar struct {
	Key          string
	DefaultValue interface{}
}

func (ev *EnvVar) GetString() string {
	if value, exists := os.LookupEnv(ev.Key); exists {
		return value
	}

	return ev.DefaultValue.(string)
}

func (ev *EnvVar) GetInt() int {
	if valueStr, exists := os.LookupEnv(ev.Key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}

	return ev.DefaultValue.(int)
}

func (ev *EnvVar) GetFloat64() float64 {
	if valueStr, exists := os.LookupEnv(ev.Key); exists {
		if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
			return value
		}
	}

	return ev.DefaultValue.(float64)
}

func (ev *EnvVar) GetBool() bool {
	if valueStr, exists := os.LookupEnv(ev.Key); exists {
		if value, err := strconv.ParseBool(valueStr); err == nil {
			return value
		}
	}

	return ev.DefaultValue.(bool)
}

func Load(structData interface{}) error {
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

	reflectValue := refavour.GetReflectValue(structData)

	for fieldName, fieldInfo := range structureFields {
		ev := EnvVar{
			Key:          fieldInfo.(*FieldInfo).Name,
			DefaultValue: nil,
		}

		var (
			value interface{}
			err   error
		)

		switch fieldInfo.(*FieldInfo).ValueType {
		case "string":
			ev.DefaultValue = fieldInfo.(*FieldInfo).DefaultValue
			value = ev.GetString()

		case "int":
			ev.DefaultValue, err = strconv.Atoi(fieldInfo.(*FieldInfo).DefaultValue)
			if err != nil {
				return err
			}

			value = ev.GetInt()

		case "float64":
			ev.DefaultValue, err = strconv.ParseFloat(fieldInfo.(*FieldInfo).DefaultValue, 64)
			if err != nil {
				return err
			}

			value = ev.GetFloat64()

		case "bool":
			ev.DefaultValue, err = strconv.ParseBool(fieldInfo.(*FieldInfo).DefaultValue)
			if err != nil {
				return err
			}

			value = ev.GetBool()

		default:
			return &UnknownValueTypeError{
				ValueType: fieldInfo.(*FieldInfo).ValueType,
			}
		}

		reflectValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}

	return nil
}
