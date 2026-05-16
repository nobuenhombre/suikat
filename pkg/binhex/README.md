# binhex

Пакет `binhex` предоставляет типы для кодирования и декодирования между строковым представлением и HEX.

## Типы

### `BinString`

Тип `string` для бинарных данных в строковом представлении.

**Методы:**

- `ToHex() HexString` — конвертировать в HEX-строку.
- `ToString() string` — вернуть строковое представление.

### `HexString`

Тип `string` для HEX-кодированных данных.

**Методы:**

- `ToBin() (BinString, error)` — декодировать из HEX в строку.
- `ToString() string` — вернуть строковое представление.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/binhex"
)

func main() {
    // --- BinString → HEX ---

    // Пример 1: кодирование строки в HEX
    original := binhex.BinString("Hello World")
    hexStr := original.ToHex()
    fmt.Println(hexStr) // 48656c6c6f20576f726c64

    // Пример 2: преобразование пустой строки
    empty := binhex.BinString("")
    fmt.Println(empty.ToHex()) // (пустая строка)

    // --- HEX → BinString ---

    // Пример 3: декодирование HEX в строку
    hexInput := binhex.HexString("48656c6c6f20576f726c64")
    binStr, err := hexInput.ToBin()
    if err != nil {
        fmt.Printf("ошибка: %v\n", err)
        return
    }
    fmt.Println(binStr) // Hello World

    // Пример 4: некорректный HEX — ошибка
    badHex := binhex.HexString("48656c6c6f20576f726c6") // нечётная длина
    _, err = badHex.ToBin()
    if err != nil {
        fmt.Printf("ошибка декодирования: %v\n", err)
        // ошибка декодирования: odd length hex string
    }

    // --- ToString ---

    // Пример 5: получение строки из типа
    fmt.Println(hexInput.ToString()) // 48656c6c6f20576f726c64
    fmt.Println(binStr.ToString())   // Hello World
}
```
