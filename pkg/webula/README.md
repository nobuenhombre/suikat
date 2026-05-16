# webula

Пакет `webula` предоставляет утилиты для работы с текстом: нормализация, очистка HTML, транслитерация, подсчёт символов и другие.

## Обзор

Пакет содержит функции для:
- Подсчёта символов в строке (юникод)
- Очистки от HTML-тегов
- Триминга среза строк
- Разбиения текста на слова
- Нормализации текста
- Определения наличия HTML
- Удаления дубликатов строк
- Нормализации URL-имён
- Нормализации алфавитных указателей
- Транслитерации русского текста

## StrLen

Подсчитывает количество символов (runes) в строке:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

fmt.Println(webula.StrLen("Hello"))    // 5
fmt.Println(webula.StrLen("Привет"))   // 6
fmt.Println(webula.StrLen("Hello Мир")) // 9
```

## StripHTML

Очищает строку от HTML-тегов:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

html := "<h1>Заголовок</h1><p>Параграф <i>курсив</i> текст</p>"
clean := webula.StripHTML(html)
fmt.Println(clean) // ЗаголовокПараграф курсив текст
```

## Trim

Тримит элементы среза строк и удаляет пустые:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

words := []string{"  Арбуз", "Дыня   ", "  Персик  "}
result := webula.Trim(words, []string{" "})
fmt.Println(result) // [Арбуз Дыня Персик]

// С несколькими символами для триминга
words = []string{"  _Арбуз=", "==Дыня__", "-Персик-"}
result = webula.Trim(words, []string{" ", "_", "=", "-"})
fmt.Println(result) // [Арбуз Дыня Персик]
```

## Words

Разбивает текст на слова, игнорируя пробелы, табы, переносы строк:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

text := "  Арбуз Дыня  \t Персик\nЯблоко  Вишня "
words := webula.Words(text)
fmt.Println(words) // [Арбуз Дыня Персик Яблоко Вишня]
```

## NormalizeText

Нормализует текст — устраняет множественные повторяющиеся символы (пробелы и др.):

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

text := "  Арбуз   Дыня  \t  Персик\nЯблоко  Вишня "

// Склейка пробелом
result := webula.NormalizeText(text, " ")
fmt.Println(result) // Арбуз Дыня Персик Яблоко Вишня

// Склейка запятой
result = webula.NormalizeText(text, ", ")
fmt.Println(result) // Арбуз, Дыня, Персик, Яблоко, Вишня
```

## IsHTML

Проверяет, содержит ли строка HTML-теги:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

fmt.Println(webula.IsHTML("<h1>Привет</h1>")) // true
fmt.Println(webula.IsHTML("Привет"))           // false
fmt.Println(webula.IsHTML("Привет &nbsp; Мир")) // true (&nbsp; — HTML-сущность)
```

## RemoveDuplicatesString

Удаляет дубликаты из среза строк, сохраняя порядок:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

words := []string{"Арбуз", "Дыня", "Фейхоа", "Арбуз", "Персик", "Арбуз", "Фейхоа"}
unique := webula.RemoveDuplicatesString(words)
fmt.Println(unique) // [Арбуз Дыня Фейхоа Персик]
```

## NormalizeNameURL

Нормализует имя для использования в URL:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

name := "Расцветали яблони и груши, шли туманы."
urlName := webula.NormalizeNameURL(name)
fmt.Println(urlName) // расцветали_яблони_и_груши_шли_туманы

// С кириллицей
name = "Привет Мир!"
urlName = webula.NormalizeNameURL(name)
fmt.Println(urlName) // привет_мир

// Со специальными символами
name = "Цена: 100*200=20000₽"
urlName = webula.NormalizeNameURL(name)
fmt.Println(urlName) // цена_100200_20000_₽
```

## NormalizeAlphabet

Нормализует фразы алфавитного указателя (разделённые запятыми):

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

