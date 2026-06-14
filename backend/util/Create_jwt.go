package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Sub   int64  `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`

	jwt.RegisteredClaims
}

func CreateJWT(payload Payload, secretKey string) (string, error) {
	payload.RegisteredClaims = jwt.RegisteredClaims{
		Subject:   string(rune(payload.Sub)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secretKey))
}
