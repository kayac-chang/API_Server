package user

import "api/model"

func (it *Repo) FindBy(key string, user *model.User) error {

	switch key {

	case "ID":
		return it.findByID(user)

	case "Token":
		return it.findByToken(user)
	}

	return model.ErrUserNotFound
}

func (it *Repo) findByID(user *model.User) error {

	// Find cache
	if _user, found := it.cache.Get(user.ID); found {

		if _user, ok := _user.(model.User); ok {

			*user = _user

			return nil
		}
	}

	// Find DB
	if err := it.db.Get(user, it.sql.findByID, user.ID); err == nil {
		return nil
	}

	return model.ErrUserNotFound
}

func (it *Repo) findByToken(user *model.User) error {

	if _user, found := it.cache.Get(user.Token); found {

		if _user, ok := _user.(model.User); ok {

			*user = _user

			return nil
		}
	}

	return model.ErrUserNotFound
}
