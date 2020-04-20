package admin

import (
	"api/framework/redis"
)

type Repo struct {
	db redis.Redis
}

func New(db redis.Redis) Repo {
	return Repo{db}
}
