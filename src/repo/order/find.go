package order

import "api/model"

func (it *Repo) FindByID(id string) (*model.Order, error) {

	if _order, found := it.cache.Get(id); found {

		if order, ok := _order.(*model.Order); ok {

			return order, nil
		}
	}

	return nil, model.ErrOrderNotFound
}
