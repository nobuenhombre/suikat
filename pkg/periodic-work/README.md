# periodic-work

Пакет `periodic-work` предоставляет механизм для периодического выполнения рабочих функций с подсчётом запусков, выполнений и graceful shutdown.

## Обзор

Пакет позволяет запускать рабочую функцию через заданные интервалы времени, отслеживать количество запусков и выполнений, а также корректно останавливать сервис.

## Service

Интерфейс сервиса:

```go
type Service interface {
    Run()              // Запуск периодического выполнения
    Stop()             // Остановка
    GracefulShutDown() // Graceful shutdown по сигналу ОС
    GetStatus() error  // Проверка статуса (Started == Executed)
}
```

## Counter

Счётчик запусков и выполнений:

```go
type Counter struct {
    Started  int64 // Количество запусков
    Executed int64 // Количество выполнений
}

func (c *Counter) IncStarted()   // Увеличить Started
func (c *Counter) IncExecuted()  // Увеличить Executed
func (c *Counter) Report()       // Вывести отчёт в log
```

## Config

Конфигурация сервиса:

```go
type Config struct {
    StopTimeout      time.Duration // Таймаут остановки
    WorkGapInterval  time.Duration // Интервал между запусками
    Worker           Worker        // Рабочая функция
    WorkerConfig     interface{}   // Параметры для Worker
}
```

## Worker

Тип рабочей функции:

```go
type Worker func(wg *sync.WaitGroup, counter *Counter, config interface{})
```

- `wg` — WaitGroup для отслеживания завершения
- `counter` — счётчик для инкремента после выполнения
- `config` — пользовательские параметры из `WorkerConfig`

## New

Создаёт сервис из конфигурации:

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

config := &periodicwork.Config{
    StopTimeout:     5 * time.Second,
    WorkGapInterval: 10 * time.Second,
    Worker:          myWorker,
    WorkerConfig:    "some-config",
}

service := periodicwork.New(config)
```

## Run

Запускает периодическое выполнение рабочей функции:

```go
import (
    "log"
    "sync"
    "time"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

func myWorker(wg *sync.WaitGroup, counter *periodicwork.Counter, config interface{}) {
    defer func() {
        counter.IncExecuted()
        wg.Done()
    }()

    log.Println("Выполнение работы...")
    // ... рабочая логика ...
}

config := &periodicwork.Config{
    StopTimeout:     5 * time.Second,
    WorkGapInterval: 10 * time.Second,
    Worker:          myWorker,
}

service := periodicwork.New(config)
service.Run()
```

## Stop

Останавливает сервис:

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

// Запуск
service.Run()

// ... через какое-то время ...

// Остановка
time.Sleep(60 * time.Second)
service.Stop()
```

## GracefulShutDown

Ожидает сигнал ОС (SIGTERM/SIGINT) и затем останавливает сервис:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

config := &periodicwork.Config{
    StopTimeout:     5 * time.Second,
    WorkGapInterval: 10 * time.Second,
    Worker:          myWorker,
}

service := periodicwork.New(config)
service.Run()

// Ожидание сигнала SIGTERM или SIGINT
service.GracefulShutDown()

log.Println("Сервис остановлен")
```

## GetStatus

Проверяет, что все запущенные задачи были выполнены:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

service.Stop()

err := service.GetStatus()
if err != nil {
    fmt.Printf("Ошибка: %v\n", err)
    // Started != Executed — некоторые задачи не завершены
} else {
    fmt.Println("Все задачи завершены успешно")
}
```

## Пример: полный цикл

```go
import (
    "fmt"
    "log"
    "sync"
    "time"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

type WorkerConfig struct {
    TaskID string
    Retries int
}

func taskWorker(wg *sync.WaitGroup, counter *periodicwork.Counter, config interface{}) {
    defer func() {
        counter.IncExecuted()
        wg.Done()
    }()

    cfg := config.(WorkerConfig)
    log.Printf("Задача %s: выполнение (retries=%d)\n", cfg.TaskID, cfg.Retries)
    time.Sleep(2 * time.Second)
}

func main() {
    config := &periodicwork.Config{
        StopTimeout:     5 * time.Second,
        WorkGapInterval: 3 * time.Second,
        Worker:          taskWorker,
        WorkerConfig:    WorkerConfig{TaskID: "task-1", Retries: 3},
    }

    service := periodicwork.New(config)
    service.Run()

    // Запуск 60 секунд
    time.Sleep(60 * time.Second)

    // Остановка
    service.Stop()

    // Проверка статуса
    err := service.GetStatus()
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
    } else {
        fmt.Println("Все задачи завершены")
    }
}
```

## Пример: мониторинг с graceful shutdown

```go
import (
    "log"
    "sync"
    "time"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

func monitorWorker(wg *sync.WaitGroup, counter *periodicwork.Counter, config interface{}) {
    defer func() {
        counter.IncExecuted()
        wg.Done()
    }()

    log.Println("Сбор метрик...")
    // Сбор метрик, отправка в мониторинг и т.д.
}

func main() {
    config := &periodicwork.Config{
        StopTimeout:     10 * time.Second,
        WorkGapInterval: 30 * time.Second,
        Worker:          monitorWorker,
    }

    service := periodicwork.New(config)
    service.Run()

    // Graceful shutdown по сигналу
    service.GracefulShutDown()
}
```

## Пример: отчёт счётчика

```go
import (
    "log"
    "sync"
    "time"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
)

func countedWorker(wg *sync.WaitGroup, counter *periodicwork.Counter, config interface{}) {
    defer func() {
        counter.IncExecuted()
        wg.Done()
    }()

    log.Println("Работа выполнена")
}

config := &periodicwork.Config{
    StopTimeout:     5 * time.Second,
    WorkGapInterval: 1 * time.Second,
    Worker:          countedWorker,
}

service := periodicwork.New(config)
service.Run()

time.Sleep(5 * time.Second)
service.Stop()

// В логе будет:
// log.Printf("Counter Started: %d, Executed: %d\n", ...)
```
