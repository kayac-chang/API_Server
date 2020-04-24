package game

import (
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"encoding/json"

	"github.com/mediocregopher/radix/v3"
)

const table = "games"

// Repo Game Repo
type Repo struct {
	redis redis.Redis
	db    postgres.DB
}

// New return Game Repo
func New(redis redis.Redis, db postgres.DB) Repo {

	return Repo{redis, db}
}

// Store store admin in game
func (it Repo) Store(game *model.Game) error {

	data, err := json.Marshal(game)
	if err != nil {
		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {

		err := conn.Do(
			radix.Cmd(nil, "HSETNX", table, game.ID, string(data)),
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

// FindByName find game by name in redis and db
func (it Repo) FindByName(name string) (*model.Game, error) {

	game := model.Game{}

	findInRedis := func() error {

		data, err := it.redis.Read("HGET", table, name)
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(data), &game); err != nil {
			return err
		}

		return nil
	}

	findInDB := func() error {

		sql := "SELECT * FROM $1 WHERE name=$2"

		if err := it.db.Ping(); err != nil {

			return err
		}

		return it.db.Get(&game, sql, table, name)
	}

	err := findInRedis()

	if err != nil && err != model.ErrNotFound {

		return nil, err
	}

	if err == model.ErrNotFound {

		if err := findInDB(); err != nil {

			return nil, err
		}
	}

	return &game, nil
}
