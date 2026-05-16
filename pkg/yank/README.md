# yank

Пакет `yank` предоставляет мощный HTTP-клиент с поддержкой аутентификации, конструкторов запросов, таймингов и лёгкого режима работы.

## Обзор

Пакет включает:
- `Client` — основной HTTP-клиент
- `Defaults` — настройки по умолчанию (URL, авторизация, заголовки)
- `Request` — конструктор HTTP-запросов
- `Response` — результат HTTP-запроса
- `LightService` — лёгкий режим для быстрых запросов
- Типы аутентификации: None, Basic, Token, Bearer, Cookie+CSRF, EDS
- Поддержка JSON, Form, Multipart контента
- Детальные тайминги запросов

## Создание клиента

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

// Создание настроек по умолчанию
defaults := yank.NewDefaults("https://api.example.com")

// Создание клиента
client := yank.New(defaults)
```

## Настройки по умолчанию (Defaults)

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

defaults := yank.NewDefaults("https://api.example.com")

// Заголовки
defaults.AddHeader("Content-Type", "application/json")
defaults.AddHeader("Accept", "application/json")

// Авторизация
defaults.AuthBasic("user", "pass")
defaults.AuthToken("my-token")
defaults.AuthBearerToken("bearer-token")
defaults.AuthCookieWithCSRFToken("cookie-name", "csrf-token")
defaults.AuthEDS("key-id", "body")
defaults.AuthNone()

// Следование редиректам
defaults.SetFollowRedirects(true)

// Transport
// defaults.SetTransport(customTransport)
```

## Request — конструктор запросов

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

request := yank.NewRequest("/api/users")
request.SetURL("https://api.example.com")
request.AddQuery("page", "1")
request.AddQuery("limit", "10")
request.AddHeader("X-Custom-Header", "value")
```

### Ответ (Response)

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

// Структура для парсинга ответа
type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

response := yank.NewResponse(&UserResponse{}, 200)
```

## HTTP-методы

### GET

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)

type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

request := yank.NewRequest("/api/users/1")
request.AddQuery("fields", "name,email")

response := yank.NewResponse(&UserResponse{}, 200)

err := client.GET(request, response, false)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}

userResponse := response.Data.(*UserResponse)
fmt.Printf("ID: %d, Name: %s, Email: %s\n", userResponse.ID, userResponse.Name, userResponse.Email)
fmt.Printf("HTTP код: %d\n", response.HTTPCode)
fmt.Printf("Тело: %s\n", string(response.Raw))
```

### POST

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)

// Запрос
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Ответ
type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

request := yank.NewRequest("/api/users")
request.SetBody(yank.NewJSON(CreateUserRequest{Name: "Иван", Email: "ivan@example.com"}))

response := yank.NewResponse(&UserResponse{}, 201)

err := client.POST(request, response, false)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}

userResponse := response.Data.(*UserResponse)
fmt.Printf("Создан пользователь: %s (ID: %d)\n", userResponse.Name, userResponse.ID)
```

### PUT

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)

type UpdateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

request := yank.NewRequest("/api/users/1")
request.SetBody(yank.NewJSON(UpdateUserRequest{Name: "Иван", Email: "ivan@example.com"}))

response := yank.NewResponse(&UserResponse{}, 200)

err := client.PUT(request, response, false)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}
```

### PATCH

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)

type PatchUserRequest struct {
    Email string `json:"email"`
}

type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

request := yank.NewRequest("/api/users/1")
request.SetBody(yank.NewJSON(PatchUserRequest{Email: "new@example.com"}))

response := yank.NewResponse(&UserResponse{}, 200)

err := client.PATCH(request, response, false)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}
```

### DELETE

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)

request := yank.NewRequest("/api/users/1")

response := yank.NewResponse(nil, 204)

err := client.DELETE(request, response, false)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}

