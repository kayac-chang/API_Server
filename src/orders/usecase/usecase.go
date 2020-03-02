package usecase

import (
	"context"
	"time"

	"github.com/KayacChang/API_Server/model"
	"github.com/KayacChang/API_Server/orders/repo"
)

type Usecase struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

const serviceName = "service"

func (it *Usecase) Create(ctx context.Context, order *model.Order) error {

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	order.ID = UUID(serviceName)

	order.State = model.Pending

	return it.repo.Insert(ctx, order)
}
