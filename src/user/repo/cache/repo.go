package cache

import (
	"api/model"
	"api/user/repo"

	"github.com/patrickmn/go-cache"
)

type repository struct {
}

var storage = cache.New(
	cache.NoExpiration,
	cache.NoExpiration,
)

func New() repo.Repository {

	return &repository{}
}

func (it *repository) findByID(user *model.User) error {

	if _user, found := storage.Get(user.ID); found {

		user = _user.(*model.User)

		return nil
	}

	return model.ErrUserNotFound
}

func (it *repository) findByToken(user *model.User) error {

	if _user, found := storage.Get(user.Token); found {

		user = _user.(*model.User)

		return nil
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

	storage.SetDefault(user.ID, user)
	storage.SetDefault(user.Token, user)

	// log.Printf("repository.cache.Store\n%s\n", json.Jsonify(storage.Items()))

	return nil
}

func (it *repository) Delete(user *model.User) error {

	storage.Delete(user.ID)
	storage.Delete(user.Token)

	// log.Printf("repository.cache.Delete\n%s\n", json.Jsonify(storage.Items()))

	return nil
}
