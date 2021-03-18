package ge

import "fmt"

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
