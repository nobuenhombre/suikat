# parallelworks

Пакет `parallelworks` предоставляет механизм для параллельного вызова функций с коллекцией результатов.

## Обзор

Пакет позволяет запустить заданную функцию-воркер параллельно для множества элементов данных и собрать результаты в `map[int]interface{}` по ID.

## Worker

Тип функции-воркера:

```go
type Worker func(data JobData) JobData
```

Принимает `JobData` и возвращает `JobData` (результат обработки).

## JobData

Структура данных для передачи в воркер и получения результата:

```go
type JobData struct {
    ID   int
    Data interface{}
}
```

- `ID` — уникальный идентификатор задачи (используется как ключ в карте результатов)
- `Data` — произвольные данные для передачи в воркер / результат обработки

## Run

Запускает воркер параллельно для всех элементов `dataList` с заданным количеством воркеров.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/parallelworks"
)

// Определение воркера
worker := func(data parallelworks.JobData) parallelworks.JobData {
    // Обработка data.Data
    result := data.Data.(int) * 2
    return parallelworks.JobData{
        ID:   data.ID,
        Data: result,
    }
}

// Исходные данные
dataList := []parallelworks.JobData{
    {ID: 1, Data: 10},
    {ID: 2, Data: 20},
    {ID: 3, Data: 30},
}

// Запуск с 4 воркерами
results := worker.Run(dataList, 4)

// Чтение результатов
for id, result := range results {
    fmt.Printf("ID %d: %v\n", id, result)
}
// Вывод:
// ID 1: 20
// ID 2: 40
// ID 3: 60
```

## Пример: обработка файлов

```go
import (
    "fmt"
    "os"
    "github.com/nobuenhombre/suikat/pkg/parallelworks"
)

type FileInfo struct {
    Name string
    Size int64
}

type FileStat struct {
    Name string
    Size int64
    Error string
}

// Воркер для получения информации о файле
worker := func(data parallelworks.JobData) parallelworks.JobData {
    name := data.Data.(string)
    info, err := os.Stat(name)
    if err != nil {
        return parallelworks.JobData{
            ID:   data.ID,
            Data: FileStat{Name: name, Error: err.Error()},
        }
    }
    return parallelworks.JobData{
        ID:   data.ID,
        Data: FileStat{Name: name, Size: info.Size()},
    }
}

// Список файлов
files := []string{"/etc/hosts", "/etc/passwd", "/nonexistent"}
dataList := make([]parallelworks.JobData, len(files))
for i, f := range files {
    dataList[i] = parallelworks.JobData{ID: i, Data: f}
}

// Параллельное выполнение
results := worker.Run(dataList, 4)

for id, stat := range results {
    s := stat.(FileStat)
    if s.Error != "" {
        fmt.Printf("Файл %s: ошибка — %s\n", s.Name, s.Error)
    } else {
        fmt.Printf("Файл %s: %d байт\n", s.Name, s.Size)
    }
}
```

## Пример: веб-запросы

```go
import (
    "fmt"
    "net/http"
    "github.com/nobuenhombre/suikat/pkg/parallelworks"
)

type URLStatus struct {
    URL     string
    Status  int
    Error   string
}

worker := func(data parallelworks.JobData) parallelworks.JobData {
    url := data.Data.(string)
    resp, err := http.Get(url)
    if err != nil {
        return parallelworks.JobData{
            ID:   data.ID,
            Data: URLStatus{URL: url, Error: err.Error()},
        }
    }
    defer resp.Body.Close()
    return parallelworks.JobData{
        ID:   data.ID,
        Data: URLStatus{URL: url, Status: resp.StatusCode},
    }
}

urls := []string{
    "https://google.com",
    "https://github.com",
    "https://golang.org",
}

dataList := make([]parallelworks.JobData, len(urls))
for i, u := range urls {
    dataList[i] = parallelworks.JobData{ID: i, Data: u}
}

results := worker.Run(dataList, 10)

for _, stat := range results {
    s := stat.(URLStatus)
    if s.Error != "" {
        fmt.Printf("%s: %s\n", s.URL, s.Error)
    } else {
        fmt.Printf("%s: %d\n", s.URL, s.Status)
    }
}
```

## Константы

```go
import "github.com/nobuenhombre/suikat/pkg/parallelworks"

// Максимальное количество воркеров
_ = parallelworks.WorkersCount // 64

// Множитель длины каналов: max(len(dataList), WorkersCountToChanLenMultiplier * workersCount)
_ = parallelworks.WorkersCountToChanLenMultiplier // 2
```

## Примечания

- Воркеры работают параллельно через `goroutine`
- Количество воркеров задаётся параметром `workersCount`
- Результаты возвращаются в виде `map[int]interface{}`, где ключ — `ID` задачи
- При большом количестве задач длина каналов ограничивается: `2 * workersCount`
