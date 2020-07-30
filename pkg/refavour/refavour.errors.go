package refavour

import (
	"fmt"
)

// Ошибка - Kind не совпадает
//---------------------------
type KindNotMatchedError struct {
	Expected string
	Actual   string
}

func (e *KindNotMatchedError) Error() string {
	return fmt.Sprintf("Kind not matched, Expect: %v, Actual: %v", e.Expected, e.Actual)
}

// Ошибка - Поле структуры не может быть установлено
//--------------------------------------------------
type CantBeSetError struct{}

func (e *CantBeSetError) Error() string {
	return "Field of Structure can't be set because it's not a Pointer"
}

// Ошибка - неправильный формат Тега
//--------------------------------------------------
type InvalidTagError struct {
	Actual   string
	Expected string
}

func (e *InvalidTagError) Error() string {
	return fmt.Sprintf("invalid tag: actual[%v], expected[%v]", e.Actual, e.Expected)
}
