package admin

import (
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"database/sql"
	"encoding/json"

	"github.com/mediocregopher/radix/v3"
)

const table = "admins"

// Repo type for persistence layer logic
type Repo struct {
	redis redis.Redis
	db    postgres.DB
}

// New create repo for tokens associate table
func New(redis redis.Redis, db postgres.DB) Repo {
	return Repo{redis, db}
}

// Store store admin in redis
func (it Repo) Store(admin *model.Admin) error {

	data, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {

		err := conn.Do(
			radix.Cmd(nil, "HSETNX", table, admin.ID, string(data)),
		)
		if err != nil {
			return err
		}

		pending := "pending:" + table
		err = conn.Do(radix.Cmd(nil, "LPUSH", pending, string(data)))
		if err != nil {
			return err
		}

		return nil
	})
}

// FindByID find admin by specify id
func (it Repo) FindByID(id string) (*model.Admin, error) {

	var err error
	admin := model.Admin{}

	findInRedis := func() error {

		res, err := it.redis.Read("HGET", table, id)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(res), &admin)
		if err != nil {
			return err
		}

		return nil
	}

	findInDB := func() error {

		sql := "SELECT * FROM admins WHERE admin_id=$1"

		if err := it.db.Ping(); err != nil {

			return err
		}

		return it.db.Get(&admin, sql, id)
	}

	if err := findInRedis(); err == nil {

		return &admin, nil
	}

	if err = findInDB(); err == nil {

		return &admin, nil
	}
	if err != sql.ErrNoRows {

		return nil, err
	}

	return nil, model.ErrNotFound
}
