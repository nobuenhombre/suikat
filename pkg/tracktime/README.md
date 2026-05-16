# tracktime

Пакет `tracktime` предоставляет структуру для замера времени выполнения блоков кода и логирования длительности.

## Обзор

Пакет позволяет измерить время между точками запуска и остановки кода и записать результат в лог.

## Tracker

Структура для хранения информации о таймере:

```go
type Tracker struct {
    Label    string
    Run      time.Time
    Finish   time.Time
    Duration time.Duration
}
```

## Start

Создаёт новый таймер с заданной меткой:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/tracktime"
)

timer := tracktime.Start("Operation name")
```

## Stop

Фиксирует время остановки и вычисляет длительность:

```go
timer.Stop()
```

## Log

Выводит информацию о таймере в лог:

```go
timer.Log()
// Вывод:
// track time:
//  - label [Operation name]
//  - run at [2024-01-01 12:00:00.000]
//  - finish at [2024-01-01 12:00:01.500]
//  - duration [1.5s]
```

## Базовое использование

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/tracktime"
)

func process() {
    timer := tracktime.Start("process function")
    defer func() {
        timer.Stop()
        timer.Log()
    }()

    // ... рабочая логика ...
    time.Sleep(100 * time.Millisecond)
}

process()
// Вывод в лог:
// track time:
//  - label [process function]
//  - run at [2024-01-01 12:00:00.000]
//  - finish at [2024-01-01 12:00:00.100]
//  - duration [100ms]
```

## Пример: обработка данных

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/tracktime"
)

func processData(data []int) []int {
    timer := tracktime.Start("processData")
    defer func() {
        timer.Stop()
        timer.Log()
    }()

    result := make([]int, len(data))
    for i, v := range data {
        result[i] = v * 2
    }

    time.Sleep(50 * time.Millisecond)

    return result
}

processData([]int{1, 2, 3, 4, 5})
```

## Пример: несколько таймеров

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/tracktime"
)

func complexOperation() {
    // Таймер для всей операции
    totalTimer := tracktime.Start("total operation")
    defer func() {
        totalTimer.Stop()
        totalTimer.Log()
    }()

    // Подзадача 1
    step1 := tracktime.Start("step 1")
    time.Sleep(100 * time.Millisecond)
    step1.Stop()
    step1.Log()

    // Подзадача 2
    step2 := tracktime.Start("step 2")
    time.Sleep(200 * time.Millisecond)
    step2.Stop()
    step2.Log()

    // Подзадача 3
    step3 := tracktime.Start("step 3")
    time.Sleep(50 * time.Millisecond)
    step3.Stop()
    step3.Log()
}
```

## Пример: ручной контроль

```go
import (
    "fmt"
    "time"
    "github.com/nobuenhombre/suikat/pkg/tracktime"
)

func main() {
    timer := tracktime.Start("manual control")

    // ... какая-то работа ...
    time.Sleep(100 * time.Millisecond)

    // Ручная остановка
    timer.Stop()

    // Доступ к полям
    fmt.Printf("Метка: %s\n", timer.Label)
    fmt.Printf("Запуск: %v\n", timer.Run)
    fmt.Printf("Завершение: %v\n", timer.Finish)
    fmt.Printf("Длительность: %v\n", timer.Duration)

    // Или логирование
    timer.Log()
}
```

## Пример: с defer для гарантии остановки

```go
import (
    "github.com/nobuenhombre/suikat/pkg/tracktime"
)

func fetchData() {
    timer := tracktime.Start("fetchData")
    defer func() {
        timer.Stop()
        timer.Log()
    }()

    // ... загрузка данных ...
}

func processData() {
    timer := tracktime.Start("processData")
    defer func() {
        timer.Stop()
        timer.Log()
    }()

    // ... обработка данных ...
}

func saveData() {
    timer := tracktime.Start("saveData")
    defer func() {
        timer.Stop()
        timer.Log()
    }()

    // ... сохранение данных ...
}

func pipeline() {
    fetchData()
    processData()
    saveData()
}
```
