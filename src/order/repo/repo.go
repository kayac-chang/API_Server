package repo

import "api/model"

type Repository interface {
	Store(order *model.Order) error
	FindBy(key string, order *model.Order) error
	Replace(order *model.Order) error
}
