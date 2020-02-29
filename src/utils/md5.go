package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(text string) string {

	hash := md5.Sum([]byte(text))

	return fmt.Sprintf("%x", hash)
}
