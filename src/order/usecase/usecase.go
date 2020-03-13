package usecase

import (
	"api/model"
	"api/order/repo"
	"api/utils"
	"time"
)

type usercase struct {
	db    repo.Repository
	cache repo.Repository
}

type Usecase interface {
	Find(order *model.Order) error
	Store(order *model.Order) error
}

func New(db, cache repo.Repository) Usecase {

	return &usercase{db, cache}
}

func (it *usercase) Find(order *model.Order) error {
	return nil
}

func (it *usercase) Store(order *model.Order) error {

	// if err := it.Find(order); err == nil {

	// 	return model.ErrExisted
	// }

	order.ID = utils.UUID()

	time := time.Now()
	order.CreatedAt = &time

	return it.db.Store(order)
}
