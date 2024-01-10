package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	cost := 14

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(hash), err
}

func VerifyHash(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
