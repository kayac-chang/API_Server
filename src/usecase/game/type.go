package game

import (
	"api/env"
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"api/repo/game"
	"api/repo/token"
	"api/utils"
	"strings"
	"time"
)

// Usecase game usecase instance
type Usecase struct {
	env   env.Env
	game  game.Repo
	token token.Repo
}

// New create game usecase instance
func New(env env.Env, redis redis.Redis, db postgres.DB) Usecase {

	return Usecase{
		env:   env,
		game:  game.New(redis, db),
		token: token.New(redis),
	}
}

// Store store game in to repo
func (it Usecase) Store(name, href, category string) (*model.Game, error) {

	game := model.Game{
		ID: utils.MD5(name),

		Name:     name,
		Href:     href,
		Category: category,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := it.game.Store(&game); err != nil {

		return nil, err
	}

	return &game, nil
}

// Update update game in to repo
func (it Usecase) Update(game *model.Game) (*model.Game, error) {

	game.UpdatedAt = time.Now()

	if err := it.game.Update(game); err != nil {

		return nil, err
	}

	return game, nil
}

// FindByID find game by id in repo
func (it Usecase) FindByID(id string) (*model.Game, error) {

	return it.game.FindByID(id)
}

// FindByName find game by name in repo
func (it Usecase) FindByName(name string) (*model.Game, error) {

	return it.game.FindByName(name)
}

// FindAll find all games in repo
func (it Usecase) FindAll() ([]model.Game, error) {

	return it.game.FindAll()
}

// Auth check the admin token
func (it Usecase) Auth(token string) error {

	token = strings.TrimPrefix(token, "Bearer ")

	_, err := it.token.Find(token)

	return err
}
