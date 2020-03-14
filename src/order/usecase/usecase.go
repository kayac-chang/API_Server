package usecase

import (
	"api/env"
	"api/model"
	"api/order/repo"
	userRepo "api/user/repo"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

type usercase struct {
	env       *env.Env
	db        repo.Repository
	cache     repo.Repository
	userCache userRepo.Repository
}

type Usecase interface {
	Find(order *model.Order) error
	Store(order *model.Order) error
	Checkout(order *model.Order) error
}

func New(env *env.Env, db, cache repo.Repository, userCache userRepo.Repository) Usecase {

	return &usercase{env, db, cache, userCache}
}

func (it *usercase) Find(order *model.Order) error {

	return it.db.FindBy("ID", order)
}

func (it *usercase) Store(order *model.Order) error {

	// order.ID = utils.UUID()
	// user := model.User{
	// 	ID: order.UserID,
	// }
	// if err := it.userCache.FindBy("ID", &user); err != nil {

	// 	return model.ErrUserNotFound
	// }

	// // Send to /api/v1/tgc/transaction/game/bet
	// _url := fmt.Sprintf("%s%s%s", agent_domain, agent_api_root, "/transaction/game/bet")

	// req := url.Values{
	// 	"account":   {user.Username},
	// 	"eventTime": {time.Now().String()},
	// 	"gamehall":  {service_id},
	// 	"gamecode":  {order.GameID},
	// 	"roundid":   {order.ID},
	// 	"amount":    {strconv.FormatUint(order.Bet, 10)},
	// 	"mtcode":    {utils.UUID()},
	// }

	// type Res struct {
	// 	Data map[string]interface{} `json:"data"`

	// 	Status struct {
	// 		Code    string    `json:"code"`
	// 		Message string    `json:"message"`
	// 		Time    time.Time `json:"datatime"`
	// 	} `json:"status"`
	// }

	// res := &Res{}

	// if err := post(_url, req, res); err != nil {

	// 	return errs.WithMessagef(err, "status: %+v", res.Status)
	// }

	// TODO: transfer to game currency
	// user.Balance = res.Data["balance"].(uint64)

	order.State = model.Pending

	time := time.Now()
	order.CreatedAt = &time

	return it.db.Store(order)
}

func (it *usercase) Checkout(order *model.Order) error {

	// order.ID = utils.UUID()
	// user := model.User{
	// 	ID: order.UserID,
	// }
	// if err := it.userCache.FindBy("ID", &user); err != nil {

	// 	return model.ErrUserNotFound
	// }

	// // Send to /api/v1/tgc/transaction/game/endround
	// _url := fmt.Sprintf("%s%s%s", agent_domain, agent_api_root, "/transaction/game/endround")

	// req := url.Values{
	// 	"account":  {user.Username},
	// 	"gamehall": {service_id},
	// 	"gamecode": {order.GameID},
	// 	"roundid":  {order.ID},
	// 	"amount":   {strconv.FormatUint(order.Bet, 10)},
	// 	"mtcode":   {utils.UUID()},
	// }

	// type Res struct {
	// 	Data map[string]interface{} `json:"data"`

	// 	Status struct {
	// 		Code    string    `json:"code"`
	// 		Message string    `json:"message"`
	// 		Time    time.Time `json:"datatime"`
	// 	} `json:"status"`
	// }

	// res := &Res{}

	// if err := post(_url, req, res); err != nil {

	// 	return errs.WithMessagef(err, "status: %+v", res.Status)
	// }

	order.State = model.Completed

	return it.db.Replace(order)
}

func post(url string, req url.Values, res interface{}) error {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(req.Encode()))
	if err != nil {
		log.Errorf("Error: %s\n", err.Error())

		return err
	}

	// TODO: request.Header.Set("organization_token" , value)

	resp, err := client.Do(request)
	if err != nil {
		log.Errorf("Error: %s\n", err.Error())

		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		log.Errorf("Error: %s\n", err.Error())

		return fmt.Errorf("Can't deserialize response: %s", url)
	}

	return nil
}
