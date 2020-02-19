package db

import (
	"database/sql"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

type DB struct {
	*sql.DB
}

func Run(url string) (*DB, error) {

	config, err := pgx.ParseURI(url)

	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDB(config)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
