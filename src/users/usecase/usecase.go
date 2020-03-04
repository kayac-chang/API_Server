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
	repo *repo.Repo
}

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

func (it *Usecase) Auth(ctx context.Context, user *model.User) (interface{}, error) {
	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// err := getUserFromAgency(ctx, user)

	// if err != nil {
	// 	return nil, err
	// }

	err := it.repo.FindByName(ctx, user)

	if err == nil {
		// If user exist
		return genToken(user)
	}

	// user not exist
	user, err = it.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return genToken(user)
}

func (it *Usecase) Create(ctx context.Context, user *model.User) (*model.User, error) {

	user.ID = utils.MD5(user.Username)

	err := it.repo.Insert(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func genToken(user *model.User) (interface{}, error) {
	exp := time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": env.ServiceID(),

		"iat": time.Now().Unix(),

		"exp": exp,

		"user": user.ID,
	})

	tokenString, err := token.SignedString(
		// TODO: Must change to secret number in .env
		[]byte(env.ServiceID()),
	)

	if err != nil {
		return nil, system.ErrGenTokenError
	}

	res := &model.Token{
		ServiceID:   env.ServiceID(),
		AccessToken: tokenString,
		Type:        "Bearer",
		Expire:      exp,
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
