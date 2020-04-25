package token

import (
	"api/framework/jwt"
	"api/model"
	"api/utils"
	"net/http"
	"time"
)

// Regist regist token business flow
func (it Usecase) Regist(username string, session string) (*model.Token, error) {

	balance, err := it.agent.CheckPlayer(username, session)
	if err != nil {

		return nil, &model.Error{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	token, err := jwt.Sign(it.env)
	if err != nil {

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Find User
	user, err := it.user.FindByID(utils.MD5(username))
	if err != nil && err != model.ErrNotFound {

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Create User
	if user == nil {

		user = &model.User{
			ID:        utils.MD5(username),
			Username:  username,
			Balance:   balance,
			Session:   session,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := it.user.Store(user); err != nil {

			return nil, &model.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	// Associate
	if err := it.token.Store(token.AccessToken, user.ID); err != nil {

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return token, nil
}