fmt.Println("Пользователь удалён")
```

## Аутентификация

### None (без авторизации)

```go
request := yank.NewRequest("/api/public")
request.AuthNone()
```

### Basic

```go
request := yank.NewRequest("/api/protected")
request.AuthBasic("username", "password")
```

### Token

```go
request := yank.NewRequest("/api/protected")
request.AuthToken("secret-token")
```

### Bearer Token

```go
request := yank.NewRequest("/api/protected")
request.AuthBearerToken("bearer-token-value")
```

### Cookie + CSRF Token

```go
request := yank.NewRequest("/api/protected")
request.AuthCookieWithCSRFToken("session-cookie", "csrf-token-value")
```

### EDS

```go
request := yank.NewRequest("/api/protected")
request.AuthEDS("key-id", "body")
```

## Типы контента

### JSON

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

request.SetBody(yank.NewJSON(User{Name: "Иван", Email: "ivan@example.com"}))
```

### Form URL-encoded

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

type FormData struct {
    Username string `form:"username"`
    Password string `form:"password"`
}

request.SetBody(yank.NewFormURLEncoded(FormData{
    Username: "ivan",
    Password: "secret",
}))
```

### Form Multipart

```go
import (
    "github.com/nobuenhombre/suikat/pkg/yank"
)

type MultipartData struct {
    File    string `form:"file"`
    Comment string `form:"comment"`
}

request.SetBody(yank.NewFormMultipart(MultipartData{
    File:    "data:text/plain;base64,SGVsbG8=",
    Comment: "Привет",
}))
```

## Лёгкий режим (Light Client)

Упрощённый интерфейс для быстрых запросов без ручного создания Request/Response:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)
light := client.Light()

// GET
var user UserResponse
statusCode, rawBody, err := light.GET("/api/users/1", &user, 200)
if err != nil {
    fmt.Println("Ошибка:", err)
}
fmt.Printf("Статус: %d, Тело: %s\n", statusCode, string(rawBody))

// POST
var createdUser UserResponse
sendData := CreateUserRequest{Name: "Иван", Email: "ivan@example.com"}
statusCode, rawBody, err = light.POST("/api/users", sendData, &createdUser, 201, mimes.JSON)
if err != nil {
    fmt.Println("Ошибка:", err)
}

// PUT
var updatedUser UserResponse
statusCode, rawBody, err = light.PUT("/api/users/1", sendData, &updatedUser, 200, mimes.JSON)

// PATCH
var patchedUser UserResponse
statusCode, rawBody, err = light.PATCH("/api/users/1", PatchUserRequest{Email: "new@example.com"}, &patchedUser, 200, mimes.JSON)

// DELETE
statusCode, rawBody, err = light.DELETE("/api/users/1", nil, nil, 204, "")
```

## Тайминги

Пакет автоматически измеряет время выполнения запроса:

```go
import (
    "fmt"
    "time"
    "github.com/nobuenhombre/suikat/pkg/yank"
)

defaults := yank.NewDefaults("https://api.example.com")
client := yank.New(defaults)

request := yank.NewRequest("/api/users")
response := yank.NewResponse(&UserResponse{}, 200)

err := client.GET(request, response, false)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}

timing := response.Timing
fmt.Printf("Подготовка запроса: %v\n", timing.PrepareRequest)
fmt.Printf("Установка соединения: %v\n", timing.Connect)
fmt.Printf("Отправка запроса: %v\n", timing.SendRequest)
fmt.Printf("Первый байт ответа: %v\n", timing.TimeToFirstByte)
fmt.Printf("Скачивание тела: %v\n", timing.DownloadContent)
fmt.Printf("Парсинг ответа: %v\n", timing.ParseResponse)
fmt.Printf("Всего: %v\n", timing.Total)
```

## Полный пример с авторизацией и таймингами

