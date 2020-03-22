package game

import (
	"api/model"

	errs "github.com/pkg/errors"
)

func (it *Repo) Store(game *model.Game) error {

	tx, err := it.db.Beginx()
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.insert, game)
	if err != nil {
		tx.Rollback()

		if it.db.IsConstraintErr(err) {

			return errs.WithMessage(model.ErrDBConstraint, err.Error())
		}

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(game, it.sql.findByID, game.ID)
	if err != nil {
		tx.Rollback()

		return err
	}

	// === Save to Cache ===
	defer it.storeCache(game)

	return tx.Commit()
}

func (it *Repo) storeCache(games ...*model.Game) {

	for _, game := range games {

		it.cache.SetDefault(game.ID, *game)
	}
}
