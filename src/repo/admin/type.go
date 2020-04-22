package admin

import (
	"api/framework/redis"
	"api/model"
)

const prefix = "admins:"

// Repo type for persistence layer logic
type Repo struct {
	db redis.Redis
}

// New create repo for tokens associate table
func New(db redis.Redis) Repo {
	return Repo{db}
}

// Store store associate table with key by admin
func (it Repo) Store(admin *model.Admin) error {

	key := prefix + admin.ID

	return it.db.Set(key, admin)
}
