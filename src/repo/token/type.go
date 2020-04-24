package token

import (
	"api/framework/redis"

	"github.com/mediocregopher/radix/v3"
)

const table = "tokens"

// Repo type for persistence layer logic
type Repo struct {
	db redis.Redis
}

// New create repo for tokens associate table
func New(db redis.Redis) Repo {

	return Repo{db}
}

// Store associate token and primary key from another table
func (it Repo) Store(token, pk string) error {

	return it.db.Write(table, func(conn radix.Conn) error {

		return conn.Do(radix.Cmd(nil, "HSETNX", table, token, pk))
	})
}
