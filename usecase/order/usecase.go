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
	"api/repo/suborder"
	"api/repo/token"
	"api/repo/user"
)

// Usecase ...
type Usecase struct {
	env      env.Env
	order    order.Repo
	suborder suborder.Repo
	game     game.Repo
	user     user.Repo
	token    token.Repo
	agent    agent.Agent
}

// New ...
func New(env env.Env, redis redis.Redis, db postgres.DB) Usecase {

	return Usecase{
		env:      env,
		order:    order.New(redis, db),
		suborder: suborder.New(redis, db),
		game:     game.New(redis, db),
		user:     user.New(redis, db),
		token:    token.New(redis),
		agent:    agent.New(env),
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

// SendOrder ...
func (it Usecase) SendOrder(user *model.User, game *model.Game, order *model.Order) (*agent.Bet, error) {

	bet := agent.Bet{
		OrderID: utils.UUID(),

		Amount: order.Bet,

		Username: user.Username,
		Gamename: game.Name,

		Session:   user.Session,
		CreatedAt: time.Now(),
	}

	balance, err := it.agent.SendBet(bet)
	if err != nil {
		return nil, err
	}

	// == Update Balance ==
	user.Balance = balance
	go it.UpdateUser(user)

	return &bet, nil
}

// SendSubOrder ...
func (it Usecase) SendSubOrder(user *model.User, game *model.Game, subOrder *model.SubOrder) (*agent.Bet, error) {

	bet := agent.Bet{
		OrderID:    subOrder.OrderID,
		SubOrderID: utils.UUID(),

		Amount: subOrder.Bet,

		Username: user.Username,
		Gamename: game.Name,

		Session:   user.Session,
		CreatedAt: time.Now(),
	}

	balance, err := it.agent.SendBet(bet)
	if err != nil {
		return nil, err
	}

	// == Update Balance ==
	user.Balance = balance
	go it.UpdateUser(user)

	return &bet, nil
}

// Checkout ...
func (it Usecase) Checkout(user *model.User, game *model.Game, order *model.Order) error {

	subOrders, err := it.suborder.FindAllInOrder(order.ID)
	if err != nil {
		return err
	}

	subOrderIDs := []string{}
	for _, subOrder := range subOrders {
		subOrderIDs = append(subOrderIDs, subOrder.ID)
	}

	bet := agent.Bet{
		OrderID:   order.ID,
		Username:  user.Username,
		Gamename:  game.Name,
		Amount:    order.Win,
		Session:   user.Session,
		CreatedAt: time.Now(),
	}

	balance, err := it.agent.SendEndRound(bet, subOrderIDs...)
	if err != nil {
		order.State = model.Issue

		return err
	}

	// == Update ==
	go it.StoreOrder(order)

	user.Balance = balance
	go it.UpdateUser(user)

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

// StoreSubOrder ...
func (it Usecase) StoreSubOrder(subOrder *model.SubOrder) error {

	subOrder.UpdatedAt = time.Now()

	return it.suborder.Store(subOrder)
}
