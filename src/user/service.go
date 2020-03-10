package user

import (
	"api/env"
	"api/model"
	"api/user/repo/cache"
	"api/user/repo/postgres"
	"api/user/usecase"

	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/fatih/structs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type handler struct {
	usecase usecase.Usecase
}

var (
	ErrUnexpectPayload = errors.New("Unexpected Request Payload")
)

func New(e *env.Env) {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	c := cache.New()
	db := postgres.New(e.Postgres.ToURL(), 30)

	it := handler{usecase: usecase.New(db, c)}

	r.Post("/token", it.POST)

	http.ListenAndServe(":8000", r)
}

func (it *handler) POST(w http.ResponseWriter, r *http.Request) {

	// == Parse Payload ==
	req := &struct {
		Game     string `json:"game"`
		Username string `json:"username"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		res := model.Response{

			Code: http.StatusBadRequest,

			Error: model.Error{
				Name:    "Unexpect Payload",
				Message: ErrUnexpectPayload.Error(),
			},
		}
		send(w, res)

		return
	}

	// == Registration ==
	user := &model.User{
		Username: req.Username,
	}
	if err := it.usecase.Regist(user); err != nil {

		res := model.Response{

			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "User Registration Failed",
				Message: err.Error(),
			},
		}
		send(w, res)

		return
	}

	// == Sign Token ==
	token, err := it.usecase.Sign(user)
	if err != nil {

		res := model.Response{

			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "User Sign Token Failed",
				Message: err.Error(),
			},
		}
		send(w, res)

		return
	}

	// == Send Response ==
	data := structs.Map(token)

	data["links"] = []model.Link{
		{Relation: "access", Method: "GET", Href: "https://<game_domain>"},
		{Relation: "reauthorize", Method: "POST", Href: "https://<service_domain>/v1/token"},
	}

	res := model.Response{

		Code: http.StatusOK,

		Data: data,
	}
	send(w, res)
}

func (it *handler) Auth(w http.ResponseWriter, r *http.Request) {

	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	if token == "" {
		res := model.Response{

			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "Authentication Failed",
				Message: model.ErrUnauthorized.Error(),
			},
		}
		send(w, res)

		return
	}

	// == Authentication ==
	user := &model.User{
		Token: token,
	}

	err := it.usecase.Auth(user)
	if err != nil {

		res := model.Response{

			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "Authentication Failed",
				Message: model.ErrUnauthorized.Error(),
			},
		}
		send(w, res)

		return
	}

	// == Send Response ==

	// Send by protobuf
}

func send(w http.ResponseWriter, data model.Response) {

	output, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("Serialization Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	w.Write(output)
}
