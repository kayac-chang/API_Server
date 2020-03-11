package repo

import "api/model"

type Repository interface {
	Store(user *model.Game) error
	FindBy(key string, user *model.Game) (*model.Game, error)
}
