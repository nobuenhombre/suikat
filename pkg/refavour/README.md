# refavour

Пакет `refavour` предоставляет интерфейс для работы с тегами структур и набор функций для работы с reflect.

## Обзор

Пакет содержит:
- Интерфейс `TagProcessor` для парсинга тегов структур
- Утилиты для проверки типов через reflect
- Функцию для чтения полей и тегов структуры

## TagProcessor

Интерфейс для парсинга тегов структур:

```go
type TagProcessor interface {
    GetFieldInfo(typeField reflect.StructField, valueField reflect.Value) (interface{}, error)
}
```

## GetReflectValue

Возвращает `reflect.Value` из `interface{}`. Если передан указатель — разыменовывает.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

type User struct {
    Name string
    Age  int
}

u := User{Name: "Alice", Age: 30}

val := refavour.GetReflectValue(u)
fmt.Println(val.NumField()) // 2

val = refavour.GetReflectValue(&u)
fmt.Println(val.NumField()) // 2 (разыменовывает указатель)
```

## CheckKind

Проверяет, что тип `interface{}` соответствует ожидаемому `reflect.Kind`.

```go
import (
    "fmt"
    "reflect"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

type User struct{ Name string }

var s string = "hello"
var u User
var arr []int
var m map[string]int

fmt.Println(refavour.CheckKind(s, reflect.String))   // nil
fmt.Println(refavour.CheckKind(u, reflect.Struct))    // nil
fmt.Println(refavour.CheckKind(arr, reflect.Slice))   // nil
fmt.Println(refavour.CheckKind(m, reflect.Map))       // nil
fmt.Println(refavour.CheckKind(s, reflect.Int))       // MismatchError
```

## CheckStructure

Проверяет, что `interface{}` является структурой.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

type Config struct {
    Host string
    Port int
}

var cfg Config
err := refavour.CheckStructure(cfg)
fmt.Println(err) // nil (структура)

err = refavour.CheckStructure("hello")
fmt.Println(err) // MismatchError (не структура)
```

## CheckMap

Проверяет, что `interface{}` является картой.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

m := map[string]int{"a": 1}
err := refavour.CheckMap(m)
fmt.Println(err) // nil (карта)

err = refavour.CheckMap("hello")
fmt.Println(err) // MismatchError (не карта)
```

## CheckSlice

Проверяет, что `interface{}` является срезом.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

s := []int{1, 2, 3}
err := refavour.CheckSlice(s)
fmt.Println(err) // nil (срез)

err = refavour.CheckSlice(42)
fmt.Println(err) // MismatchError (не срез)
```

## CheckCanBeChanged

Проверяет, что значение можно изменять.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

x := 42
err := refavour.CheckCanBeChanged(x)
fmt.Println(err) // MismatchError (значение нельзя изменить)

err = refavour.CheckCanBeChanged(&x)
fmt.Println(err) // nil (указатель можно разыменовать и изменить)
```

## FieldsInfo

Тип для хранения информации о полях структуры:

```go
type FieldsInfo map[string]interface{}
```

## GetStructureFieldsTypes

Читает поля структуры и возвращает карту с именами полей и их информацией.

```go
import (
    "fmt"
    "reflect"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Реализация TagProcessor
type JSONTagProcessor struct{}

func (p *JSONTagProcessor) GetFieldInfo(
    typeField reflect.StructField,
    valueField reflect.Value,
) (interface{}, error) {
    tag := typeField.Tag.Get("json")
    return map[string]string{
        "name": typeField.Name,
        "tag":  tag,
        "type": typeField.Type.String(),
    }, nil
}

u := User{Name: "Alice", Age: 30}

fields, err := refavour.GetStructureFieldsTypes(u, &JSONTagProcessor{})
if err != nil {
    panic(err)
}

for name, info := range fields {
    fmt.Printf("Поле %s: %v\n", name, info)
}
// Вывод:
// Поле Name: map[name:Name tag:name type:string]
// Поле Age: map[name:Age tag:age type:int]
```

## Пример: извлечение JSON-тегов

```go
import (
    "fmt"
    "reflect"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
    Debug bool   `json:"debug" default:"false"`
}

type JSONTagProcessor struct{}

func (p *JSONTagProcessor) GetFieldInfo(
    typeField reflect.StructField,
    valueField reflect.Value,
) (interface{}, error) {
    return typeField.Tag.Get("json"), nil
}

cfg := Config{Host: "localhost", Port: 8080, Debug: true}

fields, err := refavour.GetStructureFieldsTypes(cfg, &JSONTagProcessor{})
if err != nil {
    panic(err)
}

for name, tag := range fields {
    fmt.Printf("%s -> %s\n", name, tag)
}
// Вывод:
// Host -> host
// Port -> port
// Debug -> debug
```

## Пример: валидация структуры

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/refavour"
)

func validateData(data interface{}) error {
    if err := refavour.CheckStructure(data); err != nil {
        return fmt.Errorf("ожидалась структура: %w", err)
    }
    return nil
}

type User struct {
    Name string
    Age  int
}

err := validateData(User{Name: "Alice", Age: 30})
fmt.Println(err) // nil

err = validateData("hello")
fmt.Println(err) // ожидалась структура: MismatchError
```
