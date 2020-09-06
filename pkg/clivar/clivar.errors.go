package clivar

import "fmt"

// Ошибка - Неизвестный тип значения
//--------------------------------------------------
type UnknownValueTypeError struct {
	ValueType string
}

func (e *UnknownValueTypeError) Error() string {
	return fmt.Sprintf("unknown value type: %v", e.ValueType)
}
