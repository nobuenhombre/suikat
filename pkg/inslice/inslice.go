package inslice

import (
	"fmt"
	"reflect"
)

func String(a string, list *[]string) bool {
	if list == nil {
		return false
	}

	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Int(a int, list *[]int) bool {
	if list == nil {
		return false
	}

	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Int32(a int32, list *[]int32) bool {
	if list == nil {
		return false
	}

	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Int64(a int64, list *[]int64) bool {
	if list == nil {
		return false
	}

	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Float32(a float32, list *[]float32) bool {
	if list == nil {
		return false
	}

	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Float64(a float64, list *[]float64) bool {
	if list == nil {
		return false
	}

	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func IsIndexExists(index int, list interface{}) bool {
	s := reflect.ValueOf(list)
	if s.Kind() != reflect.Slice {
		return false
	}

	if s.IsNil() {
		return false
	}

	return index > 0 && index < s.Len()
}

type IndexNotExistsError struct {
	Index int
}

func (e *IndexNotExistsError) Error() string {
	return fmt.Sprintf("index [%v] not exists", e.Index)
}
