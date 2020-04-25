package token

import (
	"api/model"
	"api/model/response"
	"api/utils"

	"net/http"
)

// POST POST /tokens
func (it Handler) POST(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Content-Type #1 ==
		if r.Header.Get("Content-Type") != "application/json" {

			return &model.Error{
				Code:    http.StatusBadRequest,
				Name:    "Check Content-Type #1",
				Message: "Content-Type must be application/json",
			}
		}

		// == Check Session #2 ==
		session := r.Header.Get("session")
		if session == "" {

			return &model.Error{
				Code:    http.StatusBadRequest,
				Name:    "Check Session #2",
				Message: "Request header must contain session",
			}
		}

		// == Parse JSON #3 ==
		req, err := utils.ParseJSON(r.Body)
		if err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Parse JSON #3",
				Message: err.Error(),
			}
		}

		// == Check Request Payload #4 ==
		if err := utils.CheckPayload(req, "game", "username"); err != nil {
			err := err.(*model.Error)

			err.Name = "Check Request Payload #4"

			return err
		}

		username := req["username"].(string)
		name := req["game"].(string)

		// == Registration #5 ==
		registration := utils.Promisefy(func() (interface{}, error) {

			token, err := it.usecase.Regist(username, session)
			if err != nil {
				err := err.(*model.Error)

				err.Name = "Registration #5"

				return nil, err
			}

			return token, nil
		})

		// == Get Game Link #6 ==
		getGameLink := utils.Promisefy(func() (interface{}, error) {

			game, err := it.usecase.FindGameByName(name)
			if err != nil {

				_err := &model.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Get Game Link #6",
					Message: err.Error(),
				}

				return nil, _err
			}

			return game, nil
		})

		res, err := utils.WaitAll(registration, getGameLink)
		if err != nil {
			return err
		}

		token := res[0].(*model.Token)
		game := res[1].(*model.Game)

		href := it.getHref("/tokens")

		gameHref := game.Href + "?" + "token=" + token.AccessToken

		return response.JSON{
			Code: http.StatusCreated,

			Data: map[string]interface{}{
				"token": token,
				"links": [...]response.Link{
					{Relation: "access", Method: "GET", Href: gameHref},
					{Relation: "reauthorize", Method: "POST", Href: href},
				},
			},
		}
	}

	it.Send(w, main())
}
