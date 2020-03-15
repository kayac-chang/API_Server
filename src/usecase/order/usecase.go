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

	order.State = model.Pending

	time := time.Now()
	order.CreatedAt = &time

	return it.db.Store(order)
}

func (it *usercase) Checkout(order *model.Order) error {

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
