package token

import (
	"api/model"
	"api/model/response"
	"api/utils"

	"encoding/json"
	"net/http"

	errs "github.com/pkg/errors"
)

// POST POST /tokens
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
		token, game, err := it.business(req)
		if err != nil {
			return err
		}

		return it.genResponse(token, game)
	}

	it.Send(w, main())
}

func (it *Handler) checkHeader(r *http.Request) error {

	session := r.Header.Get("session")
	if err := utils.CheckSession(session); err != nil {
		return err
	}

	contentType := r.Header.Get("Content-Type")
	if err := utils.CheckContentType(contentType, "application/json"); err != nil {
		return err
	}

	return nil
}

func (it *Handler) parseJSON(r *http.Request) (map[string]interface{}, error) {

	req := map[string]interface{}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(err, "Error occured when parsing payload").Error(),
		}
	}

	return req, nil
}

func (it *Handler) checkPayload(req map[string]interface{}) error {

	return utils.CheckPayload(req, "game", "username")
}

func (it *Handler) business(req map[string]interface{}) (*model.Token, *model.Game, error) {

	username := req["username"].(string)
	session := req["session"].(string)
	name := req["game"].(string)

	registration := utils.Promisefy(func() (interface{}, error) {

		token, err := it.usecase.Regist(username, session)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Registration #4"

			return nil, err
		}

		return token, nil
	})

	getGameLink := utils.Promisefy(func() (interface{}, error) {

		game, err := it.usecase.FindGameByName(name)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Get Game Link #5"

			return nil, err
		}

		return game, nil
	})

	res, err := utils.WaitAll(registration, getGameLink)
	if err != nil {
		return nil, nil, err
	}

	token := res[0].(*model.Token)
	game := res[1].(*model.Game)

	return token, game, nil
}

func (it *Handler) genResponse(token *model.Token, game *model.Game) interface{} {

	href := it.getHref("/tokens")

	gameHref := game.Href + "?" + "access_token=" + token.AccessToken

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
