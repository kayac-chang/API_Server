package user

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	"api/model"
	"api/model/pb"
	"api/model/request"
	"api/model/response"
	game "api/usecase/game"
	user "api/usecase/user"

	"net/http"

	"github.com/go-chi/chi"
)

type Handler struct {
	*server.Server
	env      *env.Env
	userCase *user.Usecase
	gameCase *game.Usecase
}

func New(e *env.Env, db *postgres.DB, c *cache.Cache) {

	s := server.New(e)

	it := Handler{
		s,
		e,
		user.New(e, db, c),
		game.New(e, db, c),
	}

	s.Route("/"+e.API.Version, func(s chi.Router) {
		s.With(it.ParseJSON).Post("/token", it.POST)
		s.Get("/auth", it.Auth)
	})

	s.Listen(e.API.UserPort)
}

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	req := r.Context().Value(request.JSON).(map[string]string)

	// == Registration ==
	token, err := it.userCase.Regist(req["username"])
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

	game, err := it.gameCase.Find(req["game"])
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusInternalServerError,

			Error: model.Error{
				Name:    "Server Error",
				Message: err.Error(),
			},
		})

		return
	}

	href := it.env.Service.Domain + "/" + it.env.API.Version + "/token"

	res := map[string]interface{}{
		"token": token,
		"links": [...]response.Link{
			{Relation: "access", Method: "GET", Href: game.Href},
			{Relation: "reauthorize", Method: "POST", Href: href},
		},
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusCreated,

		Data: res,
	})
}

func (it *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(request.USER).(*model.User)

	// == Send Response ==
	it.Send(w, response.ProtoBuf{

		Code: http.StatusOK,

		Data: &pb.User{
			UserId:   user.ID,
			Username: user.Username,
			Balance:  user.Balance,
		},
	})
}
