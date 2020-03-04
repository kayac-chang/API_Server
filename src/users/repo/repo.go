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
	insert     string
	findByID   string
	findByName string
}

func New() *Repo {

	sql := &querys{
		insert:     utils.ParseFile("users/sql/insert_one.sql"),
		findByID:   utils.ParseFile("users/sql/find_by_id.sql"),
		findByName: utils.ParseFile("users/sql/find_by_name.sql"),
	}

	return &Repo{
		db.New(env.Postgres()),

		sql,
	}
}

func (db *Repo) Insert(ctx context.Context, user *model.User) error {

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	tx := db.MustBeginTx(ctx, opt)

	// === Check If Record Exist ===
	err := tx.Get(user, db.sql.findByID, user.ID)

	if err == nil {
		tx.Rollback()

		return system.ErrConflict
	}

	// === Insert ===
	_, err = tx.NamedExec(db.sql.insert, user)

	if err != nil {
		tx.Rollback()

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(user, db.sql.findByID, user.ID)

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (db *Repo) FindByName(ctx context.Context, user *model.User) error {

	return db.GetContext(
		ctx,
		user,
		db.sql.findByName,
		user.Username,
	)
}
