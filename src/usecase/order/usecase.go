package order

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/model"
	gamerepo "api/repo/game"
	orderrepo "api/repo/order"
	userrepo "api/repo/user"
	"api/utils"
	"api/utils/json"

	"database/sql"
	"fmt"
	"log"
	"time"
)

type Usecase struct {
	env   *env.Env
	order *orderrepo.Repo
	user  *userrepo.Repo
	game  *gamerepo.Repo
}

func New(env *env.Env, db *postgres.DB, c *cache.Cache) *Usecase {

	return &Usecase{
		env:   env,
		order: orderrepo.New(db, c),
		// user:  userrepo.New(db, c),
		game: gamerepo.New(db, c),
	}
}

func (it *Usecase) Create(order *model.Order) error {

	if _, err := it.game.FindByID(order.GameID); err != nil {
		return err
	}

	order.ID = utils.UUID()
	order.State = model.Pending
	order.CreatedAt = sql.NullTime{time.Now(), true}

	if err := it.sendBet(order); err != nil {
		return err
	}

	return it.order.Store("Cache", order)
}

func (it *Usecase) Checkout(orderID string, win uint64) (*model.Order, error) {

	order, err := it.order.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	order.State = model.Completed
	order.Win = win
	order.CompletedAt = sql.NullTime{time.Now(), true}

	// send end round
	if err := it.sendEndRound(order); err != nil {

		order.State = model.Issue

		it.order.Store("Cache", order)

		return nil, err
	}

	if err := it.Store(order); err != nil {

		return nil, err
	}

	return order, nil
}

func (it *Usecase) Store(order *model.Order) error {

	if err := it.order.Store("DB", order); err != nil {
		return err
	}

	it.order.RemoveCache(order)

	return nil
}

// === Private ===

func (it *Usecase) sendBet(order *model.Order) error {
	api := "/transaction/game/bet"

	user := model.User{
		ID: order.UserID,
	}
	// if err := it.user.FindBy("ID", &user); err != nil {
	// 	return err
	// }

	game, err := it.game.FindByID(order.GameID)
	if err != nil {
		return err
	}

	url := it.env.Agent.Domain + it.env.Agent.API + api

	req := map[string]interface{}{
		"account":    user.Username,
		"created_at": order.CreatedAt.Time,
		"gamename":   game.Name,
		"roundid":    order.ID,
		"amount":     order.Bet,
	}

	headers := map[string]string{
		"Content-Type":       "application/json",
		"organization-token": it.env.Agent.Token,
		"session":            user.Session,
	}

	resp, err := utils.Post(url, req, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	res := map[string]interface{}{}
	json.Parse(resp.Body, &res)

	if resp.StatusCode != 200 {
		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

		err := res["error"].(map[string]interface{})

		return fmt.Errorf("%s", err["message"])
	}

	log.Printf("Agent: [ %s ] Success !!!\nResponse:\n %s", api, json.Jsonify(res))

	data := res["data"].(map[string]interface{})
	balance := data["balance"].(float64)

	user.Balance = uint64(balance)

	// if err := it.user.Store("Cache", &user); err != nil {
	// 	return err
	// }

	return nil
}

func (it *Usecase) sendEndRound(order *model.Order) error {
	api := "/transaction/game/endround"

	user := model.User{
		ID: order.UserID,
	}
	// if err := it.user.FindBy("ID", &user); err != nil {
	// 	return err
	// }

	game, err := it.game.FindByID(order.GameID)
	if err != nil {
		return err
	}

	url := it.env.Agent.Domain + it.env.Agent.API + api

	req := map[string]interface{}{
		"account":      user.Username,
		"gamename":     game.Name,
		"roundid":      order.ID,
		"amount":       order.Win,
		"completed_at": order.CompletedAt.Time,
	}

	headers := map[string]string{
		"Content-Type":       "application/json",
		"organization-token": it.env.Agent.Token,
		"session":            user.Session,
	}

	resp, err := utils.Post(url, req, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	res := map[string]interface{}{}
	json.Parse(resp.Body, &res)

	if resp.StatusCode != 200 {
		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

		err := res["error"].(map[string]interface{})

		return fmt.Errorf("%s", err["message"])
	}

	log.Printf("Agent: [ %s ] Success !!!\nResponse:\n %s", api, json.Jsonify(res))

	data := res["data"].(map[string]interface{})
	balance := data["balance"].(float64)

	user.Balance = uint64(balance)

	// if err := it.user.Store("Cache", &user); err != nil {
	// 	return err
	// }

	return nil
}
