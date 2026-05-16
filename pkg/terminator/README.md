# terminator

Пакет `terminator` предоставляет механизм graceful shutdown сервисов при получении сигнала прерывания от операционной системы.

## Обзор

Пакет позволяет зарегистрировать несколько сервисов и при получении сигнала `SIGINT` или `SIGTERM` корректно остановить каждый из них.

## IGracefulShutDownService

Интерфейс сервиса, поддерживающего graceful shutdown:

```go
type IGracefulShutDownService interface {
    Run()
    GracefulShutDown()
}
```

## ITerminator

Интерфейс терминатора:

```go
type ITerminator interface {
    WaitOSInterruptSignalAndShutDown()
}
```

## Service

Структура терминатора:

```go
type Service struct {
    Services []terminated.IGracefulShutDownService
}
```

## New

Создаёт терминатор с заданным списком сервисов:

```go
import (
    "github.com/nobuenhombre/suikat/pkg/terminator"
    "github.com/nobuenhombre/suikat/pkg/terminator/terminated"
)

// Сервисы, которые нужно остановить
var services []terminated.IGracefulShutDownService
services = append(services, myService1)
services = append(services, myService2)

terminator := terminator.New(services)
```

## WaitOSInterruptSignalAndShutDown

Ожидает сигнал ОС и запускает graceful shutdown всех зарегистрированных сервисов:

```go
import (
    "log"
    "github.com/nobuenhombre/suikat/pkg/terminator"
    "github.com/nobuenhombre/suikat/pkg/terminator/terminated"
)

// Регистрация сервисов
var services []terminated.IGracefulShutDownService
services = append(services, service1)
services = append(services, service2)

terminator := terminator.New(services)

// Ожидание сигнала SIGINT или SIGTERM
terminator.WaitOSInterruptSignalAndShutDown()

log.Println("Все сервисы остановлены")
```

### Сигналы

| Сигнал | Команда | Описание |
|--------|---------|----------|
| `SIGINT` | `kill -2 <pid>` / `Ctrl+C` | Прерывание с клавиатуры |
| `SIGTERM` | `kill <pid>` | Запрос на завершение |

> `SIGKILL` (`kill -9`) не может быть перехвачен.

## Пример: полный цикл

```go
import (
    "log"
    "sync"
    "time"
    "github.com/nobuenhombre/suikat/pkg/terminator"
    "github.com/nobuenhombre/suikat/pkg/terminator/terminated"
)

// Реализация сервиса
type MyService struct {
    running bool
    wg      sync.WaitGroup
}

func (s *MyService) Run() {
    s.running = true
    s.wg.Add(1)
    go func() {
        defer s.wg.Done()
        for s.running {
            log.Println("Работа сервиса...")
            time.Sleep(1 * time.Second)
        }
    }()
}

func (s *MyService) GracefulShutDown() {
    log.Println("Graceful shutdown MyService...")
    s.running = false
    s.wg.Wait()
    log.Println("MyService остановлен")
}

func main() {
    // Создание сервисов
    service1 := &MyService{}
    service2 := &MyService{}

    service1.Run()
    service2.Run()

    // Создание терминатора
    services := []terminated.IGracefulShutDownService{
        service1,
        service2,
    }
    term := terminator.New(services)

    // Ожидание сигнала и graceful shutdown
    term.WaitOSInterruptSignalAndShutDown()

    log.Println("Все сервисы остановлены")
}
```

## Пример: с periodic-work

```go
import (
    "time"
    "sync"
    "github.com/nobuenhombre/suikat/pkg/periodic-work"
    "github.com/nobuenhombre/suikat/pkg/terminator"
    "github.com/nobuenhombre/suikat/pkg/terminator/terminated"
)

func myWorker(wg *sync.WaitGroup, counter *periodicwork.Counter, config interface{}) {
    defer func() {
        counter.IncExecuted()
        wg.Done()
    }()
    // рабочая логика
}

// Создание periodic-work сервиса
config := &periodicwork.Config{
    StopTimeout:     5 * time.Second,
    WorkGapInterval: 10 * time.Second,
    Worker:          myWorker,
}
periodicService := periodicwork.New(config)
periodicService.Run()

// Создание терминатора
services := []terminated.IGracefulShutDownService{
    periodicService,
}
term := terminator.New(services)

// Ожидание сигнала
term.WaitOSInterruptSignalAndShutDown()
```

## Пример: несколько сервисов

```go
import (
    "log"
    "github.com/nobuenhombre/suikat/pkg/terminator"
    "github.com/nobuenhombre/suikat/pkg/terminator/terminated"
)

type DatabaseService struct{}

func (s *DatabaseService) Run() {
    log.Println("DatabaseService запущен")
}

func (s *DatabaseService) GracefulShutDown() {
    log.Println("DatabaseService: закрытие пула соединений")
}

type CacheService struct{}

func (s *CacheService) Run() {
    log.Println("CacheService запущен")
}

func (s *CacheService) GracefulShutDown() {
    log.Println("CacheService: очистка кэша")
}

func main() {
    db := &DatabaseService{}
    cache := &CacheService{}

    db.Run()
    cache.Run()

    services := []terminated.IGracefulShutDownService{
        db,
        cache,
    }

    term := terminator.New(services)
    term.WaitOSInterruptSignalAndShutDown()

    log.Println("Все сервисы остановлены")
}
```

## Порядок shutdown

Сервисы останавливаются в порядке регистрации:

```go
services := []terminated.IGracefulShutDownService{
    serviceA, // остановится первым
    serviceB, // остановится вторым
    serviceC, // остановится третьим
}
term := terminator.New(services)
term.WaitOSInterruptSignalAndShutDown()
```
