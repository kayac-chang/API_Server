package utils

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func MD5(text string) string {

	hash := md5.Sum([]byte(text))

	return fmt.Sprintf("%x", hash)
}

func Hash(text string) string {

	hash := sha3.Sum256([]byte(text))

	return fmt.Sprintf("%x", hash)
}
