package repo

import (
	"context"
	"database/sql"

	"github.com/KayacChang/API_Server/model"
	"github.com/KayacChang/API_Server/system"
	"github.com/KayacChang/API_Server/system/db"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/utils"
)

type Repo struct {
	*db.DB
}

type Querys struct {
	insert   string
	findByID string
}

var querys Querys

func New(cfg env.PostgresConfig) *Repo {

	querys = Querys{
		insert:   utils.Parse("orders/sql/insert_one.sql"),
		findByID: utils.Parse("orders/sql/find_by_id.sql"),
	}

	return &Repo{
		db.New(cfg.ToURL()),
	}
}

func (db *Repo) Insert(ctx context.Context, order *model.Order) error {

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	tx := db.MustBeginTx(ctx, opt)

	// === Check If Record Exist ===
	err := tx.Get(order, querys.findByID, order.ID)

	if err == nil {
		tx.Rollback()

		return system.ErrConflict
	}

	// === Insert ===
	_, err = tx.NamedExec(querys.insert, order)

	if err != nil {
		tx.Rollback()

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(order, querys.findByID, order.ID)

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
