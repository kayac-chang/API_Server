package repo

import (
	"context"
	"database/sql"

	"server/model"
	"server/system"
	"server/system/db"
	"server/system/env"
	"server/utils"
)

type Repo struct {
	*db.DB

	sql *querys
}

type querys struct {
	insert      string
	findByID    string
	findByName  string
	updateToken string
}

func New() *Repo {

	sql := &querys{
		insert:      utils.ParseFile("users/sql/insert_one.sql"),
		findByID:    utils.ParseFile("users/sql/find_by_id.sql"),
		findByName:  utils.ParseFile("users/sql/find_by_name.sql"),
		updateToken: utils.ParseFile("users/sql/update_token.sql"),
	}

	return &Repo{
		db.New(env.Postgres()),

		sql,
	}
}

func (it *Repo) Insert(ctx context.Context, user *model.User) error {

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	tx := it.MustBeginTx(ctx, opt)

	// === Check If Record Exist ===
	err := tx.Get(user, it.sql.findByID, user.ID)

	if err == nil {
		tx.Rollback()

		return system.ErrConflict
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.insert, user)

	if err != nil {
		tx.Rollback()

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

func (it *Repo) Update(ctx context.Context, user *model.User) error {

	// === Update ===
	_, err := it.NamedExec(it.sql.updateToken, user)

	return err
}

func (it *Repo) FindByName(ctx context.Context, user *model.User) error {

	return it.GetContext(
		ctx,
		user,
		it.sql.findByName,
		user.Username,
	)
}