text := "Расцветали яблони и груши, шли туманы над    рекой,   Катюша   "
result := webula.NormalizeAlphabet(text)
fmt.Println(result) // Расцветали яблони и груши, шли туманы над рекой, Катюша
```

## TranslitRusLat

Транслитерирует русский текст в латиницу:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

text := "Привет Мир!"
result := webula.TranslitRusLat(text)
fmt.Println(result) // Privet Mir!

text = "Ёжик в тумане"
result = webula.TranslitRusLat(text)
fmt.Println(result) // Yozhik v tumanе
```

### Маппинг транслитерации

| Кириллица | Латиница | Кириллица | Латиница |
|-----------|----------|-----------|----------|
| А | A | а | a |
| Б | B | б | b |
| В | V | в | v |
| Г | G | г | g |
| Д | D | д | d |
| Е | E | е | e |
| Ё | YO | ё | yo |
| Ж | ZH | ж | zh |
| З | Z | з | z |
| И | I | и | i |
| Й | Y | й | y |
| К | K | к | k |
| Л | L | л | l |
| М | M | м | m |
| Н | N | н | n |
| О | O | о | o |
| П | P | п | p |
| Р | R | р | r |
| С | S | с | s |
| Т | T | т | t |
| У | U | у | u |
| Ф | F | ф | f |
| Х | H | х | h |
| Ц | C | ц | c |
| Ч | CH | ч | ch |
| Ш | SH | ш | sh |
| Щ | SCH | щ | sch |
| Ъ | (пусто) | ъ | (пусто) |
| Ы | Y | ы | y |
| Ь | (пусто) | ь | (пусто) |
| Э | E | э | e |
| Ю | YU | ю | yu |
| Я | YA | я | ya |

## Константы

Пакет содержит константы для часто используемых символов:

```go
import "github.com/nobuenhombre/suikat/pkg/webula"

// Строковые константы
_ = webula.EmptyString       // ""
_ = webula.Space             // " "
_ = webula.Underline         // "_"
_ = webula.Dot               // "."
_ = webula.NewLine           // "\n"
_ = webula.CarriageReturn    // "\r"
_ = webula.Tab               // "\t"
_ = webula.Comma             // ","
_ = webula.Colon             // ":"
_ = webula.Semicolon         // ";"
_ = webula.Gradus            // "°"
_ = webula.SingleQuote       // "'"
_ = webula.DoubleQuote       // "\""
_ = webula.QuoteLeft         // "«"
_ = webula.QuoteRight        // "»"
_ = webula.Mult              // "*"
_ = webula.Div               // "/"
_ = webula.Plus              // "+"
_ = webula.Minus             // "-"
_ = webula.Equal             // "="
_ = webula.Percent           // "%"
_ = webula.Number            // "№"
_ = webula.Exclamation       // "!"
_ = webula.RoundBracketLeft  // "("
_ = webula.RoundBracketRight // ")"

// HTML-сущности
_ = webula.HTMLSpaceInUtf8   // "\xc2\xa0" (неразрывный пробел в UTF-8)
_ = webula.HTMLSpace         // "&nbsp;"
_ = webula.HTMLMDash         // "&mdash;"
```

## Пример: полный пайплайн обработки текста

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/webula"
)

func processText(rawText string) {
    // 1. Очистка от HTML
    cleanText := webula.StripHTML(rawText)

    // 2. Нормализация
    normalized := webula.NormalizeText(cleanText, " ")

    // 3. Разбиение на слова
    words := webula.Words(normalized)

    // 4. Удаление дубликатов
    uniqueWords := webula.RemoveDuplicatesString(words)

    // 5. Транслитерация
    translit := webula.TranslitRusLat(strings.Join(uniqueWords, " "))

    // 6. Нормализация для URL
    urlName := webula.NormalizeNameURL(translit)

    fmt.Printf("Исходный: %s\n", rawText)
    fmt.Printf("Очищенный: %s\n", cleanText)
    fmt.Printf("Нормализованный: %s\n", normalized)
    fmt.Printf("Слова: %v\n", uniqueWords)
    fmt.Printf("URL: %s\n", urlName)
}

processText("<h1>Привет Мир!</h1><p>Это <b>тест</b> текста.</p>")
```
