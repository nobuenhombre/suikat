# scatter-fs

Пакет `scatter-fs` предоставляет виртуальную файловую систему, которая распределяет файлы по нескольким базовым директориям с учётом свободного места на диске.

## Обзор

Пакет позволяет создать абстракцию над несколькими директориями, при которой:
- Файлы создаются в директории с достаточным свободным местом
- Файлы открываются из всех базовых директорий
- Директории создаются во всех базовых директориях одновременно
- Доступна операция перемещения директории между базовыми директориями

## IFileSystem

Интерфейс виртуальной файловой системы:

```go
type IFileSystem interface {
    Suggest(path string) (string, error)
    ResolveRealPath(path string) (string, error)
    Create(path string) (*os.File, error)
    Open(path string) (*os.File, error)
    MkdirAll(path string, perm fs.FileMode) error
    Remove(path string) error
    RemoveAll(path string) error
    Stat(path string) (fs.FileInfo, error)
    ReadDir(path string) ([]fs.DirEntry, error)
    MoveDir(oldDir, newDir string) error
}
```

## FileSystem

Реализация виртуальной файловой системы:

```go
type FileSystem struct {
    dirs        []string // Базовые директории
    freePercent float64  // Минимальный процент свободного места
}
```

## New

Создаёт виртуальную файловую систему из списка директорий и минимального процента свободного места:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{
    "/data/storage1",
    "/data/storage2",
    "/data/storage3",
}

vfs := scatterfs.New(dirs, 10) // минимум 10% свободного места
```

## Suggest

Возвращает реальный путь для файла, выбирая директорию с достаточным свободным местом:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

realPath, err := vfs.Suggest("documents/report.pdf")
if err != nil {
    panic(err)
}
fmt.Println(realPath) // /data/storage1/documents/report.pdf (или storage2)
```

## ResolveRealPath

Находит реальный путь файла, перебирая все базовые директории:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

// Файл создан в storage2
realPath, err := vfs.ResolveRealPath("documents/report.pdf")
if err != nil {
    panic(err)
}
fmt.Println(realPath) // /data/storage2/documents/report.pdf
```

## Create

Создаёт файл в подходящей базовой директории:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

file, err := vfs.Create("documents/report.pdf")
if err != nil {
    panic(err)
}
defer file.Close()

_, err = file.WriteString("Hello, scatter-fs!")
if err != nil {
    panic(err)
}
fmt.Println("Файл создан")
```

## Open

Открывает файл из любой базовой директории:

```go
import (
    "fmt"
    "io"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

file, err := vfs.Open("documents/report.pdf")
if err != nil {
    panic(err)
}
defer file.Close()

data, err := io.ReadAll(file)
if err != nil {
    panic(err)
}
fmt.Println(string(data))
```

## MkdirAll

Создаёт директорию во всех базовых директориях:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

err := vfs.MkdirAll("documents/archive", 0755)
if err != nil {
    panic(err)
}
// Создано: /data/storage1/documents/archive
// Создано: /data/storage2/documents/archive
```

## Remove

Удаляет файл из всех базовых директорий:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

err := vfs.Remove("documents/report.pdf")
if err != nil {
    panic(err)
}
```

## RemoveAll

Рекурсивно удаляет директорию из всех базовых директорий:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

err := vfs.RemoveAll("documents/archive")
if err != nil {
    panic(err)
}
```

## Stat

Возвращает информацию о файле:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

info, err := vfs.Stat("documents/report.pdf")
if err != nil {
    panic(err)
}
fmt.Printf("Имя: %s\n", info.Name())
fmt.Printf("Размер: %d байт\n", info.Size())
fmt.Printf("Это директория: %v\n", info.IsDir())
```

## ReadDir

Читает содержимое директории из всех базовых директорий (объединяя результаты):

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2"}
vfs := scatterfs.New(dirs, 10)

entries, err := vfs.ReadDir("documents")
if err != nil {
    panic(err)
}
for _, entry := range entries {
    fmt.Printf("%s (dir: %v)\n", entry.Name(), entry.IsDir())
}
```

## MoveDir

Перемещает виртуальную директорию в другую виртуальную директорию, распределяя файлы по базовым директориям:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

dirs := []string{"/data/storage1", "/data/storage2", "/data/storage3"}
vfs := scatterfs.New(dirs, 10)

// Перемещение
err := vfs.MoveDir("documents/old", "documents/new")
if err != nil {
    panic(err)
}
// Файлы перемещены из documents/old в documents/new
// Распределены по доступным базовым директориям
// Пустые директории в old удалены
```

## Ошибки

Пакет определяет следующие ошибки:

```go
import (
    "errors"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

// Проверка ошибок

// Нет директорий с достаточным свободным местом
err := vfs.Suggest("file.txt")
if errors.Is(err, scatterfs.NoDirectoriesWithFreeSpaceAvailableError) {
    fmt.Println("Нет свободного места")
}

// Неверный путь
err = vfs.Create("")
if errors.Is(err, scatterfs.InvalidPathError) {
    fmt.Println("Неверный путь")
}

// Нельзя удалить корень
err = vfs.Remove("")
if errors.Is(err, scatterfs.CantRemoveRootError) {
    fmt.Println("Нельзя удалить корень")
}

// Директория не найдена
err = vfs.MoveDir("nonexistent", "new")
if errors.Is(err, scatterfs.DirectoryNotFoundInAnyBaseDirectoryError) {
    fmt.Println("Директория не найдена в базовых директориях")
}
```

## Пример: полное использование

```go
import (
    "fmt"
    "io"
    "github.com/nobuenhombre/suikat/pkg/scatter-fs"
)

func main() {
    dirs := []string{
        "/data/storage1",
        "/data/storage2",
        "/data/storage3",
    }

    vfs := scatterfs.New(dirs, 10)

    // 1. Создание директории
    err := vfs.MkdirAll("uploads/2024", 0755)
    if err != nil {
        panic(err)
    }

    // 2. Создание файла
    file, err := vfs.Create("uploads/2024/document.txt")
    if err != nil {
        panic(err)
    }
    _, err = file.WriteString("Content here")
    file.Close()
    if err != nil {
        panic(err)
    }

    // 3. Чтение файла
    file, err = vfs.Open("uploads/2024/document.txt")
    if err != nil {
        panic(err)
    }
    data, _ := io.ReadAll(file)
    file.Close()
    fmt.Println(string(data)) // Content here

    // 4. Информация о файле
    info, err := vfs.Stat("uploads/2024/document.txt")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Размер: %d\n", info.Size())

    // 5. Список файлов
    entries, err := vfs.ReadDir("uploads/2024")
    for _, entry := range entries {
        fmt.Printf("Файл: %s\n", entry.Name())
    }

    // 6. Удаление
    err = vfs.RemoveAll("uploads/2024")
    if err != nil {
        panic(err)
    }
}
```
