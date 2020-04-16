package token

import (
	"api/framework/redis"
	"encoding/json"
)

const prefix = "tokens:"

// Repo type for persistence layer logic
type Repo struct {
	db redis.Redis
}

// New create repo for tokens associate table
func New(db redis.Redis) Repo {

	return Repo{db}
}

// Store store associate table with key by token
func (it Repo) Store(token string, associate map[string]string) error {

	json, err := json.Marshal(associate)
	if err != nil {
		return err
	}

	key := prefix + token

	return it.db.Set(key, json)
}
