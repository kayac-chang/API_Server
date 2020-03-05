package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/model"
	"server/system"
	"server/system/env"
	"server/users/repo"
	"server/utils"
	"server/utils/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Usecase struct {
	Repo *repo.Repo
}

// TODO: AuthToken

func (it *Usecase) AuthUser(ctx context.Context, user *model.User) (interface{}, error) {
	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// err := getUserFromAgency(ctx, user)

	// if err != nil {
	// 	return nil, err
	// }

	err := it.Repo.FindByName(ctx, user)

	if err == nil {
		// If user exist
		return it.CreateToken(ctx, user)
	}

	// user not exist
	if _, err = it.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return it.CreateToken(ctx, user)
}

func (it *Usecase) CreateUser(ctx context.Context, user *model.User) (interface{}, error) {

	user.ID = utils.MD5(user.Username)

	err := it.Repo.Insert(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (it *Usecase) UpdateUser(ctx context.Context, user *model.User) (interface{}, error) {

	err := it.Repo.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (it *Usecase) CreateToken(ctx context.Context, user *model.User) (interface{}, error) {

	user.Token = utils.UUID()

	if _, err := it.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": env.ServiceID(),

		"iat": time.Now().Unix(),

		"jti": user.Token,
	})

	tokenString, err := token.SignedString(
		// TODO: Must change to secret number in .env
		[]byte(env.ServiceID()),
	)

	if err != nil {
		return nil, system.ErrGenTokenError
	}

	res := &model.Token{
		AccessToken: tokenString,
		CreatedAt:   time.Now(),
	}

	return res, nil
}

func getUserFromAgency(ctx context.Context, user *model.User) error {

	req := utils.Request{
		URL: fmt.Sprintf("/api/tgc/player/check/%s", user.Username),

		Header: map[string]string{
			"organization_token": env.ServiceID(),
		},

		Context: ctx,
	}

	res := req.Send()

	if res.StatusCode != http.StatusOK {
		return echo.ErrNotFound
	}

	defer res.Body.Close()

	var data struct {
	}

	json.Parse(res.Body, &data)

	// TODO: Assign data to User

	return nil
}
