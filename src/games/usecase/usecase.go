package usecase

import (
	"context"
	"time"

	"github.com/KayacChang/API_Server/games/entity"
	"github.com/KayacChang/API_Server/games/repo"
	"github.com/KayacChang/API_Server/utils"
)

type Usecase struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Usecase {

	return &Usecase{repo}
}

func (it *Usecase) Store(ctx context.Context, game *entity.Game) error {
	// Business
	game.ID = utils.MD5(game.Name)

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Exec
	err := it.repo.Insert(ctx, game)

	if err != nil {
		return err
	}

	return nil
}

func (it *Usecase) Find(ctx context.Context, games *[]entity.Game) error {

	// Timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err := it.repo.Find(ctx, games)

	if err != nil {
		return err
	}

	return nil
}
