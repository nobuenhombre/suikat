# os-detector

Пакет `os-detector` предоставляет константы для имён операционных систем и тип ошибки для неизвестных ОС.

## Обзор

Пакет содержит строковые константы для основных операционных систем, совместимые с именами из `runtime.GOOS`.

## Константы операционных систем

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/os-detector"
)

fmt.Println(osdetector.OSWindows) // "windows"
fmt.Println(osdetector.OSLinux)   // "linux"
fmt.Println(osdetector.OSMacOs)   // "darwin"
fmt.Println(osdetector.OSUnknown) // "unknown"
```

### Использование в условных выражениях

```go
import (
    "runtime"
    "github.com/nobuenhombre/suikat/pkg/os-detector"
)

func detectOS() string {
    goOS := runtime.GOOS
    
    switch goOS {
    case osdetector.OSWindows:
        return "Windows"
    case osdetector.OSLinux:
        return "Linux"
    case osdetector.OSMacOs:
        return "macOS"
    default:
        return osdetector.OSUnknown
    }
}

fmt.Println(detectOS())
```

### Проверка платформы

```go
import (
    "github.com/nobuenhombre/suikat/pkg/os-detector"
)

func isLinux() bool {
    return runtime.GOOS == osdetector.OSLinux
}

func isWindows() bool {
    return runtime.GOOS == osdetector.OSWindows
}

func isMacOS() bool {
    return runtime.GOOS == osdetector.OSMacOs
}
```

## UnknownOSError

Тип ошибки для неизвестной операционной системы.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/os-detector"
)

func validateOS(name string) error {
    switch name {
    case osdetector.OSWindows, osdetector.OSLinux, osdetector.OSMacOs:
        return nil
    default:
        return &osdetector.UnknownOSError{Name: name}
    }
}

err := validateOS("freebsd")
if err != nil {
    fmt.Println(err) // Unknown OS (name = freebsd)
}
```
