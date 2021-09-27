package sfu

import "fmt"

type UnknownTypeError struct {
	Type string
}

func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type error (type = %v)", e.Type)
}

type PrivateStructFieldError struct {
	Name string
}

func (e *PrivateStructFieldError) Error() string {
	return fmt.Sprintf("private field of struct (%v)", e.Name)
}
