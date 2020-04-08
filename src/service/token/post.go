package token

import (
	"api/model"
	"api/model/response"

	"encoding/json"
	"net/http"
)

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	// == Check Session ID ==
	session := r.Header.Get("session")
	if session == "" {

		it.Send(w, response.JSON{
			Code: http.StatusBadRequest,

			Error: model.Error{
				Name:    "Missing Session ID",
				Message: "Required header session id for authentication",
			},
		})

		return
	}

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

	// == Registration ==
	username := req["username"]
	token, err := it.userCase.Regist(username, session)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "User Registration Failed",
				Message: err.Error(),
			},
		})

		return
	}

	gamename := req["game"]
	game, err := it.gameCase.FindByName(gamename)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusNotFound,

			Error: model.Error{
				Name:    gamename + " Not Found",
				Message: err.Error(),
			},
		})

		return
	}

	gameHref := game.Href + "?" + "access_token=" + token.AccessToken
	selfHref := it.env.Service.Domain + "/" + it.env.API.Version + "/token"

	res := map[string]interface{}{
		"token": token,
		"links": [...]response.Link{
			{Relation: "access", Method: "GET", Href: gameHref},
			{Relation: "reauthorize", Method: "POST", Href: selfHref},
		},
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusCreated,

		Data: res,
	})
}
