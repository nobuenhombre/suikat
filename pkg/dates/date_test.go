package dates

import (
	"reflect"
	"testing"
	"time"
)

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
}

func TestDiff(t *testing.T) {
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
