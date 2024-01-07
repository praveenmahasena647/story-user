package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), 10)
}

var PasswordErr = func(hashed []byte, p string) error { // I used to do YAVASCRIPT
	return bcrypt.CompareHashAndPassword(hashed, []byte(p))
}
