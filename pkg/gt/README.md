# gt

Пакет `gt` предоставляет удобную обёртку над `text/template` и `html/template` для загрузки шаблонов из директорий с поддержкой рекурсивного обхода подкаталогов.

## Основные возможности

- Загрузка шаблонов из одной или нескольких директорий
- Рекурсивный парсинг подкаталогов
- Поддержка пользовательских функций через `template.FuncMap`
- Два типа: одиночный путь (`Path`) и коллекция путей (`Paths`)

## Типы шаблонов

| Тип       | Маска файлов     | Назначение            |
|-----------|------------------|-----------------------|
| `HTMLPath`   | `*.gohtml`       | HTML-шаблоны          |
| `TextPath`   | `*.go.tpl`       | Текстовые шаблоны     |
| `HTMLPaths`  | `*.gohtml`       | Несколько HTML-путьев |
| `TextPaths`  | `*.go.tpl`       | Несколько текстовых   |

## Структура директорий

```
templates/
├── base.gohtml
├── partials/
│   ├── header.gohtml
│   └── footer.gohtml
```

## HTML-шаблоны (одиночный путь)

```go
import "github.com/nobuenhombre/suikat/pkg/gt"

tpl := gt.HTMLPath("templates")

// С пользовательскими функциями
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
}
result, err := tpl.HTML("base", map[string]string{"Title": "Hello"}, funcMap)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)

// Без функций
result, err = tpl.HTML("base", data)
```

## Текстовые шаблоны (одиночный путь)

```go
import "github.com/nobuenhombre/suikat/pkg/gt"

tpl := gt.TextPath("templates")

funcMap := template.FuncMap{
    "upper": strings.ToUpper,
}
result, err := tpl.Text("email", map[string]string{"Name": "Alice"}, funcMap)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

## Несколько путей (HTML)

```go
import "github.com/nobuenhombre/suikat/pkg/gt"

paths := gt.NewHTMLPaths()
paths.AddPath(gt.Path("templates"))
paths.AddPath(gt.Path("email-templates"))

result, err := paths.HTML("base", data, funcMap)
if err != nil {
    log.Fatal(err)
}
```

## Несколько путей (текст)

```go
import "github.com/nobuenhombre/suikat/pkg/gt"

paths := gt.NewTextPaths()
paths.AddPath(gt.Path("templates"))
paths.AddPath(gt.Path("config-templates"))

result, err := paths.Text("config", data, funcMap)
if err != nil {
    log.Fatal(err)
}
```

## Подкаталоги

Метод `GetSubDirectories` получает все подкаталоги рекурсивно:

```go
import "github.com/nobuenhombre/suikat/pkg/gt"

dirs, err := gt.Path("templates").GetSubDirectories()
// dirs: ["templates/partials", "templates/email", ...]
```

## Проверка ошибок

```go
import (
    "github.com/nobuenhombre/suikat/pkg/gt"
    "errors"
)

// Проверка на отсутствие путей
if errors.Is(err, gt.NoPathsDefinedError) {
    log.Println("Нет определённых путей для шаблонов")
}
```
