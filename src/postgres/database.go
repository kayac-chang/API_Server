package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "service"
	password = "123456"
	sslmode  = "disable"
)

type DB struct {
	*sqlx.DB
}

type Tx struct {
	*sqlx.Tx
}

type Query func(Tx) error

func (db DB) Transact(fn Query) error {

	tx, err := db.Beginx()

	if err != nil {
		return err
	}

	err = fn(Tx{tx})

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (db DB) MustPreparex(query string) *sqlx.Stmt {

	stmt, err := db.Preparex(query)

	if err != nil {
		log.Fatal(err)
	}

	return stmt
}

func New(name string) DB {

	info := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, sslmode,
	)

	db := sqlx.MustConnect("postgres", info)

	fmt.Println("Connection Successed...")

	return DB{db}
}
