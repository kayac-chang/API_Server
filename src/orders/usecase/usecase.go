package usecase

import (
	"context"
	"time"

	"server/model"
	"server/orders/repo"
	"server/utils"
)

type Usecase struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

func (it *Usecase) Create(ctx context.Context, order *model.Order) error {

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// TODO: send to /api/v1/tgc/transaction/game/bet, ask for betting

	order.ID = utils.UUID()
	order.State = model.Pending

	// TODO: Move to redis
	return it.repo.Insert(ctx, order)
}
