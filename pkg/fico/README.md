# fico

Пакет `fico` предоставляет доступ к содержимому файлов: чтение, запись, сжатие (GZ/BR), кодирование (Base64/HEX/DataURI), работа с пачкой файлов.

## Зависимости

- `github.com/mholt/archiver/v3` — архивирование (GZ, Brotli)
- `github.com/andybalholm/brotli` — сжатие Brotli
- `github.com/nobuenhombre/suikat/pkg/fina` — получение частей имени файла
- `github.com/nobuenhombre/suikat/pkg/mimes` — MIME-типы по расширению
- `github.com/nobuenhombre/suikat/pkg/ge` — ошибки

## Константы

| Константа | Значение |
|-----------|----------|
| `ExtGZ` | `.gz` |
| `ExtBR` | `.br` |
| `ExtB64` | `.b64` |
| `ExtHEX` | `.hex` |

## Тип `TxtFile`

```go
type TxtFile string
```

Путь к файлу в виде строки.

### Методы чтения

- `ReadBytes() ([]byte, error)` — прочитать содержимое как `[]byte`.
- `Read() (string, error)` — прочитать содержимое как `string`.
- `ReadAsB64String() (string, error)` — прочитать и закодировать в Base64.
- `ReadAsHexString() (string, error)` — прочитать и закодировать в HEX.
- `DataURI() (string, error)` — получить data URI (MIME + Base64).

### Методы записи

- `Write(content string) error` — записать строку в файл.
- `WriteJSON(content interface{}) error` — записать JSON.
- `WriteGZ(content string) error` — записать сжатый GZ.
- `WriteBR(content string) error` — записать сжатый Brotli.
- `WriteAndCompress(content string) error` — записать оригинал + GZ + BR.

### Методы трансформации

- `GZ() error` — прочитать файл и создать `.gz` копию.
- `BR() error` — прочитать файл и создать `.br` копию.
- `B64() error` — прочитать файл и создать `.b64` копию.
- `Hex() error` — прочитать файл и создать `.hex` копию.

## Тип `TxtFilesPack`

```go
type TxtFilesPack map[string]string
```

Пачка файлов: имя файла → содержимое.

**Методы:**
- `Read() error` — прочитать все файлы.
- `Write() error` — записать все файлы.
- `WriteGZ() error` — записать все файлы как GZ.
- `Remove() error` — удалить все файлы.

## Функция `StrBytes`

```go
func StrBytes(in, glue string, isUpper bool) string
```

Конвертирует строку в HEX с разделителем `glue`. Если `isUpper == true` — верхний регистр.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/fico"
)

func main() {
    // --- Пример 1: Чтение файла ---

    f := fico.TxtFile("/tmp/test.txt")

    // Запись
    err := f.Write("Hello, World!")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // Чтение
    content, err := f.Read()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", content) // Hello, World!

    // Чтение как []byte
    data, err := f.ReadBytes()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", data)

    // --- Пример 2: JSON ---

    type Config struct {
        Port int    `json:"port"`
        Host string `json:"host"`
    }

    cfg := Config{Port: 8080, Host: "localhost"}
    err = f.WriteJSON(cfg)
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // --- Пример 3: Сжатие GZ ---

    err = f.Write("data to compress")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    err = f.GZ()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Создан файл /tmp/test.txt.gz

    // --- Пример 4: Сжатие BR ---

    err = f.BR()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Создан файл /tmp/test.txt.br

    // --- Пример 5: WriteAndCompress ---

    err = f.WriteAndCompress("data")
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Созданы: test.txt, test.txt.gz, test.txt.br

    // --- Пример 6: Base64 ---

    err = f.B64()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Создан файл /tmp/test.txt.b64

    // Чтение Base64
    b64str, err := f.ReadAsB64String()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", b64str)

    // --- Пример 7: HEX ---

    err = f.Hex()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    // Создан файл /tmp/test.txt.hex

    hexstr, err := f.ReadAsHexString()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", hexstr)

    // --- Пример 8: DataURI ---

    uri, err := f.DataURI()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%s\n", uri)
    // data:text/plain;base64,...

    // --- Пример 9: TxtFilesPack ---

    pack := fico.TxtFilesPack{
        "/tmp/file1.txt": "content 1",
        "/tmp/file2.txt": "content 2",
    }

    err = pack.Write()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    err = pack.Read()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Printf("%v\n", pack)

    err = pack.WriteGZ()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    err = pack.Remove()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }

    // --- Пример 10: StrBytes ---

    hex := fico.StrBytes("abc", " ", true)
    fmt.Printf("%s\n", hex) // 61 62 63
}
```
