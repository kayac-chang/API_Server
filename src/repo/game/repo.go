package game

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
	insert      string
	findByID    string
	findWhereID string
	findAll     string
}

func New(db *postgres.DB, c *cache.Cache) *Repo {

	return &Repo{
		db:    db,
		cache: c,
		sql: querys{
			insert:      utils.ParseFile("sql/game/insert_one.sql"),
			findByID:    utils.ParseFile("sql/game/find_by_id.sql"),
			findWhereID: utils.ParseFile("sql/game/find_where_id.sql"),
			findAll:     utils.ParseFile("sql/game/find_all.sql"),
		},
	}
}
