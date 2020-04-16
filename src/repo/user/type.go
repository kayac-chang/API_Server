package user

import (
	"api/framework/redis"
)

const prefix = "users:"

type Repo struct {
	db redis.Redis
}

func New(db redis.Redis) Repo {

	return Repo{db}
}
