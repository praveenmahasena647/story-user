package middleware

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(p []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(p, 10)
}
