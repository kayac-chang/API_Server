package token

import (
	"api/framework/jwt"
	"api/model"
	"net/http"
	"time"

	errs "github.com/pkg/errors"
)

// Regist regist token business flow
func (it Usecase) Regist(username string, session string) (*model.Token, error) {

	_balance, _err := it.agent.CheckPlayer(username, session)
	if _err != nil {

		msg := "Request username authorized failed"

		err := &model.Error{
			Code:    http.StatusUnauthorized,
			Message: errs.WithMessage(_err, msg).Error(),
		}

		return nil, err
	}

	// TODO: maybe fix to uint64 is a bad idea
	balance := uint64(_balance)

	token, _err := jwt.Sign(it.env)
	if _err != nil {

		msg := "Error occured when generating JWT token"

		err := &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(_err, msg).Error(),
		}

		return nil, err
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
	if _err := it.user.Store(user); _err != nil {

		msg := "Error occured when storing user"

		err := &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(_err, msg).Error(),
		}

		return nil, err
	}

	// Associate
	associate := map[string]string{
		"session": session,
		"user":    username,
	}
	if _err := it.token.Store(token.AccessToken, associate); _err != nil {

		msg := "Error occured when storing token associate data"

		err := &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(_err, msg).Error(),
		}

		return nil, err
	}

	return &token, nil
}
