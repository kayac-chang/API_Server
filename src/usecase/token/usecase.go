package token

import (
	"api/agent"
	"api/env"
	"api/framework/redis"
	"api/model"

	"api/repo/game"
	"api/repo/token"
	"api/repo/user"
)

type Usecase struct {
	env   env.Env
	user  user.Repo
	token token.Repo
	game  game.Repo
	agent agent.Agent
}

func New(env env.Env, redis redis.Redis) Usecase {

	user := user.New(redis)
	token := token.New(redis)
	game := game.New(redis)
	agent := agent.New(env)

	return Usecase{env, user, token, game, agent}
}

func (it Usecase) FindGameByName(name string) (*model.Game, error) {

	return it.game.Find(name)
}
