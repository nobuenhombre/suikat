# execution-tree

Пакет `execution-tree` предоставляет систему выполнения деревьев задач (execution trees). Каждая нода содержит функцию-исполнитель (`NodeFunc`), которая определяет, к какой следующей ноде перейти.

## Основные типы

### `NodeKey`

Тип `string` — уникальный идентификатор ветки.

### `NodeKeyExit`

Константа `"exit"` — сигнал завершения выполнения.

### `NodeFunc`

Тип функции: `func(input interface{}) NodeKey`. Выполняет ноду и возвращает ключ следующей ноды.

### `Node`

Структура ноды дерева:
- `Executor NodeFunc` — функция-исполнитель.
- `ExecutorInput interface{}` — данные, передаваемые в `NodeFunc`.
- `Branches map[NodeKey]*Node` — мапа возможных следующих нод.

**Методы:**
- `AddBranch(key NodeKey, branchNode *Node)` — добавить ветку.
- `Run() error` — выполнить ноду и перейти к следующей на основе результата `NodeFunc`.

### `IExecutionTree`

Интерфейс с методами `AddBranch()`, `Run()`.

## Функция `NewNode`

```go
func NewNode(executor NodeFunc, executorInput interface{}) *Node
```

Создаёт новую ноду.

## Ошибки

### `ExecutorNotSetError`

Возвращается, если `Executor == nil`.

```go
type ExecutorNotSetError struct {
    Key string
}
```

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/execution-tree"
)

func main() {
    // --- Пример 1: Линейная цепочка нод ---

    type Sum int
    var testSum executiontree.Sum

    executorA := func(input interface{}) executiontree.NodeKey {
        sum := input.(*executiontree.Sum)
        *sum += 1
        return "B"
    }

    executorB := func(input interface{}) executiontree.NodeKey {
        sum := input.(*executiontree.Sum)
        *sum += 2
        return "C"
    }

    executorC := func(input interface{}) executiontree.NodeKey {
        sum := input.(*executiontree.Sum)
        *sum += 3
        return executiontree.NodeKeyExit
    }

    nodeA := executiontree.NewNode(executorA, &testSum)
    nodeB := executiontree.NewNode(executorB, &testSum)
    nodeC := executiontree.NewNode(executorC, &testSum)

    nodeA.AddBranch("B", nodeB)
    nodeB.AddBranch("C", nodeC)

    err := nodeA.Run()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    fmt.Printf("Sum: %d\n", testSum) // Sum: 6 (1 + 2 + 3)
}

// --- Пример 2: Ветвление ---

func main() {
    router := func(input interface{}) executiontree.NodeKey {
        if input.(string) == "A" {
            return "branchA"
        }
        return "branchB"
    }

    processA := func(input interface{}) executiontree.NodeKey {
        fmt.Printf("Branch A: %v\n", input)
        return executiontree.NodeKeyExit
    }

    processB := func(input interface{}) executiontree.NodeKey {
        fmt.Printf("Branch B: %v\n", input)
        return executiontree.NodeKeyExit
    }

    root := executiontree.NewNode(router, "go to B")
    branchA := executiontree.NewNode(processA, nil)
    branchB := executiontree.NewNode(processB, nil)

    root.AddBranch("branchA", branchA)
    root.AddBranch("branchB", branchB)

    err := root.Run()
    // Вывод: Branch B: go to B
}

// --- Пример 3: Ошибка — executor не установлен ---

func main() {
    node := executiontree.NewNode(nil, nil)
    err := node.Run()
    if err != nil {
        // executor is not set
        fmt.Printf("ошибка: %v\n", err)
    }
}

// --- Пример 4: Ошибка — нодa не найдена ---

func main() {
    badNode := executiontree.NewNode(
        func(input interface{}) executiontree.NodeKey {
            return "nonexistent"
        },
        nil,
    )

    err := badNode.Run()
    if err != nil {
        // NotFoundError: key = "nonexistent"
        fmt.Printf("ошибка: %v\n", err)
    }
}
```
