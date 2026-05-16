# less-rest-client

Пакет `less-rest-client` предоставляет HTTP-клиент с поддержкой различных форматов запросов и типов аутентификации.

## Обзор

Клиент `Client` поддерживает:
- Форматы тела запроса: JSON, form-urlencoded, multipart/form-data
- Типы аутентификации: None, Basic (login/password), Token, Bearer Token, Cookie + CSRF, EDS
- HTTP-методы: GET, POST, PUT, PATCH, DELETE
- Проверку ожидаемого HTTP-статуса ответа
- Отладочную информацию о последнем запросе

## Создание клиента

```go
import (
    "github.com/nobuenhombre/suikat/pkg/less-rest-client"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

client := &lessrestclient.Client{
    URL: "https://api.example.com",
}
client.Init()

// Или через конструктор
var lrc lessrestclient.LRC = lessrestclient.New(&lessrestclient.Client{
    URL: "https://api.example.com",
})
```

## Аутентификация

### Без аутентификации

```go
client := &lessrestclient.Client{
    URL:      "https://api.example.com",
    AuthType: lessrestclient.AuthTypeNone,
}
```

### Basic Auth (логин + пароль)

```go
client := &lessrestclient.Client{
    URL:      "https://api.example.com",
    AuthType: lessrestclient.AuthTypeLoginPass,
    Username: "user",
    Password: "pass",
}
```

### Token

```go
client := &lessrestclient.Client{
    URL:      "https://api.example.com",
    AuthType: lessrestclient.AuthTypeToken,
    Token:    "my-token-value",
}
```

### Bearer Token

```go
client := &lessrestclient.Client{
    URL:         "https://api.example.com",
    AuthType:    lessrestclient.AuthTypeBearerToken,
    BearerToken: "eyJhbGciOiJIUzI1NiIs...",
}
```

### Cookie + CSRF Token

```go
client := &lessrestclient.Client{
    URL:               "https://api.example.com",
    AuthType:          lessrestclient.AuthTypeCookieWithCSRFToken,
    Cookie:            "sessionid=abc123",
    XCSRFToken:        "csrf-token-value",
}
```

### EDS (цифровая подпись)

```go
client := &lessrestclient.Client{
    URL:         "https://api.example.com",
    AuthType:    lessrestclient.AuthTypeEDS,
    SignKeyID:   "key-id-123",
    SignBody:    "signed-body-data",
}
```

## Форматы запросов

### JSON

```go
client := &lessrestclient.Client{
    URL:         "https://api.example.com",
    ContentType: mimes.JSON,
}

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var result User
statusCode, respBody, err := client.POST("/users", User{Name: "Alice", Age: 30}, &result, 200)
if err != nil {
    panic(err)
}
fmt.Printf("Status: %d\n", statusCode)
fmt.Printf("Response: %s\n", string(respBody))
```

### Form URL-encoded

```go
import (
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/less-rest-client"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

client := &lessrestclient.Client{
    URL:         "https://api.example.com",
    ContentType: mimes.FormUrlencoded,
}

formData := url.Values{
    "username": {"admin"},
    "password": {"secret"},
}

var result map[string]interface{}
statusCode, respBody, err := client.POST("/login", formData, &result, 200)
```

### Multipart Form Data (с файлами)

```go
import (
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/less-rest-client"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

client := &lessrestclient.Client{
    URL:         "https://api.example.com",
    ContentType: mimes.FormMultipartData,
}

formData := url.Values{
    "name":  {"report"},
    "@file": {"/path/to/file.pdf"}, // префикс @ = файл
}

var result map[string]interface{}
statusCode, respBody, err := client.POST("/upload", formData, &result, 200)
```

## HTTP-методы

### GET

```go
var result map[string]interface{}
statusCode, respBody, err := client.GET("/api/users/1", &result, 200)
```

### POST

```go
type RequestBody struct {
    Title string `json:"title"`
    Body  string `json:"body"`
}

var result map[string]interface{}
statusCode, respBody, err := client.POST("/api/posts", RequestBody{Title: "Hello", Body: "World"}, &result, 201)
```

### PUT

```go
var result map[string]interface{}
statusCode, respBody, err := client.PUT("/api/users/1", map[string]interface{}{"name": "Alice"}, &result, 200)
```

### PATCH

```go
var result map[string]interface{}
statusCode, respBody, err := client.PATCH("/api/users/1", map[string]interface{}{"email": "alice@example.com"}, &result, 200)
```

### DELETE

```go
var result map[string]interface{}
statusCode, respBody, err := client.DELETE("/api/users/1", &result, 200)
```

## Дополнительные заголовки

```go
client.SetAdditionalHeaders(map[string]string{
    "X-Custom-Header": "value",
    "X-Request-ID":    "12345",
})

// Очистка дополнительных заголовков
client.CleanAdditionalHeaders(map[string]string{})
```

## Debug-информация

```go
// После выполнения запроса
debugJSON, err := client.GetLastRequestDebugInfo()
if err != nil {
    panic(err)
}
fmt.Println(debugJSON)
// {
//   "URL": "https://api.example.com/api/users",
//   "Method": "GET",
//   "Headers": {...},
//   "RequestBody": "",
//   "ResponseBody": "{\"name\":\"Alice\"}"
// }
```

## Ошибки

Пакет определяет множество типов ошибок:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/less-rest-client"
)

statusCode, _, err := client.GET("/api/users", &result, 200)

if err != nil {
    switch e := err.(type) {
    case *lessrestclient.WrongStatusCodeError:
        fmt.Printf("Ожидался статус %d, получен %d\n", e.Expected, e.Actual)
        fmt.Printf("Тело запроса: %s\n", e.RequestRawBody)
        fmt.Printf("Тело ответа: %s\n", e.ResponseRawBody)

    case *lessrestclient.ClientError:
        fmt.Printf("Ошибка клиента: %v\n", e.Parent)

    case *lessrestclient.UnMarshalingError:
        fmt.Printf("Ошибка десериализации: %v\n", e.Parent)

    case *lessrestclient.MarshalingError:
        fmt.Printf("Ошибка сериализации: %v\n", e.Parent)
    }
}
```

### Типы ошибок

| Ошибка | Описание |
|--------|----------|
| `UnknownAuthTypeError` | Неизвестный тип аутентификации |
| `UnknownContentTypeError` | Неизвестный формат контента |
| `MarshalingError` | Ошибка сериализации в JSON |
| `UnMarshalingError` | Ошибка десериализации из JSON |
| `CreateRequestError` | Ошибка создания HTTP-запроса |
| `ClientError` | Ошибка выполнения HTTP-запроса |
| `ReadResponseBodyError` | Ошибка чтения тела ответа |
| `WrongStatusCodeError` | Несоответствие фактического и ожидаемого статуса |
| `FormDataError` | Ошибка формирования формы |
| `BodyFormMultipartDataError` | Ошибка multipart-формы |
| `BodyFormUrlencodedError` | Ошибка urlencoded-формы |
| `BodyJSONError` | Ошибка JSON-тела |
| `IsNotPointerError` | Аргумент не является указателем |

## Константы

```go
import "github.com/nobuenhombre/suikat/pkg/less-rest-client"

// Типы аутентификации
_ = lessrestclient.AuthTypeNone          // 0
_ = lessrestclient.AuthTypeLoginPass     // 1
_ = lessrestclient.AuthTypeToken         // 2
_ = lessrestclient.AuthTypeBearerToken   // 3
_ = lessrestclient.AuthTypeCookieWithCSRFToken // 4
_ = lessrestclient.AuthTypeEDS           // 5
```
