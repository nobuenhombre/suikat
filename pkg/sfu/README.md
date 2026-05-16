# sfu

Пакет `sfu` конвертирует Go-структуры в `url.Values` (form-encoded данные) для использования с HTTP-запросами.

## Обзор

Пакет рекурсивно обходит поля структуры и преобразует их в пары ключ-значение в формате `url.Values`. Поддерживаются вложенные структуры и срезы.

Поля структуры должны иметь тег `form:"имя_ключа"`.

## Convert

Конвертирует структуру в `url.Values`:

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type User struct {
    Name string `form:"name"`
    Age  int64  `form:"age"`
}

user := &User{
    Name: "Alice",
    Age:  30,
}

form := &url.Values{}
err := sfu.Convert(user, "", form)
if err != nil {
    panic(err)
}

fmt.Println(form.Get("name")) // Alice
fmt.Println(form.Get("age"))  // 30
```

### Простая структура

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Config struct {
    Host     string  `form:"host"`
    Port     int64   `form:"port"`
    Debug    bool    `form:"debug"`
    Coeff    float64 `form:"coeff"`
}

cfg := &Config{
    Host:  "localhost",
    Port:  8080,
    Debug: true,
    Coeff: 1.5,
}

form := &url.Values{}
err := sfu.Convert(cfg, "", form)
if err != nil {
    panic(err)
}

for key, values := range *form {
    fmt.Printf("%s: %v\n", key, values)
}
// Вывод:
// host: [localhost]
// port: [8080]
// debug: [true]
// coeff: [1.5]
```

### Вложенная структура

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Address struct {
    Country string `form:"country"`
    City    string `form:"city"`
    Street  string `form:"street"`
}

type User struct {
    Name    string   `form:"name"`
    Address Address  `form:"address"`
}

user := &User{
    Name: "Alice",
    Address: Address{
        Country: "Russia",
        City:    "Moscow",
        Street:  "Lenina",
    },
}

form := &url.Values{}
err := sfu.Convert(user, "", form)
if err != nil {
    panic(err)
}

// address[country]: Russia
// address[city]: Moscow
// address[street]: Lenina
for key, values := range *form {
    fmt.Printf("%s: %v\n", key, values)
}
```

### Вложенные структуры с префиксом

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Credentials struct {
    Login string `form:"login"`
    Pass  string `form:"pass"`
}

type Auth struct {
    Credentials Credentials `form:"credentials"`
}

auth := &Auth{
    Credentials: Credentials{
        Login: "admin",
        Pass:  "secret",
    },
}

form := &url.Values{}
// parent = "auth" добавляет префикс к ключам
err := sfu.Convert(auth, "auth", form)
if err != nil {
    panic(err)
}

for key, values := range *form {
    fmt.Printf("%s: %v\n", key, values)
}
// Вывод:
// auth[credentials][login]: [admin]
// auth[credentials][pass]: [secret]
```

### Срез простых типов

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Settings struct {
    Robots   []string `form:"robots"`
    Timeouts []int64  `form:"timeouts"`
    Sizes    []float64 `form:"sizes"`
    Allows   []bool   `form:"allows"`
}

settings := &Settings{
    Robots:   []string{"Mail.ru", "Yandex-Bot", "Google-Bot"},
    Timeouts: []int64{100, 200, 300},
    Sizes:    []float64{1.5, 2.5, 3.5},
    Allows:   []bool{true, false, true},
}

form := &url.Values{}
err := sfu.Convert(settings, "", form)
if err != nil {
    panic(err)
}

for key, values := range *form {
    fmt.Printf("%s: %v\n", key, values)
}
// Вывод:
// robots[0]: [Mail.ru]
// robots[1]: [Yandex-Bot]
// robots[2]: [Google-Bot]
// timeouts[0]: [100]
// timeouts[1]: [200]
// timeouts[2]: [300]
// sizes[0]: [1.5]
// sizes[1]: [2.5]
// sizes[2]: [3.5]
// allows[0]: [true]
// allows[1]: [false]
// allows[2]: [true]
```

### Срез вложенных структур

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Tunnel struct {
    IPFrom  string `form:"ipFrom"`
    IPTo    string `form:"ipTo"`
    Enabled bool   `form:"enabled"`
}

type Network struct {
    Name    string   `form:"name"`
    Tunnels []Tunnel `form:"tunnels"`
}

network := &Network{
    Name: "production",
    Tunnels: []Tunnel{
        {IPFrom: "10.0.0.1", IPTo: "10.0.0.2", Enabled: true},
        {IPFrom: "10.0.0.3", IPTo: "10.0.0.4", Enabled: false},
    },
}

form := &url.Values{}
err := sfu.Convert(network, "", form)
if err != nil {
    panic(err)
}

for key, values := range *form {
    fmt.Printf("%s: %v\n", key, values)
}
// Вывод:
// name: [production]
// tunnels[0][ipFrom]: [10.0.0.1]
// tunnels[0][ipTo]: [10.0.0.2]
// tunnels[0][enabled]: [true]
// tunnels[1][ipFrom]: [10.0.0.3]
// tunnels[1][ipTo]: [10.0.0.4]
// tunnels[1][enabled]: [false]
```

