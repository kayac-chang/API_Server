package accounts

import (
	"context"
	"log"
	"net/http"

	"github.com/KayacChang/API_Server/accounts/entity"
	"github.com/KayacChang/API_Server/accounts/repo"
	"github.com/KayacChang/API_Server/accounts/usecase"
	"github.com/KayacChang/API_Server/system"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/web"
	"github.com/labstack/echo/v4"
)

// New create accounts service
func New(cfg env.Config) {
	server := web.NewServer()

	logic := usecase.New(repo.New(cfg.Postgres))

	server.POST("/accounts", post(logic), bind)

	log.Fatal(server.StartTLS(":8081", ".private/cert.pem", ".private/key.pem"))
}

func bind(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		var account entity.Account

		err := ctx.Bind(&account)

		if err != nil {
			/*
				422 Unprocessable Entity
				the server understands the content type of the request entity,
				and the syntax of the request entity is correct,
				but it was unable to process the contained instructions.
			*/
			return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		ctx.Set("account", &account)

		return next(ctx)
	}
}

func post(logic *usecase.Usecase) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		c := ctx.Request().Context()

		if c == nil {
			c = context.Background()
		}

		account := ctx.Get("account").(*entity.Account)

		return exec(
			ctx,
			http.StatusCreated,
			func() (interface{}, error) { return account, logic.Store(c, account) },
		)
	}
}

type execFunc func() (interface{}, error)

func exec(ctx echo.Context, code int, fn execFunc) error {

	res, err := fn()

	if err != nil {

		code := system.GetStatusCode(err)

		res := web.ResponseError{
			Message: err.Error(),
		}

		return ctx.JSON(code, res)
	}

	return ctx.JSON(code, res)
}
