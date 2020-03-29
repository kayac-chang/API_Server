package admin

import (
	"api/model"
	"api/utils"
	"time"

	errs "github.com/pkg/errors"
)

func (it *Usecase) Store(req map[string]string) (*model.Admin, error) {

	email := req["email"]
	username := req["username"]
	password := req["password"]
	org := req["organization"]

	// Long Task
	if err := utils.CheckMail(email); err != nil {

		return nil, errs.Wrap(err, "Check Email failed")
	}

	if err := checkPassword(password); err != nil {

		return nil, errs.Wrap(err, "Password Not Valid")
	}

	admin := model.Admin{
		ID:        newID(email),
		Username:  username,
		Password:  utils.Hash(password),
		Email:     email,
		Org:       org,
		CreatedAt: time.Now(),
	}

	if err := it.repo.Store("DB", &admin); err != nil {

		return nil, err
	}

	return &admin, nil
}

// == private ==
func newID(text string) string {
	return utils.MD5(text)
}
