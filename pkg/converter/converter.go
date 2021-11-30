package converter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type ParserError struct {
	ParserType string
	Value      string
	Parent     error
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("ParserType [%v], Value [%v], Error [%v]", e.ParserType, e.Value, e.Parent.Error())
}

func StringToInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Int",
			Value:      s,
			Parent:     err,
		})
	}

	return n, nil
}

func StringToInt8(s string) (int8, error) {
	n, err := strconv.ParseInt(s, 10, 8)

	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Int8",
			Value:      s,
			Parent:     err,
		})
	}

	return int8(n), nil
}

func StringToInt16(s string) (int16, error) {
	n, err := strconv.ParseInt(s, 10, 16)

	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Int16",
			Value:      s,
			Parent:     err,
		})
	}

	return int16(n), nil
}

func StringToInt32(s string) (int32, error) {
	n, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Int32",
			Value:      s,
			Parent:     err,
		})
	}

	return int32(n), nil
}

func StringToInt64(s string) (int64, error) {
	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Int64",
			Value:      s,
			Parent:     err,
		})
	}

	return n, nil
}

func StringToBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)

	if err != nil {
		return false, ge.Pin(&ParserError{
			ParserType: "Bool",
			Value:      s,
			Parent:     err,
		})
	}

	return b, nil
}

func StringToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)

	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Float32",
			Value:      s,
			Parent:     err,
		})
	}

	return float32(f), nil
}

func StringToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, ge.Pin(&ParserError{
			ParserType: "Float64",
			Value:      s,
			Parent:     err,
		})
	}

	return f, nil
}

func StringToTime(s, format string) (time.Time, error) {
	t, err := time.Parse(format, s)

	if err != nil {
		return time.Time{}, ge.Pin(&ParserError{
			ParserType: "time.Time",
			Value:      s,
			Parent:     err,
		})
	}

	return t, nil
}

func StringToIntSlice(s, sep string) ([]int, error) {
	sList := strings.Split(s, sep)

	out := make([]int, 0, len(sList))

	for i, sValue := range sList {
		iValue, err := StringToInt(sValue)
		if err != nil {
			return []int{}, ge.Pin(&ParserError{
				ParserType: "IntSlice",
				Value:      fmt.Sprintf("s=%v sep=%v index=%v", s, sep, i),
				Parent:     err,
			})
		}

		out = append(out, iValue)
	}

	return out, nil
}
