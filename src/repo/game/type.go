package game

import (
	"api/framework/redis"
)

// Repo Game Repo
type Repo struct {
	db redis.Redis
}

// New return Game Repo
func New(db redis.Redis) Repo {

	return Repo{db}
}
