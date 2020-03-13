package repo

import "api/model"

type Repository interface {
	Store(user *model.User) error
	FindBy(key string, user *model.User) error
	Remove(user *model.User) error
}
