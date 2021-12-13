// Package chacha provides primitives and functions
// to create a complex tree of checks when one check depends on another check
package chacha

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

// Validator describes the validation function
type Validator func(params ...interface{}) (interface{}, error)

// DontCheckChildrens a special marker that stops checking child branches
type DontCheckChildrens struct{}

// Tree describes the validation tree
type Tree struct {
	Validator Validator
	Data      interface{}
	Result    error
	Valid     bool
	Childrens []Tree
}

// CheckChildrens returns a flag - whether it is necessary to perform checks of child branches
func (t *Tree) CheckChildrens() bool {
	checkChildrens := true

	if t.Data != nil {
		valData := reflect.ValueOf(t.Data)
		checkChildrens = valData.Type().String() != "*chacha.DontCheckChildrens"
	}

	return checkChildrens
}

// Validate Recursive function - traverses all branches of the validation tree and calls the validator function.
// The result of the check is saved in the Valid field.
func (t *Tree) Validate(params ...interface{}) {
	t.Data, t.Result = t.Validator(params...)
	if t.Result != nil {
		t.Valid = false
	} else {
		t.Valid = true
		if t.CheckChildrens() {
			for index := range t.Childrens {
				params[0] = t.Data
				t.Childrens[index].Validate(params...)
				t.Valid = t.Valid && t.Childrens[index].Valid
			}
		}
	}
}

// ShowErrors display validation errors on the screen
func (t *Tree) ShowErrors() {
	if t.Result != nil {
		invalidColor := color.New(color.FgRed).SprintFunc()
		fmt.Printf(" %v %v\n", invalidColor("♥"), t.Result)
	}

	if t.CheckChildrens() {
		for index := range t.Childrens {
			t.Childrens[index].ShowErrors()
		}
	}
}
