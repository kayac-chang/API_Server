package game

import "api/model"

const prefix = "games:"

func (it Repo) Find(gamename string) (*model.Game, error) {

	game := model.Game{}

	if err := it.db.Get(prefix+gamename, &game); err != nil {
		return nil, err
	}

	return &game, nil
}
