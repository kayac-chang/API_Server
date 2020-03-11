package usecase

import (
	"api/game/repo"
	"api/model"
)

type usercase struct {
	db    repo.Repository
	cache repo.Repository
}

type Usecase interface {
	Find(game *model.Game) error
	Store(game *model.Game) error
}

func New(db, cache repo.Repository) Usecase {

	return &usercase{db, cache}
}
