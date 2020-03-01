package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KayacChang/API_Server/games/entity"
	"github.com/KayacChang/API_Server/system"
	"github.com/KayacChang/API_Server/system/db"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/utils"
)

type Repo struct {
	*db.DB
}

type Querys struct {
	insert     string
	findByName string
	findAll    string
}

var querys Querys

func New(cfg env.PostgresConfig) *Repo {

	querys = Querys{
		insert:     utils.Parse("games/sql/insert_one.sql"),
		findByName: utils.Parse("games/sql/find_by_name.sql"),
		findAll:    utils.Parse("games/sql/find_all.sql"),
	}

	return &Repo{
		db.New(cfg.ToURL()),
	}
}

func (db *Repo) Insert(ctx context.Context, game *entity.Game) error {

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}

	tx := db.MustBeginTx(ctx, opt)

	// === Check If Record Exist ===
	err := tx.Get(game, querys.findByName, game.Name)

	if err == nil {
		tx.Rollback()

		return system.ErrConflict
	}

	// === Insert ===
	_, err = tx.NamedExec(querys.insert, game)

	if err != nil {
		fmt.Println("this...1")
		tx.Rollback()

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(game, querys.findByName, game.Name)

	if err != nil {
		fmt.Println("this...2")
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (db *Repo) Find(ctx context.Context, games *[]entity.Game) error {

	return db.Select(games, querys.findAll)
}
