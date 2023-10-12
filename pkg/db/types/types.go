// Package types provides types to work with SQL databases
package types

import "time"

// SQLLoggerFunc provides the log func used by generated queries.
type SQLLoggerFunc func(sql string, du time.Duration, sqlParams ...any)

type SQLLogger interface {
	WriteLog(s string, du time.Duration, optionsAndArgs ...any)
}

type DBConfig interface {
	GetDSN() string
}

type DBLog struct {
	Log SQLLoggerFunc
}

func (dbl *DBLog) WriteLog(sql string, du time.Duration, sqlParams ...any) {
	dbl.Log(sql, du, sqlParams...)
}
