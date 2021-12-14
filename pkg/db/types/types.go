// Package types provides types to work with SQL databases
package types

import "time"

// SQLLoggerFunc provides the log func used by generated queries.
type SQLLoggerFunc func(sql string, du time.Duration, sqlParams ...interface{})

type SQLLogger interface {
	WriteLog(sql string, du time.Duration, sqlParams ...interface{})
}

type DBConfig interface {
	GetDSN() string
}

type DBLog struct {
	Log SQLLoggerFunc
}

func (dbl *DBLog) WriteLog(sql string, du time.Duration, sqlParams ...interface{}) {
	dbl.Log(sql, du, sqlParams...)
}
