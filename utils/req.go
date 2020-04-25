package utils

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Request struct {
	URL     string
	Header  map[string]string
	Context context.Context
}

func (it Request) Send() *http.Response {

	req, err := http.NewRequest("GET", it.URL, nil)

	if err != nil {
		log.Fatal(err)
	}

	for key, val := range it.Header {
		req.Header.Set(key, val)
	}

	if it.Context != nil {
		req = req.WithContext(it.Context)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
