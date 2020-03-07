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
type List struct {
	Validator Validator
	Data      interface{}
	Result    error
	Valid     bool
	Childrens []List
}

// Проверять детей
func (l *List) CheckChildrens() bool {
	checkChildrens := true
	if l.Data != nil {
		valData := reflect.ValueOf(l.Data)
		checkChildrens = valData.Type().String() != "*chacha.DontCheckChildrens"
	}
	return checkChildrens
}

// Валидировать рекурсивно дерево
func (l *List) Validate(params ...interface{}) {
	l.Data, l.Result = l.Validator([]interface{}(params)...)
	if l.Result != nil {
		l.Valid = false
	} else {
		l.Valid = true
		if l.CheckChildrens() {
			for index := range l.Childrens {
				params[0] = l.Data
				l.Childrens[index].Validate([]interface{}(params)...)
				l.Valid = l.Valid && l.Childrens[index].Valid
			}
		}
	}
}

func (l *List) ShowErrors() {
	if l.Result != nil {
		invalidColor := color.New(color.FgRed).SprintFunc()
		fmt.Printf(" %v %v\n", invalidColor("♥"), l.Result)
	}
	if l.CheckChildrens() {
		for index := range l.Childrens {
			l.Childrens[index].ShowErrors()
		}
	}
}
