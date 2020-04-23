package admin

import (
	"api/env"
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"api/repo/admin"
	"api/utils"
)

type Usecase struct {
	env  env.Env
	repo admin.Repo
}

func New(env env.Env, redis redis.Redis, db postgres.DB) Usecase {

	return Usecase{
		env:  env,
		repo: admin.New(redis, db),
	}
}

func (it Usecase) Find(email string) (*model.Admin, error) {

	id := utils.MD5(email)

	return it.repo.FindByID(id)
}
