package game

import (
	"api/model"
	"api/model/response"
	"net/http"
)

// GET ...
func (it Handler) GET(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Authorization #1 ==
		token := r.Header.Get("Authorization")
		if err := it.usecase.Auth(token); err != nil {

			return &model.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Check Authorization #1",
				Message: err.Error(),
			}
		}

		// == Find Game #2 ==
		id := it.URLParam(r, "id")
		game, err := it.usecase.FindByID(id)
		if err != nil {

			return &model.Error{
				Code:    http.StatusNotFound,
				Name:    "Find Game #2",
				Message: err.Error(),
			}
		}

		return response.JSON{
			Code: http.StatusOK,

			Data: game,
		}
	}

	it.Send(w, main())
}

// GETALL ...
func (it Handler) GETALL(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Find Games #1 ==
		games, err := it.usecase.FindAll()
		if err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Find Games #1",
				Message: err.Error(),
			}
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

		return response.JSON{
			Code: http.StatusOK,

			Data: res,
		}
	}

	it.Send(w, main())
}
