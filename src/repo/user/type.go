package user

import (
	"api/framework/redis"
)

type Repo struct {
	*redis.Redis
}

func New(redis *redis.Redis) *Repo {

	return &Repo{redis}
}
