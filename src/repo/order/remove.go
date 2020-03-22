package order

import "api/model"

func (it *Repo) RemoveCache(order *model.Order) {

	it.cache.Delete(order.ID)
}
