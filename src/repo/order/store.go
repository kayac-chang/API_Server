package order

import (
	"api/model"
	"time"

	errs "github.com/pkg/errors"
)

func (it *Repo) Store(dest string, order *model.Order) error {

	switch dest {

	case "DB":
		return it.storeDB(it.sql.insert, order)

	case "Cache":
		return it.storeCache(order)
	}

	return nil
}

func (it *Repo) storeCache(order *model.Order) error {

	it.cache.Set(order.ID, order, 1*time.Hour)

	// fmt.Printf("%s\n", json.Jsonify(it.cache.Items()))

	return nil
}

func (it *Repo) Replace(order *model.Order) error {

	return it.storeDB(it.sql.updateByID, order)
}

func (it *Repo) storeDB(sql string, order *model.Order) error {

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
