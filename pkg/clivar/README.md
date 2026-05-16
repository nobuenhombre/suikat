# clivar

Пакет `clivar` предоставляет синтаксический сахар для работы с `flag` через теги структур. Позволяет описать, как заполнить поля структуры из командной строки, используя простой формат тегов.

## Зависимости

- `github.com/nobuenhombre/suikat/pkg/refavour` — рефлексия для работы со структурами
- `github.com/nobuenhombre/suikat/pkg/ge` — ошибки

## Формат тега `cli`

```
NAME[description key]:valueType=defaultValue
```

| Часть | Описание |
|-------|----------|
| `NAME` | Имя флага для `flag`-пакета |
| `description key` | Описание флага |
| `valueType` | Тип значения: `string`, `int`, `float64`, `bool`, `SliceStrings` |
| `defaultValue` | Значение по умолчанию |

## Функция `Load`

```go
func Load(structData interface{}) error
```

Заполняет поля структуры из командной строки на основе тегов `cli`. Вызывает `flag.Parse()` внутри.

## Примеры

```go
package main

import (
    "flag"
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/clivar"
)

// --- Пример 1: Все типы ---

type CommandLineParams struct {
    Path           string   `cli:"PATH[Path to file]:string=/some/default/path"`
    Port           int      `cli:"PORT[Port for server]:int=8080"`
    Coefficient    float64  `cli:"COEFFICIENT[Coefficient transmutation]:float64=75.31"`
    MakeSomeAction bool     `cli:"MSA[Do make some action?]:bool=false"`
    Languages      []string `cli:"LANGUAGES[Languages]:SliceStrings=ru,en,de,fr"`
}

func main() {
    var params CommandLineParams

    err := clivar.Load(&params)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // Вывод значений
    fmt.Printf("Path: %s\n", params.Path)
    fmt.Printf("Port: %d\n", params.Port)
    fmt.Printf("Coefficient: %f\n", params.Coefficient)
    fmt.Printf("MakeSomeAction: %t\n", params.MakeSomeAction)
    fmt.Printf("Languages: %v\n", params.Languages)
}

// --- Пример 2: Использование через командную строку ---

// Запуск:
//   go run main.go -PATH="/custom/path" -PORT=3000 -MSA=true
//
// Результат:
//   Path: /custom/path
//   Port: 3000
//   MakeSomeAction: true
//   Languages: [ru en de fr] (по умолчанию)

// --- Пример 3: Только обязательные поля ---

type SimpleConfig struct {
    Host string `cli:"HOST[Server host]:string=localhost"`
    Port int    `cli:"PORT[Server port]:int=0"`
}

func main() {
    var cfg SimpleConfig
    err := clivar.Load(&cfg)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s:%d\n", cfg.Host, cfg.Port)
}

// --- Пример 4: Ошибка в теге ---

// Если тег не соответствует формату, Load вернёт ge.MismatchError.
// Неправильный тег: "PORT:8080" (пропущено описание в [])
// Правильный тег: "PORT[Port]:int=8080"
```

## Типы значений

| ValueType | Go-тип | defaultValue формат |
|-----------|--------|---------------------|
| `string` | `string` | Строка |
| `int` | `int` | Целое число |
| `float64` | `float64` | Число с точкой |
| `bool` | `bool` | `true` / `false` |
| `SliceStrings` | `[]string` | Строка через запятую |
