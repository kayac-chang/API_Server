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
