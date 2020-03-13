package postgres

import (
	"api/framework/postgres"
	"api/model"
	"api/order/repo"
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
}

func New(url string, timeout int) repo.Repository {

	return &repository{
		db:      postgres.New(url, timeout),
		timeout: time.Duration(timeout) * time.Second,

		sql: querys{
			insert:   utils.ParseFile("sql/order/insert_one.sql"),
			findByID: utils.ParseFile("sql/order/find_by_id.sql"),
		},
	}
}

func (it *repository) withTimeout() (context.Context, context.CancelFunc) {

	return context.WithTimeout(context.Background(), it.timeout)
}

func (it *repository) findByID(ctx context.Context, order *model.Order) error {

	err := it.db.GetContext(ctx, order, it.sql.findByID, order.ID)

	if err == sql.ErrNoRows {
		return model.ErrUserNotFound
	}

	return err
}

func (it *repository) FindBy(key string, order *model.Order) error {

	ctx, cancel := it.withTimeout()
	defer cancel()

	switch key {
	case "ID":
		return it.findByID(ctx, order)
	}

	return model.ErrUserNotFound
}

func (it *repository) Store(order *model.Order) error {

	ctx, cancel := it.withTimeout()
	defer cancel()

	opt := &sql.TxOptions{Isolation: sql.LevelSerializable}
	tx, err := it.db.BeginTxx(ctx, opt)
	if err != nil {
		return err
	}

	// === Insert ===
	_, err = tx.NamedExec(it.sql.insert, order)
	if err != nil {
		tx.Rollback()

		if it.db.IsConstraintErr(err) {

			return errs.WithMessage(model.ErrDBConstraint, err.Error())
		}

		return err
	}

	// === Get Inserted Data ===
	err = tx.Unsafe().Get(order, it.sql.findByID, order.ID)
	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
