package order

import (
	"api/framework/redis"
	"api/model"
	"encoding/json"

	"github.com/mediocregopher/radix/v3"
)

const table = "orders"

// Repo ...
type Repo struct {
	redis redis.Redis
}

// New ...
func New(redis redis.Redis) Repo {

	return Repo{redis}
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
