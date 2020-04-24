package token

import (
	"api/framework/redis"

	"github.com/mediocregopher/radix/v3"
)

const table = "tokens"

// Repo type for persistence layer logic
type Repo struct {
	redis redis.Redis
}

// New create repo for tokens associate table
func New(redis redis.Redis) Repo {

	return Repo{redis}
}

// Store associate token and primary key from another table
func (it Repo) Store(token, pk string) error {

	return it.redis.Write(table, func(conn radix.Conn) error {

		return conn.Do(radix.Cmd(nil, "HSETNX", table, token, pk))
	})
}

// Find get primary key with token
func (it Repo) Find(token string) (string, error) {

	return it.redis.Read("HGET", table, token)
}
