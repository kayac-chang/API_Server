package game

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/model"
	repo "api/repo/game"
	"api/utils"
	"time"
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

func (it *Usecase) Find(name string) (*model.Game, error) {

	return it.repo.FindByID(utils.MD5(name))
}

func (it *Usecase) FindAll() ([]*model.Game, error) {

	return it.repo.FindAll()
}

func (it *Usecase) Store(name, href, category string) (*model.Game, error) {

	game := model.Game{
		ID: utils.MD5(name),

		Name:     name,
		Href:     href,
		Category: category,

		CreatedAt: time.Now(),
	}

	if err := it.repo.Store(&game); err != nil {
		return nil, err
	}

	return &game, nil
}
