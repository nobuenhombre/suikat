# dates

Пакет `dates` предоставляет функции для работы с датами и временем: константы форматов, вычисление разницы, диапазоны времени, получение таймзон.

## Зависимости

- `time/tzdata` — встроенные данные таймзон

## Форматы дат (константы)

| Константа | Формат |
|-----------|--------|
| `DateFormatDashYYYYMMDD` | `2006-01-02` |
| `DateFormatPointDDMMYYYY` | `02.01.2006` |
| `DateFormatPointDDMMYY` | `02.01.06` |
| `DateFormatDashDDMMYYYY` | `02-01-2006` |
| `DateFormatDashDDMMYY` | `02-01-06` |
| `DateTimeFormatDashYYYYMMDDHHmmss` | `2006-01-02 15:04:05` |
| `DateTimeFormat1C` | `2006-01-02T15:04:05` |
| `DateTimeFormatQiwi` | `02012006 15:04:05` |
| `DateFormatYYYYMMDD` | `20060102` |
| `DateTimeFormatYYYYMMDDHHmmss` | `20060102150405` |

## Тип `DateTimeDiff`

```go
type DateTimeDiff struct {
    Year  int
    Month int
    Day   int
    Hour  int
    Min   int
    Sec   int
}
```

## Функции

### Разница между датами

```go
func Diff(a, b time.Time) *DateTimeDiff
```

Вычисляет разницу между `a` и `b`. Возвращает `DateTimeDiff` с нормализованными отрицательными значениями.

### Начало дня/недели

```go
func BeginOfDay(t time.Time) time.Time
func BeginOfPrevDay(t time.Time) time.Time
func BeginOfNextDay(t time.Time) time.Time
func BeginOfPrevWeek(t time.Time) time.Time
func BeginOfNextWeek(t time.Time) time.Time
```

### Диапазоны времени

```go
func BeforePeriod(t time.Time, period int64, measure time.Duration) time.Time
func AfterPeriod(t time.Time, period int64, measure time.Duration) time.Time
func GetMonthRange(year int, month time.Month) (time.Time, time.Time)
func GetLastMonthRange(now time.Time) (time.Time, time.Time)
func GetLastWeekRange(now time.Time) (time.Time, time.Time)
func GetLast3dRange(now time.Time) (time.Time, time.Time)
func GetLast24HoursRange(now time.Time) (time.Time, time.Time)
func GetLastPeriodRange(now time.Time, period time.Duration) (time.Time, time.Time)
```

### Таймзоны

```go
func GetMoscowLocation() *time.Location
func GetSamaraLocation() *time.Location
func GetBerlinLocation() *time.Location
```

## Примеры

```go
package main

import (
    "fmt"
    "time"
    "github.com/nobuenhombre/suikat/pkg/dates"
)

func main() {
    // --- Пример 1: Разница между датами ---

    t1, _ := time.Parse(dates.DateTimeFormatDashYYYYMMDDHHmmss, "2026-01-01 10:00:00")
    t2, _ := time.Parse(dates.DateTimeFormatDashYYYYMMDDHHmmss, "2026-06-15 14:30:45")

    diff := dates.Diff(t1, t2)
    fmt.Printf("%d лет, %d мес, %d дн, %d ч, %d мин, %d сек\n",
        diff.Year, diff.Month, diff.Day, diff.Hour, diff.Min, diff.Sec)
    // 0 лет, 5 мес, 14 дн, 4 ч, 30 мин, 45 сек

    // --- Пример 2: Начало дня ---

    t, _ := time.Parse(dates.DateTimeFormatDashYYYYMMDDHHmmss, "2026-05-16 14:30:45")
    fmt.Println(dates.BeginOfDay(t))
    // 2026-05-16 00:00:00 +0000 UTC

    fmt.Println(dates.BeginOfNextDay(t))
    // 2026-05-17 00:00:00 +0000 UTC

    fmt.Println(dates.BeginOfPrevDay(t))
    // 2026-05-15 00:00:00 +0000 UTC

    // --- Пример 3: Диапазон месяца ---

    start, end := dates.GetMonthRange(2026, time.May)
    fmt.Printf("Май 2026: %v — %v\n", start, end)
    // Май 2026: 2026-05-01 00:00:00 +0000 UTC — 2026-05-31 23:59:59.999999999 +0000 UTC

    // --- Пример 4: Последний месяц ---

    now := time.Now()
    start, end := dates.GetLastMonthRange(now)
    fmt.Printf("Прошлый месяц: %v — %v\n", start, end)

    // --- Пример 5: Диапазон за N периодов ---

    start, end := dates.GetLast24HoursRange(now)
    fmt.Printf("Последние 24 часа: %v — %v\n", start, end)

    start, end = dates.GetLast3dRange(now)
    fmt.Printf("Последние 3 дня: %v — %v\n", start, end)

    // --- Пример 6: Таймзоны ---

    moscow := dates.GetMoscowLocation()
    berlin := dates.GetBerlinLocation()

    tMoscow := time.Date(2026, 5, 16, 12, 0, 0, 0, moscow)
    tBerlin := time.Date(2026, 5, 16, 12, 0, 0, 0, berlin)

    fmt.Printf("Москва: %v\n", tMoscow)
    fmt.Printf("Берлин: %v\n", tBerlin)

    // --- Пример 7: BeforePeriod / AfterPeriod ---

    t := time.Now()
    fmt.Println(dates.BeforePeriod(t, 7, 24*time.Hour))  // 7 дней назад
    fmt.Println(dates.AfterPeriod(t, 30, 24*time.Hour))  // 30 дней вперёд
}
```

## Поддиректория `i18n/ru`

Пакет `datesi18nru` — русская локализация дней недели и месяцев:

```go
import "github.com/nobuenhombre/suikat/pkg/dates/i18n/ru"

// Длинные названия дней
datesi18nru.GetLongDayName(time.Monday) // Понедельник

// Короткие названия дней
datesi18nru.GetShortDayName(time.Monday) // Пн

// Длинные названия месяцев
datesi18nru.GetLongMonthName(time.May) // Май

// Длинные названия месяцев с окончанием
datesi18nru.GetLongMonthNameExt(time.May) // Мая
```
