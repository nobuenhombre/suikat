// Package clivar provides syntactic sugar for working with the flag package.
// Allows you to describe using structure tags how to use flag to fill in this structure.
package clivar

import (
	"flag"

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
		err := fieldInfo.(*FieldInfo).readFlag(fieldName, tempMap)
		if err != nil {
			return ge.Pin(err)
		}
	}

	flag.Parse()

	reflectValue := refavour.GetReflectValue(structData)

	for fieldName, fieldInfo := range structureFields {
		err := fieldInfo.(*FieldInfo).fillStructField(fieldName, tempMap, reflectValue)
		if err != nil {
			return ge.Pin(err)
		}
	}

	return nil
}
