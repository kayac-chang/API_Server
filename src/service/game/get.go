package game

import (
	"api/model"
	"api/model/response"

	"net/http"
)

func (it *Handler) GET(w http.ResponseWriter, r *http.Request) {

	name := it.URLParam(r, "name")
	game, err := it.usecase.Find(name)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusNotFound,

			Error: model.Error{
				Name:    "Game not found",
				Message: err.Error(),
			},
		})

		return
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusOK,

		Data: game,
	})
}

func (it *Handler) GET_ALL(w http.ResponseWriter, r *http.Request) {

	games, err := it.usecase.FindAll()
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
		Code: http.StatusOK,

		Data: res,
	})
}
