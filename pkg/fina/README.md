# fina

Пакет `fina` предоставляет функции для работы с частями имён файлов: получение информации, добавление префикса/суффикса, смена расширения.

## Основные типы

### `FilePartsInfo`

```go
type FilePartsInfo struct {
    Path           string
    Name           string
    Ext            string
    NameWithoutExt string
}
```

- `Path` — путь к директории (с разделителем в конце).
- `Name` — полное имя файла.
- `Ext` — расширение файла.
- `NameWithoutExt` — имя без расширения.

**Методы:**
- `Prefix(prefix string) string` — добавить префикс: `/path/demo-file.ext`
- `Suffix(suffix string) string` — добавить суффикс: `/path/file-suffix.ext`
- `PS(prefix, suffix string) string` — префикс + суффикс: `/path/demo-file-omed.ext`
- `NewExt(ext string) string` — новая расширение: `/path/file.newext`
- `PrefixWithNewExt(prefix, ext string) string` — префикс + новое расширение
- `SuffixWithNewExt(suffix, ext string) string` — суффикс + новое расширение
- `PSWithNewExt(prefix, suffix, ext string) string` — всё вместе

## Функция `GetFilePartsInfo`

```go
func GetFilePartsInfo(file string) *FilePartsInfo
```

Возвращает информацию о частях файла.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/fina"
)

func main() {
    fpi := fina.GetFilePartsInfo("/some/path/file.txt")

    fmt.Printf("Path: %s\n", fpi.Path)          // /some/path/
    fmt.Printf("Name: %s\n", fpi.Name)           // file.txt
    fmt.Printf("Ext: %s\n", fpi.Ext)             // .txt
    fmt.Printf("NameWithoutExt: %s\n", fpi.NameWithoutExt) // file

    // --- Пример 1: Префикс ---

    newName := fpi.Prefix("demo-")
    fmt.Println(newName) // /some/path/demo-file.txt

    // --- Пример 2: Суффикс ---

    newName = fpi.Suffix("-backup")
    fmt.Println(newName) // /some/path/file-backup.txt

    // --- Пример 3: Префикс + суффикс ---

    newName = fpi.PS("demo-", "-v2")
    fmt.Println(newName) // /some/path/demo-file-v2.txt

    // --- Пример 4: Новое расширение ---

    newName = fpi.NewExt(".md")
    fmt.Println(newName) // /some/path/file.md

    // --- Пример 5: Префикс + новое расширение ---

    newName = fpi.PrefixWithNewExt("demo-", ".md")
    fmt.Println(newName) // /some/path/demo-file.md

    // --- Пример 6: Суффикс + новое расширение ---

    newName = fpi.SuffixWithNewExt("-backup", ".md")
    fmt.Println(newName) // /some/path/file-backup.md

    // --- Пример 7: Всё вместе ---

    newName = fpi.PSWithNewExt("demo-", "-v2", ".md")
    fmt.Println(newName) // /some/path/demo-file-v2.md
}
```
