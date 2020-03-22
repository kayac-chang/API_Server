package game

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	"api/model"
	"api/model/request"
	"api/model/response"
	game "api/usecase/game"

	"net/http"
)

type Handler struct {
	*server.Server
	env  *env.Env
	game *game.Usecase
}

func New(s *server.Server, e *env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		e,
		game.New(e, db, c),
	}

}

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	req := r.Context().Value(request.JSON).(map[string]string)

	// == Parse Payload ==
	game, err := it.game.Store(req["name"], req["href"], req["category"])
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusInternalServerError,

			Error: model.Error{
				Name:    "Game Create Error",
				Message: err.Error(),
			},
		})

		return
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusCreated,

		Data: game,
	})
}

func (it *Handler) GET(w http.ResponseWriter, r *http.Request) {

	games, err := it.game.FindAll()
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusInternalServerError,

			Error: model.Error{
				Name:    "Server Error",
				Message: err.Error(),
			},
		})

		return
	}

	// == Send Response ==
	links := []response.Link{}
	for _, game := range games {

		links = append(links, response.Link{
			Relation: game.Name,
			Method:   "GET",
			Href:     game.Href,
		})
	}

	href := it.env.Service.Domain + "/" + it.env.API.Version + "/games"
	links = append(links, response.Link{
		Relation: "self",
		Method:   "GET",
		Href:     href,
	})

	res := map[string]interface{}{
		"links": links,
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusCreated,

		Data: res,
	})
}