```go
import (
    "encoding/json"
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

func main() {
    // 1. Настройки
    defaults := yank.NewDefaults("https://api.example.com")
    defaults.AddHeader("Accept", "application/json")
    defaults.AuthBearerToken("my-bearer-token")

    client := yank.New(defaults)

    // 2. Создаём пользователя
    type CreateUserRequest struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    type UserResponse struct {
        ID    int    `json:"id"`
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    postRequest := yank.NewRequest("/api/users")
    postRequest.SetBody(yank.NewJSON(CreateUserRequest{
        Name:  "Иван",
        Email: "ivan@example.com",
    }))

    postResponse := yank.NewResponse(&UserResponse{}, 201)

    err := client.POST(postRequest, postResponse, false)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    user := postResponse.Data.(*UserResponse)
    fmt.Printf("Создан пользователь: %s (ID: %d)\n", user.Name, user.ID)
    fmt.Printf("Тайминг: %v\n", postResponse.Timing.Total)

    // 3. Получаем пользователя
    getRequest := yank.NewRequest(fmt.Sprintf("/api/users/%d", user.ID))

    getResponse := yank.NewResponse(&UserResponse{}, 200)

    err = client.GET(getRequest, getResponse, false)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    fetchedUser := getResponse.Data.(*UserResponse)
    bodyJSON, _ := json.MarshalIndent(fetchedUser, "", "  ")
    fmt.Printf("Пользователь: %s\n", string(bodyJSON))

    // 4. Обновляем пользователя
    patchRequest := yank.NewRequest(fmt.Sprintf("/api/users/%d", user.ID))
    patchRequest.SetBody(yank.NewJSON(map[string]string{"email": "new@example.com"}))

    patchResponse := yank.NewResponse(&UserResponse{}, 200)

    err = client.PATCH(patchRequest, patchResponse, false)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    // 5. Удаляем пользователя
    deleteRequest := yank.NewRequest(fmt.Sprintf("/api/users/%d", user.ID))

    deleteResponse := yank.NewResponse(nil, 204)

    err = client.DELETE(deleteRequest, deleteResponse, false)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    fmt.Println("Пользователь удалён")
}
```

## Полный пример с Light Client

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/yank"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

func main() {
    defaults := yank.NewDefaults("https://api.example.com")
    defaults.AddHeader("Accept", "application/json")
    defaults.AuthBearerToken("my-bearer-token")

    client := yank.New(defaults)
    light := client.Light()

    // GET
    var user UserResponse
    statusCode, rawBody, err := light.GET("/api/users/1", &user, 200)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Printf("Статус: %d, Тело: %s\n", statusCode, string(rawBody))

    // POST
    var createdUser UserResponse
    statusCode, rawBody, err = light.POST("/api/users", CreateUserRequest{Name: "Иван", Email: "ivan@example.com"}, &createdUser, 201, mimes.JSON)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

    // PUT
    var updatedUser UserResponse
    statusCode, rawBody, err = light.PUT("/api/users/1", CreateUserRequest{Name: "Иван", Email: "ivan@example.com"}, &updatedUser, 200, mimes.JSON)

    // DELETE
    statusCode, rawBody, err = light.DELETE("/api/users/1", nil, nil, 204, "")
}
```

## Структура пакета

```
yank/
├── auth.go                    -- базовый тип Auth
├── auth_basic.go              -- Basic Auth
├── auth_bearer-token.go       -- Bearer Token Auth
├── auth_cookie-with-csrf-token.go -- Cookie + CSRF
├── auth_eds.go                -- EDS Auth
├── auth_none.go               -- Без авторизации
├── auth_token.go              -- Token Auth
├── base-content.go            -- базовый тип Content
├── base-content_bytes.go      -- бинарный контент
├── base-content_form-multipart-data.go -- Multipart
├── base-content_form-url-encoded.go -- Form URL-encoded
├── base-content_json.go       -- JSON контент
├── body-buffer.go             -- буфер тела
├── client.go                  -- основной Client
├── client-errors.go           -- типы ошибок
├── defaults.go                -- настройки по умолчанию
├── http.go                    -- HTTPRequest/HTTPResponse
├── interfaces.go              -- интерфейсы конструкторов
├── light-client.go            -- лёгкий режим
├── raw-content.go             -- сырой контент
├── request.go                 -- Request
├── response.go                -- Response
├── timing.go                  -- тайминги
├── client_test.go             -- тесты
└── light-client_test.go       -- тесты light-client
```
