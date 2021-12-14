// Package ge provides
// en: error handling functions, allows you to check the original type of error and find its exact position in the code
// ru: функции работы с ошибками, позволяет проверить исходный тип ошибки и найти ее точное положение в коде
package ge

import (
	"fmt"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/inslice"
)

// IdentityError
// en: Description of the error identifying its place of origin
// ru: Описание ошибки идентифицирующей место ее происхождения
type IdentityError struct {
	// Message
	// en: error text message
	// ru: текстовое сообщение ошибки
	Message string

	// Parent
	// en: parental error
	// ru: родительская ошибка
	Parent error

	// Params
	// en: error parameters - variables in the presence of which the error occurred
	// - will help you understand what happened
	// ru: параметры ошибки - переменные в присутствии которых произошла ошибка - помогут понять что произошло
	Params Params

	// Way
	// en: path - shows the specific place in the code where the error occurred
	// ru: путь - показывает конкретное место в коде, где произошла ошибка
	Way *Way
}

// New
// en: Creating a new error
// ru: Создание новой ошибки
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

// Pin
// en: Attach the error that occurred
// ru: Прикрепить произошедшую ошибку
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

// Unwrap
// en: support errors Wrapper interface
// ru: поддержка интерфейса Wrapper для ошибок
func (e *IdentityError) Unwrap() error {
	return e.Parent
}

// Error
// en: error text formation
// ru: формирование текста ошибки
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
