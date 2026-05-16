# hash

Пакет `hash` предоставляет функции для хеширования строк.

## Функции

### `MD5(s string) string`

Вычисляет MD5-хеш строки. Возвращает hex-строку.

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/hash"
)

func main() {
    h := hash.MD5("hello")
    fmt.Println(h) // 5d41402abc4b2a76b9719d911017c592
}
```
