package user

import (
	"api/env"
	"api/framework/server"
	"api/model"
	"api/model/pb"
	"api/model/response"
	"api/user/repo/cache"
	"api/user/repo/postgres"
	"api/user/usecase"

	"encoding/json"
	"net/http"
	"strings"

	"github.com/fatih/structs"
)

type handler struct {
	*server.Server

	usecase usecase.Usecase
}

func New(e *env.Env) {

	s := server.New()

	c := cache.New()
	db := postgres.New(e.Postgres.ToURL(), 30)

	it := handler{
		s,
		usecase.New(db, c),
	}

	s.Post("/token", it.POST)
	s.Get("/auth", it.Auth)

	http.ListenAndServe(":8000", s)
}

func (it *handler) POST(w http.ResponseWriter, r *http.Request) {

	// == Parse Payload ==
	req := &struct {
		Game     string `json:"game"`
		Username string `json:"username"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		res := response.JSON{

			Code: http.StatusBadRequest,

			Error: model.Error{
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		}
		it.SendJSON(w, res)

		return
	}

	// == Registration ==
	user := &model.User{
		Username: req.Username,
	}
	if err := it.usecase.Regist(user); err != nil {

		res := response.JSON{

			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "User Registration Failed",
				Message: err.Error(),
			},
		}
		it.SendJSON(w, res)

		return
	}

	// == Sign Token ==
	token, err := it.usecase.Sign(user)
	if err != nil {

		res := response.JSON{

			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "User Sign Token Failed",
				Message: err.Error(),
			},
		}
		it.SendJSON(w, res)

		return
	}

	// == Send Response ==
	data := structs.Map(token)

	data["links"] = []model.Link{
		{Relation: "access", Method: "GET", Href: "https://<game_domain>"},
		{Relation: "reauthorize", Method: "POST", Href: "https://<service_domain>/v1/token"},
	}

	res := response.JSON{

		Code: http.StatusCreated,

		Data: data,
	}
	it.SendJSON(w, res)
}

func (it *handler) Auth(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	if token == "" {

		res := response.ProtoBuf{

			Code: http.StatusUnauthorized,

			Data: &pb.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Authentication Failed",
				Message: model.ErrUnauthorized.Error(),
			},
		}
		it.SendProtoBuf(w, res)

		return
	}

	token = strings.Split(token, " ")[1]

	// == Authentication ==
	user := model.User{
		Token: token,
	}

	err := it.usecase.Auth(&user)
	if err != nil {

		res := response.ProtoBuf{

			Code: http.StatusUnauthorized,

			Data: &pb.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Authentication Failed",
				Message: model.ErrUnauthorized.Error(),
			},
		}
		it.SendProtoBuf(w, res)

		return
	}

	// == Send Response ==

	res := response.ProtoBuf{

		Code: http.StatusOK,

		Data: &pb.User{
			UserId:   user.ID,
			Username: user.Username,
			Balance:  user.Balance,
		},
	}
	it.SendProtoBuf(w, res)
}
