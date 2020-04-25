package token

import (
	"api/model"
)

func (it Usecase) Auth(token string) (string, error) {

	return it.token.Find(token)
}

func (it Usecase) FindUserByID(id string) (*model.User, error) {

	user, err := it.user.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
