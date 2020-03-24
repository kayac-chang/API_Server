package user

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/model"
	"api/utils/json"
	"fmt"
	"log"

	repo "api/repo/user"
	"api/utils"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type Usecase struct {
	env  *env.Env
	repo *repo.Repo
}

func New(env *env.Env, db *postgres.DB, c *cache.Cache) *Usecase {

	return &Usecase{
		env:  env,
		repo: repo.New(db, c),
	}
}

func (it *Usecase) Regist(username string, session string) (*model.Token, error) {

	// Send to
	balance, err := it.sendToCheckPlayer(username, session)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:       utils.MD5(username),
		Username: username,
	}

	if err := it.repo.FindBy("ID", &user); err != nil {

		if err != nil && err != model.ErrUserNotFound {
			return nil, err
		}

		user.CreatedAt = time.Now()

		if err = it.repo.Store("DB", &user); err != nil {
			return nil, err
		}
	}

	// == Sign Token ==
	it.repo.RemoveCache(&user)

	token, err := it.sign(&user)
	if err != nil {
		return nil, err
	}

	user.Token = token.AccessToken
	user.Balance = balance
	user.Session = session

	it.repo.Store("Cache", &user)

	return token, nil
}

// === Private ===

func (it *Usecase) sign(user *model.User) (*model.Token, error) {

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

	res := model.Token{
		AccessToken: tokenString,
		Type:        "Bearer",
		ServiceID:   it.env.Service.ID,
		CreatedAt:   createdTime,
	}

	return &res, nil
}

func (it *Usecase) sendToCheckPlayer(username string, session string) (uint64, error) {
	api := "/player/check/" + username

	url := it.env.Agent.Domain + it.env.Agent.API + api

	req := map[string]interface{}{}

	headers := map[string]string{
		"Content-Type":       "application/json",
		"organization-token": it.env.Agent.Token,
		"session":            session,
	}

	resp, err := utils.Post(url, req, headers)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	res := map[string]interface{}{}
	json.Parse(resp.Body, &res)

	if resp.StatusCode != 200 {
		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

		err := res["error"].(map[string]interface{})

		return 0, fmt.Errorf("%s", err["message"])
	}

	log.Printf("Agent: [ %s ] Success !!!\nResponse:\n %s", api, json.Jsonify(res))

	data := res["data"].(map[string]interface{})
	balance := data["balance"].(map[string]interface{})
	amount := balance["balance"].(float64)

	return uint64(amount), nil
}
