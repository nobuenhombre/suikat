// Package mariadb provides constructor to connect to mariadb with standard sql
package mariadb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nobuenhombre/suikat/pkg/db/types"
	"github.com/nobuenhombre/suikat/pkg/ge"
)

// Protocol = UNIXSocket, Address = "/tmp/mysql.sock"
// Protocol = TCP, Address = "localhost:5555"
const (
	ProtocolTCP        = "tcp"
	ProtocolUNIXSocket = "unix"
)

// Config describe connection params to database
type Config struct {
	Protocol string
	Address  string
	Schema   string
	User     string
	Password string
	Charset  string
}

// GetDSN
// return DSN string for db connection params
func (cfg *Config) GetDSN() string {
	dsn := fmt.Sprintf(
		"%v:%v@%v(%v)/%v?charset=%v&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Protocol,
		cfg.Address,
		cfg.Schema,
		cfg.Charset,
	)

	return dsn
}

// DBQuery
// describe interface
type DBQuery interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
	types.SQLLogger
}

// Conn
// database connection and log
type Conn struct {
	*sql.DB
	*types.DBLog
}

// New
// create new database connection
func New(cfg *Config, log types.SQLLoggerFunc) (DBQuery, error) {
	dsn := cfg.GetDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"dsn": dsn})
	}

	err = db.Ping()
	if err != nil {
		return nil, ge.Pin(err)
	}

	return &Conn{
		DB: db,
		DBLog: &types.DBLog{
			Log: log,
		},
	}, nil
}
