package token

import (
	"api/model"
	"api/model/response"
	"api/utils"

	"encoding/json"
	"net/http"

	errs "github.com/pkg/errors"
)

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		if err := it.checkHeader(r); err != nil {
			err := err.(*model.Error)

			err.Name = "Check Request #1"

			return err
		}

		req, err := it.parseJSON(r)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Parse JSON #2"

			return err
		}

		if err := it.checkPayload(req); err != nil {
			err := err.(*model.Error)

			err.Name = "Check Request Payload #3"

			return err
		}

		req["session"] = r.Header.Get("session")
		results, err := it.business(req)
		if err != nil {
			return err
		}

		token := results[0].(*model.Token)
		game := results[1].(*model.Game)

		return it.genResponse(game, token)
	}

	it.Send(w, main())
}

func (it *Handler) checkHeader(r *http.Request) error {
	session := r.Header.Get("session")
	contentType := r.Header.Get("Content-Type")

	return it.token.CheckHeader(session, contentType)
}

func (it *Handler) parseJSON(r *http.Request) (map[string]string, error) {

	req := map[string]string{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(err, "Error occured when parsing payload").Error(),
		}
	}

	return req, nil
}

func (it *Handler) checkPayload(req map[string]string) error {

	gamename := req["game"]
	username := req["username"]

	return it.token.CheckPayload(gamename, username)
}

func (it *Handler) business(req map[string]string) ([]interface{}, error) {

	registration := utils.Promisefy(func() (interface{}, error) {

		username := req["username"]
		session := req["session"]

		token, err := it.token.Regist(username, session)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Registration #4"

			return nil, err
		}

		return token, nil
	})

	getGameLink := utils.Promisefy(func() (interface{}, error) {

		name := req["game"]

		game, err := it.game.FindByName(name)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Get Game Link #5"

			return nil, err
		}

		return game, nil
	})

	return utils.WaitAll(registration, getGameLink)
}

func (it *Handler) genResponse(game *model.Game, token *model.Token) interface{} {

	href := it.getHref("/tokens")

	return response.JSON{
		Code: http.StatusCreated,

		Data: map[string]interface{}{
			"token": token,
			"links": [...]response.Link{
				{Relation: "access", Method: "GET", Href: game.Href},
				{Relation: "reauthorize", Method: "POST", Href: href},
			},
		},
	}
}
