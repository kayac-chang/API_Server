package game

import (
	"api/model"

	errs "github.com/pkg/errors"
)

func (it *Repo) FindByID(id string) (*model.Game, error) {

	game := model.Game{}

	// Find in Cache
	// TODO

	// Find in DB
	if err := it.db.Get(&game, it.sql.findByID, id); err != nil {
		return nil, errs.WithMessagef(err, id+" not existed")
	}

	// === Save to Cache ===
	// TODO

	return &game, nil
}

func (it *Repo) FindByName(name string) (*model.Game, error) {

	game := model.Game{}

	// Find in Cache
	// TODO

	// Find in DB
	if err := it.db.Get(&game, it.sql.findByName, name); err != nil {
		return nil, err
	}

	// === Save to Cache ===
	// TODO

	return &game, nil
}

func (it *Repo) FindAll() ([]*model.Game, error) {

	games := []*model.Game{}

	if err := it.db.Select(&games, it.sql.findAll); err != nil {
		return nil, err
	}

	// === Save to Cache ===
	// TODO

	return games, nil
}
