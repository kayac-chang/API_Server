package utils

import (
	"bytes"
	"encoding/json"
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

func Stringify(value interface{}) string {

	bytes, err := json.Marshal(value)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func Jsonify(obj interface{}) string {

	src, err := json.Marshal(obj)

	if err != nil {
		panic(err)
	}

	dst := &bytes.Buffer{}

	err = json.Indent(dst, src, "", "  ")

	if err != nil {
		panic(err)
	}

	return dst.String()
}
