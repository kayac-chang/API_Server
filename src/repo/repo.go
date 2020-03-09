package repo

import "user/model"

type Repository interface {
	Store(user *model.User) error
	FindBy(key string, user *model.User) error
	Delete(user *model.User) error
}
