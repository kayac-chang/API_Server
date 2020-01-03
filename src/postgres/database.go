package postgres

import (
	"fmt"

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

func New(name string) *sqlx.DB {

	info := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, sslmode,
	)

	db := sqlx.MustConnect("postgres", info)

	fmt.Println("Connection Successed...")

	return db
}
