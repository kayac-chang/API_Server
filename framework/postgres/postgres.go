package postgres

import (
	"api/env"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func New(env env.Env) DB {

	db := sqlx.MustOpen("pgx", env.Postgres.ToURL())

	return DB{db}
}
