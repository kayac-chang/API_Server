package cache

import (
	"api/framework/cache"
	"api/model"
	"api/order/repo"
)

type repository struct {
	*cache.Cache
}

func New() repo.Repository {

	return &repository{cache.Get()}
}

func (it *repository) findByID(order *model.Order) error {

	if _order, found := it.Get(order.ID); found {

		if _order, ok := _order.(*model.Order); ok {

			order = _order

			return nil
		}
	}

	return model.ErrGameNotFound
}

func (it *repository) FindBy(key string, order *model.Order) error {

	switch key {
	case "ID":
		return it.findByID(order)
	}

	return model.ErrGameNotFound
}

func (it *repository) Store(order *model.Order) error {

	it.SetDefault(order.ID, order)

	// log.Printf("repository.cache.Store\n%s\n", json.Jsonify(storage.Items()))

	return nil
}

func (it *repository) Replace(order *model.Order) error {

	return nil
}
