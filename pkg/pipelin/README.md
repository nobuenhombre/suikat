# pipelin

Пакет `pipelin` предоставляет механизм конвейерной обработки данных через цепочку функций.

## Обзор

Пакет позволяет связать несколько функций (Work) в конвейер, где выход каждой функции является входом для следующей. Все функции выполняются параллельно через goroutine.

```
in -> [A] -> [B] -> [C] -> [D] -> out
```

## Work

Тип функции-шага конвейера:

```go
type Work func(in, out chan interface{})
```

- `in` — канал для входящих данных
- `out` — канал для исходящих данных

## WorkPipeline

Срез функций-шагов:

```go
type WorkPipeline []Work
```

## Run

Запускает конвейер. Все шаги выполняются параллельно.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/pipelin"
)

// Шаг 1: генерация чисел
func generate(in, out chan interface{}) {
    defer close(out)
    for i := 0; i < 10; i++ {
        out <- i
    }
}

// Шаг 2: удвоение
func double(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        num := data.(int)
        out <- num * 2
    }
}

// Шаг 3: форматирование
func format(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        num := data.(int)
        out <- fmt.Sprintf("Result: %d", num)
    }
}

// Создание конвейера
pipeline := pipelin.WorkPipeline{
    generate,
    double,
    format,
}

pipeline.Run()
```

## Пример: обработка строк

```go
import (
    "fmt"
    "strings"
    "github.com/nobuenhombre/suikat/pkg/pipelin"
)

// Шаг 1: генерация строк
func stringGen(in, out chan interface{}) {
    defer close(out)
    words := []string{"hello", "world", "golang", "pipeline"}
    for _, w := range words {
        out <- w
    }
}

// Шаг 2: upper case
func toUpper(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        s := data.(string)
        out <- strings.ToUpper(s)
    }
}

// Шаг 3: добавление префикса
func addPrefix(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        s := data.(string)
        out <- ">>> " + s
    }
}

// Шаг 4: вывод результата
func printResult(in, out chan interface{}) {
    for data := range in {
        fmt.Println(data.(string))
    }
}

pipeline := pipelin.WorkPipeline{
    stringGen,
    toUpper,
    addPrefix,
    printResult,
}

pipeline.Run()
// Вывод:
// >>> HELLO
// >>> WORLD
// >>> GOLANG
// >>> PIPELINE
```

## Пример: фильтрация данных

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/pipelin"
)

// Генерация чисел
func numberGen(in, out chan interface{}) {
    defer close(out)
    for i := 0; i < 20; i++ {
        out <- i
    }
}

// Фильтрация чётных
func filterEven(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        num := data.(int)
        if num%2 == 0 {
            out <- num
        }
    }
}

// Возведение в квадрат
func square(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        num := data.(int)
        out <- num * num
    }
}

// Вывод
func printResult(in, out chan interface{}) {
    for data := range in {
        fmt.Println(data.(int))
    }
}

pipeline := pipelin.WorkPipeline{
    numberGen,
    filterEven,
    square,
    printResult,
}

pipeline.Run()
// Вывод: 0, 4, 16, 36, 64, 100, 144, 196, 256, 324
```

## Пример: агрегация

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/pipelin"
)

// Генерация чисел
func numberGen(in, out chan interface{}) {
    defer close(out)
    for i := 1; i <= 100; i++ {
        out <- i
    }
}

// Суммирование
func sum(in, out chan interface{}) {
    defer close(out)
    total := 0
    for data := range in {
        num := data.(int)
        total += num
    }
    out <- total
}

// Результат
func printResult(in, out chan interface{}) {
    for data := range in {
        fmt.Printf("Сумма: %d\n", data.(int))
    }
}

pipeline := pipelin.WorkPipeline{
    numberGen,
    sum,
    printResult,
}

pipeline.Run()
// Вывод: Сумма: 5050
```

## Пример: сложная обработка

```go
import (
    "fmt"
    "strings"
    "github.com/nobuenhombre/suikat/pkg/pipelin"
)

type User struct {
    Name string
    Age  int
}

type FilteredUser struct {
    Name string
    Age  int
}

// Генерация пользователей
func userGen(in, out chan interface{}) {
    defer close(out)
    users := []User{
        {Name: "Alice", Age: 25},
        {Name: "Bob", Age: 17},
        {Name: "Charlie", Age: 30},
        {Name: "Dave", Age: 15},
        {Name: "Eve", Age: 22},
    }
    for _, u := range users {
        out <- u
    }
}

// Фильтрация по возрасту (>= 18)
func filterAdults(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        u := data.(User)
        if u.Age >= 18 {
            out <- FilteredUser{Name: u.Name, Age: u.Age}
        }
    }
}

// Форматирование
func formatResult(in, out chan interface{}) {
    defer close(out)
    for data := range in {
        u := data.(FilteredUser)
        out <- fmt.Sprintf("%s (%d лет)", u.Name, u.Age)
    }
}

// Вывод
func printResult(in, out chan interface{}) {
    for data := range in {
        fmt.Println(data.(string))
    }
}

pipeline := pipelin.WorkPipeline{
    userGen,
    filterAdults,
    formatResult,
    printResult,
}

pipeline.Run()
// Вывод:
// Alice (25 лет)
// Charlie (30 лет)
// Eve (22 лет)
```

## Примечания

- Каждый шаг получает канал с буфером размера 1
- Функция `Run()` блокируется до завершения всех шагов
- Каждая функция должна закрывать свой выходной канал (`close(out)`) перед завершением
- Порядок шагов важен — данные текут слева направо
