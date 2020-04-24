package user

import (
	"api/framework/redis"
	"api/model"
	"encoding/json"

	"github.com/mediocregopher/radix/v3"
)

const table = "users"

// Repo ...
type Repo struct {
	db redis.Redis
}

// New ...
func New(db redis.Redis) Repo {

	return Repo{db}
}

// Store store user into redis
func (it Repo) Store(user *model.User) error {

	data, err := json.Marshal(user)
	if err != nil {

		return err
	}

	return it.db.Write(table, func(conn radix.Conn) error {

		err := conn.Do(
			radix.Cmd(nil, "HSET", table, user.ID, string(data)),
		)
		if err != nil {
			return err
		}

		pending := "pending:" + table
		err = conn.Do(
			radix.Cmd(nil, "LPUSH", pending, string(data)),
		)
		if err != nil {
			return err
		}

		return nil
	})
}
