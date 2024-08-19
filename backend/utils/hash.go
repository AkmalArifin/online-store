package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, retrievedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(retrievedPassword), []byte(password))

	return err == nil
}
