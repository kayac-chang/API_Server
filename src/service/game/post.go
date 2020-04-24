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
		contentType := r.Header.Get("Content-Type")
		if err := utils.CheckContentType(contentType, "application/json"); err != nil {

			err.Name = "Check Content-Type #2"

			return err
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

		name := req["name"].(string)
		href := req["href"].(string)
		category := req["category"].(string)

		// == Create Game #5 ==
		game, err := it.usecase.Store(name, href, category)
		if err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Create Game #5",
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
