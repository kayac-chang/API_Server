package usecase

import (
	"api/game/repo"
	"api/model"
	"api/utils"
)

type usercase struct {
	db    repo.Repository
	cache repo.Repository
}

type Usecase interface {
	Find(game *model.Game) error
	Store(game *model.Game) error
	FindAll(games *[]model.Game) error
}

func New(db, cache repo.Repository) Usecase {

	return &usercase{db, cache}
}

func (it *usercase) Find(game *model.Game) error {

	return it.db.FindBy("ID", game)
}

func (it *usercase) FindAll(games *[]model.Game) error {

	return it.db.FindAll(games)
}

func (it *usercase) Store(game *model.Game) error {

	game.ID = utils.MD5(game.Name)

	if err := it.Find(game); err == nil {

		return model.ErrExisted
	}

	return it.db.Store(game)
}
