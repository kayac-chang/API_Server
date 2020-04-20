package admin

import (
	"api/env"
	"api/framework/redis"
	"api/repo/admin"
)

type Usecase struct {
	env  env.Env
	repo admin.Repo
}

func New(env env.Env, db redis.Redis) Usecase {

	return Usecase{
		env:  env,
		repo: admin.New(db),
	}
}
