package token

import (
	"api/agent"
	"api/env"

	repo "api/repo/user"
)

type Usecase struct {
	env   *env.Env
	repo  *repo.Repo
	agent *agent.Agent
}

func New(env *env.Env, repo *repo.Repo, agent *agent.Agent) *Usecase {

	return &Usecase{env, repo, agent}
}
