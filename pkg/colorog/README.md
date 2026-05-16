# colorog

Пакет `colorog` предоставляет цветное логирование с настраиваемой палитрой цветов.

## Зависимости

- `github.com/fatih/color` — ANSI-цветная печать

## Основные типы

### `Palette`

Палитра цветов:
- `TimeColor` — цвет времени (magenta по умолчанию)
- `SuccessColor` — цвет успеха (green)
- `InfoColor` — цвет информации (cyan)
- `MessageColor` — цвет сообщения (white)
- `ErrorColor` — цвет ошибки (yellow)
- `FatalColor` — цвет фатальной ошибки (red)
- `PanicColor` — цвет паники (white на красном фоне)

### `ColoredLog`

Структура логгера:
- `Palette` — палитра цветов
- `ShowTime` — показывать ли время в логах
- `TimeFormat` — формат времени (по умолчанию `"2006-01-02 15:04:05"`)

**Методы:**
- `Success(...)` / `Successf(format, ...)` / `Successln(...)` — зелёные сообщения
- `Error(...)` / `Errorf(format, ...)` / `Errorln(...)` — жёлтые сообщения
- `Info(...)` / `Infof(format, ...)` / `Infoln(...)` — голубые сообщения
- `Message(...)` / `Messagef(format, ...)` / `Messageln(...)` — белые сообщения
- `Fatal(...)` / `Fatalf(format, ...)` / `Fatalln(...)` — красные сообщения (exit 1)
- `Panic(...)` / `Panicf(format, ...)` / `Panicln(...)` — белые на красном (panic)

## Функция `NewColoredLog`

```go
func NewColoredLog(showTime bool, timeFormat string) *ColoredLog
```

Создаёт цветной логгер и устанавливает его как output для `log` пакета.

## Примеры

```go
package main

import (
    "github.com/nobuenhombre/suikat/pkg/colorog"
)

func main() {
    // --- Пример 1: Логгер с временем ---

    logger := colorog.NewColoredLog(true, "2006-01-02 15:04:05")

    logger.Info("сервер запущен")
    // [cyan]сервер запущен[/cyan]

    logger.Success("запрос выполнен успешно")
    // [green]запрос выполнен успешно[/green]

    logger.Error("не удалось подключиться к БД")
    // [yellow]не удалось подключиться к БД[/yellow]

    logger.Message("обычное сообщение")
    // [white]обычное сообщение[/white]

    // --- Пример 2: С форматированием ---

    port := 8080
    logger.Infof("запуск на порту %d", port)
    // [cyan]запуск на порту 8080[/cyan]

    items := 42
    logger.Successf("обработано %d элементов", items)
    // [green]обработано 42 элементов[/green]

    // --- Пример 3: Логгер без времени ---

    logger2 := colorog.NewColoredLog(false, "")
    logger2.Info("без времени")

    // --- Пример 4: Fatal и Panic ---

    // logger.Fatal("фатальная ошибка")
    // // [red]фатальная ошибка[/red] + os.Exit(1)

    // logger.Panic("паника!")
    // // white на красном + panic
}
```
