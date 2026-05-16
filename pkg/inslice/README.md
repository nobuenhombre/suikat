# inslice

Пакет `inslice` предоставляет функции для поиска элементов в срезах (slices) и проверки существования индексов.

## Обзор

Пакет содержит типизированные функции поиска значения в срезе с использованием указателей на срез. Все функции безопасно обрабатывают `nil`-срез, возвращая `false`.

## Функции

### Поиск элементов

#### `String`

Проверяет, содержится ли строка в срезе строк.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

names := []string{"alice", "bob", "charlie"}

fmt.Println(inslice.String("bob", &names)) // true
fmt.Println(inslice.String("dave", &names)) // false
```

#### `Int`

Проверяет, содержится ли int в срезе int.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

numbers := []int{1, 2, 3, 4, 5}

fmt.Println(inslice.Int(3, &numbers)) // true
fmt.Println(inslice.Int(10, &numbers)) // false
```

#### `Int32`

Проверяет, содержится ли int32 в срезе int32.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

values := []int32{100, 200, 300}

fmt.Println(inslice.Int32(200, &values)) // true
fmt.Println(inslice.Int32(400, &values)) // false
```

#### `Int64`

Проверяет, содержится ли int64 в срезе int64.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

values := []int64{1000, 2000, 3000}

fmt.Println(inslice.Int64(2000, &values)) // true
fmt.Println(inslice.Int64(4000, &values)) // false
```

#### `Float32`

Проверяет, содержится ли float32 в срезе float32.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

values := []float32{1.1, 2.2, 3.3}

fmt.Println(inslice.Float32(2.2, &values)) // true
fmt.Println(inslice.Float32(4.4, &values)) // false
```

#### `Float64`

Проверяет, содержится ли float64 в срезе float64.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

values := []float64{1.11, 2.22, 3.33}

fmt.Println(inslice.Float64(2.22, &values)) // true
fmt.Println(inslice.Float64(4.44, &values)) // false
```

### Проверка индексов

#### `IsIndexExists`

Проверяет, существует ли указанный индекс в срезе любого типа. Возвращает `false`, если переданный аргумент не является срезом.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

data := []string{"a", "b", "c"}

fmt.Println(inslice.IsIndexExists(1, data)) // true
fmt.Println(inslice.IsIndexExists(5, data)) // false
fmt.Println(inslice.IsIndexExists(-1, data)) // false

// Работает с любым типом среза
fmt.Println(inslice.IsIndexExists(0, []int{10, 20})) // true
fmt.Println(inslice.IsIndexExists(0, 42)) // false (не срез)
```

#### `CheckIndex`

Проверяет, существует ли указанный индекс в срезе. Если индекс не существует — возвращает ошибку `*IndexNotExistsError`.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

data := []string{"a", "b", "c"}

err := inslice.CheckIndex(1, data)
fmt.Println(err) // nil

err = inslice.CheckIndex(5, data)
fmt.Println(err) // index [5] not exists

err = inslice.CheckIndex(-1, data)
fmt.Println(err) // index [-1] not exists
```

#### `IndexNotExistsError`

Структура ошибки, возвращаемой функцией `CheckIndex`.

```go
import "github.com/bookworker06JAN1979/Sources/golang.app/suikat/pkg/inslice"

data := []int{1, 2, 3}
err := inslice.CheckIndex(10, data)

if err != nil {
    if idxErr, ok := err.(*inslice.IndexNotExistsError); ok {
        fmt.Printf("Индекс %d не существует\n", idxErr.Index)
        // Вывод: Индекс 10 не существует
    }
}
```
