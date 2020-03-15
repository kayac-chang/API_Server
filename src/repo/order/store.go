package order

import (
	"api/model"

	errs "github.com/pkg/errors"
)

func (it *Repo) Store(order *model.Order) error {

	return it.store(it.sql.insert, order)
}

func (it *Repo) Replace(order *model.Order) error {

	return it.store(it.sql.updateByID, order)
}

func (it *Repo) store(sql string, order *model.Order) error {

	tx, err := it.db.Beginx()
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(sql, order)
	if err != nil {
		tx.Rollback()

		if it.db.IsConstraintErr(err) {

			return errs.WithMessage(model.ErrDBConstraint, err.Error())
		}

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(order, it.sql.findByID, order.ID)
	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
