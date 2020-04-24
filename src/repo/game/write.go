package game

import (
	"api/model"
	"encoding/json"

	"github.com/mediocregopher/radix/v3"
)

// Store store game in db
func (it Repo) Store(game *model.Game) error {

	data, err := json.Marshal(game)
	if err != nil {
		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {

		insert := "insert:" + table
		err = conn.Do(
			radix.Cmd(nil, "LPUSH", insert, string(data)),
		)
		if err != nil {
			return err
		}

		return nil
	})
}

// Update update game in db
func (it Repo) Update(game *model.Game) error {

	data, err := json.Marshal(game)
	if err != nil {
		return err
	}

	return it.redis.Write(table, func(conn radix.Conn) error {

		update := "update:" + table
		err = conn.Do(
			radix.Cmd(nil, "LPUSH", update, string(data)),
		)
		if err != nil {
			return err
		}

		return nil
	})
}
