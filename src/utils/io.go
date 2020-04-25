package utils

import (
	"api/model"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	errs "github.com/pkg/errors"
)

func ParseFile(filename string) string {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func Fetch(url string) (map[string]interface{}, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	res := map[string]interface{}{}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("Error: %s\n", err.Error())

		return nil, err
	}

	return res, nil
}

func Post(url string, body interface{}, headers map[string]string) (*http.Response, error) {

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())

		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))

	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return client.Do(req)
}

func ParseJSON(body io.ReadCloser) (map[string]interface{}, error) {

	req := map[string]interface{}{}

	if err := json.NewDecoder(body).Decode(&req); err != nil {

		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: errs.WithMessage(err, "Error occured when parsing payload").Error(),
		}
	}

	return req, nil
}
