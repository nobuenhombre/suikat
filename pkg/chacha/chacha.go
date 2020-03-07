package chacha

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

// Тип проверяльщик
type Validator func(params ...interface{}) (interface{}, error)

type DontCheckChildrens struct{}

// Дерево проверок
type Tree struct {
	Validator Validator
	Data      interface{}
	Result    error
	Valid     bool
	Childrens []Tree
}

// Проверять детей
func (t *Tree) CheckChildrens() bool {
	checkChildrens := true

	if t.Data != nil {
		valData := reflect.ValueOf(t.Data)
		checkChildrens = valData.Type().String() != "*chacha.DontCheckChildrens"
	}

	return checkChildrens
}

// Валидировать рекурсивно дерево
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
