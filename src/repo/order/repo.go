package order

import (
	"api/framework/cache"
	"api/framework/postgres"
	"api/utils"
)

type Repo struct {
	db    *postgres.DB
	cache *cache.Cache
	sql   querys
}

type querys struct {
	insert     string
	findByID   string
	updateByID string
}

func New(db *postgres.DB, c *cache.Cache) *Repo {

	return &Repo{
		db:    db,
		cache: c,
		sql: querys{
			insert:     utils.ParseFile("sql/order/insert_one.sql"),
			findByID:   utils.ParseFile("sql/order/find_by_id.sql"),
			updateByID: utils.ParseFile("sql/order/update_by_id.sql"),
		},
	}
}
