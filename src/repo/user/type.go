package user

import (
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mediocregopher/radix/v3"
)

const table = "users"

// Repo ...
type Repo struct {
	redis redis.Redis
	db    postgres.DB
}

// New ...
func New(redis redis.Redis, db postgres.DB) Repo {

	return Repo{redis, db}
}

// Store store user into redis
func (it Repo) Store(user *model.User) error {

	data, err := json.Marshal(user)
	if err != nil {

		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {

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

// FindByID find user by id in redis
func (it Repo) FindByID(id string) (*model.User, error) {

	var err error
	user := model.User{}

	findInRedis := func() error {

		res, err := it.redis.Read("HGET", table, id)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(res), &user)
		if err != nil {
			return err
		}

		return nil
	}

	findInDB := func() error {

		sql := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", table)

		if err := it.db.Ping(); err != nil {
			return err
		}

		return it.db.Get(&user, sql, id)
	}

	if err = findInRedis(); err == nil {

		return &user, nil
	}
	if err != model.ErrNotFound {

		return nil, err
	}

	if err = findInDB(); err == nil {

		return &user, nil
	}
	if err != sql.ErrNoRows {

		return nil, err
	}

	return nil, model.ErrNotFound
}
