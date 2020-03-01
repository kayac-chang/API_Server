package usecase

import (
	"context"
	"time"

	"github.com/KayacChang/API_Server/accounts/entity"
	"github.com/KayacChang/API_Server/accounts/repo"
	"github.com/KayacChang/API_Server/utils"
)

type Usecase struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

func (it *Usecase) Store(ctx context.Context, account *entity.Account) error {
	// Business
	account.ID = utils.Hash(
		account.Email + account.Username,
	)

	account.Password = utils.HashAndSalt(account.Password)

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Exec
	return it.repo.Insert(ctx, account)
}
