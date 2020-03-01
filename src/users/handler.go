package users

import (
	"context"
	"log"
	"net/http"

	"github.com/KayacChang/API_Server/model"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/web"
	"github.com/KayacChang/API_Server/users/repo"
	"github.com/KayacChang/API_Server/users/usecase"
	"github.com/labstack/echo/v4"
)

// New create accounts service
func New(cfg env.Config) {
	server := web.NewServer()

	logic := usecase.New(repo.New(cfg.Postgres))

	server.Use(web.Bind("user", newUser))

	server.POST("/token", auth(logic))

	log.Fatal(server.StartTLS(":8081", ".private/cert.pem", ".private/key.pem"))
}

func newUser() interface{} {

	return &model.User{}
}

func auth(logic *usecase.Usecase) echo.HandlerFunc {

	handler := func(ctx echo.Context) (interface{}, error) {

		c := ctx.Request().Context()

		if c == nil {
			c = context.Background()
		}

		user := ctx.Get("user").(*model.User)

		return logic.Auth(c, user)
	}

	return web.Send(http.StatusOK, handler)
}
