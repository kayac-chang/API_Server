package token

import (
	"api/agent"
	"api/env"
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"

	"api/repo/game"
	"api/repo/token"
	"api/repo/user"
)

// Usecase ...
type Usecase struct {
	env   env.Env
	user  user.Repo
	token token.Repo
	game  game.Repo
	agent agent.Agent
}

// New ...
func New(env env.Env, redis redis.Redis, db postgres.DB) Usecase {

	user := user.New(redis, db)
	token := token.New(redis)
	game := game.New(redis, db)
	agent := agent.New(env)

	return Usecase{env, user, token, game, agent}
}

// FindGameByName ...
func (it Usecase) FindGameByName(name string) (*model.Game, error) {

	return it.game.FindByName(name)
}
