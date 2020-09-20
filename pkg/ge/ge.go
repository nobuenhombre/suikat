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
	return fmt.Sprintf("Place [%v], Error [%v]", e.Place, e.Parent.Error())
}
