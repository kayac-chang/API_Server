package admin

import (
	"api/model"
	"fmt"
)

func (it *Repo) FindByID(id string) (*model.Admin, error) {

	admin := model.Admin{}

	// Find in Cache
	// TODO

	// Find in DB
	if err := it.db.Get(&admin, it.sql.findByID, id); err != nil {
		return nil, err
	}

	// === Save to Cache ===
	// TODO

	return &admin, nil
}

func (it *Repo) FindByToken(token string) (*model.Admin, error) {

	// Find in Cache
	if _admin, founded := it.cache.Get(token); founded {

		if admin, ok := _admin.(model.Admin); ok {

			return &admin, nil
		}
	}

	return nil, fmt.Errorf("Admin not found by this token")
}
