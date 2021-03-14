package types

// Log provides the log func used by generated queries.
type SQLLoggerFunc func(sql string, sqlParams ...interface{})

type SQLLogger interface {
	WriteLog(sql string, sqlParams ...interface{})
}

type DBConfig interface {
	GetDSN() string
}

type DBLog struct {
	Log SQLLoggerFunc
}

func (dbl *DBLog) WriteLog(sql string, sqlParams ...interface{}) {
	dbl.Log(sql, sqlParams...)
}
