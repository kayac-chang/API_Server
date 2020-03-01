package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func Hash(text string) string {

	hash := sha3.Sum256([]byte(text))

	return fmt.Sprintf("%x", hash)
}

func HashAndSalt(text string) string {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(text),
		bcrypt.MinCost,
	)

	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func Compare(hash string, plain string) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(plain),
	)

	if err != nil {
		return false
	}

	return true
}
