package repo

import "api/model"

type Repository interface {
	Store(game *model.Game) error
	FindBy(key string, game *model.Game) error
	FindAll(games *[]model.Game) error
}
