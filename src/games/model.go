package games

import (
	"github.com/KayacChang/API_Server/pg"
)

type Game struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Href string `json:"href" db:"href"`
}

type Table struct {
	selectAll func(*[]Game) error

	selectByID func(*Game, string) error

	insertOne func(*Game) error
}

func TableFrom(db pg.DB) Table {

	querys := db.Prepare(map[string]string{
		"selectAll":  "SELECT id, name, href FROM games",
		"selectByID": "SELECT id, name, href FROM games WHERE id = $1",
		"insertOne":  "INSERT INTO games (name, href) VALUES ($1, $2) RETURNING id, name, href",
	})

	return Table{

		selectAll: func(res *[]Game) error {
			return querys["selectAll"].Select(res)
		},

		selectByID: func(res *Game, id string) error {
			return querys["selectByID"].Get(res, id)
		},

		insertOne: func(g *Game) error {

			return db.Transact(func(tx pg.Tx) error {

				return querys["insertOne"].QueryRowx(g.Name, g.Href).StructScan(g)
			})
		},
	}
}
