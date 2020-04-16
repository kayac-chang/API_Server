package user

import (
	"api/model"
	"encoding/json"
)

// Store store user into redis
func (it Repo) Store(user model.User) error {

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := prefix + user.Username

	return it.db.Set(key, json)
}
