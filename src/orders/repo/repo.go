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
	insert   string
	findByID string
}

func New() *Repo {

	sql := &querys{
		insert:   utils.ParseFile("orders/sql/insert_one.sql"),
		findByID: utils.ParseFile("orders/sql/find_by_id.sql"),
	}

	return &Repo{
		db.New(env.Postgres()),

		sql,
	}
}

func (db *Repo) Insert(ctx context.Context, order *model.Order) error {

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	tx := db.MustBeginTx(ctx, opt)

	// === Check If Record Exist ===
	err := tx.Get(order, db.sql.findByID, order.ID)

	if err == nil {
		tx.Rollback()

		return system.ErrConflict
	}

	// === Insert ===
	_, err = tx.NamedExec(db.sql.insert, order)

	if err != nil {
		tx.Rollback()

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(order, db.sql.findByID, order.ID)

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
