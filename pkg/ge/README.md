# ge

Пакет `ge` — расширенная система ошибок с отслеживанием координат (Way), параметров (Params) и идентификации места происхождения.

## Основные типы

### `IdentityError`

Основная структура ошибки:

```go
type IdentityError struct {
    CreatedAt time.Time
    Message   string
    Parent    error
    Params    Params
    Way       *Way
}
```

- `CreatedAt` — время создания.
- `Message` — текстовое сообщение.
- `Parent` — родительская ошибка.
- `Params` — параметры (переменные в момент ошибки).
- `Way` — координаты в коде (пакет, функция, файл, строка).

**Методы:**
- `Error() string` — текстовое представление.
- `Unwrap() error` — поддержка `errors.Unwrap()`.
- `RootError() error` — корневая ошибка.

### `Way`

Координаты ошибки:

```go
type Way struct {
    Package string
    Caller  string
    File    string
    Line    int
}
```

**Методы:**
- `View() string` — форматированная строка: `caller package file line N`

### `Params`

```go
type Params map[string]interface{}
```

**Методы:**
- `View() string` — строковое представление: `(key = value), ...`

## Функции

### `New(message string, params ...Params) error`

Создаёт новую ошибку с сообщением.

### `Pin(parent error, params ...Params) error`

Прикрепляет родительскую ошибку.

## Встроенные ошибки

### `NotFoundError`

```go
type NotFoundError struct {
    Key string
}
```
Ошибка — ключ не найден.

### `NotReleasedError`

```go
type NotReleasedError struct {
    Name string
}
```
Ошибка — функционал не реализован (для заглушек).

### `RegExpIsNotCompiledError`

```go
type RegExpIsNotCompiledError struct{}
```
Ошибка — регулярное выражение не компилируется.

### `UndefinedSwitchCaseError`

```go
type UndefinedSwitchCaseError struct {
    Var interface{}
}
```
Ошибка — неизвестный вариант switch.

### `MismatchError`

```go
type MismatchError struct {
    ComparedItems string
    Expected      interface{}
    Actual        interface{}
}
```
Ошибка — несовпадение значений.

### `UnknownTypeError`

```go
type UnknownTypeError struct {
    Type string
}
```
Ошибка — неизвестный тип.

### `TypeAssertionError`

```go
type TypeAssertionError struct {
    TargetType string
}
```
Ошибка — ошибка type assertion.

### `PrivateStructFieldError`

```go
type PrivateStructFieldError struct {
    Name string
}
```
Ошибка — приватное поле структуры.

### `CantBeSetError`

```go
type CantBeSetError struct{}
```
Ошибка — нельзя записать в поле (не указатель).

### `LimitCountExhaustedError`

```go
type LimitCountExhaustedError struct{}
```
Ошибка — число попыток исчерпано.

### `VarRequiredError`

```go
type VarRequiredError struct {
    VarName string
}
```
Ошибка — переменная необходима.

### `HandlerRequiredError`

```go
type HandlerRequiredError struct {
    HandlerName string
}
```
Ошибка — handler необходим.

### `ServiceRequiredError`

```go
type ServiceRequiredError struct {
    ServiceName string
}
```
Ошибка — сервис необходим.

### `UndefinedTypeError`

```go
type UndefinedTypeError struct {
    Type reflect.Type
}
```
Ошибка — неопределённый тип.

## Устаревшие ошибки

### `IdentityPlaceError`

```go
type IdentityPlaceError struct {
    Place  string
    Parent error
}
```
Устаревший тип ошибки с местом.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/ge"
)

func main() {
    // --- Пример 1: Создание ошибки ---

    err := ge.New("что-то пошло не так", ge.Params{"key": "value", "count": 42})
    if err != nil {
        fmt.Printf("%v\n", err)
        // CreatedAt: 2026-05-16 12:00:00
        // Way: main() github.com/nobuenhombre/suikat/pkg/ge ge.go 42
        // Params: (key = value), (count = 42)
        // Message: что-то пошло не так
    }

    // --- Пример 2: Pin ---

    baseErr := fmt.Errorf("базовая ошибка")
    pinnedErr := ge.Pin(baseErr, ge.Params{"step": "processing"})
    if pinnedErr != nil {
        fmt.Printf("%v\n", pinnedErr)
        // CreatedAt: 2026-05-16 12:00:00
        // Way: main() ...
        // Params: (step = processing)
        // ParentError:
        //     базовая ошибка
    }

    // --- Пример 3: NotFoundError ---

    notFound := &ge.NotFoundError{Key: "user123"}
    fmt.Printf("%v\n", notFound) // not found (key = user123)

    // Проверка
    if ge.ErrorIs(notFound, &ge.NotFoundError{}) {
        fmt.Println("not found")
    }

    // --- Пример 4: MismatchError ---

    mismatch := &ge.MismatchError{
        ComparedItems: "types",
        Expected:      "int",
        Actual:        "string",
    }
    fmt.Printf("%v\n", mismatch) // wrong types, expected [int], actual [string]

    // --- Пример 5: UndefinedSwitchCaseError ---

    undefined := &ge.UndefinedSwitchCaseError{Var: "unknown"}
    fmt.Printf("%v\n", undefined) // udefined switch case [unknown]

    // --- Пример 6: RootError ---

    innerErr := fmt.Errorf("inner error")
    outerErr := ge.Pin(innerErr)
    root := outerErr.(*ge.IdentityError).RootError()
    fmt.Printf("root: %v\n", root) // inner error
}
```
