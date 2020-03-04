package utils

import (
	"io/ioutil"
)

func ParseFile(filename string) string {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}
