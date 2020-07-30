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

type CSVFieldInfo struct {
	Type  reflect.Type
	Order int
}

type CSVTag struct {
	Tag string
}

func NewTagProcessor() refavour.TagProcessor {
	return &CSVTag{
		Tag: "order",
	}
}

func (tag *CSVTag) GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error) {
	tagData, tagDataErr := strconv.Atoi(typeField.Tag.Get(tag.Tag))
	if tagDataErr != nil {
		return nil, tagDataErr
	}

	return &CSVFieldInfo{
		Type:  valueField.Type(),
		Order: tagData,
	}, nil
}
