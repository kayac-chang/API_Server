package game

import (
	"api/env"
	"api/framework/redis"
	"api/repo/game"
)

type Usecase struct {
	env  env.Env
	repo game.Repo
}

func New(env env.Env, db redis.Redis) Usecase {

	repo := game.New(db)

	return Usecase{env, repo}
}
