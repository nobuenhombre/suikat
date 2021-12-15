package envar

import (
	"reflect"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

const (
	TagEnvExample      = "NAME:valueType=defaultValue"
	CountPartsTagData  = 2
	CountPartsNameType = 2
)

// Tag Examples
//==============================================================
//type AppConfig struct {
//	Path           string  `env:"PATH:string=/some/default/path"`
//	Port           int     `env:"PORT:int=8080"`
//	Coefficient    float64 `env:"COEFFICIENT:float64=75.31"`
//	MakeSomeAction bool    `env:"MSA:bool=false"`
//}

type FieldInfo struct {
	Type         reflect.Type
	Name         string
	ValueType    string
	DefaultValue string
}

type TagInfo struct {
	Tag string
}

func NewTagProcessor() refavour.TagProcessor {
	return &TagInfo{
		Tag: "env",
	}
}

func (tag *TagInfo) GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error) {
	tagData := typeField.Tag.Get(tag.Tag)

	partsTagData := strings.Split(tagData, "=")
	if len(partsTagData) != CountPartsTagData {
		return nil, &ge.MismatchError{
			Actual:   tagData,
			Expected: TagEnvExample,
		}
	}

	nameType := partsTagData[0]
	valueStr := partsTagData[1]

	partsNameType := strings.Split(nameType, ":")
	if len(partsNameType) != CountPartsNameType {
		return nil, &ge.MismatchError{
			Actual:   tagData,
			Expected: TagEnvExample,
		}
	}

	name := partsNameType[0]
	valueType := partsNameType[1]

	return &FieldInfo{
		Type:         valueField.Type(),
		Name:         name,
		ValueType:    valueType,
		DefaultValue: valueStr,
	}, nil
}
