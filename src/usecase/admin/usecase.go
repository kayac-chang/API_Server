package admin

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	repo "api/repo/admin"
)

type Usecase struct {
	env  env.Env
	repo *repo.Repo
}

func New(env env.Env, db *postgres.DB, c *cache.Cache) *Usecase {

	return &Usecase{
		env:  env,
		repo: repo.New(db, c),
	}
}
