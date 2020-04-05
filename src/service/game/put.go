package game

import (
	"api/model"
	"api/model/response"
	"encoding/json"
	"net/http"
)

func (it *Handler) PUT(w http.ResponseWriter, r *http.Request) {

	if err := it.authenticate(r); err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusUnauthorized,

			Error: response.Error{
				Name:    "Unauthorized",
				Message: err.Error(),
			},
		})

		return
	}

	// == Parse Payload ==
	req := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusBadRequest,

			Error: response.Error{
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		})

		return
	}

	name := it.URLParam(r, "name")

	// == Update Game ==
	game, err := it.game.Update(name, req)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusInternalServerError,

			Error: response.Error{
				Name:    "Game Update Error",
				Message: err.Error(),
			},
		})

		return
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusAccepted,

		Data: game,
	})
}
