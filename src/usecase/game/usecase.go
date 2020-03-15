package game

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/model"
	repo "api/repo/game"
	"api/utils"
)

type Usecase struct {
	env  *env.Env
	repo *repo.Repo
}

func New(env *env.Env, db *postgres.DB, c *cache.Cache) *Usecase {

	return &Usecase{
		env:  env,
		repo: repo.New(db, c),
	}
}

func (it *Usecase) Find(names ...string) ([]*model.Game, error) {

	ids := make([]string, len(names))

	for _, name := range names {
		ids = append(ids, utils.MD5(name))
	}

	games, err := it.repo.FindByID(ids)

	if err != nil {
		return nil, err
	}

	return games, nil
}

func (it *Usecase) Store() error {

	return nil
}
