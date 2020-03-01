package token

import (
	"log"
	"net/http"
	"time"

	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// New create accounts service
func New(cfg env.Config) {

	server := web.NewServer()

	server.POST("/token", login)

	log.Fatal(server.StartTLS(":8081", ".private/cert.pem", ".private/key.pem"))
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
