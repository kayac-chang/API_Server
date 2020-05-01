package agent

import (
	"api/env"
	"api/model"
	"api/utils"
	"api/utils/json"
	"fmt"
	"net/http"
	"time"

	"log"
)

// Agent ...
type Agent struct {
	Domain string
	API    string
	Token  string
}

// Bet ...
type Bet struct {
	Roundid string

	Username string
	Gamename string

	Amount  float64
	Session string

	CreatedAt time.Time
}

// New ...
func New(env env.Env) Agent {

	return Agent{
		Domain: env.Agent.Domain,
		API:    env.Agent.API,
		Token:  env.Agent.Token,
	}
}

// CheckPlayer ...
func (it Agent) CheckPlayer(username string, session string) (float64, error) {

	api := "/player/check/" + username

	url := it.Domain + it.API + api

	headers := map[string]string{
		"Content-Type":       "application/json",
		"organization-token": it.Token,
		"session":            session,
	}

	log.Printf("Agent [ %s ] Post\n header: %s", api, json.Jsonify(headers))

	resp, err := utils.Post(url, nil, headers)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	res := map[string]interface{}{}
	json.Parse(resp.Body, &res)

	if resp.StatusCode != 200 {
		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

		msg, err := getErrorMsg(res)
		if err != nil {
			return 0, &model.Error{
				Code:    resp.StatusCode,
				Message: err.Error(),
			}
		}

		return 0, &model.Error{
			Code:    resp.StatusCode,
			Message: msg,
		}
	}

	log.Printf("Agent [ %s ] Success\nResponse:\n %s", api, json.Jsonify(res))

	return getBalance(res)
}

// SendBet ...
func (it Agent) SendBet(bet Bet) (float64, error) {

	api := "/transaction/game/bet"

	url := it.Domain + it.API + api

	req := map[string]interface{}{
		"account":    bet.Username,
		"created_at": bet.CreatedAt.String(),
		"gamename":   bet.Gamename,
		"roundid":    bet.Roundid,
		"amount":     bet.Amount,
	}

	headers := map[string]string{
		"Content-Type":       "application/json",
		"organization-token": it.Token,
		"session":            bet.Session,
	}

	resp, err := utils.Post(url, req, headers)
	if err != nil {

		return 0, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	defer resp.Body.Close()

	res := map[string]interface{}{}
	json.Parse(resp.Body, &res)

	if resp.StatusCode != 200 {
		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

		msg, err := getErrorMsg(res)
		if err != nil {
			return 0, &model.Error{
				Code:    resp.StatusCode,
				Message: err.Error(),
			}
		}

		return 0, &model.Error{
			Code:    resp.StatusCode,
			Message: msg,
		}
	}

	log.Printf("Agent: [ %s ] Success !!!\nResponse:\n %s", api, json.Jsonify(res))

	return getBalance(res)
}

// SendEndRound ...
func (it Agent) SendEndRound(bet Bet) (float64, error) {

	api := "/transaction/game/endround"

	url := it.Domain + it.API + api

	req := map[string]interface{}{
		"account":      bet.Username,
		"gamename":     bet.Gamename,
		"roundid":      bet.Roundid,
		"amount":       bet.Amount,
		"completed_at": bet.CreatedAt.String(),
	}

	headers := map[string]string{
		"Content-Type":       "application/json",
		"organization-token": it.Token,
		"session":            bet.Session,
	}

	resp, err := utils.Post(url, req, headers)
	if err != nil {

		return 0, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	defer resp.Body.Close()

	res := map[string]interface{}{}
	json.Parse(resp.Body, &res)

	if resp.StatusCode != 200 {
		log.Printf("Agent: [ %s ] Failed...\n Error:\n %s", api, json.Jsonify(res))

		msg, err := getErrorMsg(res)
		if err != nil {
			return 0, &model.Error{
				Code:    resp.StatusCode,
				Message: err.Error(),
			}
		}

		return 0, &model.Error{
			Code:    resp.StatusCode,
			Message: msg,
		}
	}

	log.Printf("Agent: [ %s ] Success !!!\nResponse:\n %s", api, json.Jsonify(res))

	return getBalance(res)
}

func getErrorMsg(res map[string]interface{}) (string, error) {

	data, ok := res["error"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Unexpected response structure from agent: %+v", res)
	}

	msg, ok := data["message"].(string)
	if !ok {
		return "", fmt.Errorf("Unexpected response structure from agent: %+v", res)
	}

	return msg, nil
}

func getBalance(res map[string]interface{}) (float64, error) {

	data, ok := res["data"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Unexpected response structure from agent: %+v", res)
	}

	balance, ok := data["balance"].(float64)
	if !ok {
		return 0, fmt.Errorf("Unexpected response structure from agent: %+v", res)
	}

	return balance, nil
}
