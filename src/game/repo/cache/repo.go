package cache

import (
	"api/framework/cache"
	"api/game/repo"
	"api/model"
)

type repository struct {
	*cache.Cache
}

func New() repo.Repository {

	return &repository{cache.Get()}
}

func (it *repository) findByID(game *model.Game) error {

	if _game, found := it.Get(game.ID); found {

		if _game, ok := _game.(*model.Game); ok {

			game = _game

			return nil
		}
	}

	return model.ErrGameNotFound
}

func (it *repository) FindBy(key string, game *model.Game) error {

	switch key {
	case "ID":
		return it.findByID(game)
	}

	return model.ErrGameNotFound
}

func (it *repository) Store(game *model.Game) error {

	it.SetDefault(game.ID, game)

	// log.Printf("repository.cache.Store\n%s\n", json.Jsonify(storage.Items()))

	return nil
}
