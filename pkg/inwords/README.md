# inwords

Пакет `inwords` конвертирует числовые значения (денежные суммы) в текстовое представление на русском языке с правильным склонением.

## Обзор

Основная функция `Format` принимает число с плавающей запятой (например, сумму в рублях) и возвращает строку с написанным словами числом и склонёнными единицами валюты.

Поддерживаются суммы до триллионов включительно.

## Format

Конвертирует числовую сумму в текстовое представление.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/inwords"
)

// Базовое использование
text, err := inwords.Format(25230.33)
fmt.Println(text)
// Вывод: двадцать пять тысяч двести тридцать рублей 33 копейки

// Простая сумма
text, err = inwords.Format(230.45)
fmt.Println(text)
// Вывод: двести тридцать рублей 45 копеек

// Ноль
text, err = inwords.Format(0.0)
fmt.Println(text)
// Вывод: ноль рублей 00 копеек

// С тысячами
text, err = inwords.Format(3333.01)
fmt.Println(text)
// Вывод: три тысячи тридцать три рубля 01 копейка
```

### Склонение единиц валюты

Функция автоматически склоняет названия валют в зависимости от числа:

```go
import "github.com/nobuenhombre/suikat/pkg/inwords"

// «рубль» (оканчивается на 1, кроме 11)
inwords.Format(1.00)   // один рубль 00 копеек
inwords.Format(21.00)  // двадцать один рубль 00 копеек
inwords.Format(101.00) // сто один рубль 00 копеек

// «рубля» (оканчивается на 2-4, кроме 12-14)
inwords.Format(2.00)   // два рубля 00 копеек
inwords.Format(22.00)  // двадцать два рубля 00 копеек
inwords.Format(102.00) // сто два рубля 00 копеек

// «рублей» (остальные случаи)
inwords.Format(5.00)   // пять рублей 00 копеек
inwords.Format(11.00)  // одиннадцать рублей 00 копеек
inwords.Format(25.00)  // двадцать пять рублей 00 копеек
inwords.Format(111.00) // сто одиннадцать рублей 00 копеек
```

### Копейки

Аналогичное склонение применяется к копейкам:

```go
import "github.com/nobuenhombre/suikat/pkg/inwords"

inwords.Format(1.01)   // один рубль 01 копейка
inwords.Format(1.02)   // один рубль 02 копейки
inwords.Format(1.05)   // один рубль 05 копеек
inwords.Format(1.11)   // один рубль 11 копеек
```

### Большие числа

```go
import "github.com/nobuenhombre/suikat/pkg/inwords"

// Миллионы
inwords.Format(1500000.50)
// пять миллионов полсотни тысяч рублей 50 копеек

// Миллиарды
inwords.Format(1000000000.00)
// один миллиард рублей 00 копеек

// Триллионы
inwords.Format(1000000000000.00)
// один триллион рублей 00 копеек
```

### Обработка ошибок

В нормальных случаях функция не возвращает ошибок. Ошибка может возникнуть только при некорректном внутренном состоянии.

```go
import "github.com/nobuenhombre/suikat/pkg/inwords"

text, err := inwords.Format(123.45)
if err != nil {
    // обработка ошибки
    panic(err)
}
fmt.Println(text)
```

## Константы

```go
import "github.com/nobuenhombre/suikat/pkg/inwords"

// Пол для чисел (по умолчанию используется GenderMale)
_ = inwords.GenderMale   // "male"
_ = inwords.GenderFemale // "female"

// Длина класса цифр
_ = inwords.ClassLength // 3

// Максимальное количество классов (поддержка до триллионов)
_ = inwords.MaxNumberClasses // 5
```
