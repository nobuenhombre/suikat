package clivar

import (
	"flag"
	"reflect"
	"strconv"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

type CliVar struct {
	Key          string
	Description  string
	DefaultValue interface{}
}

func (cli *CliVar) GetString() *string {
	return flag.String(
		cli.Key,
		cli.DefaultValue.(string),
		cli.Description,
	)
}

func (cli *CliVar) GetInt() *int {
	return flag.Int(
		cli.Key,
		cli.DefaultValue.(int),
		cli.Description,
	)
}

func (cli *CliVar) GetFloat64() *float64 {
	return flag.Float64(
		cli.Key,
		cli.DefaultValue.(float64),
		cli.Description,
	)
}

func (cli *CliVar) GetBool() *bool {
	return flag.Bool(
		cli.Key,
		cli.DefaultValue.(bool),
		cli.Description,
	)
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

	tempMap := make(map[string]interface{})

	for fieldName, fieldInfo := range structureFields {
		ev := CliVar{
			Key:          fieldInfo.(*FieldInfo).Name,
			Description:  fieldInfo.(*FieldInfo).Description,
			DefaultValue: nil,
		}

		var err error

		switch fieldInfo.(*FieldInfo).ValueType {
		case "string":
			ev.DefaultValue = fieldInfo.(*FieldInfo).DefaultValue
			tempMap[fieldName] = ev.GetString()

		case "int":
			ev.DefaultValue, err = strconv.Atoi(fieldInfo.(*FieldInfo).DefaultValue)
			if err != nil {
				return err
			}

			tempMap[fieldName] = ev.GetInt()

		case "float64":
			ev.DefaultValue, err = strconv.ParseFloat(fieldInfo.(*FieldInfo).DefaultValue, 64)
			if err != nil {
				return err
			}

			tempMap[fieldName] = ev.GetFloat64()

		case "bool":
			ev.DefaultValue, err = strconv.ParseBool(fieldInfo.(*FieldInfo).DefaultValue)
			if err != nil {
				return err
			}

			tempMap[fieldName] = ev.GetBool()

		default:
			return &UnknownValueTypeError{
				ValueType: fieldInfo.(*FieldInfo).ValueType,
			}
		}
	}

	flag.Parse()

	reflectValue := refavour.GetReflectValue(structData)

	for fieldName, fieldInfo := range structureFields {
		switch fieldInfo.(*FieldInfo).ValueType {
		case "string":
			reflectValue.FieldByName(fieldName).Set(reflect.ValueOf(*(tempMap[fieldName].(*string))))

		case "int":
			reflectValue.FieldByName(fieldName).Set(reflect.ValueOf(*(tempMap[fieldName].(*int))))

		case "float64":
			reflectValue.FieldByName(fieldName).Set(reflect.ValueOf(*(tempMap[fieldName].(*float64))))

		case "bool":
			reflectValue.FieldByName(fieldName).Set(reflect.ValueOf(*(tempMap[fieldName].(*bool))))

		default:
			return &UnknownValueTypeError{
				ValueType: fieldInfo.(*FieldInfo).ValueType,
			}
		}
	}

	return nil
}
