package token

import (
	"api/framework/jwt"
	"api/model"
	"net/http"

	errs "github.com/pkg/errors"
)

func (it *Usecase) Regist(username string, session string) (*model.Token, error) {

	// Check user with agent
	balance, err := it.agent.CheckPlayer(username, session)
	if err != nil {
		msg := "Request username authorized failed"

		return nil, &model.Error{
			Code:    http.StatusUnauthorized,
			Name:    "Unauthorized",
			Message: errs.WithMessage(err, msg).Error(),
		}
	}

	// Generate JWT token
	token, err := it.sign()
	if err != nil {
		msg := "Error occured when generating JWT token"

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Name:    "Generate Token Failed",
			Message: errs.WithMessage(err, msg).Error(),
		}
	}

	// Store user in Redis with key users:token

	return token, nil
}

func (it *Usecase) sign() (*model.Token, error) {

	return jwt.Sign(it.env)
}
