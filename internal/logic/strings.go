package logic

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func StringIfElse(condition bool, onTrue, onFalse string) string {
	if condition {
		return onTrue
	}
	return onFalse
}

func HashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func RandomString(length int) (string, error) {
	data := make([]byte, length)
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(data), nil
}
