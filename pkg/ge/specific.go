package ge

import "fmt"

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
