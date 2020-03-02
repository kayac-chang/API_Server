package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/KayacChang/API_Server/system/env"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
)

func UUID(name string) string {

	uuid := uuid.NewV5(env.DomainKey(), name)

	return uuid.String()
}

func MD5(text string) string {

	hash := md5.Sum([]byte(text))

	return fmt.Sprintf("%x", hash)
}

func Hash(text string) string {

	hash := sha3.Sum256([]byte(text))

	return fmt.Sprintf("%x", hash)
}
