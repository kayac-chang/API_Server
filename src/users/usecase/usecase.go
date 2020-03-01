package usecase

import (
	"context"
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

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

func (it *Usecase) Create(ctx context.Context, user *model.User) error {

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Business
	user.ID = utils.MD5(user.Username)

	user.Password = utils.Hash(user.Password)

	// Exec
	return it.repo.Insert(ctx, user)
}

const secret = "secret"
const serviceName = "sunny.com"

type response struct {
	Token  string `json:"access_token"`
	Type   string `json:"token_type"`
	Expire int64  `json:"expires_in"`
}

func (it *Usecase) Auth(ctx context.Context, user *model.User) (*response, error) {

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Business
	user.Password = utils.Hash(user.Password)

	// Auth
	err := it.repo.FindByUserAndPW(ctx, user)

	if err != nil {

		return nil, system.ErrAuthFailure
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
