# repeater

Пакет `repeater` предоставляет механизм повторного вызова функции с ограничением количества попыток и таймаутом между повторами.

## Обзор

Пакет полезен при взаимодействии с внешними API, где запрос может не выполниться с первого раза (временные сбои, длительные операции). Репитер повторяет вызов воркера до тех пор, пока чекер не вернёт `true` или не будет исчерпан лимит попыток.

## WorkerResult

Результат выполнения воркера:

```go
type WorkerResult struct {
    OutData interface{}
    Err     error
}
```

## Worker

Тип функции-воркера:

```go
type Worker func(data interface{}) WorkerResult
```

Принимает входные данные и возвращает результат.

## Checker

Тип функции-чекера (проверка условия завершения):

```go
type Checker func(wr WorkerResult) (bool, error)
```

Возвращает:
- `true` — условие выполнено, можно вернуть результат
- `false` — условие не выполнено, повторить попытку
- `error` — ошибка проверки

## Config

Конфигурация репитера:

```go
type Config struct {
    LimitCount int64         // Максимальное количество попыток
    Timeout    time.Duration // Таймаут между попытками
}
```

## Run

Запускает воркер с повторами:

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/repeater"
)

// Воркер: попытка выполнить запрос
func myWorker(data interface{}) repeater.WorkerResult {
    // ... выполнение запроса ...
    result := "some-data"
    err := error(nil)
    return repeater.WorkerResult{
        OutData: result,
        Err:     err,
    }
}

// Чекер: проверка результата
func myChecker(wr repeater.WorkerResult) (bool, error) {
    if wr.Err != nil {
        return false, nil // повторить
    }
    return true, nil // успех
}

config := repeater.Config{
    LimitCount: 5,
    Timeout:    1 * time.Second,
}

result, err := myWorker.Run("input-data", myChecker, config)
if err != nil {
    panic(err)
}
fmt.Println(result) // some-data
```

## Пример: запрос к API с повторами

```go
import (
    "fmt"
    "net/http"
    "time"
    "github.com/nobuenhombre/suikat/pkg/repeater"
)

// URL для запроса
type APIRequest struct {
    URL string
}

// Воркер: HTTP-запрос
func apiWorker(data interface{}) repeater.WorkerResult {
    req := data.(APIRequest)
    resp, err := http.Get(req.URL)
    if err != nil {
        return repeater.WorkerResult{Err: err}
    }
    defer resp.Body.Close()

    return repeater.WorkerResult{
        OutData: resp.StatusCode,
    }
}

// Чекер: проверка статуса
func apiChecker(wr repeater.WorkerResult) (bool, error) {
    if wr.Err != nil {
        return false, nil // повторить
    }
    status := wr.OutData.(int)
    return status == 200, nil // повторять, пока не получим 200
}

req := APIRequest{URL: "https://api.example.com/status"}

config := repeater.Config{
    LimitCount: 10,
    Timeout:    2 * time.Second,
}

statusCode, err := apiWorker.Run(req, apiChecker, config)
if err != nil {
    fmt.Printf("Не удалось получить ответ: %v\n", err)
} else {
    fmt.Printf("Статус: %d\n", statusCode)
}
```

## Пример: ожидание завершения асинхронной операции

```go
import (
    "fmt"
    "time"
    "github.com/nobuenhombre/suikat/pkg/repeater"
)

// Имитация проверки статуса задачи
type TaskStatus struct {
    ID     string
    Status string
}

func checkTaskWorker(data interface{}) repeater.WorkerResult {
    taskID := data.(string)
    // Имитация проверки
    status := "processing" // или "done" в реальности
    if status == "done" {
        return repeater.WorkerResult{OutData: "Задача выполнена"}
    }
    return repeater.WorkerResult{OutData: "Задача выполняется..."}
}

func checkTaskChecker(wr repeater.WorkerResult) (bool, error) {
    result := wr.OutData.(string)
    return result == "Задача выполнена", nil
}

config := repeater.Config{
    LimitCount: 20,
    Timeout:    5 * time.Second,
}

result, err := checkTaskWorker.Run("task-123", checkTaskChecker, config)
if err != nil {
    fmt.Printf("Ожидание завершено: %v\n", err)
} else {
    fmt.Println(result) // Задача выполнена
}
```

## Пример: обработка ошибок чекера

```go
import (
    "fmt"
    "time"
    "github.com/nobuenhombre/suikat/pkg/repeater"
)

func safeWorker(data interface{}) repeater.WorkerResult {
    return repeater.WorkerResult{
        OutData: "result",
    }
}

func strictChecker(wr repeater.WorkerResult) (bool, error) {
    if wr.OutData == nil {
        return false, fmt.Errorf("данные отсутствуют")
    }
    return true, nil
}

config := repeater.Config{
    LimitCount: 3,
    Timeout:    1 * time.Second,
}

result, err := safeWorker.Run("data", strictChecker, config)
if err != nil {
    fmt.Printf("Ошибка: %v\n", err)
}
```

## Ошибки

При исчерпании лимита попыток возвращается `*ge.LimitCountExhaustedError`:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/ge"
    "github.com/nobuenhombre/suikat/pkg/repeater"
)

result, err := myWorker.Run("data", myChecker, config)
if err != nil {
    if _, ok := err.(*ge.LimitCountExhaustedError); ok {
        fmt.Println("Исчерпан лимит попыток")
    } else {
        fmt.Printf("Ошибка: %v\n", err)
    }
}
```
