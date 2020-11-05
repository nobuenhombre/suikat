package clivar

import (
	"reflect"
	"strings"

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
		return nil, &refavour.InvalidTagError{
			Actual:   tagData,
			Expected: TagCliExample,
		}
	}

	nameType := partsTagData[0]
	valueStr := partsTagData[1]

	partsNameType := strings.Split(nameType, ":")
	if len(partsNameType) != CountPartsNameType {
		return nil, &refavour.InvalidTagError{
			Actual:   tagData,
			Expected: TagCliExample,
		}
	}

	nameDescription := partsNameType[0]
	valueType := partsNameType[1]

	partsNameDescription := strings.Split(nameDescription, "[")
	if len(partsNameDescription) != CountPartsNameDescription {
		return nil, &refavour.InvalidTagError{
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
