# fitree

Пакет `fitree` предоставляет рекурсивное сканирование дерева каталогов с файлами.

## Константы

| Константа | Значение |
|-----------|----------|
| `StartDepth` | `1` — начальная глубина сканирования |

## Основные типы

### `TreeNodeStruct`

```go
type TreeNodeStruct struct {
    Path         string
    Name         string
    Depth        int
    Files        []os.FileInfo
    FilesCount   int
    SubDirs      []os.FileInfo
    SubDirsCount int
}
```

- `Path` — путь к директории.
- `Name` — имя директории.
- `Depth` — глубина в дереве.
- `Files` — файлы в директории.
- `FilesCount` — количество файлов.
- `SubDirs` — поддиректории.
- `SubDirsCount` — количество поддиректорий.

**Методы:**
- `Fill(path string, depth int) error` — заполнить структуру при сканировании.

### `TreeNodeListStruct`

```go
type TreeNodeListStruct struct {
    List    []TreeNodeStruct
    Reverse map[string]int
}
```

- `List` — список узлов.
- `Reverse` — мапа путь → индекс.

**Методы:**
- `Add(node TreeNodeStruct)` — добавить узел.
- `Scan(path string, depth int, ignoreErr bool) error` — рекурсивно сканировать дерево.
- `GetNode(index int) (*TreeNodeStruct, error)` — получить узел по индексу.

## Ошибки

### `NodeIndexDontExistsError`

```go
type NodeIndexDontExistsError struct {
    Index int
}
```

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/fitree"
)

func main() {
    // --- Пример 1: Сканирование дерева ---

    list := fitree.TreeNodeListStruct{}

    err := list.Scan("/path/to/directory", fitree.StartDepth, false)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // Обход всех узлов
    for i, node := range list.List {
        fmt.Printf("[%d] %s (глубина: %d, файлов: %d, поддиректорий: %d)\n",
            i, node.Path, node.Depth, node.FilesCount, node.SubDirsCount)

        // Обход файлов
        for _, f := range node.Files {
            fmt.Printf("  файл: %s (%d байт)\n", f.Name(), f.Size())
        }

        // Обход поддиректорий
        for _, d := range node.SubDirs {
            fmt.Printf("  поддиректория: %s\n", d.Name())
        }
    }

    // --- Пример 2: Получение узла по индексу ---

    node, err := list.GetNode(0)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("первый узел: %s\n", node.Path)

    // --- Пример 3: Игнорирование ошибок ---

    list2 := fitree.TreeNodeListStruct{}
    err = list2.Scan("/path/to/directory", fitree.StartDepth, true)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Ошибки сканирования игнорируются
}
```
