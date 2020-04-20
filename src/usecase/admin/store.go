package admin

import (
	"api/model"
	"api/utils"
	"time"
)

func (it Usecase) Store(email, username, password string) (*model.Admin, error) {

	admin := model.Admin{
		ID:        utils.MD5(email),
		Username:  username,
		Password:  utils.Hash(password),
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := it.repo.Store(&admin); err != nil {
		return nil, err
	}

	return &admin, nil
}
