package game

import (
	"api/framework/redis"
	"api/model"
)

const prefix = "games:"

// Repo Game Repo
type Repo struct {
	db redis.Redis
}

// New return Game Repo
func New(db redis.Redis) Repo {

	return Repo{db}
}

// Find game by name
func (it Repo) Find(name string) (*model.Game, error) {

	game := model.Game{}

	if err := it.db.Get(prefix+name, &game); err != nil {
		return nil, err
	}

	return &game, nil
}
