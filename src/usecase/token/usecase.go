package token

import (
	"api/agent"
	"api/env"
	"api/framework/redis"

	"api/repo/token"
	"api/repo/user"
)

type Usecase struct {
	env   env.Env
	user  user.Repo
	token token.Repo
	agent agent.Agent
}

func New(env env.Env, db redis.Redis) Usecase {

	user := user.New(db)
	token := token.New(db)
	agent := agent.New(env)

	return Usecase{env, user, token, agent}
}
