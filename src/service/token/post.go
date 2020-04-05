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

	checkHeader := func() error {
		session := r.Header.Get("session")
		contentType := r.Header.Get("Content-Type")

		return it.token.CheckHeader(session, contentType)
	}

	payload := func() (map[string]string, error) {

		req := map[string]string{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

			return nil, &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Parse JSON #2",
				Message: errs.WithMessage(err, "Error occured when parsing payload").Error(),
			}
		}

		// Check Payload
		gamename := req["game"]
		username := req["username"]

		if err := it.token.CheckPayload(gamename, username); err != nil {
			return nil, err
		}

		return req, nil
	}

	business := func(req map[string]string) ([]interface{}, error) {
		// Registration
		registration := utils.Promisefy(func() (interface{}, error) {

			username := req["username"]
			session := r.Header.Get("session")

			return it.token.Regist(username, session)
		})

		// Get Game Link
		getGameLink := utils.Promisefy(func() (interface{}, error) {

			name := req["game"]

			return it.gameCase.FindByName(name)
		})

		return utils.WaitAll(registration, getGameLink)
	}

	genResponse := func(game *model.Game, token *model.Token) interface{} {
		TAG := "/tokens"

		href := it.getHref(TAG)

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

	main := func() interface{} {
		if err := checkHeader(); err != nil {
			return err
		}

		req, err := payload()
		if err != nil {
			return err
		}

		results, err := business(req)
		if err != nil {
			return err
		}

		token := results[0].(*model.Token)
		game := results[1].(*model.Game)

		return genResponse(game, token)
	}

	it.Send(w, main())
}
