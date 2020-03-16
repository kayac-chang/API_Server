package server

import (
	"api/framework/cache"
	"api/model"
	"api/model/pb"
	"api/model/request"
	"api/model/response"
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func (it *Server) URLParam(r *http.Request, key string) string {

	return chi.URLParam(r, key)
}

func (it *Server) ParseJSON(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		req := map[string]string{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

			res := response.JSON{
				Code: http.StatusBadRequest,

				Error: model.Error{
					Name:    "Unexpect Payload",
					Message: model.ErrUnexpectPayload.Error(),
				},
			}

			it.Send(w, res)

			return
		}

		next.ServeHTTP(w, bind(r, req))
	}

	return http.HandlerFunc(fn)
}

func (it *Server) User(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 {

			it.Send(w, response.ProtoBuf{
				Code: http.StatusUnauthorized,

				Data: &pb.Error{
					Code:    http.StatusUnauthorized,
					Name:    "Authentication Failed",
					Message: model.ErrUnauthorized.Error(),
				},
			})

			return
		}

		user := model.User{
			Token: auth[1],
		}
		if err := findByToken(cache.Get(), &user); err != nil {

			it.Send(w, response.ProtoBuf{
				Code: http.StatusUnauthorized,

				Data: &pb.Error{
					Code:    http.StatusUnauthorized,
					Name:    "Authentication Failed",
					Message: err.Error(),
				},
			})

			return
		}

		next.ServeHTTP(w, bind(r, &user))
	}

	return http.HandlerFunc(fn)
}

func (it *Server) Order(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		reqByte, err := ioutil.ReadAll(r.Body)
		if err != nil {

			it.Send(w, response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Unexpect Payload",
					Message: err.Error(),
				},
			})

			return
		}

		req := pb.Order{}
		if err := proto.Unmarshal(reqByte, &req); err != nil {

			it.Send(w, response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Unexpect Payload",
					Message: err.Error(),
				},
			})

			return
		}

		order := model.Order{
			ID: req.GetOrderId(),

			GameID: req.GetGameId(),
			UserID: req.GetUserId(),

			State: model.ToState(req.GetState()),
			Bet:   req.GetBet(),

			CreatedAt:   Time(req.GetCreatedAt()),
			CompletedAt: Time(req.GetCompletedAt()),
		}

		next.ServeHTTP(w, bind(r, &order))
	}

	return http.HandlerFunc(fn)
}

func Time(ts *timestamp.Timestamp) sql.NullTime {

	if ts == nil {
		return sql.NullTime{}
	}

	time, err := ptypes.Timestamp(ts)

	if err != nil {
		panic(err)
	}

	return sql.NullTime{time, true}
}

func findByToken(cache *cache.Cache, user *model.User) error {

	if _user, found := cache.Get(user.Token); found {

		if _user, ok := _user.(model.User); ok {

			*user = _user

			return nil
		}
	}

	return model.ErrUserNotFound
}

func bind(r *http.Request, val interface{}) *http.Request {

	ctx := r.Context()

	switch val := val.(type) {

	case map[string]string:
		ctx = context.WithValue(ctx, request.JSON, val)

	case *model.User:
		ctx = context.WithValue(ctx, request.USER, val)

	case *model.Order:
		ctx = context.WithValue(ctx, request.ORDER, val)

	default:
		log.Fatalf("Unsupport Type: %t\n", val)
	}

	return r.WithContext(ctx)
}
