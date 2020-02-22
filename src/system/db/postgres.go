package db

import (
	"github.com/KayacChang/API_Server/system/log"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// DB sqlx.DB wrapper
type DB struct {
	*sqlx.DB
}

// New Create DB instance
func New(dataSource string) *DB {

	db, err := sqlx.Connect("pgx", dataSource)

	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	log.Info("Connect to Postgres success...\n")

	return &DB{db}
}
