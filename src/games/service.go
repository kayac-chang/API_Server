package games

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func Serve(r *httprouter.Router, db *sqlx.DB) *httprouter.Router {

	r.GET("/games", get(db))

	return r
}

func get(db *sqlx.DB) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		games := []game{}

		err := db.Select(&games, "SELECT * FROM games")

		if err != nil {
			log.Fatalln(err)
		}

		//	Response to Client
		w.WriteHeader(http.StatusOK)

		sendJSON(w, games)
	}
}
