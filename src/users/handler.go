package users

import (
	"context"
	"log"
	"net/http"

	"server/model"
	"server/system/web"
	"server/users/repo"
	"server/users/usecase"

	"github.com/labstack/echo/v4"
)

// New create accounts service
func New() {
	server := web.NewServer()

	logic := &usecase.Usecase{
		Repo: repo.New(),
	}

	server.Use(web.Bind("user", newUser))

	server.POST("/token", auth(logic))

	// TODO: search token
	// TODO: cancel the token

	// TODO: GET /users/?access_token=<token>

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

		return logic.AuthUser(c, user)
	}

	return web.Send(http.StatusOK, handler)
}
