package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KayacChang/API_Server/system/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ======================================

type ResponseError struct {
	Message string `json:"message"`
}

// ======================================

// NewServer return Server instance
func NewServer() *echo.Echo {

	server := echo.New()

	server.Use(logger)

	server.Use(middleware.Recover())

	server.HTTPErrorHandler = errorHandler

	return server
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
