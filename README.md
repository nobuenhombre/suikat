# suikat

[![Build Status](https://travis-ci.com/nobuenhombre/suikat.svg?branch=master)](https://app.travis-ci.com/github/nobuenhombre/suikat)
[![Go Report Card](https://goreportcard.com/badge/github.com/nobuenhombre/suikat)](https://goreportcard.com/report/github.com/nobuenhombre/suikat)
[![codecov](https://codecov.io/gh/nobuenhombre/suikat/branch/master/graph/badge.svg)](https://codecov.io/gh/nobuenhombre/suikat)
[![GitHub Release](https://img.shields.io/github/release/nobuenhombre/suikat.svg)](https://github.com/nobuenhombre/suikat/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/nobuenhombre/suikat.svg)](https://pkg.go.dev/github.com/nobuenhombre/suikat)

# Suikat

Набор Go-пакетов для различных задач: работа с файлами, сетью, текстом, ошибками, рефлексией, шифрованием и др.

## Пакеты

### Работа с файлами и директориями

| Пакет | Описание |
|-------|----------|
| **[adapt](pkg/adapt/)** | Адаптеры для работы с различными источниками данных |
| **[binhex](pkg/binhex/)** | Кодирование/декодирование в HEX |
| **[cascade](pkg/cascade/)** | Цепочки обработчиков данных |
| **[chunks](pkg/chunks/)** | Разбиение строк и слайсов на чанки |
| **[converter](pkg/converter/)** | Конвертация типов (строки, числа и др.) |
| **[fico](pkg/fico/)** | Работа с содержимым текстовых файлов: чтение, запись, сжатие (GZ, BR), кодирование (Base64, HEX), Data URI |
| **[fina](pkg/fina/)** | Работа с частями имён файлов: префиксы, суффиксы, смена расширения |
| **[figlu](pkg/figlu/)** | Склеивание содержимого файлов из директории или списка путей в единый файл/строку с фильтрацией по расширению |
| **[fitree](pkg/fitree/)** | Рекурсивное сканирование дерева каталогов и получение списка файлов |
| **[futi](pkg/futi/)** | Утилиты: проверка существования, создание temp-файлов, копирование, перемещение, удаление |
| **[imgs](pkg/imgs/)** | Получение размеров изображений по пути к файлу |
| **[mimes](pkg/mimes/)** | Константы MIME-типов и функция получения MIME по расширению файла |
| **[scatterfs](pkg/scatterfs/)** | Виртуальная файловая система с распределением по нескольким базовым директориям на основе свободного места |

### Сетевые пакеты

| Пакет | Описание |
|-------|----------|
| **[credentials](pkg/credentials/)** | Хранение и управление учётными данными |
| **[interceptor](pkg/interceptor/)** | HTTP-interceptor для перехвата и модификации HTTP-запросов и ответов с маршрутизацией |
| **[less-rest-client](pkg/less-rest-client/)** | Облегчённый HTTP-клиент для REST API |
| **[sfu](pkg/sfu/)** | Преобразование структур Go в URL-encoded form data |
| **[yank](pkg/yank/)** | Полнофункциональный HTTP-клиент: аутентификация (Basic, Bearer, Cookie+CSRF, EDSS), контент (JSON, form-urlencoded, multipart), тайминги, буферизация |

### Работа с текстом

| Пакет | Описание |
|-------|----------|
| **[chacha](pkg/chacha/)** | Валидатор узлов дерева |
| **[clivar](pkg/clivar/)** | Парсинг командной строки (CLI variables) |
| **[colorog](pkg/colorog/)** | Цветное логирование |
| **[csvarser](pkg/csvarser/)** | Парсер CSV в структуру на рефлексии |
| **[dates](pkg/dates/)** | Работа с датами |
| **[inwords](pkg/inwords/)** | Преобразование чисел в текстовое представление на русском языке (числительные, склонения) |
| **[replacer](pkg/replacer/)** | Замена строк по правилам: регулярные выражения и точные строковые совпадения |
| **[webula](pkg/webula/)** | Обработка текстов для web: нормализация, транслитерация, очистка HTML, URL-имена, алфавитные указатели |

### Ошибки и обработка

| Пакет | Описание |
|-------|----------|
| **[ge](pkg/ge/)** | Расширенная система работы с ошибками: точное место возникновения, параметры контекста, проверка типа ошибки. Встроенные ошибки: NotFoundError, MismatchError, LimitCountExhaustedError, RegExpIsNotCompiledError и др. |

### Рефлексия и структуры

| Пакет | Описание |
|-------|----------|
| **[refavour](pkg/refavour/)** | Работа с тегами структур и reflect: проверка типов, чтение полей, TagProcessor интерфейс |

### Шифрование

| Пакет | Описание |
|-------|----------|
| **[secreto](pkg/secreto/)** | Шифрование/расшифровка данных: AES-GCM, scrypt KDF. Функции: Encrypt/Decrypt (с паролем), EncryptSimple/DecryptSimple (с прямым ключом), GenerateSimpleKey |

### Работа с базами данных

| Пакет | Описание |
|-------|----------|
| **[db/types](pkg/db/types/)** | Общие типы для SQL-баз: `SQLLoggerFunc`, `SQLLogger`, `DBConfig`, `DBLog` |
| **[db/connectors/maria-db](pkg/db/connectors/maria-db/)** | Коннектор к MariaDB через `database/sql` с драйвером mysql. Поддержка TCP и Unix-сокетов, `parseTime=true` |
| **[db/connectors/postgres-pgx-db](pkg/db/connectors/postgres-pgx-db/)** | Коннектор к PostgreSQL через pgx/pgxpool. Поддержка PgBouncer, кэширование запросов, пакетные запросы (`SendBatch`), проверка constraint-ошибок |

### Утилиты

| Пакет | Описание |
|-------|----------|
| **[envar](pkg/envar/)** | Работа с переменными окружения |
| **[execution-tree](pkg/execution-tree/)** | Дерево выполнения (execution tree) для построения конвейеров обработки с ветвлением |
| **[fifo](pkg/fifo/)** | Работа с именованными каналами (FIFO / named pipes) в Linux/Unix |
| **[gt](pkg/gt/)** | Текстовые и HTML-шаблоны: рендеринг одиночных и множественных шаблонов |
| **[hash](pkg/hash/)** | Вычисление MD5-хеша строки |
| **[inslice](pkg/inslice/)** | Поиск значений в срезах и проверка индексов |
| **[jsonconfig](pkg/jsonconfig/)** | Загрузка конфигурационных файлов в формате JSON |
| **[os-detector](pkg/os-detector/)** | Определение операционной системы |
| **[osexec](pkg/osexec/)** | Выполнение команд в зависимости от ОС (Unix/Windows) |
| **[parallelworks](pkg/parallelworks/)** | Параллельное выполнение работы с использованием горутин |
| **[periodicwork](pkg/periodicwork/)** | Периодическое выполнение работ с graceful shutdown |
| **[pipelin](pkg/pipelin/)** | Конвейеры обработки данных с каналами и горутинами |
| **[repeater](pkg/repeater/)** | Повторный вызов функции с лимитом попыток и таймаутом (для нестабильных внешних API) |
| **[terminator](pkg/terminator/)** | Graceful shutdown сервисов при получении OS-сигналов (SIGINT, SIGTERM) |
| **[tracktime](pkg/tracktime/)** | Измерение и логирование времени выполнения участков кода |

## Зависимости проекта

Проект использует следующие внешние зависимости:
- `github.com/andybalholm/brotli` — Brotli сжатие
- `github.com/mholt/archiver/v3` — GZ/Brotli архивация
- `github.com/microcosm-cc/bluemonday` — санитизация HTML
- `golang.org/x/crypto/scrypt` — scrypt KDF
