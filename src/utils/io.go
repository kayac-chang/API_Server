package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ParseFile(filename string) string {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func Fetch(url string, res interface{}) error {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())

		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(res)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())

		return fmt.Errorf("Can't deserialize response: %s", url)
	}

	return nil
}
