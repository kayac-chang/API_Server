package usecase

import (
	"api/env"
	"api/model"
	"api/user/repo"
	"api/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
)

type usercase struct {
	env   *env.Env
	db    repo.Repository
	cache repo.Repository
}

type Usecase interface {
	Find(user *model.User) error
	Store(user *model.User) error
	Regist(user *model.User) error
	Sign(user *model.User) (*model.Token, error)
	Auth(user *model.User) error
}

func New(env *env.Env, db, cache repo.Repository) Usecase {

	return &usercase{env, db, cache}
}

func fetch(url string, res interface{}) error {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)

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

func (it *usercase) Auth(user *model.User) error {

	return it.cache.FindBy("Token", user)
}

func (it *usercase) Regist(user *model.User) error {

	// Send to /api/v1/tgc/player/check/:account
	// url := fmt.Sprintf("%s%s%s%s",
	// 	it.env.Agent.Domain,
	// 	it.env.Agent.API,
	// 	"/player/check/",
	// 	user.Username,
	// )

	// type Res struct {
	// 	Balance struct {
	// 		Balance  float64 `json:"balance"`
	// 		Currency string  `json:"currency"`
	// 	} `json:"balance"`

	// 	Status struct {
	// 		Code    string    `json:"code"`
	// 		Message string    `json:"message"`
	// 		Time    time.Time `json:"datatime"`
	// 	} `json:"status"`
	// }

	// res := &Res{}
	// if err := fetch(url, res); err != nil {

	// 	return errs.WithMessagef(err, "status: %+v", res.Status)
	// }

	if err := it.Find(user); err != nil {

		if err != nil && err != model.ErrUserNotFound {
			return err
		}

		if err = it.Store(user); err != nil {
			return err
		}
	}

	// TODO: Transform Balance into game coin
	user.Balance = 600270

	return nil
}

func (it *usercase) Find(user *model.User) error {

	user.ID = utils.MD5(user.Username)

	return it.db.FindBy("ID", user)
}

func (it *usercase) Store(user *model.User) error {

	user.ID = utils.MD5(user.Username)
	user.CreatedAt = time.Now()

	return it.db.Store(user)
}

func (it *usercase) Sign(user *model.User) (*model.Token, error) {

	if err := it.cache.FindBy("ID", user); err == nil {
		log.Printf("found user in cache, remove it\n")

		it.cache.Remove(user)
	}

	createdTime := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": it.env.Service.ID,
		"iat": createdTime.Unix(),
		"jti": utils.UUID(),
	})

	tokenString, err := token.SignedString(it.env.Secret)
	if err != nil {
		return nil, err
	}

	user.Token = tokenString
	it.cache.Store(user)

	res := &model.Token{
		AccessToken: tokenString,
		Type:        "Bearer",
		ServiceID:   it.env.Service.ID,
		CreatedAt:   createdTime,
	}

	return res, nil
}
