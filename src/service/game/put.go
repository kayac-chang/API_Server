package game

import (
	"api/model"
	"api/model/response"
	"encoding/json"
	"net/http"
)

func (it *Handler) PUT(w http.ResponseWriter, r *http.Request) {

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

	name := it.URLParam(r, "name")

	// == Update Game ==
	game, err := it.usecase.Update(name, req)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusInternalServerError,

			Error: model.Error{
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
