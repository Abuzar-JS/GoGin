package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "supersecret"

func GenerateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"nbf":   time.Date(2024, 01, 01, 12, 0, 0, 0, time.UTC).Unix()})

	tokenStr, err := token.SignedString([]byte(secret))

	return tokenStr, err
}
