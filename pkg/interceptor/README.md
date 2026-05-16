# interceptor

Пакет `interceptor` предоставляет HTTP-маршрутизатор с поддержкой регулярных выражений в URI, утилиты для формирования HTTP-ответов и работы с загруженными файлами.

## Обзор

Пакет включает:
- **HTTPRouter** — простой маршрутизатор с поддержкой методов HTTP и паттернов URI
- **HTTPAnswer** — удобный способ формирования HTTP-ответов с поддержкой gzip, кэширования, ETag и CORS
- **HTTPRegexp** — сопоставление URI-частей с регулярными выражениями
- **HTTPRoute** — описание маршрута (метод + URI + обработчик)
- **Утилиты** — работа с загруженными файлами

## HTTPRouter

Маршрутизатор с поддержкой регулярных выражений в путях URI.

### Инициализация и регистрация маршрутов

```go
import (
    "net/http"
    "github.com/nobuenhombre/suikat/pkg/interceptor"
)

router := &interceptor.HTTPRouter{}
router.Init()

// Регистрация маршрута с точным совпадением
router.HandleFunc("GET", "/api/status", func(w http.ResponseWriter, r *http.Request) error {
    answer := &interceptor.HTTPAnswer{
        ResponseCode: http.StatusOK,
        Content:      map[string]string{"status": "ok"},
    }
    return answer.Send(w, r)
})

// Регистрация маршрута с регулярным выражением
router.HTTPRegexp.Add(interceptor.Word, "([\\w]+)")
router.HandleFunc("GET", "/api/users/:word", func(w http.ResponseWriter, r *http.Request) error {
    return answer.Send(w, r)
})

http.ListenAndServe(":8080", router)
```

### Предопределённые паттерны

```go
import "github.com/nobuenhombre/suikat/pkg/interceptor"

// Доступные предопределённые паттерны:
// interceptor.Word   — ([\w]+)  : слово (буквы, цифры, подчёркивание)
// interceptor.Number — ([\d]+)  : число (цифры)
// interceptor.Any    — .*       : любой текст

router.HTTPRegexp.AddPredefined([]string{interceptor.Word, interceptor.Number})

// Теперь можно использовать :word и :number в URI
router.HandleFunc("GET", "/api/items/:number", handler)
```

### Настройка ответа 404

```go
router.DefaultNotFoundAnswer = interceptor.HTTPAnswer{
    ResponseCode: http.StatusNotFound,
    Content:      "Not Found",
}
```

### Логирование ошибок

```go
import "go.uber.org/zap"

logger, _ := zap.NewProduction()
router.Logger = logger
```

## HTTPAnswer

Универсальный ответ HTTP с поддержкой gzip, кэширования, ETag и CORS.

### Базовое использование

```go
import (
    "net/http"
    "github.com/nobuenhombre/suikat/pkg/interceptor"
)

// Ответ JSON
answer := &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      map[string]interface{}{"name": "John", "age": 30},
}
answer.Send(w, r)
// Content-Type: application/json

// Ответ строкой
answer = &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      "Hello, World!",
}
answer.Send(w, r)
// Content-Type: text/html

// Пустой ответ
answer = &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      nil,
}
answer.Send(w, r)
// Content-Type: text/plain
```

### Ответ файлом

```go
answer := &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content: interceptor.FileData{
        Name:     "report.pdf",
        Size:     102400,
        Data:     fileBytes,
        Download: true, // true = заголовок Content-Disposition (скачивание)
    },
    ContentType: "application/pdf",
}
answer.Send(w, r)
```

### Сжатие gzip

```go
answer := &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      largeData,
    GZipped:      true,
    GZipLevel:    gzip.BestCompression,
}
answer.Send(w, r)
// Заголовок: Content-Encoding: gzip
```

### Кэширование в браузере

```go
// Включить кэширование
answer := &interceptor.HTTPAnswer{
    ResponseCode:         http.StatusOK,
    Content:              data,
    BrowserCached:        true,
    BrowserCacheLifeTime: interceptor.BrowserCacheLifeTimeWeek, // 604800 секунд (1 неделя)
}
answer.Send(w, r)

// Отключить кэширование
answer.BrowserCached = false
answer.Send(w, r)
```

### ETag

```go
answer := &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      data,
    ETagUsed:     true,
    ETag:         "abc123",
}
answer.Send(w, r)

// Клиент отправляет: If-None-Match: abc123
// Сервер возвращает: 304 Not Modified
```

### CORS-заголовки

```go
answer := &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      data,
    HTTPCORSHeaders: interceptor.HTTPCORSHeaders{
        AccessControlAllowOrigin:  "*",
        AccessControlAllowHeaders: "Origin, Content-Type, Authorization",
        AccessControlAllowMethods: "GET, POST, PUT, DELETE",
    },
}
answer.Send(w, r)
```

### Кастомный Content-Type

```go
answer := &interceptor.HTTPAnswer{
    ResponseCode: http.StatusOK,
    Content:      data,
    ContentType:  "application/xml",
    Encoding:     interceptor.EncodingCharsetUTF8,
}
answer.Send(w, r)
// Content-Type: application/xml; charset=utf-8
```

## HTTPRegexp

Сопоставление частей URI с регулярными выражениями.

### Предопределённые паттерны

```go
import "github.com/nobuenhombre/suikat/pkg/interceptor"

predefined := interceptor.Predefined()
// predefined["word"]   = "([\\w]+)"
// predefined["number"]  = "([\\d]+)"
// predefined["any"]     = ".*"
```

### Произвольные паттерны

```go
reg := &interceptor.HTTPRegexp{}
reg.Init()
reg.Add("email", "([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,})")
reg.Add("uuid", "([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})")

// Проверка совпадения
routeParts := []string{"/api", "/email", "/uuid"}
requestParts := []string{"/api", "user@example.com", "550e8400-e29b-41d4-a716-446655440000"}

matched := reg.MatchURIParts(routeParts, requestParts)
fmt.Println(matched) // true
```

## HTTPRoute

Описание маршрута и проверка совпадений.

```go
import (
    "net/http"
    "github.com/nobuenhombre/suikat/pkg/interceptor"
)

route := &interceptor.HTTPRoute{
    Method: "POST",
    URI:    "/api/users",
    F: func(w http.ResponseWriter, r *http.Request) error {
        // обработчик
        return nil
    },
}

// Проверка метода
fmt.Println(route.MatchMethod(&http.Request{Method: "POST"})) // true
fmt.Println(route.MatchMethod(&http.Request{Method: "GET"})) // false
```

## Утилиты

### Загрузка файлов

```go
import (
    "net/http"
    "github.com/nobuenhombre/suikat/pkg/interceptor"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) error {
    file, err := interceptor.GetFile("upload", r)
    if err != nil {
        return err
    }

    fmt.Printf("Имя файла: %s\n", file.Name)
    fmt.Printf("Размер: %d байт\n", file.Size)
    fmt.Printf("MIME: %s\n", file.Mime)
    fmt.Printf("Данные: %d байт\n", len(file.Data))

    return nil
}
```

### Константы

```go
import "github.com/nobuenhombre/suikat/pkg/interceptor"

// Время жизни кэша в браузере (1 неделя)
_ = interceptor.BrowserCacheLifeTimeWeek // 604800

// Кодировка
_ = interceptor.EncodingCharsetUTF8 // "utf-8"

// Максимальный размер формы в памяти (10 МБ)
_ = interceptor.MaxFormSizeMemory // 10 << 20
```
