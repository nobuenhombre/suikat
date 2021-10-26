package postgrespgxdb

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

const (
	ConstraintErrorCode = "23505"
)

func IsConstraintError(err error, name string) bool {
	var e *pgconn.PgError
	if errors.As(err, &e) {
		return e.Code == ConstraintErrorCode && e.ConstraintName == name
	}

	return false
}

type AffectedLessThenNecessaryError struct {
	Affected  int64
	Necessary int64
}

func (e *AffectedLessThenNecessaryError) Error() string {
	return fmt.Sprintf("Affected [%v] Less then Necessary [%v]", e.Affected, e.Necessary)
}
