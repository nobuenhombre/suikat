# futi

Пакет `futi` предоставляет утилиты для работы с файлами: проверка существования, создание временных файлов, копирование, перемещение, удаление.

## Основные функции

### Проверка

```go
func FileExists(filename string) bool
func DirExists(dirname string) bool
```

### Создание временных файлов

```go
func CreateTempFile(dir, pattern string, data *[]byte) (string, error)
```

Создаёт временный файл с данными. Возвращает путь к файлу.

### Копирование

```go
func Copy(inFile, outFile string) error
```

Копирует файл. Если `inFile == outFile` — ничего не делает. Проверяет, что `inFile` — обычный файл.

### Перемещение

```go
func Move(inFile, outFile string) error
```

Перемещает файл (Copy + Remove).

### Удаление

```go
func Delete(fileName string) error
```

Удаляет файл.

## Ошибки

### `DeleteFileError`

```go
type DeleteFileError struct {
    FileName string
    Parent   error
}
```

### `CreateTempFileError`

```go
type CreateTempFileError struct {
    Dir     string
    Pattern string
    Data    *[]byte
    Parent  error
}
```

### `IsNotRegularFileError`

```go
type IsNotRegularFileError struct {
    File string
}
```

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/futi"
)

func main() {
    // --- Пример 1: Проверка существования ---

    if futi.FileExists("/tmp/test.txt") {
        fmt.Println("файл существует")
    }

    if futi.DirExists("/tmp") {
        fmt.Println("директория существует")
    }

    // --- Пример 2: Создание временного файла ---

    data := []byte("hello world")
    path, err := futi.CreateTempFile("/tmp", "test-*", &data)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("создан файл: %s\n", path)

    // --- Пример 3: Копирование ---

    err = futi.Copy("/tmp/source.txt", "/tmp/dest.txt")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // --- Пример 4: Перемещение ---

    err = futi.Move("/tmp/source.txt", "/tmp/moved.txt")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // --- Пример 5: Удаление ---

    err = futi.Delete("/tmp/moved.txt")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
}
```
