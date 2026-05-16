# chacha

Пакет `chacha` предоставляет примитивы для создания сложных деревьев проверок (validation trees), где каждая последующая проверка зависит от результата предыдущей.

## Зависимости

- `github.com/fatih/color` — цветной вывод ошибок

## Основные типы

### `Validator`

Тип функции: `func(params ...interface{}) (interface{}, error)`. Возвращает данные и ошибку. Если ошибка `nil` — проверка прошла успешно.

### `DontCheckChildrens`

Специальный маркер: если `Data` имеет тип `*DontCheckChildrens`, то дочерние ветки не проверяются.

### `Tree`

Структура узла дерева проверок:
- `Validator Validator` — функция-валидатор.
- `Data interface{}` — результат выполнения валидатора.
- `Result error` — ошибка валидации.
- `Valid bool` — результат валидации (true = успешно).
- `Childrens []Tree` — дочерние узлы.

**Методы:**
- `CheckChildrens() bool` — нужно ли проверять дочерние ветки.
- `Validate(params ...interface{})` — рекурсивно обходит дерево, вызывает валидаторы, обновляет `Valid`.
- `ShowErrors()` — выводит ошибки в stdout красным цветом.

## Примеры

```go
package main

import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/chacha"
)

func main() {
    // --- Пример 1: Простая цепочка проверок ---

    // Валидатор 1: проверяет, что входные данные — строка
    validateString := func(params ...interface{}) (interface{}, error) {
        input := params[0].(string)
        if len(input) == 0 {
            return nil, fmt.Errorf("input is empty")
        }
        return input, nil
    }

    // Валидатор 2: проверяет длину строки
    validateLength := func(params ...interface{}) (interface{}, error) {
        str := params[0].(string) // результат предыдущей проверки
        if len(str) < 3 {
            return nil, fmt.Errorf("string too short: %d chars", len(str))
        }
        return str, nil
    }

    // Валидатор 3: проверяет, что строка содержит цифру
    validateHasDigit := func(params ...interface{}) (interface{}, error) {
        str := params[0].(string)
        hasDigit := false
        for _, c := range str {
            if c >= '0' && c <= '9' {
                hasDigit = true
                break
            }
        }
        if !hasDigit {
            return nil, fmt.Errorf("string must contain a digit")
        }
        return str, nil
    }

    // Строим дерево
    tree := chacha.Tree{
        Validator: validateString,
        Childrens: []chacha.Tree{
            {
                Validator: validateLength,
                Childrens: []chacha.Tree{
                    {Validator: validateHasDigit},
                },
            },
        },
    }

    // Успешная валидация
    tree.Validate("hello123")
    fmt.Printf("Valid: %v\n", tree.Valid) // Valid: true

    // Ошибка на 1-м уровне
    tree2 := chacha.Tree{
        Validator: validateString,
    }
    tree2.Validate("")
    fmt.Printf("Valid: %v\n", tree2.Valid) // Valid: false
    tree2.ShowErrors()
    //  ♥ input is empty

    // Ошибка на 2-м уровне
    tree3 := chacha.Tree{
        Validator: validateString,
        Childrens: []chacha.Tree{
            {Validator: validateLength},
        },
    }
    tree3.Validate("ab")
    fmt.Printf("Valid: %v\n", tree3.Valid) // Valid: false
    tree3.ShowErrors()
    //  ♥ string too short: 2 chars

    // Ошибка на 3-м уровне
    tree4 := chacha.Tree{
        Validator: validateString,
        Childrens: []chacha.Tree{
            {
                Validator: validateLength,
                Childrens: []chacha.Tree{
                    {Validator: validateHasDigit},
                },
            },
        },
    }
    tree4.Validate("abc")
    fmt.Printf("Valid: %v\n", tree4.Valid) // Valid: false
    tree4.ShowErrors()
    //  ♥ string must contain a digit

    // --- Пример 2: Остановка проверки дочерних веток ---

    stopValidation := func(params ...interface{}) (interface{}, error) {
        return &chacha.DontCheckChildrens{}, nil
    }

    childValidator := func(params ...interface{}) (interface{}, error) {
        return nil, fmt.Errorf("this should not be called")
    }

    tree5 := chacha.Tree{
        Validator: stopValidation,
        Childrens: []chacha.Tree{
            {Validator: childValidator},
        },
    }
    tree5.Validate(nil)
    fmt.Printf("Valid: %v\n", tree5.Valid) // Valid: true
    // ShowErrors ничего не выведет — дочерние ветки не проверялись
}
```
