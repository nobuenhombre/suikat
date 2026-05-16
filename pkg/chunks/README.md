# chunks

Пакет `chunks` предоставляет функции для разбиения одномерных слайсов на части, реверсирования и обмена половинами.

## Функции

### `SplitInt64(in []int64, limit int) [][]int64`

Разбивает `[]int64` на части размером `limit`. Последний кусок может быть меньше.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/chunks"
)

// Пример: разбиение на куски по 3 элемента
data := []int64{1, 2, 3, 4, 5, 6, 7}
result := chunks.SplitInt64(data, 3)
// result == [][]int64{
//     {1, 2, 3},
//     {4, 5, 6},
//     {7},
// }
fmt.Printf("%v\n", result)
```

### `SplitStr(s string, limit int) []string`

Разбивает строку на части длиной `limit`. Учитывает многобайтовые символы (Unicode).

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/chunks"
)

// Пример: разбиение строки на куски по 5 символов
s := "Hello, World!"
result := chunks.SplitStr(s, 5)
// result == []string{"Hello", ", Wor", "ld!"}
fmt.Printf("%v\n", result)

// Пустая строка
result2 := chunks.SplitStr("", 5)
// result2 == nil

// Строка короче limit
result3 := chunks.SplitStr("Hi", 10)
// result3 == []string{"Hi"}
```

### `SplitBytes(in []byte, limit int) [][]byte`

Разбивает `[]byte` на части размером `limit`. Последний кусок может быть меньше.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/chunks"
)

data := []byte{1, 2, 3, 4, 5, 6, 7}
result := chunks.SplitBytes(data, 3)
// result == [][]byte{{1, 2, 3}, {4, 5, 6}, {7}}
fmt.Printf("%v\n", result)
```

### `ReverseFullBytes(in []byte) []byte`

Реверсирует слайс `[]byte` в обратном порядке. Рекурсивная реализация.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/chunks"
)

data := []byte{1, 2, 3, 4, 5}
result := chunks.ReverseFullBytes(data)
// result == []byte{5, 4, 3, 2, 1}
fmt.Printf("%v\n", result)

// Пустой слайс
result2 := chunks.ReverseFullBytes([]byte{})
// result2 == []byte{}
```

### `SwapHalfBytes(in []byte) ([]byte, error)`

Меняет две половины `[]byte` местами. Длина должна быть чётной.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/chunks"
)

// Пример: чётная длина
data := []byte{1, 2, 3, 4, 5, 6}
result, err := chunks.SwapHalfBytes(data)
// result == []byte{4, 5, 6, 1, 2, 3}
// err == nil
fmt.Printf("%v\n", result)

// Пример: нечётная длина — ошибка
badData := []byte{1, 2, 3}
_, err = chunks.SwapHalfBytes(badData)
// err != nil — cant divide slice by two parts len(in)=3
fmt.Printf("ошибка: %v\n", err)
```
