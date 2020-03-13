package cache

import (
	"api/framework/cache"
	"api/model"
	"api/user/repo"
)

type repository struct {
	*cache.Cache
}

func New() repo.Repository {

	return &repository{cache.Get()}
}

func (it *repository) findByID(user *model.User) error {

	if _user, found := it.Get(user.ID); found {

		if _user, ok := _user.(model.User); ok {

			*user = _user

			return nil
		}
	}

	return model.ErrUserNotFound
}

func (it *repository) findByToken(user *model.User) error {

	if _user, found := it.Get(user.Token); found {

		if _user, ok := _user.(model.User); ok {

			*user = _user

			return nil
		}
	}

	return model.ErrUserNotFound
}

func (it *repository) FindBy(key string, user *model.User) error {

	switch key {
	case "ID":
		return it.findByID(user)
	case "Token":
		return it.findByToken(user)
	}

	return model.ErrUserNotFound
}

func (it *repository) Store(user *model.User) error {

	it.SetDefault(user.ID, *user)
	it.SetDefault(user.Token, *user)

	// log.Printf("repository.cache.Store\n%s\n", json.Jsonify(storage.Items()))

	return nil
}

func (it *repository) Remove(user *model.User) error {

	it.Delete(user.ID)
	it.Delete(user.Token)

	// log.Printf("repository.cache.Delete\n%s\n", json.Jsonify(storage.Items()))

	return nil
}
