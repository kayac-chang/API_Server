package game

import (
	"api/model"
	"api/model/response"
	"api/utils"
	"net/http"
)

// POST create new game
func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

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

		// == Check Content-Type #2 ==
		if r.Header.Get("Content-Type") != "application/json" {

			return &model.Error{
				Code:    http.StatusBadRequest,
				Name:    "Check Content-Type #2",
				Message: "Content-Type must be application/json",
			}
		}

		// == Parse JSON #3 ==
		req, err := utils.ParseJSON(r.Body)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Parse JSON #3"

			return err
		}

		// == Check Request Payload #4 ==
		if err := it.checkPayload(req); err != nil {
			err := err.(*model.Error)

			err.Name = "Check Request Payload #4"

			return err
		}

		// == Create Game Exist #5 ==
		name := req["name"].(string)
		_, err = it.usecase.FindByName(name)
		if err == nil {

			return &model.Error{
				Code:    http.StatusConflict,
				Name:    "Create Game Exist #5",
				Message: "Request game already existed",
			}
		}
		if err != model.ErrNotFound {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Create Game Exist #5",
				Message: err.Error(),
			}
		}

		// == Create Game #6 ==
		href := req["href"].(string)
		category := req["category"].(string)
		game, err := it.usecase.Store(name, href, category)
		if err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Create Game #6",
				Message: err.Error(),
			}
		}

		return response.JSON{
			Code: http.StatusCreated,

			Data: game,
		}
	}

	it.Send(w, main())
}

func (it Handler) checkPayload(req map[string]interface{}) error {

	err := utils.CheckPayload(req, "name", "href", "category")
	if err != nil {

		return err
	}

	href := req["href"].(string)
	if !utils.IsValidURL(href) {

		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "request href is not valid url",
		}
	}

	return nil
}
