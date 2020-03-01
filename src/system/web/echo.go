package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KayacChang/API_Server/system"
	"github.com/KayacChang/API_Server/system/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ======================================

type ResponseError struct {
	Message string `json:"message"`
}

type ExecFunc func(ctx echo.Context) (interface{}, error)

// ======================================

// NewServer return Server instance
func NewServer() *echo.Echo {

	server := echo.New()

	server.Use(logger)

	server.Use(middleware.Recover())

	server.HTTPErrorHandler = errorHandler

	return server
}

type GenFunc func() interface{}

func Bind(key string, fn GenFunc) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			data := fn()

			err := ctx.Bind(data)

			if err != nil {
				/*
					422 Unprocessable Entity
					the server understands the content type of the request entity,
					and the syntax of the request entity is correct,
					but it was unable to process the contained instructions.
				*/
				return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
			}

			ctx.Set(key, data)

			return next(ctx)
		}
	}
}

func Send(code int, fn ExecFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		res, err := fn(ctx)

		if err != nil {

			code := system.GetStatusCode(err)

			res := ResponseError{
				Message: err.Error(),
			}

			return ctx.JSON(code, res)
		}

		return ctx.JSON(code, res)
	}
}

// ======================================

func logger(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		logging(c).Info("incoming request")

		return next(c)
	}
}

func errorHandler(err error, c echo.Context) {

	report, ok := err.(*echo.HTTPError)

	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	logging(c).Error(report.Message)

	c.HTML(report.Code, report.Message.(string))
}

func logging(c echo.Context) *log.Entry {

	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}
