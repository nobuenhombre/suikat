package ge

import "fmt"

type NotReleasedError struct {
	Name string
}

func (e *NotReleasedError) Error() string {
	return fmt.Sprintf("Not Released Method (name = %v)", e.Name)
}

type IdentityPlaceError struct {
	Place  string
	Parent error
}

func (e *IdentityPlaceError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf("Place [%v], Error [%v]", e.Place, e.Parent.Error())
	}

	return fmt.Sprintf("Place [%v]", e.Place)
}

type IdentityParams map[string]interface{}

type IdentityError struct {
	Package string
	Caller  string
	Place   string
	Params  map[string]interface{}
	Message string
	Parent  error
}

func (e *IdentityError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf(
			"Package[%v].Caller[%v].Place[%v].Params[%#v].Message[%v].Error[%v]",
			e.Package, e.Caller, e.Place, e.Params, e.Message, e.Parent.Error(),
		)
	}

	return fmt.Sprintf(
		"Package[%v].Caller[%v].Place[%v].Params[%#v].Message[%v]",
		e.Package, e.Caller, e.Place, e.Params, e.Message,
	)
}
