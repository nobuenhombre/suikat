# replacer

Пакет `replacer` предоставляет механизм замены строк по набору правил с поддержкой регулярных выражений и точного совпадения.

## Обзор

Пакет позволяет определить набор правил замены и применить их к одному или множеству строк. Каждое правило может быть либо точным совпадением (`string`), либо регулярным выражением (`regexp`).

## ReplaceRule

Правило замены:

```go
type ReplaceRule struct {
    SourceType  string // "string" или "regexp"
    Source      string // Искомый текст или паттерн
    Replacement string // Строка замены
}

func (rr *ReplaceRule) Compile() // Компиляция регулярного выражения
```

## Типы источников

```go
import "github.com/nobuenhombre/suikat/pkg/replacer"

_ = replacer.SourceTypeRegexp   // "regexp"
_ = replacer.SourceTypeString   // "string"
```

## Data и ReplaceData

Структуры для пакетной обработки:

```go
type Data []ReplaceData
type ReplaceData struct {
    ID     int64
    Before string
    After  string
}
```

## Rules

Срез правил:

```go
type Rules []ReplaceRule
```

## ApplyRule

Применяет одно правило к строке:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

// Точное совпадение
rule := replacer.ReplaceRule{
    SourceType:  replacer.SourceTypeString,
    Source:      "%%name%%",
    Replacement: "Mr.",
}
rule.Compile()

result, err := replacer.ApplyRule("Hello %%name%% World", rule)
if err != nil {
    panic(err)
}
fmt.Println(result) // Hello Mr. World

// Регулярное выражение
regexpRule := replacer.ReplaceRule{
    SourceType:  replacer.SourceTypeRegexp,
    Source:      "data-uri\\(.*?\\)",
    Replacement: "uri",
}
regexpRule.Compile()

result, err = replacer.ApplyRule("Call data-uri(/img/welcome.jpg) This Site!", regexpRule)
if err != nil {
    panic(err)
}
fmt.Println(result) // Call uri This Site!
```

## ApplyRules

Применяет набор правил к одной строке последовательно:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

rules := &replacer.Rules{
    {
        SourceType:  replacer.SourceTypeString,
        Source:      "%%name%%",
        Replacement: "Mr.",
    },
    {
        SourceType:  replacer.SourceTypeRegexp,
        Source:      "data-uri\\(.*?\\)",
        Replacement: "uri",
    },
}

// Компиляция всех правил
for i := range *rules {
    (*rules)[i].Compile()
}

before := "Hello %%name%% World, data-uri(/img/welcome.jpg)!"
after, err := replacer.ApplyRules(before, rules)
if err != nil {
    panic(err)
}
fmt.Println(after) // Hello Mr. World, uri!
```

## ReplaceAll

Применяет правила ко всем элементам Data:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

data := &replacer.Data{
    {ID: 1, Before: "hello %%name%% world"},
    {ID: 2, Before: "Call data-uri(/img/welcome.jpg) This Site!"},
    {ID: 3, Before: "Board hmm Band %%name%%"},
}

rules := &replacer.Rules{
    {
        SourceType:  replacer.SourceTypeString,
        Source:      "%%name%%",
        Replacement: "Mr.",
    },
    {
        SourceType:  replacer.SourceTypeRegexp,
        Source:      "data-uri\\(.*?\\)",
        Replacement: "uri",
    },
    {
        SourceType:  replacer.SourceTypeRegexp,
        Source:      "([A-Z])\\w+",
        Replacement: "BZZ",
    },
}

// Компиляция правил
for i := range *rules {
    (*rules)[i].Compile()
}

outData, err := replacer.ReplaceAll(data, rules)
if err != nil {
    panic(err)
}

for _, item := range *outData {
    fmt.Printf("ID %d: %s -> %s\n", item.ID, item.Before, item.After)
}
// Вывод:
// ID 1: hello %%name%% world -> hello Mr. world
// ID 2: Call data-uri(/img/welcome.jpg) This Site! -> Call uri BZZ!
// ID 3: Board hmm Band %%name%% -> BZZ hmm BZZ Mr.
```

## Пример: обработка HTML

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

htmlData := &replacer.Data{
    {ID: 1, Before: "<img src='data-uri(/img/photo.jpg)' alt='Photo'>"},
    {ID: 2, Before: "<a href='data-uri(/css/style.css)'>Link</a>"},
}

rules := &replacer.Rules{
    {
        SourceType:  replacer.SourceTypeRegexp,
        Source:      "data-uri\\(.*?\\)",
        Replacement: "[ASSET]",
    },
}

for i := range *rules {
    (*rules)[i].Compile()
}

outData, err := replacer.ReplaceAll(htmlData, rules)
if err != nil {
    panic(err)
}

for _, item := range *outData {
    fmt.Println(item.After)
}
// Вывод:
// <img src='[ASSET]' alt='Photo'>
// <a href='[ASSET]'>Link</a>
```

## Пример: обработка текста

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

text := "Контакт: contact@example.com, телефон: +7-999-123-4567"

rules := &replacer.Rules{
    {
        SourceType:  replacer.SourceTypeRegexp,
        Source:      "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}",
        Replacement: "[EMAIL]",
    },
    {
        SourceType:  replacer.SourceTypeRegexp,
        Source:      "\\+?[0-9]{1,4}[-.][0-9]{3,4}[-.][0-9]{3,4}[-.][0-9]{3,4}",
        Replacement: "[PHONE]",
    },
}

for i := range *rules {
    (*rules)[i].Compile()
}

after, err := replacer.ApplyRules(text, rules)
if err != nil {
    panic(err)
}
fmt.Println(after) // Контакт: [EMAIL], телефон: [PHONE]
```

## Пример: последовательная обработка

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

// Шаг 1: замена плейхолдеров
text := "Hello %%name%%, welcome to %%place%%!"
rules1 := &replacer.Rules{
    {SourceType: replacer.SourceTypeString, Source: "%%name%%", Replacement: "Alice"},
    {SourceType: replacer.SourceTypeString, Source: "%%place%%", Replacement: "Wonderland"},
}

for i := range *rules1 {
    (*rules1)[i].Compile()
}

text, err := replacer.ApplyRules(text, rules1)
if err != nil {
    panic(err)
}
fmt.Println(text) // Hello Alice, welcome to Wonderland!

// Шаг 2: форматирование заголовка
rules2 := &replacer.Rules{
    {SourceType: replacer.SourceTypeString, Source: "w", Replacement: "W"},
}

for i := range *rules2 {
    (*rules2)[i].Compile()
}

text, err = replacer.ApplyRules(text, rules2)
if err != nil {
    panic(err)
}
fmt.Println(text) // Hello Alice, welcome to Wonderland!
```

## Ошибки

При использовании некомпилированного regexp-правила возвращается ошибка:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/ge"
    "github.com/nobuenhombre/suikat/pkg/replacer"
)

rule := replacer.ReplaceRule{
    SourceType:  replacer.SourceTypeRegexp,
    Source:      "data-uri\\(.*?\\)",
    Replacement: "uri",
}
// Не вызвали Compile()!

_, err := replacer.ApplyRule("text", rule)
if err != nil {
    if _, ok := err.(*ge.RegExpIsNotCompiledError); ok {
        fmt.Println("Регулярное выражение не скомпилировано")
    }
}
```
