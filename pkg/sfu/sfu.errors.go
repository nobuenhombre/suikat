package sfu

import "fmt"

type UnknownTypeError struct {
	Type string
}

func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type error (type = %v)", e.Type)
}
