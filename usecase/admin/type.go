package admin

import (
	"api/env"
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"api/repo/admin"
	"api/repo/token"
	"api/utils"
	"time"
)

type Usecase struct {
	env   env.Env
	admin admin.Repo
	token token.Repo
}

func New(env env.Env, redis redis.Redis, db postgres.DB) Usecase {

	return Usecase{
		env:   env,
		admin: admin.New(redis, db),
		token: token.New(redis),
	}
}

func (it Usecase) Find(email string) (*model.Admin, error) {

	id := utils.MD5(email)

	return it.admin.FindByID(id)
}

func (it Usecase) Store(email, username, password string) (*model.Admin, error) {

	admin := model.Admin{
		ID:        utils.MD5(email),
		Username:  username,
		Password:  utils.Hash(password),
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := it.admin.Store(&admin); err != nil {

		return nil, err
	}

	return &admin, nil
}

func (it Usecase) Associate(token *model.Token, admin *model.Admin) error {

	return it.token.Store(token.AccessToken, admin.ID)
}
