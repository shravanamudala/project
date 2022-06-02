package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(passowrd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(passowrd), 16)
	return string(hash)
}
