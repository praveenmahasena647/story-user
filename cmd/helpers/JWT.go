package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(e string) (string, error) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"e":   e,
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})
	return token.SignedString([]byte("yayayaa"))
}

func DecodeJWT(j string) (string, error) {
	var emailID string
	var token, err = jwt.Parse(j, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("its a Eoorror")
		}
		return []byte("yayayaa"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) || claims["exp"] == nil {
			return emailID, errors.New("its A Err ")
		}
		emailID = claims["e"].(string)
	}
	return emailID, err
}
