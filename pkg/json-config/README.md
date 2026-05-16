# json-config

Пакет `json-config` предоставляет функцию для загрузки конфигурационных файлов в формате JSON.

## Обзор

Единственная функция `Load` читает файл по пути и десериализует его содержимое в указанную Go-структуру.

## Load

Загружает JSON-конфиг из файла и заполняет переданную структуру.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/json-config"
)

// Определение структуры конфига
type Config struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Database struct {
        Name string `json:"name"`
        User string `json:"user"`
        Pass string `json:"pass"`
    } `json:"database"`
}

// Загрузка конфига
var cfg Config
err := jsonconfig.Load("config.json", &cfg)
if err != nil {
    panic(err)
}

fmt.Printf("Host: %s\n", cfg.Host)
fmt.Printf("Port: %d\n", cfg.Port)
fmt.Printf("DB: %s\n", cfg.Database.Name)
```

### Пример JSON-файла

```json
{
    "host": "localhost",
    "port": 8080,
    "database": {
        "name": "mydb",
        "user": "admin",
        "pass": "secret"
    }
}
```

### Загрузка в map

```go
import "github.com/nobuenhombre/suikat/pkg/json-config"

var cfg map[string]interface{}
err := jsonconfig.Load("config.json", &cfg)
if err != nil {
    panic(err)
}
```

### Загрузка в []byte с последующей обработкой

```go
import (
    "encoding/json"
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/fico"
)

// Альтернативный подход: чтение вручную
txtFile := fico.TxtFile("config.json")
data, err := txtFile.Read()
if err != nil {
    panic(err)
}

var cfg Config
err = json.Unmarshal([]byte(data), &cfg)
if err != nil {
    panic(err)
}

fmt.Println(cfg)
```
