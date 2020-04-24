package utils

import (
	"crypto/md5"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func UUID() string {

	uuid := uuid.NewV4()

	return uuid.String()
}

func MD5(text string) string {

	hash := md5.Sum([]byte(text))

	return fmt.Sprintf("%x", hash)
}

func Hash(text string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func CompareHash(hash string, compare string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(compare))

	if err != nil {
		return err
	}

	return nil
}
