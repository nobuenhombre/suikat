package ge

import (
	"fmt"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/inslice"
)

func New(message string, params ...Params) error {
	var p Params

	if inslice.IsIndexExists(0, params) {
		p = params[0]
	}

	return &IdentityError{
		Message: message,
		Params:  p,
		Way:     getWay(),
	}
}

func Pin(parent error, params ...Params) error {
	var p Params

	if inslice.IsIndexExists(0, params) {
		p = params[0]
	}

	return &IdentityError{
		Parent: parent,
		Params: p,
		Way:    getWay(),
	}
}

type IdentityError struct {
	Message string
	Parent  error
	Params  Params
	Way     *Way
}

func (e *IdentityError) Unwrap(err error) error {
	return e.Parent
}

func (e *IdentityError) Error() string {
	wayStr := ""
	if e.Way != nil {
		wayStr = fmt.Sprintf("Way[ %v ], ", e.Way.View())
	}

	paramsStr := ""
	if e.Params != nil {
		paramsStr = fmt.Sprintf("Params[ %v ]", e.Params.View())
	}

	parentStr := ""
	if e.Parent != nil {
		parentStr = fmt.Sprintf("ParentError[ %v ], ", e.Parent.Error())
	}

	messageStr := ""
	if len(e.Message) > 0 {
		messageStr = fmt.Sprintf("Message[ %v ], ", e.Message)
	}

	return strings.TrimSuffix(fmt.Sprintf("%v%v%v%v", wayStr, paramsStr, parentStr, messageStr), ", ")
}
