package postgres

import (
	"api/framework/postgres"
	"api/game/repo"
	"api/model"
	"api/utils"

	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	errs "github.com/pkg/errors"
)

type repository struct {
	db      *postgres.DB
	timeout time.Duration
	sql     querys
}

type querys struct {
	insert   string
	findByID string
	findAll  string
}

func New(url string, timeout int) repo.Repository {

	return &repository{
		db:      postgres.New(url, timeout),
		timeout: time.Duration(timeout) * time.Second,

		sql: querys{
			insert:   utils.ParseFile("sql/game/insert_one.sql"),
			findByID: utils.ParseFile("sql/game/find_by_id.sql"),
			findAll:  utils.ParseFile("sql/game/find_all.sql"),
		},
	}
}

func (it *repository) withTimeout() (context.Context, context.CancelFunc) {

	return context.WithTimeout(context.Background(), it.timeout)
}

func (it *repository) findByID(ctx context.Context, game *model.Game) error {

	err := it.db.GetContext(ctx, game, it.sql.findByID, game.ID)

	if err == sql.ErrNoRows {
		return model.ErrUserNotFound
	}

	return err
}

func (it *repository) FindAll(games *[]model.Game) error {

	err := it.db.Select(games, it.sql.findAll)

	if err == sql.ErrNoRows {
		return model.ErrUserNotFound
	}

	return err
}

func (it *repository) FindBy(key string, game *model.Game) error {

	ctx, cancel := it.withTimeout()
	defer cancel()

	switch key {
	case "ID":
		return it.findByID(ctx, game)
	}

	return model.ErrUserNotFound
}

func (it *repository) Store(game *model.Game) error {

	ctx, cancel := it.withTimeout()
	defer cancel()

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}
	tx, err := it.db.BeginTxx(ctx, opt)
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.insert, game)
	if err != nil {
		tx.Rollback()

		if it.db.IsConstraintErr(err) {

			return errs.WithMessage(model.ErrDBConstraint, err.Error())
		}

		return err
	}

	// === Get Inserted Data ===
	err = tx.Get(game, it.sql.findByID, game.ID)
	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
