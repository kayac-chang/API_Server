package token

import (
	"api/model"
)

func (it *Usecase) Auth(token string) (*model.User, error) {

	user := model.User{
		Token: token,
	}

	// if err := it.repo.FindBy("Token", &user); err != nil {
	// 	return nil, err
	// }

	return &user, nil
}
