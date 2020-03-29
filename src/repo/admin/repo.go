package admin

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
	insert   string
	findByID string
}

func New(db *postgres.DB, c *cache.Cache) *Repo {

	return &Repo{
		db:    db,
		cache: c,
		sql: querys{
			insert:   parse("insert_one.sql"),
			findByID: parse("find_by_id.sql"),
		},
	}
}

// === Private ===
func parse(file string) string {

	folder := "sql/admin/"

	return utils.ParseFile(folder + file)
}
