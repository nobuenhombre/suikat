# adapt

Пакет `adapt` предоставляет функции для проверки типа `interface{}` и преобразования в базовые типы: `bool`, `int`, `string`. Все функции преобразования возвращают ошибку `ge.MismatchError`, если тип не совпадает.

## Зависимости

- `github.com/nobuenhombre/suikat/pkg/ge` — `ge.MismatchError`, `ge.Pin`

## Функции

### `Check(val reflect.Value, expectType string) error`

Сравнивает тип `reflect.Value` с ожидаемой строкой. Если типы не совпадают — возвращает `ge.MismatchError`.

```go
import (
    "reflect"
    "github.com/nobuenhombre/suikat/pkg/adapt"
    "github.com/nobuenhombre/suikat/pkg/ge"
)

// Пример 1: тип совпадает — ошибки нет
err := adapt.Check(reflect.ValueOf(42), "int")
// err == nil

// Пример 2: тип не совпадает — ошибка
err = adapt.Check(reflect.ValueOf("hello"), "int")
// err == &ge.MismatchError{
//     ComparedItems: "val.Type().String() vs expectType",
//     Expected:      "int",
//     Actual:        "string",
// }
```

### `Bool(v interface{}) (bool, error)`

Преобразует `interface{}` в `bool`. Если значение не является `bool` — возвращает ошибку.

```go
import "github.com/nobuenhombre/suikat/pkg/adapt"

// Успешное преобразование
b, err := adapt.Bool(true)
// b == true, err == nil

// Ошибка: int нельзя преобразовать в bool
b, err = adapt.Bool(42)
// b == false, err != nil (MismatchError)
```

### `Int(v interface{}) (int, error)`

Преобразует `interface{}` в `int`. Если значение не является `int` — возвращает ошибку.

```go
import "github.com/nobuenhombre/suikat/pkg/adapt"

// Успешное преобразование
n, err := adapt.Int(123)
// n == 123, err == nil

// Ошибка: string нельзя преобразовать в int
n, err = adapt.Int("123")
// n == 0, err != nil (MismatchError)
```

### `String(v interface{}) (string, error)`

Преобразует `interface{}` в `string`. Если значение не является `string` — возвращает ошибку.

```go
import "github.com/nobuenhombre/suikat/pkg/adapt"

// Успешное преобразование
s, err := adapt.String("hello")
// s == "hello", err == nil

// Ошибка: int нельзя преобразовать в string
s, err = adapt.String(42)
// s == "", err != nil (MismatchError)
```

### `IsNil(i interface{}) bool`

Проверяет, является ли `interface{}` nil. Для интерфейсов, указателей, мап, каналов, слайсов и функций — проверяет внутреннее значение. Для остальных типов — всегда возвращает `false`.

```go
import "github.com/nobuenhombre/suikat/pkg/adapt"

// Пример 1: nil-интерфейс
var s *string = nil
adapt.IsNil(s) // true

// Пример 2: nil-слайс
var sl []int = nil
adapt.IsNil(sl) // true

// Пример 3: nil-мапа
var m map[string]int = nil
adapt.IsNil(m) // true

// Пример 4: non-nil слайс
sl2 := []int{}
adapt.IsNil(sl2) // false (слайс существует, но пустой)

// Пример 5: nil-интерфейс
var i interface{} = nil
adapt.IsNil(i) // true
```

## Пример: проверка данных из конфигурации

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/adapt"
)

func main() {
    // Допустим, данные пришли как interface{} (например, из map[string]interface{})
    config := map[string]interface{}{
        "port":     8080,
        "debug":    true,
        "name":     "my-service",
        "maxConn":  42,
    }

    // Безопасно извлекаем значения с проверкой типов
    port, err := adapt.Int(config["port"])
    if err != nil {
        fmt.Printf("ошибка port: %v\n", err)
        return
    }
    fmt.Printf("port: %d\n", port) // port: 8080

    debug, err := adapt.Bool(config["debug"])
    if err != nil {
        fmt.Printf("ошибка debug: %v\n", err)
        return
    }
    fmt.Printf("debug: %t\n", debug) // debug: true

    name, err := adapt.String(config["name"])
    if err != nil {
        fmt.Printf("ошибка name: %v\n", err)
        return
    }
    fmt.Printf("name: %s\n", name) // name: my-service

    // Проверка на nil
    missing := config["missing"]
    if adapt.IsNil(missing) {
        fmt.Println("missing: не задано") // missing: не задано
    }
}
```
