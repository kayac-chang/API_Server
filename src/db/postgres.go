package db

import (
	"log"

	"github.com/KayacChang/API_Server/entity"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func Run(env *entity.Env) (*DB, error) {

	db, err := sqlx.Connect("pgx", env.Postgres.ToURL())

	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	log.Printf("Connect to Postgres success...\n")

	return &DB{db}, nil
}
