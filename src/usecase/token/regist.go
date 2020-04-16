package token

import (
	"api/framework/jwt"
	"api/model"
	"net/http"
	"time"

	errs "github.com/pkg/errors"
)

func (it *Usecase) Regist(username string, session string) (*model.Token, error) {

	balance, err := it.agent.CheckPlayer(username, session)
	if err != nil {

		msg := "Request username authorized failed"

		return nil, &model.Error{
			Code:    http.StatusUnauthorized,
			Message: errs.WithMessage(err, msg).Error(),
		}
	}

	token, err := jwt.Sign(it.env)
	if err != nil {

		msg := "Error occured when generating JWT token"

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(err, msg).Error(),
		}
	}

	// Create User
	user := model.User{
		Username: username,
		Balance:  balance,
		Data: model.Data{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := it.user.Store(user); err != nil {

		msg := "Error occured when storing user"

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(err, msg).Error(),
		}
	}

	// Associate
	associate := map[string]string{
		"session": session,
		"user":    username,
	}
	if err := it.token.Store(token.AccessToken, associate); err != nil {

		msg := "Error occured when storing token associate data"

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(err, msg).Error(),
		}
	}

	return token, nil
}
