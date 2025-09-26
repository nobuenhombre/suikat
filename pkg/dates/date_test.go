package dates

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	type inputParams struct {
		DateA time.Time
		DateB time.Time
	}

	type datesDiffTest struct {
		in  inputParams
		out *DateTimeDiff
	}

	var datesDiffTests = []datesDiffTest{
		{
			in: inputParams{
				DateA: time.Date(2015, 5, 1, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 6, 2, 1, 1, 1, 1, time.UTC),
			},
			out: &DateTimeDiff{
				Year:  1,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
			},
		},
		{
			in: inputParams{
				DateA: time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			out: &DateTimeDiff{
				Year:  0,
				Month: 0,
				Day:   30,
				Hour:  0,
				Min:   0,
				Sec:   0,
			},
		},
		{
			in: inputParams{
				DateA: time.Date(2016, 2, 2, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			out: &DateTimeDiff{
				Year:  0,
				Month: 0,
				Day:   28,
				Hour:  0,
				Min:   0,
				Sec:   0,
			},
		},
		{
			in: inputParams{
				DateA: time.Date(2015, 2, 11, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			out: &DateTimeDiff{
				Year:  0,
				Month: 11,
				Day:   1,
				Hour:  0,
				Min:   0,
				Sec:   0,
			},
		},
		{
			in: inputParams{
				DateA: time.Date(2016, 2, 11, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			out: &DateTimeDiff{
				Year:  0,
				Month: 0,
				Day:   30,
				Hour:  0,
				Min:   0,
				Sec:   0,
			},
		},
		{
			in: inputParams{
				DateA: time.Date(2016, 2, 11, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 1, 12, 0, 0, 0, 0, GetMoscowLocation()),
			},
			out: &DateTimeDiff{
				Year:  0,
				Month: 0,
				Day:   30,
				Hour:  3,
				Min:   0,
				Sec:   0,
			},
		},
		{
			in: inputParams{
				DateA: time.Date(2016, 2, 11, 0, 0, 0, 0, time.UTC),
				DateB: time.Date(2016, 1, 12, 0, 7, 5, 0, GetSamaraLocation()),
			},
			out: &DateTimeDiff{
				Year:  0,
				Month: 0,
				Day:   30,
				Hour:  3,
				Min:   52,
				Sec:   55,
			},
		},
	}

	for i := 0; i < len(datesDiffTests); i++ {
		test := &datesDiffTests[i]
		out := Diff(test.in.DateA, test.in.DateB)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Diff(%#v, %#v), Expected %#v, Actual %#v",
				test.in.DateA, test.in.DateB, test.out, out,
			)
		}
	}
}

func TestGetUTC(t *testing.T) {
	nowCurrentLocale := time.Now()
	nowUTCLocale := nowCurrentLocale.UTC()
	nowUTC2UTCLocale := nowUTCLocale.UTC().UTC()

	t.Logf("NCL   = %v\n", nowCurrentLocale)
	t.Logf("NCL.L   = %v\n", nowCurrentLocale.Location())
	t.Logf("NUL   = %v\n", nowUTCLocale)
	t.Logf("NUL.L   = %v\n", nowUTCLocale.Location())
	t.Logf("NU2UL = %v\n", nowUTC2UTCLocale)
	t.Logf("NU2UL.L = %v\n", nowUTC2UTCLocale.Location())
	t.Logf("NULL   = %v\n", nowUTCLocale.Local())
	t.Logf("NULL.L   = %v\n", nowUTCLocale.Local().Location())
	t.Logf("NCLU   = %v\n", nowCurrentLocale.UTC())
	t.Logf("NCLU.L   = %v\n", nowCurrentLocale.UTC().Location())

	if !reflect.DeepEqual(nowUTCLocale, nowUTC2UTCLocale) {
		t.Errorf(
			"Expected %#v,\n Actual %#v",
			nowUTC2UTCLocale, nowUTCLocale,
		)
	}
}

type convertDateTest struct {
	in  time.Time
	out time.Time
}

func TestBeginOfDay(t *testing.T) {
	var beginOfDayTests = []convertDateTest{
		{
			in:  time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			out: time.Date(2015, 5, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for i := 0; i < len(beginOfDayTests); i++ {
		test := &beginOfDayTests[i]
		out := BeginOfDay(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"BeginOfDay(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

func TestBeginOfPrevDay(t *testing.T) {
	var beginOfPrevDayTests = []convertDateTest{
		{
			in:  time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			out: time.Date(2015, 4, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	for i := 0; i < len(beginOfPrevDayTests); i++ {
		test := &beginOfPrevDayTests[i]
		out := BeginOfPrevDay(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"BeginOfPrevDay(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

func TestBeginOfNextDay(t *testing.T) {
	var beginOfNextDayTests = []convertDateTest{
		{
			in:  time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			out: time.Date(2015, 5, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	for i := 0; i < len(beginOfNextDayTests); i++ {
		test := &beginOfNextDayTests[i]
		out := BeginOfNextDay(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"BeginOfNextDay(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

func TestBeginOfPrevWeek(t *testing.T) {
	var beginOfPrevWeekTests = []convertDateTest{
		{
			in:  time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			out: time.Date(2015, 4, 24, 0, 0, 0, 0, time.UTC),
		},
	}

	for i := 0; i < len(beginOfPrevWeekTests); i++ {
		test := &beginOfPrevWeekTests[i]
		out := BeginOfPrevWeek(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"BeginOfPrevWeek(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

func TestBeginOfNextWeek(t *testing.T) {
	var beginOfNextWeekTests = []convertDateTest{
		{
			in:  time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			out: time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
		},
	}

	for i := 0; i < len(beginOfNextWeekTests); i++ {
		test := &beginOfNextWeekTests[i]
		out := BeginOfNextWeek(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"BeginOfNextWeek(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

type shiftDateTest struct {
	in      time.Time
	period  int64
	measure time.Duration
	out     time.Time
}

func TestBeforePeriod(t *testing.T) {
	var beforePeriodTests = []shiftDateTest{
		{
			in:      time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			period:  10,
			measure: time.Minute,
			out:     time.Date(2015, 5, 1, 16, 35, 12, 87, time.UTC),
		},
	}

	for i := 0; i < len(beforePeriodTests); i++ {
		test := &beforePeriodTests[i]
		out := BeforePeriod(test.in, test.period, test.measure)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"BeforePeriod(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

func TestAfterPeriod(t *testing.T) {
	var afterPeriodTests = []shiftDateTest{
		{
			in:      time.Date(2015, 5, 1, 16, 45, 12, 87, time.UTC),
			period:  10,
			measure: time.Minute,
			out:     time.Date(2015, 5, 1, 16, 55, 12, 87, time.UTC),
		},
	}

	for i := 0; i < len(afterPeriodTests); i++ {
		test := &afterPeriodTests[i]
		out := AfterPeriod(test.in, test.period, test.measure)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"AfterPeriod(%#v), Expected %#v, Actual %#v",
				test.in, test.out, out,
			)
		}
	}
}

func TestGetMonthRange(t *testing.T) {
	start, end := GetMonthRange(2025, time.September)
	require.Equal(t, time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC), start)
	require.Equal(t, time.Date(2025, time.September, 30, 23, 59, 59, 999999999, time.UTC), end)
}
