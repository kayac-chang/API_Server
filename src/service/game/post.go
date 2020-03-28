package game

import (
	"api/model"
	"api/model/response"

	"encoding/json"
	"net/http"
)

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	// == Parse Payload ==
	req := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusBadRequest,

			Error: model.Error{
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		})

		return
	}

	// == Store Game ==
	game, err := it.usecase.Store(req["name"], req["href"], req["category"])
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
