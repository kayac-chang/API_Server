package jwt

import (
	"api/env"
	"api/model"
	"api/utils"

	"time"

	"github.com/dgrijalva/jwt-go"
)

func Sign(env env.Env) (*model.Token, error) {

	createdTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": env.Service.ID,
		"iat": createdTime.Unix(),
		"jti": utils.UUID(),
	})

	tokenString, err := token.SignedString(env.Secret)
	if err != nil {

		return nil, err
	}

	res := model.Token{
		AccessToken: tokenString,
		Type:        "Bearer",
		ServiceID:   env.Service.ID,
		CreatedAt:   createdTime,
	}

	return &res, nil
}
