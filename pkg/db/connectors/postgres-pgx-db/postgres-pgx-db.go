package postgrespgxdb

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nobuenhombre/suikat/pkg/db/types"
	"github.com/nobuenhombre/suikat/pkg/ge"
)

// StatementCacheMode values
// prepare - default
// describe - use it for PGBouncer connections
const (
	SCMPrepare  = "prepare"
	SCMDescribe = "describe"
)

// Start Config
//-------------------------------------------------
type Config struct {
	Host               string
	Port               string
	Name               string
	User               string
	Password           string
	SSLMode            string
	BinaryParameters   string // lib/pq setting for prepared statements in pgbouncer
	StatementCacheMode string
	MaxConnections     string
}

func (cfg *Config) GetDSN() string {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	q := url.Values{}

	if len(cfg.SSLMode) > 0 {
		q.Add("sslmode", cfg.SSLMode)
	}

	// lib/pq setting for prepared statements in pgbouncer
	// for PGX don't set this parameter
	if len(cfg.BinaryParameters) > 0 {
		q.Add("binary_parameters", cfg.BinaryParameters)
	}

	if len(cfg.StatementCacheMode) > 0 {
		q.Add("statement_cache_mode", cfg.StatementCacheMode)
	}

	if len(cfg.MaxConnections) > 0 {
		q.Add("pool_max_conns", cfg.MaxConnections)
	}

	dsn = dsn + "?" + q.Encode()

	return dsn
}

//-------------------------------------------------

// Start Connect
//-------------------------------------------------
type DBQuery interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Close()
	types.SQLLogger
}

type Conn struct {
	*pgxpool.Pool
	*types.DBLog
}

func New(cfg *Config, log types.SQLLoggerFunc) (DBQuery, error) {
	dsn := cfg.GetDSN()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"dsn": dsn})
	}

	// prepared statements for pgbouncer
	// https://blog.bullgare.com/2019/06/pgbouncer-and-prepared-statements/
	config.ConnConfig.PreferSimpleProtocol = true

	connectPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, ge.Pin(err)
	}

	return &Conn{
		Pool: connectPool,
		DBLog: &types.DBLog{
			Log: log,
		},
	}, nil
}
