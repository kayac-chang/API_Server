package utils

import (
	"io/ioutil"
	"log"
)

func Parse(filename string) string {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
