package admin

import (
	"api/model"
	"fmt"

	errs "github.com/pkg/errors"
)

func (it *Repo) Store(dest string, admin *model.Admin) error {

	switch dest {

	case "DB":
		return it.storeDB(admin)

	case "Cache":
		it.storeCache(admin)

		return nil
	}

	return fmt.Errorf("Storage %s not support\n", dest)
}

func (it *Repo) storeCache(admin *model.Admin) {

	it.cache.SetDefault(admin.ID, *admin)
	it.cache.SetDefault(admin.Token, *admin)
}

func (it *Repo) storeDB(admin *model.Admin) error {

	tx, err := it.db.Beginx()
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.insert, admin)
	if err != nil {
		tx.Rollback()

		if it.db.IsConstraintErr(err) {

			return errs.WithMessage(model.ErrDBConstraint, err.Error())
		}

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(admin, it.sql.findByID, admin.ID)
	if err != nil {
		tx.Rollback()

		return err
	}

	// === Save to Cache ===
	// TODO

	return tx.Commit()
}
