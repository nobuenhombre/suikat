package sfu

import (
	"reflect"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

// Tag Examples
//==============================================================
//type SomeFormData struct {
//	Path           string      `form:"path"`
//	Port           int         `form:"port"`
//	Coefficient    float64     `form:"coefficient"`
//	MakeSomeAction bool        `form:"msa"`
//  OtherStruct    OtherStruct `form:"otherStruct"`
//}

type FieldInfo struct {
	Type  reflect.Type
	Name  string
	Value reflect.Value
}

type TagInfo struct {
	Tag string
}

func NewTagProcessor() refavour.TagProcessor {
	return &TagInfo{
		Tag: "form",
	}
}

func (tag *TagInfo) GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error) {
	tagData := typeField.Tag.Get(tag.Tag)

	return &FieldInfo{
		Type:  valueField.Type(),
		Name:  tagData,
		Value: valueField,
	}, nil
}