### Пример использования с HTTP-запросом

```go
import (
    "net/http"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/less-rest-client"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type LoginRequest struct {
    Username string `form:"username"`
    Password string `form:"password"`
}

// Подготовка данных
login := &LoginRequest{
    Username: "admin",
    Password: "secret",
}

form := &url.Values{}
err := sfu.Convert(login, "", form)
if err != nil {
    panic(err)
}

// Отправка через less-rest-client
client := &lessrestclient.Client{
    URL:         "https://api.example.com",
    ContentType: mimes.FormUrlencoded,
}

var result map[string]interface{}
statusCode, respBody, err := client.POST("/login", form, &result, 200)
```

## Ошибки

### Не структура

```go
import (
    "github.com/nobuenhombre/suikat/pkg/ge"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

var value int = 42
form := &url.Values{}

err := sfu.Convert(value, "", form)
if err != nil {
    if _, ok := err.(*ge.MismatchError); ok {
        fmt.Println("Ожидалась структура")
    }
}
```

### Не указатель

```go
import (
    "github.com/nobuenhombre/suikat/pkg/ge"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Config struct {
    Host string `form:"host"`
}

cfg := Config{Host: "localhost"} // не указатель!
form := &url.Values{}

err := sfu.Convert(&cfg, "", form) // указатель на структуру
if err != nil {
    fmt.Println(err)
}

// Ошибка будет, если передать указатель на значение:
// err := sfu.Convert(cfg, "", form) // cfg не указатель
```

### Неизвестный тип поля

```go
import (
    "time"
    "github.com/nobuenhombre/suikat/pkg/ge"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type BadConfig struct {
    Created time.Time `form:"created"` // time.Time не поддерживается
}

bad := &BadConfig{Created: time.Now()}
form := &url.Values{}

err := sfu.Convert(bad, "", form)
if err != nil {
    if _, ok := err.(*ge.UnknownTypeError); ok {
        fmt.Println("Неизвестный тип поля")
    }
}
```

## Поддерживаемые типы

| Тип Go | Преобразование |
|--------|---------------|
| `string` | Строковое значение |
| `int64` | `fmt.Sprintf("%v", value)` |
| `float64` | `fmt.Sprintf("%v", value)` |
| `bool` | `fmt.Sprintf("%v", value)` |
| `struct` | Рекурсивная конвертация |
| `[]string` | Срез строк |
| `[]int64` | Срез int64 |
| `[]float64` | Срез float64 |
| `[]bool` | Срез bool |
| `[]struct` | Срез структур (рекурсивно) |

## Пример: полная форма регистрации

```go
import (
    "fmt"
    "net/url"
    "github.com/nobuenhombre/suikat/pkg/sfu"
)

type Address struct {
    Country string `form:"country"`
    City    string `form:"city"`
    Zip     string `form:"zip"`
}

type RegistrationForm struct {
    Email    string   `form:"email"`
    Password string   `form:"password"`
    Name     string   `form:"name"`
    Age      int64    `form:"age"`
    IsActive bool     `form:"is_active"`
    Address  Address  `form:"address"`
    Tags     []string `form:"tags"`
}

form := &RegistrationForm{
    Email:    "alice@example.com",
    Password: "secret123",
    Name:     "Alice",
    Age:      30,
    IsActive: true,
    Address: Address{
        Country: "Russia",
        City:    "Moscow",
        Zip:     "101000",
    },
    Tags: []string{"admin", "verified"},
}

values := &url.Values{}
err := sfu.Convert(form, "", values)
if err != nil {
    panic(err)
}

// Вывод всех полей
for key, vals := range *values {
    fmt.Printf("%s=%s\n", key, vals[0])
}
// Вывод:
// email=alice@example.com
// password=secret123
// name=Alice
// age=30
// is_active=true
// address[country]=Russia
// address[city]=Moscow
// address[zip]=101000
// tags[0]=admin
// tags[1]=verified
```
