# osexec

Пакет `osexec` предоставляет утилиты для выполнения команд операционной системы с автоматическим определением платформы.

## Обзор

Пакет содержит функции для выполнения команд на Unix-системах (Linux, macOS) и Windows, а также универсальную функцию `OSRun`, которая автоматически определяет ОС и вызывает соответствующую функцию.

## OSRun

Выполняет команду с автоматическим определением операционной системы.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/osexec"
)

// На Linux/macOS
output, err := osexec.OSRun("ls", []string{"-la", "/tmp"})
if err != nil {
    panic(err)
}
fmt.Println(output)

// На Windows
output, err = osexec.OSRun("dir", []string{"C:\\tmp"})
if err != nil {
    panic(err)
}
fmt.Println(output)
```

## RunUnix

Выполняет команду на Unix-системах (Linux, macOS).

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/osexec"
)

// Получение информации о системе
output, err := osexec.RunUnix("uname", []string{"-a"})
if err != nil {
    panic(err)
}
fmt.Println(output)

// Проверка свободного места
output, err = osexec.RunUnix("df", []string{"-h", "/"})
if err != nil {
    panic(err)
}
fmt.Println(output)

// Получение uptime
output, err = osexec.RunUnix("uptime", nil)
if err != nil {
    panic(err)
}
fmt.Println(output)
```

## RunWindows

Выполняет команду на Windows через `CMD /c`.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/osexec"
)

// Информация о системе
output, err := osexec.RunWindows("systeminfo", nil)
if err != nil {
    panic(err)
}
fmt.Println(output)

// Список файлов
output, err = osexec.RunWindows("dir", []string{"C:\\", "/b"})
if err != nil {
    panic(err)
}
fmt.Println(output)

// Ping
output, err = osexec.RunWindows("ping", []string{"-n", "3", "google.com"})
if err != nil {
    panic(err)
}
fmt.Println(output)
```

## RunError

Структура ошибки, возвращаемая при неудачном выполнении команды.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/osexec"
)

output, err := osexec.RunUnix("nonexistent-command", nil)
if err != nil {
    if runErr, ok := err.(*osexec.RunError); ok {
        fmt.Printf("Команда: %s\n", runErr.Command)
        fmt.Printf("Аргументы: %v\n", runErr.Args)
        fmt.Printf("Стандартный вывод: %s\n", runErr.StdOut)
        fmt.Printf("Ошибка: %v\n", runErr.Parent)
    }
}
```

### Обработка ошибки от OSRun

```go
import (
    "fmt"
    osDetector "github.com/nobuenhombre/suikat/pkg/os-detector"
    "github.com/nobuenhombre/suikat/pkg/osexec"
)

output, err := osexec.OSRun("ls", []string{"/nonexistent"})
if err != nil {
    switch e := err.(type) {
    case *osexec.RunError:
        fmt.Printf("RunError: команда '%s' завершилась ошибкой\n", e.Command)
        fmt.Printf("StdOut: %s\n", e.StdOut)
        fmt.Printf("Parent: %v\n", e.Parent)

    case *osDetector.UnknownOSError:
        fmt.Printf("Неизвестная ОС: %s\n", e.Name)
    }
}
```

## Пример: кроссплатформенная проверка доступности команды

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/osexec"
)

func checkCommand(cmd string) bool {
    switch cmd {
    case "ls", "uname", "df", "uptime":
        // Unix-команды
        _, err := osexec.RunUnix(cmd, nil)
        return err == nil
    case "dir", "systeminfo", "ping":
        // Windows-команды
        _, err := osexec.RunWindows(cmd, nil)
        return err == nil
    default:
        // Универсальный вызов
        _, err := osexec.OSRun(cmd, nil)
        return err == nil
    }
}

fmt.Println(checkCommand("ls"))  // true на Linux
fmt.Println(checkCommand("dir")) // true на Windows
```
