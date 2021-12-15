package clivar

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

const (
	TagCliExample             = "NAME[description key]:valueType=defaultValue"
	CountPartsTagData         = 2
	CountPartsNameType        = 2
	CountPartsNameDescription = 2
)

// Tag Examples
//==============================================================
//type CommandLineParams struct {
//	Path           string  `cli:"PATH[Path to file]:string=/some/default/path"`
//	Port           int     `cli:"PORT[Port for server]:int=8080"`
//	Coefficient    float64 `cli:"COEFFICIENT[Coefficient transmutation]:float64=75.31"`
//	MakeSomeAction bool    `cli:"MSA[Do make some action?]:bool=false"`
//}

type FieldInfo struct {
	Type         reflect.Type
	Name         string
	ValueType    string
	DefaultValue string
	Description  string
}

func (f *FieldInfo) readFlag(name string, store map[string]interface{}) (err error) {
	ev := CliVar{
		Key:          f.Name,
		Description:  f.Description,
		DefaultValue: nil,
	}

	switch f.ValueType {
	case "string":
		ev.DefaultValue = f.DefaultValue
		store[name] = ev.GetString()

	case "int":
		ev.DefaultValue, err = strconv.Atoi(f.DefaultValue)
		if err != nil {
			return err
		}

		store[name] = ev.GetInt()

	case "float64":
		ev.DefaultValue, err = strconv.ParseFloat(f.DefaultValue, 64)
		if err != nil {
			return err
		}

		store[name] = ev.GetFloat64()

	case "bool":
		ev.DefaultValue, err = strconv.ParseBool(f.DefaultValue)
		if err != nil {
			return err
		}

		store[name] = ev.GetBool()

	default:
		return ge.Pin(&ge.UndefinedSwitchCaseError{
			Var: f.ValueType,
		})
	}

	return nil
}

func (f *FieldInfo) fillStructField(name string, store map[string]interface{}, data reflect.Value) (err error) {
	switch f.ValueType {
	case "string":
		data.FieldByName(name).Set(reflect.ValueOf(*(store[name].(*string))))

	case "int":
		data.FieldByName(name).Set(reflect.ValueOf(*(store[name].(*int))))

	case "float64":
		data.FieldByName(name).Set(reflect.ValueOf(*(store[name].(*float64))))

	case "bool":
		data.FieldByName(name).Set(reflect.ValueOf(*(store[name].(*bool))))

	default:
		return ge.Pin(&ge.UndefinedSwitchCaseError{
			Var: f.ValueType,
		})
	}

	return nil
}

type TagInfo struct {
	Tag string
}

func NewTagProcessor() refavour.TagProcessor {
	return &TagInfo{
		Tag: "cli",
	}
}

func (tag *TagInfo) GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error) {
	tagData := typeField.Tag.Get(tag.Tag)

	partsTagData := strings.Split(tagData, "=")
	if len(partsTagData) != CountPartsTagData {
		return nil, &ge.MismatchError{
			Actual:   tagData,
			Expected: TagCliExample,
		}
	}

	nameType := partsTagData[0]
	valueStr := partsTagData[1]

	partsNameType := strings.Split(nameType, ":")
	if len(partsNameType) != CountPartsNameType {
		return nil, &ge.MismatchError{
			Actual:   tagData,
			Expected: TagCliExample,
		}
	}

	nameDescription := partsNameType[0]
	valueType := partsNameType[1]

	partsNameDescription := strings.Split(nameDescription, "[")
	if len(partsNameDescription) != CountPartsNameDescription {
		return nil, &ge.MismatchError{
			Actual:   tagData,
			Expected: TagCliExample,
		}
	}

	name := partsNameDescription[0]
	description := strings.TrimRight(partsNameDescription[1], "]")

	return &FieldInfo{
		Type:         valueField.Type(),
		Name:         name,
		ValueType:    valueType,
		DefaultValue: valueStr,
		Description:  description,
	}, nil
}
