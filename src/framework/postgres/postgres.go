package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func New(url string, timeout int) *DB {

	_timeout := time.Duration(timeout) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), _timeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "pgx", url)

	if err != nil {
		panic(err)
	}

	return &DB{db}
}

func (it *DB) IsConstraintErr(err error) bool {

	pgerr, ok := err.(pgx.PgError)
	if !ok {
		return false
	}

	errlist := [...]string{
		pgerrcode.IntegrityConstraintViolation,
		pgerrcode.RestrictViolation,
		pgerrcode.NotNullViolation,
		pgerrcode.ForeignKeyViolation,
		pgerrcode.UniqueViolation,
		pgerrcode.CheckViolation,
		pgerrcode.ExclusionViolation,
	}

	for _, code := range errlist {
		if code == pgerr.Code {
			return true
		}
	}

	return false
}
