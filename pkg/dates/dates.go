package dates

import "time"

const (
	DateFormatDashYYYYMMDD           = "2006-01-02"
	DateFormatPointDDMMYYYY          = "02.01.2006"
	DateTimeFormatDashYYYYMMDDHHmmss = "2006-01-02 15:04:05"
	DateTimeFormat1C                 = "2006-01-02T15:04:05"
)

const (
	WeekDays  = 7
	MonthDays = 31
	YearDays  = 365
)

type DateTimeDiff struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Sec   int
}

func (dd *DateTimeDiff) InSeconds() int64 {
	out := int64(0)
	out += int64(dd.Sec)
	out += int64(dd.Min * 60)
	out += int64(dd.Hour * 60 * 60)
	out += int64(dd.Day * 60 * 60 * 24)

	return out
}

func Diff(a, b time.Time) *DateTimeDiff {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := y2 - y1
	month := int(M2 - M1)
	day := d2 - d1
	hour := h2 - h1
	min := m2 - m1
	sec := s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return &DateTimeDiff{
		Year:  year,
		Month: month,
		Day:   day,
		Hour:  hour,
		Min:   min,
		Sec:   sec,
	}
}

func BeginOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func BeginOfPrevDay(t time.Time) time.Time {
	prevDay := t.AddDate(0, 0, -1)
	return BeginOfDay(prevDay)
}

func BeginOfNextDay(t time.Time) time.Time {
	nextDay := t.AddDate(0, 0, 1)
	return BeginOfDay(nextDay)
}

func BeginOfPrevWeek(t time.Time) time.Time {
	prevDay := t.AddDate(0, 0, -7)
	return BeginOfDay(prevDay)
}

func BeginOfNextWeek(t time.Time) time.Time {
	nextDay := t.AddDate(0, 0, 7)
	return BeginOfDay(nextDay)
}

func BeforePeriod(t time.Time, period int64, measure time.Duration) time.Time {
	return t.Add(time.Duration(-1*period) * measure)
}

func AfterPeriod(t time.Time, period int64, measure time.Duration) time.Time {
	return t.Add(time.Duration(period) * measure)
}
