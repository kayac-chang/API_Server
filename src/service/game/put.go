package game

import (
	"api/model"
	"api/model/response"
	"api/utils"
	"net/http"
)

// PUT ...
func (it Handler) PUT(w http.ResponseWriter, r *http.Request) {

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

		// == Check Game ID Existed #5 ==
		gameID := it.URLParam(r, "id")
		game, err := it.usecase.FindByID(gameID)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Check Game ID Existed #5"

			return err
		}

		// == Update Game By ID #6 ==
		game.Name = req["name"].(string)
		game.Href = req["href"].(string)
		game.Category = req["category"].(string)

		if game, err = it.usecase.Update(game); err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Update Game By ID #6",
				Message: err.Error(),
			}
		}

		return response.JSON{
			Code: http.StatusAccepted,

			Data: game,
		}
	}

	it.Send(w, main())
}
