# imgs

Пакет `imgs` предоставляет утилиту для получения размеров изображений по пути к файлу.

## Функции

### GetDimension

Возвращает ширину и высоту изображения, не загружая всё изображение в память. Использует `image.DecodeConfig` для быстрого чтения только заголовка файла.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/imgs"
)

width, height, err := imgs.GetDimension("photo.jpg")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Размер: %dx%d\n", width, height)
```

### Поддерживаемые форматы

Пакет поддерживает все форматы, зарегистрированные в `image/png`, `image/jpeg`, `image/gif` и других, доступных через стандартную библиотеку `image`.

```go
// Работает с любыми форматами, зарегистрированными в image
width, height, err := imgs.GetDimension("photo.png")   // PNG
width, height, err = imgs.GetDimension("photo.jpg")    // JPEG
width, height, err = imgs.GetDimension("photo.gif")    // GIF
width, height, err = imgs.GetDimension("photo.webp")   // WebP (требуется image/webp)
```

### Проверка существования файла

```go
import (
    "os"
    "github.com/nobuenhombre/suikat/pkg/imgs"
)

if _, err := os.Stat("photo.jpg"); os.IsNotExist(err) {
    log.Fatal("Файл не существует")
}

width, height, err := imgs.GetDimension("photo.jpg")
if err != nil {
    // Обработка ошибок: файл не найден, повреждён или неподдерживаемый формат
    log.Printf("Ошибка: %v", err)
}
```
