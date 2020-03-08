package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user/model"
	"user/repo"
	"user/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	errs "github.com/pkg/errors"
)

var hmacSampleSecret = []byte("my_secret_key")
var service_id = "service"

var agent_domain = "http://localhost:3000"
var agent_api_root = "/api/v1/tgc"

type usercase struct {
	db    repo.Repository
	cache repo.Repository
}

type Usecase interface {
	Find(user *model.User) error
	Store(user *model.User) error
	Regist(user *model.User) error
	Sign(user *model.User) (*model.Token, error)
}

func New(db, cache repo.Repository) Usecase {

	return &usercase{db, cache}
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

func (it *usercase) Regist(user *model.User) error {

	// Send to /api/v1/tgc/player/check/:account
	url := fmt.Sprintf("%s%s%s%s", agent_domain, agent_api_root, "/player/check/", user.Username)

	type Res struct {
		Balance struct {
			Balance  float64 `json:"balance"`
			Currency string  `json:"currency"`
		} `json:"balance"`

		Status struct {
			Code    string    `json:"code"`
			Message string    `json:"message"`
			Time    time.Time `json:"datatime"`
		} `json:"status"`
	}

	res := &Res{}

	if err := fetch(url, res); err != nil {

		return errs.WithMessagef(err, "status: %+v", res.Status)
	}

	if err := it.Find(user); err != nil {

		if err != nil && err != model.ErrUserNotFound {
			return err
		}

		if err = it.Store(user); err != nil {
			return err
		}
	}

	// TODO: Transform Balance into game coin
	user.Balance = res.Balance.Balance

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

	err := it.cache.FindBy("ID", user)

	if err == nil {

		fmt.Println("found user in cache, remove it")

		it.cache.Delete(user)
	}

	createdTime := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": service_id,
		"iat": createdTime.Unix(),
		"jti": utils.UUID(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {

		return nil, err
	}

	user.Token = tokenString
	it.cache.Store(user)

	res := &model.Token{
		AccessToken: tokenString,
		Type:        "Bearer",
		ServiceID:   service_id,
		CreatedAt:   createdTime,
	}

	return res, nil
}
