package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/KayacChang/API_Server/model"
	"github.com/KayacChang/API_Server/system"
	"github.com/KayacChang/API_Server/users/repo"
	"github.com/KayacChang/API_Server/utils"
	"github.com/dgrijalva/jwt-go"
)

type Usecase struct {
	repo *repo.Repo
}

const secret = "secret"
const serviceName = "sunny.com"

type response struct {
	Token  string `json:"access_token"`
	Type   string `json:"token_type"`
	Expire int64  `json:"expires_in"`
}

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

func (it *Usecase) Auth(ctx context.Context, user *model.User) (*response, error) {

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Check

	// if !exist(ctx, user) {
	// 	return nil, system.ErrAuthFailure
	// }

	user.ID = utils.MD5(user.Username)

	user.Password = utils.Hash(user.Password)

	err := it.repo.FindByID(ctx, user)

	if err != nil {

		err = it.repo.Insert(ctx, user)

		if err != nil {

			return nil, err
		}
	}

	// Generate JWT Token
	exp := time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": serviceName,

		"iat": time.Now().Unix(),

		"exp": exp,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {

		return nil, system.ErrGenTokenError
	}

	res := &response{
		Token:  tokenString,
		Type:   "Bearer",
		Expire: exp,
	}

	return res, nil
}

func exist(ctx context.Context, user *model.User) bool {

	req := utils.Request{

		URL: fmt.Sprintf("/api/tgc/player/check/%s", user.Username),

		Header: map[string]string{
			"organization_token": serviceName,
		},

		Context: ctx,
	}

	res := req.Send()

	defer res.Body.Close()

	var data struct {
		Status struct {
			Message string `json:"message"`
		} `json:"status"`
	}

	err := json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}

	return data.Status.Message == "Success"
}
