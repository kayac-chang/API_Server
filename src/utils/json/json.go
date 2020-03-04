package json

import (
	"bytes"
	"encoding/json"
	"io"
)

func Parse(reader io.Reader, obj interface{}) {

	err := json.NewDecoder(reader).Decode(&obj)

	if err != nil {
		panic(err)
	}
}

func Stringify(value interface{}) string {

	bytes, err := json.Marshal(value)

	if err != nil {
		panic(err)
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
