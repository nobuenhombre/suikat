# figlu

Пакет `figlu` предоставляет функции для склеивания содержимого файлов из директории.

## Основные типы

### `Path`

```go
type Path string
```

Путь к директории.

**Методы:**
- `GlueContent(onlyExt string, ignoreScanErr bool) (string, error)` — вернуть склеенное содержимое всех файлов.
- `Glue(outFile TxtFile, onlyExt string, ignoreScanErr bool) error` — записать склеенное содержимое в файл.

### `PathList`

```go
type PathList []Path
```

Список путей.

**Методы:**
- `GlueContent(onlyExt string, ignoreScanErr bool) (string, error)` — склеить содержимое всех путей.
- `Glue(outFile TxtFile, onlyExt string, ignoreScanErr bool) error` — записать склеенное содержимое в файл.

## Параметры

| Параметр | Описание |
|----------|----------|
| `onlyExt` | Фильтр по расширению (например, `.css`). Если пустая строка — все файлы. |
| `ignoreScanErr` | Игнорировать ошибки сканирования. |

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/fico"
    "github.com/nobuenhombre/suikat/pkg/figlu"
)

func main() {
    // --- Пример 1: Склейка всех файлов из директории ---

    dirPath := figlu.Path("/path/to/directory")

    // Склейка всего
    content, err := dirPath.GlueContent("", false)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("склеено: %d байт\n", len(content))

    // --- Пример 2: Фильтр по расширению ---

    // Склейка только .css файлов
    cssContent, err := dirPath.GlueContent(".css", false)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("CSS: %d байт\n", len(cssContent))

    // --- Пример 3: Запись в файл ---

    outFile := fico.TxtFile("/tmp/bundle.css")
    err = dirPath.Glue(outFile, ".css", false)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("записано в %s\n", outFile)

    // --- Пример 4: PathList ---

    paths := figlu.PathList{
        figlu.Path("/path/to/dir1"),
        figlu.Path("/path/to/dir2"),
    }

    content, err = paths.GlueContent("", false)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("склеено из всех директорий: %d байт\n", len(content))

    // --- Пример 5: Игнорирование ошибок ---

    content, err = dirPath.GlueContent("", true)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Ошибки сканирования игнорируются
}
```
