package games

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KayacChang/API_Server/postgres"
	"github.com/julienschmidt/httprouter"
)

const (
	selectAll  = "SELECT id, name, href FROM games"
	selectByID = selectAll + " WHERE id = $1"
	insertOne  = "INSERT INTO games (name, href) VALUES ($1, $2) RETURNING id, name, href"
)

func Serve(r *httprouter.Router, db postgres.DB) *httprouter.Router {

	r.GET("/games",
		read(db, selectAll))

	r.GET("/games/:id",
		read(db, selectByID))

	r.POST("/games",
		create(db, insertOne))

	return r
}

func read(db postgres.DB, query string) httprouter.Handle {

	stmt := db.MustPreparex(query)

	exec := func(games *[]Game, id string) error {

		if id == "" {
			return stmt.Select(games)
		}

		return stmt.Select(games, id)
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		games := []Game{}

		//	Transation
		err := exec(&games, p.ByName("id"))

		if err != nil {
			// TODO
			log.Fatal(err)
		}

		//	Response
		w.WriteHeader(http.StatusOK)

		sendJSON(w, games)
	}
}

func create(db postgres.DB, query string) httprouter.Handle {

	stmt := db.MustPreparex(query)

	exec := func(g *Game) error {

		return db.Transact(func(tx postgres.Tx) error {

			return stmt.QueryRowx(g.Name, g.Href).StructScan(g)
		})
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		game := Game{}

		//	Parse
		err := json.NewDecoder(r.Body).Decode(&game)

		if err != nil {
			//TODO
			log.Fatal(err)
		}

		//	Transation
		err = exec(&game)

		if err != nil {
			//TODO
			log.Fatal(err)
		}

		//	Log
		log.Printf(`
		From %s:
		Insert one record into [ db.postgre.games ]
		record => %#v
		`, r.URL.Path, game)

		//	Response
		w.WriteHeader(http.StatusCreated)

		sendJSON(w, game)
	}
}
