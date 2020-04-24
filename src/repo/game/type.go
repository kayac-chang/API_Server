package game

import (
	"api/framework/postgres"
	"api/framework/redis"
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
