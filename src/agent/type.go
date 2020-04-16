package agent

import (
	"api/env"
	"api/utils"
	"api/utils/json"

	"fmt"
	"log"
)

type Agent struct {
	Domain string
	API    string
	Token  string
}

func New(env env.Env) Agent {
	return Agent{
		Domain: env.Agent.Domain,
		API:    env.Agent.API,
		Token:  env.Agent.Token,
	}
}

func (it Agent) CheckPlayer(username string, session string) (uint64, error) {

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
		log.Printf("Agent [ %s ] Failed\n Error:\n %s", api, json.Jsonify(res))

		return 0, fmt.Errorf("%s", res["message"])
	}

	log.Printf("Agent [ %s ] Success\nResponse:\n %s", api, json.Jsonify(res))

	data := res["data"].(map[string]interface{})
	amount := data["balance"].(float64)

	return uint64(amount), nil
}
