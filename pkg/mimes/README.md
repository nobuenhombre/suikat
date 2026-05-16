# mimes

Пакет `mimes` предоставляет константы MIME-типов и функцию для определения MIME-типа по расширению файла.

## Обзор

Пакет содержит более 70 констант MIME-типов для различных форматов файлов (изображения, документы, архивы, шрифты, видео, аудио и т.д.), а также функцию `GetByExt` для получения MIME-типа по расширению файла.

## GetByExt

Возвращает MIME-тип по расширению файла. Если расширение не найдено — возвращает `application/octet-stream`.

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

fmt.Println(mimes.GetByExt(".png"))    // image/png
fmt.Println(mimes.GetByExt(".jpg"))    // image/jpeg
fmt.Println(mimes.GetByExt(".pdf"))    // application/pdf
fmt.Println(mimes.GetByExt(".json"))   // application/json
fmt.Println(mimes.GetByExt(".zip"))    // application/zip
fmt.Println(mimes.GetByExt(".txt"))    // text/plain
fmt.Println(mimes.GetByExt(".xyz"))    // application/octet-stream (неизвестный тип)
```

## Константы MIME-типов

### Изображения

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.WindowsBitmapGraphics     // "image/bmp"
_ = mimes.GraphicsInterchangeFormat  // "image/gif"
_ = mimes.JPEGImages                 // "image/jpeg"
_ = mimes.PortableNetworkGraphics    // "image/png"
_ = mimes.ScalableVectorGraphics     // "image/svg+xml"
_ = mimes.TaggedImageFileFormat      // "image/tiff"
_ = mimes.WebPImage                  // "image/webp"
_ = mimes.Icon                       // "image/vnd.microsoft.icon"
```

### Документы

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

// PDF
_ = mimes.AdobePortableDocumentFormat // "application/pdf"

// Microsoft Office
_ = mimes.MicrosoftWord              // "application/msword"
_ = mimes.MicrosoftWordOpenXML       // "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
_ = mimes.MicrosoftExcel             // "application/vnd.ms-excel"
_ = mimes.MicrosoftExcelOpenXML      // "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
_ = mimes.MicrosoftPowerPoint        // "application/vnd.ms-powerpoint"
_ = mimes.MicrosoftPowerPointOpenXML // "application/vnd.openxmlformats-officedocument.presentationml.presentation"
_ = mimes.MicrosoftVisio             // "application/vnd.visio"

// OpenDocument
_ = mimes.OpenDocumentText           // "application/vnd.oasis.opendocument.text"
_ = mimes.OpenDocumentSpreadsheet    // "application/vnd.oasis.opendocument.spreadsheet"
_ = mimes.OpenDocumentPresentation   // "application/vnd.oasis.opendocument.presentation"

// Другое
_ = mimes.AbiWordDocument            // "application/x-abiword"
_ = mimes.AmazonKindleEBook          // "application/vnd.amazon.ebook"
_ = mimes.ElectronicPublication      // "application/epub+zip"
_ = mimes.RichTextFormat             // "application/rtf"
_ = mimes.AppleInstallerPackage      // "application/vnd.apple.installer+xml"
```

### Архивы

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.ZIPArchive                  // "application/zip"
_ = mimes.BZipArchive                 // "application/x-bzip"
_ = mimes.BZip2Archive                // "application/x-bzip2"
_ = mimes.GZipCompressedArchive       // "application/gzip"
_ = mimes.RARArchive                  // "application/vnd.rar"
_ = mimes.TapeArchive                 // "application/x-tar"
_ = mimes.SevenZipArchive             // "application/x-7z-compressed"
_ = mimes.ArchiveDocument             // "application/x-freearc"
_ = mimes.BinaryData                  // "application/octet-stream" (общий бинарный тип)
```

### Аудио

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.AACAudio                    // "audio/aac"
_ = mimes.MP3Audio                    // "audio/mpeg"
_ = mimes.OGGAudio                    // "audio/ogg"
_ = mimes.OpusAudio                   // "audio/opus"
_ = mimes.MusicalInstrumentDigitalInterface // "audio/midi"
_ = mimes.WaveformAudio               // "audio/wav"
_ = mimes.WEBMAudio                   // "audio/webm"
```

### Видео

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.AudioVideoInterleave        // "video/x-msvideo"
_ = mimes.MPEGVideo                   // "video/mpeg"
_ = mimes.OGGVideo                    // "video/ogg"
_ = mimes.MPEGTransportStream         // "video/mp2t"
_ = mimes.WEBMVideo                   // "video/webm"
```

### Шрифты

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.OpenTypeFont                // "font/otf"
_ = mimes.TrueTypeFont                // "font/ttf"
_ = mimes.WebOpenFontFormat           // "font/woff"
_ = mimes.WebOpenFontFormat2          // "font/woff2"
_ = mimes.MicrosoftEmbeddedOpenTypeFont // "application/vnd.ms-fontobject"
```

### Текстовые форматы

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.Text                        // "text/plain"
_ = mimes.HyperTextMarkupLanguage     // "text/html"
_ = mimes.CascadingStyleSheets        // "text/css"
_ = mimes.CommaSeparatedValues        // "text/csv"
_ = mimes.JavaScript                  // "text/javascript"
_ = mimes.ICalendar                   // "text/calendar"
_ = mimes.XHTML                       // "application/xhtml+xml"
_ = mimes.XML                         // "application/xml"
_ = mimes.JSON                        // "application/json"
_ = mimes.JSONLD                      // "application/ld+json"
_ = mimes.WebManifest                 // "application/manifest+json"
```

### Формы

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.FormUrlencoded    // "application/x-www-form-urlencoded"
_ = mimes.FormMultipartData // "multipart/form-data"
```

### Скрипты

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.CShellScript      // "application/x-csh"
_ = mimes.BourneShellScript // "application/x-sh"
_ = mimes.PHPHypertextPreprocessor // "application/php"
_ = mimes.JavaArchive       // "application/java-archive"
```

### Другое

```go
import "github.com/nobuenhombre/suikat/pkg/mimes"

_ = mimes.AdobeFlashDocument    // "application/x-shockwave-flash"
```

## Пример: определение MIME-типа для загрузки файла

```go
import (
    "fmt"
    "path/filepath"
    "github.com/nobuenhombre/suikat/pkg/mimes"
)

func handleUpload(filename string, data []byte) {
    ext := filepath.Ext(filename)
    mime := mimes.GetByExt(ext)
    
    fmt.Printf("Файл: %s\n", filename)
    fmt.Printf("Расширение: %s\n", ext)
    fmt.Printf("MIME-тип: %s\n", mime)
    fmt.Printf("Размер: %d байт\n", len(data))
}

handleUpload("photo.jpg", imageData)
// Файл: photo.jpg
// Расширение: .jpg
// MIME-тип: image/jpeg
// Размер: 12345 байт
```
