package ge

import (
	"errors"
	"fmt"
	"reflect"
)

// en: the most common errors
// ru: самые часто встречающиеся ошибки

// NotFoundError
// en: error - key not found
// ru: ошибка - ключ не найден
type NotFoundError struct {
	Key string
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found (key = %v)", e.Key)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *NotFoundError) Is(target error) bool {
	var val *NotFoundError
	if !errors.As(target, &val) {
		return false
	}

	if val.Key != e.Key {
		return false
	}

	return true
}

// NotReleasedError
// en: error - some functionality is not implemented. for use in function stubs.
// ru: Ошибка - какой-то функционал не реализован. Для использования в заглушках функций.
type NotReleasedError struct {
	Name string
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *NotReleasedError) Error() string {
	return fmt.Sprintf("not released method (name = %v)", e.Name)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *NotReleasedError) Is(target error) bool {
	var val *NotReleasedError
	if !errors.As(target, &val) {
		return false
	}

	if val.Name != e.Name {
		return false
	}

	return true
}

// RegExpIsNotCompiledError
// en: error - the regular expression is not compiled
// ru: ошибка - регулярное выражение не компилируется
type RegExpIsNotCompiledError struct {
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *RegExpIsNotCompiledError) Error() string {
	return "regexp is not compiled"
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *RegExpIsNotCompiledError) Is(target error) bool {
	var val *RegExpIsNotCompiledError

	return errors.As(target, &val)
}

// UndefinedSwitchCaseError
// en: error - unknown variant from a known list
// ru: ошибка - неизвестный вариант из известного списка
type UndefinedSwitchCaseError struct {
	Var interface{}
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *UndefinedSwitchCaseError) Error() string {
	return fmt.Sprintf("udefined switch case [%v]", e.Var)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *UndefinedSwitchCaseError) Is(target error) bool {
	var val *UndefinedSwitchCaseError
	if !errors.As(target, &val) {
		return false
	}

	if val.Var != e.Var {
		return false
	}

	return true
}

// MismatchError
// en: error - something didn't match with something
// ru: ошибка - что-то с чем-то не совпало
type MismatchError struct {
	ComparedItems string
	Expected      interface{}
	Actual        interface{}
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *MismatchError) Error() string {
	return fmt.Sprintf("wrong %v, expected [%v], actual [%v]", e.ComparedItems, e.Expected, e.Actual)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *MismatchError) Is(target error) bool {
	var val *MismatchError
	if !errors.As(target, &val) {
		return false
	}

	if !(val.ComparedItems == e.ComparedItems && val.Expected == e.Expected && val.Actual == e.Actual) {
		return false
	}

	return true
}

// UnknownTypeError
// en: error - unknown type
// ru: ошибка - неизвестный тип
type UnknownTypeError struct {
	Type string
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("unknown type error (type = %v)", e.Type)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *UnknownTypeError) Is(target error) bool {
	var val *UnknownTypeError
	if !errors.As(target, &val) {
		return false
	}

	if val.Type != e.Type {
		return false
	}

	return true
}

// TypeAssertionError
// en: error - unknown type
// ru: ошибка - неизвестный тип
type TypeAssertionError struct {
	TargetType string
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *TypeAssertionError) Error() string {
	return fmt.Sprintf("type assertion error (target type = %v)", e.TargetType)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *TypeAssertionError) Is(target error) bool {
	var val *TypeAssertionError
	if !errors.As(target, &val) {
		return false
	}

	if val.TargetType != e.TargetType {
		return false
	}

	return true
}

// PrivateStructFieldError
// en: error - field in struct is private
// ru: ошибка - поле структуры приватное
type PrivateStructFieldError struct {
	Name string
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *PrivateStructFieldError) Error() string {
	return fmt.Sprintf("private field of struct (%v)", e.Name)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *PrivateStructFieldError) Is(target error) bool {
	var val *PrivateStructFieldError
	if !errors.As(target, &val) {
		return false
	}

	if val.Name != e.Name {
		return false
	}

	return true
}

// CantBeSetError
// en: error - you cannot write data to a field or structure because it is not a pointer
// ru: ошибка - записать данные в поле или структуру нельзя потому что это не ссылка
type CantBeSetError struct{}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *CantBeSetError) Error() string {
	return "field of structure can't be set because it's not a pointer"
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *CantBeSetError) Is(target error) bool {
	var val *CantBeSetError

	return errors.As(target, &val)
}

// LimitCountExhaustedError
// en: error - the number of attempts has been exhausted
// ru: ошибка - число попыток исчерпано
type LimitCountExhaustedError struct {
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *LimitCountExhaustedError) Error() string {
	return "the number of attempts has been exhausted"
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *LimitCountExhaustedError) Is(target error) bool {
	var val *LimitCountExhaustedError

	return errors.As(target, &val)
}

// ServiceRequiredError
// en: error - service required
// ru: ошибка - требуется сервис
type ServiceRequiredError struct {
	ServiceName string
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *ServiceRequiredError) Error() string {
	return fmt.Sprintf("service required (name = %v)", e.ServiceName)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *ServiceRequiredError) Is(target error) bool {
	var val *ServiceRequiredError
	if !errors.As(target, &val) {
		return false
	}

	if val.ServiceName != e.ServiceName {
		return false
	}

	return true
}

// UndefinedTypeError
// en: error - undefined type
// ru: ошибка - неопределенный тип
type UndefinedTypeError struct {
	Type reflect.Type
}

// Error
// en: error text formation
// ru: формирование текста ошибки
func (e *UndefinedTypeError) Error() string {
	return fmt.Sprintf("undefined type error = %v", e.Type)
}

// Is
// en: compare with target error
// ru: сравнение с другой ошибкой
func (e *UndefinedTypeError) Is(target error) bool {
	var val *UndefinedTypeError
	if !errors.As(target, &val) {
		return false
	}

	if val.Type != e.Type {
		return false
	}

	return true
}
