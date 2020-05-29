package order

import (
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mediocregopher/radix/v3"
)

const table = "orders"

// Repo ...
type Repo struct {
	redis redis.Redis
	db    postgres.DB
}

// New ...
func New(redis redis.Redis, db postgres.DB) Repo {

	return Repo{redis, db}
}

// FindByID ...
func (it Repo) FindByID(id string) (*model.Order, error) {
	var err error
	order := model.Order{}

	findInRedis := func() error {
		res, err := it.redis.Read("HGET", table, id)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(res), &order)
		if err != nil {
			return err
		}

		return nil
	}

	findInDB := func() error {
		sql := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", table)

		if err := it.db.Ping(); err != nil {
			return err
		}

		return it.db.Get(&order, sql, id)
	}

	if err = findInRedis(); err == nil {
		return &order, nil
	}
	if err != model.ErrNotFound {
		return nil, err
	}
	if err = findInDB(); err == nil {
		return &order, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	return nil, model.ErrNotFound
}

// Store ...
func (it Repo) Store(order *model.Order) error {

	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {
		err := conn.Do(
			radix.Cmd(nil, "HSET", table, order.ID, string(data)),
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
