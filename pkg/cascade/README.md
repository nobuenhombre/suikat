# cascade

Пакет `cascade` предоставляет систему каскадного выполнения задач (task cascade). Каждая задача содержит функцию-исполнитель (`ExecutorFunc`), которая определяет, какая следующая задача будет выполнена.

## Зависимости

- `github.com/nobuenhombre/suikat/pkg/ge` — `ge.Pin` для оборачивания ошибок

## Основные типы

### `Key`

Тип `string` — уникальный идентификатор задачи. Получается из имени функции-исполнителя через `runtime.FuncForPC`.

### `ExitKey`

Константа `"exit"` — сигнал завершения выполнения каскада.

### `ExecutorFunc`

Тип функции: `func(data any) Key`. Выполняет задачу и возвращает ключ следующей задачи.

**Методы:**
- `Name() Key` — возвращает полное имя функции (включая package path) через `reflect` и `runtime`.

### `Task`

Структура задачи:
- `key` — уникальный идентификатор (из имени функции).
- `data` — данные, передаваемые в `ExecutorFunc`.
- `exec` — функция-исполнитель.
- `next` — мапа возможных следующих задач: `map[Key]*Task`.

**Методы:**
- `Key() Key` — вернуть идентификатор задачи.
- `LinkNext(next *Task) error` — связать текущую задачу со следующей. Возвращает `ErrNextTaskNotProvided` если `next == nil`.
- `Run() error` — выполнить задачу и перейти к следующей на основе результата `ExecutorFunc`.

### `ITask`

Интерфейс с методами `Key()`, `LinkNext()`, `Run()`.

## Ошибки

- `ErrExecutorNotProvided` — не предоставлена функция-исполнитель.
- `ErrNextTaskNotProvided` — передан nil-задача при связывании.
- `ErrNextTaskNotFound` — возвращённый ключ не найден в мапе `next`.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/cascade"
)

func main() {
    // --- Пример 1: Линейная цепочка задач ---

    // Функции-исполнители
    step1 := func(data any) cascade.Key {
        fmt.Printf("Шаг 1: данные = %v\n", data)
        return "step2" // или cascade.ExitKey для завершения
    }

    step2 := func(data any) cascade.Key {
        fmt.Printf("Шаг 2: данные = %v\n", data)
        return cascade.ExitKey
    }

    // Создаём задачи
    t1, err := cascade.NewTask("hello", step1)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    t2, err := cascade.NewTask(42, step2)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // Связываем задачи
    err = t1.LinkNext(t2)
    if err != nil {
        fmt.Printf("ошибка связывания: %v\n", err)
        return
    }

    // Запускаем каскад
    err = t1.Run()
    // Вывод:
    // Шаг 1: данные = hello
    // Шаг 2: данные = 42
    // err == nil

    // --- Пример 2: Ветвление ---

    var choiceCount int

    processA := func(data any) cascade.Key {
        fmt.Printf("Обработка A: %v\n", data)
        return cascade.ExitKey
    }

    processB := func(data any) cascade.Key {
        fmt.Printf("Обработка B: %v\n", data)
        return cascade.ExitKey
    }

    router := func(data any) cascade.Key {
        choiceCount++
        // Ветвление по данным
        if data.(string) == "A" {
            return processA.Name()
        }
        return processB.Name()
    }

    taskA, _ := cascade.NewTask(nil, processA)
    taskB, _ := cascade.NewTask(nil, processB)

    mainTask, _ := cascade.NewTask("go to B", router)
    mainTask.LinkNext(taskA)
    mainTask.LinkNext(taskB)

    err = mainTask.Run()
    // Вывод: Обработка B: go to B

    // --- Пример 3: Цепочка с несколькими ветвями ---

    // Создаём задачу, которая возвращает несуществующий ключ
    badTask, _ := cascade.NewTask(nil, func(data any) cascade.Key {
        return "nonexistent"
    })

    err = badTask.Run()
    // err == ErrNextTaskNotFound
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
    }

    // --- Пример 4: Проверка ключа задачи ---

    fmt.Printf("ключ taskA: %v\n", taskA.Key())
    // ключ taskA: cascade.processA (полное имя функции)
}
```
