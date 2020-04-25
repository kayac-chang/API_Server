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
		order: order.New(redis),
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

// SendBet ...
func (it Usecase) SendBet(user *model.User, game *model.Game, order *model.Order) (float64, error) {

	bet := agent.Bet{
		Roundid:   utils.UUID(),
		Username:  user.Username,
		Gamename:  game.Name,
		Amount:    float64(order.Bet),
		Session:   user.Session,
		CreatedAt: time.Now(),
	}

	balance, err := it.agent.SendBet(bet)
	if err != nil {
		return 0, err
	}

	order.ID = bet.Roundid
	order.State = model.Pending
	order.CreatedAt = bet.CreatedAt

	return balance, nil
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

// func (it Usecase) Checkout(orderID string, win uint64) (*model.Order, error) {

// 	order, err := it.order.FindByID(orderID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	order.State = model.Completed
// 	order.Win = win
// 	order.CompletedAt = sql.NullTime{time.Now(), true}

// 	// send end round
// 	if err := it.sendEndRound(order); err != nil {

// 		order.State = model.Issue

// 		it.order.Store("Cache", order)

// 		return nil, err
// 	}

// 	if err := it.Store(order); err != nil {

// 		return nil, err
// 	}

// 	return order, nil
// }

// === Private ===

// func (it *Usecase) sendEndRound(order *model.Order) error {
// 	api := "/transaction/game/endround"

// 	// user := model.User{
// 	// 	ID: order.UserID,
// 	// }
// 	// if err := it.user.FindBy("ID", &user); err != nil {
// 	// 	return err
// 	// }

// 	game, err := it.game.FindByID(order.GameID)
// 	if err != nil {
// 		return err
// 	}

// 	url := it.env.Agent.Domain + it.env.Agent.API + api

// 	req := map[string]interface{}{
// 		// "account":      user.Username,
// 		"gamename":     game.Name,
// 		"roundid":      order.ID,
// 		"amount":       order.Win,
// 		"completed_at": order.CompletedAt.Time,
// 	}

// 	headers := map[string]string{
// 		"Content-Type":       "application/json",
// 		"organization-token": it.env.Agent.Token,
// 		// "session":            user.Session,
// 	}

// 	resp, err := utils.Post(url, req, headers)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	res := map[string]interface{}{}
// 	json.Parse(resp.Body, &res)

// 	if resp.StatusCode != 200 {
// 		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

// 		err := res["error"].(map[string]interface{})

// 		return fmt.Errorf("%s", err["message"])
// 	}

// 	log.Printf("Agent: [ %s ] Success !!!\nResponse:\n %s", api, json.Jsonify(res))

// 	// data := res["data"].(map[string]interface{})
// 	// balance := data["balance"].(float64)

// 	// user.Balance = uint64(balance)

// 	// if err := it.user.Store("Cache", &user); err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }
