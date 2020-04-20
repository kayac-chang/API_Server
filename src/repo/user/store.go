package user

import (
	"api/model"
)

// Store store user into redis
func (it Repo) Store(user model.User) error {

	key := prefix + user.Username

	return it.db.Set(key, user)
}
