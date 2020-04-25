package utils

import (
	"strconv"
	"time"
)

func Number(str string) int {

	num, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		panic(err)
	}

	return int(num)
}

func StrToTime(str string, unit time.Duration) time.Duration {

	return time.Duration(Number(str)) * unit
}
