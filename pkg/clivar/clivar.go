// Package clivar provides syntactic sugar for working with the flag package.
// Allows you to describe using structure tags how to use flag to fill in this structure.
package clivar

import (
	"flag"
	"reflect"
	"strconv"

	"github.com/nobuenhombre/suikat/pkg/ge"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

// CliVar describe tag for struct receiver
type CliVar struct {
	Key          string
	Description  string
	DefaultValue interface{}
}

// GetString read string value from flag
func (cli *CliVar) GetString() *string {
	return flag.String(
		cli.Key,
		cli.DefaultValue.(string),
		cli.Description,
	)
}

// GetInt read int value from flag
func (cli *CliVar) GetInt() *int {
	return flag.Int(
		cli.Key,
		cli.DefaultValue.(int),
		cli.Description,
	)
}

// GetFloat64 read float64 value from flag
func (cli *CliVar) GetFloat64() *float64 {
	return flag.Float64(
		cli.Key,
		cli.DefaultValue.(float64),
		cli.Description,
	)
}

// GetBool read bool value from flag
func (cli *CliVar) GetBool() *bool {
	return flag.Bool(
		cli.Key,
		cli.DefaultValue.(bool),
		cli.Description,
	)
}

// Load field of target struct from flag like described in tags
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
			return ge.Pin(&ge.UndefinedSwitchCaseError{
				Var: fieldInfo.(*FieldInfo).ValueType,
			})
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
			return ge.Pin(&ge.UndefinedSwitchCaseError{
				Var: fieldInfo.(*FieldInfo).ValueType,
			})
		}
	}

	return nil
}
