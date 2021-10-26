package postgrespgxdb

import (
	"fmt"

	"github.com/jackc/pgconn"
)

const (
	ConstraintErrorCode = "23505"
)

func IsConstraintError(err error, name string) bool {
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return false
	}

	return pgErr.Code == ConstraintErrorCode && pgErr.ConstraintName == name
}

type AffectedLessThenNecessaryError struct {
	Affected  int64
	Necessary int64
}

func (e *AffectedLessThenNecessaryError) Error() string {
	return fmt.Sprintf("Affected [%v] Less then Necessary [%v]", e.Affected, e.Necessary)
}
