package utils

import (
	"encoding/json"
	"io"
	"reflect"
)

func Parse(src interface{}, data interface{}) (err error) {

	switch _src := src.(type) {

	case io.Reader:
		err = parseFromReader(_src, data)

	}

	if err != nil {
		return err
	}

	return nil
}

func parseFromReader(src io.Reader, data interface{}) error {

	err := json.NewDecoder(src).Decode(&data)

	if err != nil {
		return err
	}

	return nil
}

func getFieldNames(data interface{}) []string {

	t := reflect.TypeOf(data)

	r := make([]string, t.NumField())

	for i := range r {

		f := t.Field(i)

		r = append(r, f.Name)
	}

	return r
}
