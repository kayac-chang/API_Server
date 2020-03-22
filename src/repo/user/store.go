package user

import (
	"api/model"
	"log"

	errs "github.com/pkg/errors"
)

func (it *Repo) Store(dest string, user *model.User) error {

	switch dest {

	case "DB":
		return it.storeDB(user)

	case "Cache":
		return it.storeCache(user)

	default:
		log.Fatalf("Unsupport Storage: %s\n", dest)
	}

	return nil
}

func (it *Repo) storeDB(user *model.User) error {

	tx, err := it.db.Beginx()
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.insert, user)
	if err != nil {
		tx.Rollback()

		if it.db.IsConstraintErr(err) {

			return errs.WithMessage(model.ErrDBConstraint, err.Error())
		}

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(user, it.sql.findByID, user.ID)
	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (it *Repo) storeCache(user *model.User) error {

	it.cache.SetDefault(user.ID, *user)
	it.cache.SetDefault(user.Token, *user)

	return nil
}
