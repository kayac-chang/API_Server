package user

import (
	"api/framework/cache"
	"api/framework/postgres"
	"api/model"
	"api/utils"
)

type Repo struct {
	db    *postgres.DB
	cache *cache.Cache
	sql   querys
}

type querys struct {
	insert   string
	findByID string
}

func New(db *postgres.DB, c *cache.Cache) *Repo {

	return &Repo{
		db:    db,
		cache: c,
		sql: querys{
			insert:   utils.ParseFile("sql/user/insert_one.sql"),
			findByID: utils.ParseFile("sql/user/find_by_id.sql"),
		},
	}
}

func (it *Repo) RemoveCache(user *model.User) {

	it.cache.Delete(user.ID)
	it.cache.Delete(user.Token)
}
