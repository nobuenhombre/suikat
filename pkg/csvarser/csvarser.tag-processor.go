package csvarser

import (
	"reflect"
	"strconv"

	"github.com/nobuenhombre/suikat/pkg/refavour"
)

// Tag Examples
//==============================================================
//type SomeCSV struct {
//	A pgtype.Int4    `order:"0"`
//	B pgtype.Int8    `order:"1"`
//	C pgtype.Varchar `order:"2"`
//}

type FieldInfo struct {
	Type  reflect.Type
	Order int
}

type TagInfo struct {
	Tag string
}

func NewTagProcessor() refavour.TagProcessor {
	return &TagInfo{
		Tag: "order",
	}
}

func (tag *TagInfo) GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error) {
	tagData, tagDataErr := strconv.Atoi(typeField.Tag.Get(tag.Tag))
	if tagDataErr != nil {
		return nil, tagDataErr
	}

	return &FieldInfo{
		Type:  valueField.Type(),
		Order: tagData,
	}, nil
}
