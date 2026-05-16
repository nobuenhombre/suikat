# credentials

Пакет `credentials` предоставляет функции для хранения и чтения содержимого файлов из переменных окружения. Файл кодируется в base64 и хранится в переменной окружения, затем может быть прочитан как `[]byte`, `string` или записан во временный файл.

## Зависимости

- `encoding/base64` — кодирование/декодирование
- `github.com/nobuenhombre/suikat/pkg/futi` — создание временных файлов

## Основные типы

### `Data`

```go
type Data struct {
    Key string
}
```

### `EnvCred`

Интерфейс:
- `GetBytes() ([]byte, error)` — получить содержимое как `[]byte`.
- `GetString() (string, error)` — получить содержимое как `string`.
- `GetFile() (string, error)` — записать содержимое во временный файл, вернуть путь.

## Функция `New`

```go
func New(key string) EnvCred
```

Создаёт `EnvCred` с именем переменной окружения.

## Примеры

```go
package main

import (
    "fmt"
    "os"
    "github.com/nobuenhombre/suikat/pkg/credentials"
)

func main() {
    // Предположим, у нас есть credentials.json:
    // {"token": "abc123", "secret": "xyz789"}

    // В Docker/GKE это хранится как base64 в переменной окружения
    // В примере устанавливаем вручную:
    import "encoding/base64"
    original := `{"token": "abc123", "secret": "xyz789"}`
    os.Setenv("MY_CREDS", base64.StdEncoding.EncodeToString([]byte(original)))

    // --- Пример 1: Получение []byte ---

    creds := credentials.New("MY_CREDS")
    data, err := creds.GetBytes()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", data) // {"token": "abc123", "secret": "xyz789"}

    // --- Пример 2: Получение string ---

    str, err := creds.GetString()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", str) // {"token": "abc123", "secret": "xyz789"}

    // --- Пример 3: Получение временного файла ---

    filePath, err := creds.GetFile()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("файл: %s\n", filePath) // /tmp/credentials123456
    // Файл можно использовать напрямую в приложении
}
```
