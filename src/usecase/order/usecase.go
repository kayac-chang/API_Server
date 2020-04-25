package order

import (
	"api/agent"
	"api/env"
	"api/framework/postgres"
	"api/framework/redis"
	"api/model"
	"api/utils"
	"strings"
	"time"

	"api/repo/game"
	"api/repo/order"
	"api/repo/token"
	"api/repo/user"
)

// Usecase ...
type Usecase struct {
	env   env.Env
	order order.Repo
	game  game.Repo
	user  user.Repo
	token token.Repo
	agent agent.Agent
}

// New ...
func New(env env.Env, redis redis.Redis, db postgres.DB) Usecase {

	return Usecase{
		env:   env,
		order: order.New(redis, db),
		game:  game.New(redis, db),
		user:  user.New(redis, db),
		token: token.New(redis),
		agent: agent.New(env),
	}
}

// Auth ...
func (it Usecase) Auth(token string) error {

	token = strings.TrimPrefix(token, "Bearer ")

	_, err := it.token.Find(token)

	return err
}

// FindUserByID ...
func (it Usecase) FindUserByID(id string) (*model.User, error) {

	user, err := it.user.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindGameByID ...
func (it Usecase) FindGameByID(id string) (*model.Game, error) {

	game, err := it.game.FindByID(id)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// FindOrderByID ...
func (it Usecase) FindOrderByID(id string) (*model.Order, error) {

	order, err := it.order.FindByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// SendBet ...
func (it Usecase) SendBet(user *model.User, game *model.Game, order *model.Order) error {

	bet := agent.Bet{
		Roundid:   utils.UUID(),
		Username:  user.Username,
		Gamename:  game.Name,
		Amount:    order.Bet,
		Session:   user.Session,
		CreatedAt: time.Now(),
	}

	balance, err := it.agent.SendBet(bet)
	if err != nil {
		return err
	}

	// == Update Balance ==
	user.Balance = balance
	defer it.UpdateUser(user)

	// == Store Order ==
	order.ID = bet.Roundid
	order.State = model.Pending
	order.CreatedAt = bet.CreatedAt
	defer it.StoreOrder(order)

	return nil
}

// Checkout ...
func (it Usecase) Checkout(user *model.User, game *model.Game, order *model.Order) error {

	bet := agent.Bet{
		Roundid:   order.ID,
		Username:  user.Username,
		Gamename:  game.Name,
		Amount:    order.Win,
		Session:   user.Session,
		CreatedAt: time.Now(),
	}

	balance, err := it.agent.SendEndRound(bet)
	if err != nil {
		order.State = model.Issue

		return err
	}

	// == Update ==
	order.State = model.Completed
	order.CompletedAt = time.Now()
	defer it.StoreOrder(order)

	user.Balance = balance
	defer it.UpdateUser(user)

	return nil
}

// StoreOrder ...
func (it Usecase) StoreOrder(order *model.Order) error {

	order.UpdatedAt = time.Now()

	return it.order.Store(order)
}

// UpdateUser ...
func (it Usecase) UpdateUser(user *model.User) error {

	user.UpdatedAt = time.Now()

	return it.user.Store(user)
}
