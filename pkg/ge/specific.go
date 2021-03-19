package ge

import "fmt"

type NotReleasedError struct {
	Name string
}

func (e *NotReleasedError) Error() string {
	return fmt.Sprintf("not released method (name = %v)", e.Name)
}

type RegExpIsNotCompiledError struct {
}

func (e *RegExpIsNotCompiledError) Error() string {
	return "regexp is not compiled"
}

type UndefinedSwitchCaseError struct {
	Var interface{}
}

func (e *UndefinedSwitchCaseError) Error() string {
	return fmt.Sprintf("udefined switch case [%v]", e.Var)
}

type MismatchError struct {
	ComparedItems string
	Expected      interface{}
	Actual        interface{}
}

func (e *MismatchError) Error() string {
	return fmt.Sprintf("wrong %v, expected [%v], actual [%v]", e.ComparedItems, e.Expected, e.Actual)
}
