package csvarser

import (
	"fmt"
)

// Ошибка Парсер не найден
//------------------------
type ParserNotFoundError struct {
	FieldType string
}

func (e *ParserNotFoundError) Error() string {
	return fmt.Sprintf("Parser not found for Type [%v]", e.FieldType)
}

// Ошибка Поле отсутсвует в СЛайсе
//--------------------------------
type FieldNotExistsInSliceError struct {
	FieldName string
	FieldType string
	Index     int
}

func (e *FieldNotExistsInSliceError) Error() string {
	return fmt.Sprintf("Field %v (%v) not Exists in Slice index=[%v]", e.FieldName, e.FieldType, e.Index)
}
