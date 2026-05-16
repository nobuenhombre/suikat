# csvarser

Пакет `csvarser` предоставляет CSV-парсер, который заполняет поля структуры из строкового слайса на основе тегов `order`.

## Зависимости

- `github.com/nobuenhombre/suikat/pkg/refavour` — рефлексия для работы со структурами

## Основные типы

### `CsvParser`

```go
type CsvParser struct {
    TypeParsers map[string]ParserFunc
}
```

**Методы:**
- `Init()` — инициализировать `TypeParsers`.
- `AddTypeParser(dataType string, parser ParserFunc)` — добавить парсер для типа.
- `FillStructFromSlice(structData interface{}, sliceData []string) error` — заполнить структуру из слайса строк.

### `ParserFunc`

```go
type ParserFunc func(s string) (interface{}, error)
```

### `FieldInfo`

```go
type FieldInfo struct {
    Type  reflect.Type
    Order int
}
```

## Ошибки

### `ParserNotFoundError`

Возвращается, если не найден парсер для типа поля.

```go
type ParserNotFoundError struct {
    FieldType string
}
```

### `FieldNotExistsInSliceError`

Возвращается, если поле отсутствует в слайсе.

```go
type FieldNotExistsInSliceError struct {
    FieldName string
    FieldType string
    Index     int
}
```

## Формат тега `order`

```go
type SomeCSV struct {
    A int    `order:"0"`
    B string `order:"1"`
    C float64 `order:"2"`
}
```

## Примеры

```go
package main

import (
    "fmt"
    "strconv"
    "github.com/nobuenhombre/suikat/pkg/csvarser"
)

// --- Пример 1: Парсинг CSV-строки в структуру ---

type Person struct {
    ID   int    `order:"0"`
    Name string `order:"1"`
    Age  int    `order:"2"`
}

func main() {
    parser := &csvarser.CsvParser{}
    parser.Init()

    // Добавляем парсеры для типов
    parser.AddTypeParser("int", func(s string) (interface{}, error) {
        n, err := strconv.Atoi(s)
        return n, err
    })

    parser.AddTypeParser("string", func(s string) (interface{}, error) {
        return s, nil
    })

    // CSV-строка: "1,John,30"
    csvRow := []string{"1", "John", "30"}

    var person Person
    err := parser.FillStructFromSlice(&person, csvRow)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    fmt.Printf("%d %s %d\n", person.ID, person.Name, person.Age)
    // 1 John 30
}

// --- Пример 2: Ошибка — парсер не найден ---

type Config struct {
    Port int    `order:"0"`
    Host string `order:"1"`
}

func main() {
    parser := &csvarser.CsvParser{}
    parser.Init()

    // Не добавили парсер для "int" — ошибка!
    var cfg Config
    err := parser.FillStructFromSlice(&cfg, []string{"8080", "localhost"})
    if err != nil {
        // Parser not found for Type [int]
        fmt.Printf("ошибка: %v\n", err)
    }
}

// --- Пример 3: Ошибка — поле отсутствует в слайсе ---

type Simple struct {
    A int `order:"0"`
    B int `order:"1"`
}

func main() {
    parser := &csvarser.CsvParser{}
    parser.Init()
    parser.AddTypeParser("int", func(s string) (interface{}, error) {
        n, _ := strconv.Atoi(s)
        return n, nil
    })

    var s Simple
    err := parser.FillStructFromSlice(&s, []string{"1"})
    if err != nil {
        // Field B (int) not Exists in Slice index=[1]
        fmt.Printf("ошибка: %v\n", err)
    }
}
```
