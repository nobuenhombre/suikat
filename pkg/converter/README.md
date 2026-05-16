# converter

Пакет `converter` предоставляет функции для преобразования строк в другие типы: int, float64, bool, time.Duration и т.д. Все функции возвращают `*ParserError` при ошибке.

## Зависимости

- `github.com/nobuenhombre/suikat/pkg/ge` — `ge.Pin`
- `github.com/xhit/go-str2duration/v2` — парсинг duration

## Ошибка `ParserError`

```go
type ParserError struct {
    ParserType string
    Value      string
    Parent     error
}
```

## Функции

### Числовые преобразования

| Функция | Возвращает |
|---------|-----------|
| `StringToInt(s string) (int, error)` | int |
| `StringToInt8(s string) (int8, error)` | int8 |
| `StringToInt16(s string) (int16, error)` | int16 |
| `StringToInt32(s string) (int32, error)` | int32 |
| `StringToInt64(s string) (int64, error)` | int64 |

### Другие преобразования

| Функция | Возвращает |
|---------|-----------|
| `StringToBool(s string) (bool, error)` | bool |
| `StringToFloat32(s string) (float32, error)` | float32 |
| `StringToFloat64(s string) (float64, error)` | float64 |
| `StringToTime(s, format string) (time.Time, error)` | time.Time |
| `StringToIntSlice(s, sep string) ([]int, error)` | []int |
| `StringToDuration(s string) (time.Duration, error)` | time.Duration |

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/converter"
)

func main() {
    // --- Пример 1: StringToInt ---

    n, err := converter.StringToInt("42")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%d\n", n) // 42

    // Неправильный формат
    _, err = converter.StringToInt("abc")
    if err != nil {
        // ParserType [Int], Value [abc], Error [...]
        fmt.Printf("ошибка: %v\n", err)
    }

    // --- Пример 2: StringToInt64 ---

    n64, err := converter.StringToInt64("9223372036854775807")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%d\n", n64)

    // --- Пример 3: StringToBool ---

    b, err := converter.StringToBool("true")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%t\n", b) // true

    // --- Пример 4: StringToFloat64 ---

    f, err := converter.StringToFloat64("3.14159")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%.5f\n", f) // 3.14159

    // --- Пример 5: StringToTime ---

    t, err := converter.StringToTime("2026-05-16", "2006-01-02")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%v\n", t) // 2026-05-16 00:00:00 +0000 UTC

    // --- Пример 6: StringToIntSlice ---

    slice, err := converter.StringToIntSlice("1,2,3,4,5", ",")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%v\n", slice) // [1 2 3 4 5]

    // --- Пример 7: StringToDuration ---

    // Поддерживаемые форматы:
    // - "30s" → 30 секунд
    // - "1:30" → 1 минута 30 секунд
    // - "1:30:45" → 1 час 30 минут 45 секунд

    d, err := converter.StringToDuration("1:30:45")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%v\n", d) // 1h30m45s

    d2, err := converter.StringToDuration("30s")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%v\n", d2) // 30s
}
```
