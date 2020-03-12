package game

import (
	"api/env"
	"api/framework/server"
	"api/game/repo/cache"
	"api/game/repo/postgres"
	"api/game/usecase"
	"api/model"
	"api/model/response"
	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
)

type handler struct {
	*server.Server

	usecase usecase.Usecase
}

func New(e *env.Env) {

	s := server.New()

	c := cache.New()
	db := postgres.New(e.Postgres.ToURL(), 30)

	it := handler{
		s,
		usecase.New(db, c),
	}

	s.Post("/games", it.POST)
	s.Get("/games", it.GET)

	http.ListenAndServe(":8000", s)
}

func (it *handler) POST(w http.ResponseWriter, r *http.Request) {

	// == Parse Payload ==
	game := model.Game{}

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {

		log.Errorf("%s\n", err.Error())

		res := response.JSON{

			Code: http.StatusBadRequest,

			Error: model.Error{
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		}
		it.SendJSON(w, res)

		return
	}

	err = it.usecase.Store(&game)
	if err != nil {

		log.Errorf("%s\n", err.Error())

		var res response.JSON

		switch err {

		case model.ErrExisted:
			res = response.JSON{
				Code: http.StatusConflict,

				Error: model.Error{
					Name:    "Existed",
					Message: err.Error(),
				},
			}

		default:
			res = response.JSON{
				Code: http.StatusInternalServerError,

				Error: model.Error{
					Name:    "Server Error",
					Message: err.Error(),
				},
			}
		}

		it.SendJSON(w, res)

		return
	}

	// == Send Response ==
	res := response.JSON{

		Code: http.StatusCreated,

		Data: game,
	}
	it.SendJSON(w, res)
}

func (it *handler) GET(w http.ResponseWriter, r *http.Request) {

	games := []model.Game{}

	err := it.usecase.FindAll(&games)
	if err != nil {

		log.Errorf("%s\n", err.Error())

		var res response.JSON

		switch err {

		default:
			res = response.JSON{
				Code: http.StatusInternalServerError,

				Error: model.Error{
					Name:    "Server Error",
					Message: err.Error(),
				},
			}
		}

		it.SendJSON(w, res)

		return
	}

	// == Send Response ==

	data := map[string]interface{}{}

	links := []model.Link{}

	// self Links
	links = append(links, model.Link{
		Relation: "self",
		Method:   "GET",
		Href:     "/games",
	})

	for _, game := range games {

		link := model.Link{
			Relation: game.Name,
			Method:   "GET",
			Href:     game.Href,
		}

		links = append(links, link)
	}

	data["links"] = &links

	res := response.JSON{

		Code: http.StatusCreated,

		Data: data,
	}
	it.SendJSON(w, res)
}
