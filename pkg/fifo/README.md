# fifo

Пакет `fifo` предоставляет работу с именованными каналами (FIFO) в Unix.

## Основные типы

### `Service`

Интерфейс для работы с FIFO:

```go
type Service interface {
    Create() error
    IsExists() (bool, error)
    Delete() error
    OpenToRead() error
    OpenToWrite() error
    Close() error
    Write(msg string) error
    Read(rcv MessageReceiver) error
}
```

### `MessageReceiver`

```go
type MessageReceiver func(string) error
```

Функция-обработчик прочитанных сообщений.

## Функция `New`

```go
func New(name string) Service
```

Создаёт сервис для FIFO с указанным именем.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/fifo"
)

func main() {
    fifoName := "/tmp/my-fifo"

    // --- Пример 1: Создание и удаление ---

    f := fifo.New(fifoName)

    // Создаём FIFO
    err := f.Create()
    if err != nil {
        fmt.Printf("ошибка создания: %v\n", err)
        return
    }

    // Проверяем существование
    exists, err := f.IsExists()
    if err != nil {
        fmt.Printf("ошибка проверки: %v\n", err)
        return
    }
    fmt.Printf("существует: %t\n", exists) // true

    // Удаляем FIFO
    err = f.Delete()
    if err != nil {
        fmt.Printf("ошибка удаления: %v\n", err)
        return
    }

    // --- Пример 2: Запись и чтение ---

    // Создаём FIFO
    f1 := fifo.New(fifoName)
    err = f1.Create()
    if err != nil {
        fmt.Printf("ошибка создания: %v\n", err)
        return
    }

    // Writer
    writer := fifo.New(fifoName)
    err = writer.OpenToWrite()
    if err != nil {
        fmt.Printf("ошибка открытия на запись: %v\n", err)
        return
    }

    go func() {
        for i := 1; i <= 5; i++ {
            writer.Write(fmt.Sprintf("message %d", i))
        }
        writer.Close()
    }()

    // Reader
    reader := fifo.New(fifoName)
    err = reader.OpenToRead()
    if err != nil {
        fmt.Printf("ошибка открытия на чтение: %v\n", err)
        return
    }

    rcv := func(msg string) error {
        fmt.Printf("получено: %s\n", msg)
        return nil
    }

    err = reader.Read(rcv)
    if err != nil {
        fmt.Printf("ошибка чтения: %v\n", err)
        return
    }

    reader.Close()
    f1.Delete()
}
```
