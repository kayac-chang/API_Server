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
	insert          string
	findByID        string
	findByUserAndPW string
}

var querys Querys

func New(cfg env.PostgresConfig) *Repo {

	querys = Querys{
		insert:   utils.Parse("users/sql/insert_one.sql"),
		findByID: utils.Parse("users/sql/find_by_id.sql"),
	}

	return &Repo{
		db.New(cfg.ToURL()),
	}
}

func (db *Repo) Insert(ctx context.Context, user *model.User) error {

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	tx := db.MustBeginTx(ctx, opt)

	// === Check If Record Exist ===
	err := tx.Get(user, querys.findByID, user.ID)

	if err == nil {
		tx.Rollback()

		return system.ErrConflict
	}

	// === Insert ===
	_, err = tx.NamedExec(querys.insert, user)

	if err != nil {
		tx.Rollback()

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(user, querys.findByID, user.ID)

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (db *Repo) FindByID(ctx context.Context, user *model.User) error {

	return db.GetContext(
		ctx,
		user,
		querys.findByID,
		user.ID,
	)
}
