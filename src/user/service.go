package user

import (
	"api/env"
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
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/protobuf/proto"
)

type handler struct {
	usecase usecase.Usecase
}

var ()

func NewServer() *chi.Mux {

	server := chi.NewRouter()
	server.Use(middleware.RequestID)
	server.Use(middleware.RealIP)
	server.Use(middleware.Logger)
	server.Use(middleware.Recoverer)

	return server
}

func New(e *env.Env) {

	server := NewServer()

	c := cache.New()
	db := postgres.New(e.Postgres.ToURL(), 30)
	it := handler{usecase.New(db, c)}

	server.Post("/token", it.POST)
	server.Get("/auth", it.Auth)

	http.ListenAndServe(":8000", server)
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
		sendJSON(w, res)

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
		sendJSON(w, res)

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
		sendJSON(w, res)

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
	sendJSON(w, res)
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
		sendProtoBuf(w, res)

		return
	}

	token = strings.Split(token, " ")[1]

	// == Authentication ==
	user := &model.User{
		Token: token,
	}

	user, err := it.usecase.Auth(user)
	if err != nil {

		res := response.ProtoBuf{

			Code: http.StatusUnauthorized,

			Data: &pb.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Authentication Failed",
				Message: model.ErrUnauthorized.Error(),
			},
		}
		sendProtoBuf(w, res)

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
	sendProtoBuf(w, res)
}

func sendJSON(w http.ResponseWriter, data response.JSON) {

	output, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("Serialization Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	w.Write(output)
}

func sendProtoBuf(w http.ResponseWriter, res response.ProtoBuf) {

	out, err := proto.Marshal(res.Data)
	if err != nil {
		w.Write([]byte("Serialization Error"))
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(res.Code)
	w.Write(out)
}
