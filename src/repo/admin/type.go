package admin

import (
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"database/sql"
	"fmt"
)

const table = "admins"

// Repo type for persistence layer logic
type Repo struct {
	redis redis.Redis
	db    postgres.DB
}

// New create repo for tokens associate table
func New(redis redis.Redis, db postgres.DB) Repo {
	return Repo{redis, db}
}

// Store store admin in redis
func (it Repo) Store(admin *model.Admin) error {

	query := fmt.Sprintf(`
		INSERT INTO %s 
			(admin_id, email, username, password, created_at, updated_at) 
		VALUES 
			(:admin_id, :email, :username, :password, :created_at, :updated_at)
	`, table)

	_, err := it.db.NamedExec(query, admin)

	return err
}

// FindByID find admin by specify id
func (it Repo) FindByID(id string) (*model.Admin, error) {

	admin := model.Admin{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE admin_id=$1", table)

	if err := it.db.Ping(); err != nil {
		return nil, err
	}

	if err := it.db.Get(&admin, query, id); err != nil {

		if err == sql.ErrNoRows {

			return nil, model.ErrNotFound
		}

		return nil, err
	}

	return &admin, nil
}
