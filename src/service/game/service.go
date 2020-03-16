package game

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	"api/model"
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

// func (it *handler) POST(w http.ResponseWriter, r *http.Request) {

// 	// == Parse Payload ==
// 	game := model.Game{}

// 	err := json.NewDecoder(r.Body).Decode(&game)
// 	if err != nil {

// 		log.Errorf("%s\n", err.Error())

// 		res := response.JSON{

// 			Code: http.StatusBadRequest,

// 			Error: model.Error{
// 				Name:    "Unexpect Payload",
// 				Message: model.ErrUnexpectPayload.Error(),
// 			},
// 		}
// 		it.SendJSON(w, res)

// 		return
// 	}

// 	err = it.usecase.Store(&game)
// 	if err != nil {

// 		log.Errorf("%s\n", err.Error())

// 		var res response.JSON

// 		switch err {

// 		case model.ErrExisted:
// 			res = response.JSON{
// 				Code: http.StatusConflict,

// 				Error: model.Error{
// 					Name:    "Existed",
// 					Message: err.Error(),
// 				},
// 			}

// 		default:
// 			res = response.JSON{
// 				Code: http.StatusInternalServerError,

// 				Error: model.Error{
// 					Name:    "Server Error",
// 					Message: err.Error(),
// 				},
// 			}
// 		}

// 		it.SendJSON(w, res)

// 		return
// 	}

// 	// == Send Response ==
// 	res := response.JSON{

// 		Code: http.StatusCreated,

// 		Data: game,
// 	}
// 	it.SendJSON(w, res)
// }

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
