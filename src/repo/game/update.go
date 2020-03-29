package game

import (
	"api/model"

	errs "github.com/pkg/errors"
)

func (it *Repo) Update(game *model.Game) error {

	tx, err := it.db.Beginx()
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.updateByID, game)
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
	// TODO

	return tx.Commit()
}
