package games

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KayacChang/API_Server/pg"
	"github.com/KayacChang/API_Server/utils"
	"github.com/julienschmidt/httprouter"
)

func Serve(r *httprouter.Router, db pg.DB) {

	tb := TableFrom(db)

	r.GET("/games", read(tb))

	r.GET("/games/:id", read(tb))

	r.POST("/games", insert(tb))
}

func read(tb Table) httprouter.Handle {

	selectBy := func(id string) *Game {

		res := Game{}

		err := tb.selectByID(&res, id)

		if err != nil {
			// TODO
			log.Fatal(err)
		}

		return &res
	}

	selectAll := func() *[]Game {

		res := []Game{}

		err := tb.selectAll(&res)

		if err != nil {
			// TODO
			log.Fatal(err)
		}

		return &res
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		send := utils.Prepare(w)

		//	Response
		if id := p.ByName("id"); id != "" {

			send(http.StatusOK, selectBy(id))

		} else {

			send(http.StatusOK, selectAll())
		}
	}
}

func insert(tb Table) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		send := utils.Prepare(w)

		res := Game{}

		//	Parse
		err := json.NewDecoder(r.Body).Decode(&res)

		if err != nil {
			//TODO
			log.Fatal(err)
		}

		//	Transation
		err = tb.insertOne(&res)

		if err != nil {
			//TODO
			log.Fatal(err)
		}

		//	Log
		log.Printf(`
		From %s:
		Insert one record into [ db.postgre.games ]
		record => %#v
		`, r.URL.Path, res)

		//	Response
		send(http.StatusCreated, res)
	}
}
