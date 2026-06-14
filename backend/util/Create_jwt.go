package util

import (
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int64  `json:"sub"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`

	IssuedAt  int64 `json:"iat"`
	ExpiresAt int64 `json:"exp"`
}

func CreateJWT(payload Payload, secretKey string) (string, error) {

	claims := jwt.MapClaims{
		"sub":           payload.Sub,
		"name":          payload.Name,
		"email":         payload.Email,
		"is_shop_owner": payload.IsShopOwner,
		"iat":           payload.IssuedAt,
		"exp":           payload.ExpiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

/*
func CreateJWT(data Payload) (string, error) {
	head := Header{Alg: "HS256", Typ: "JWT"}
	headerBytes, err := json.Marshal(head)
	if err != nil {
		return "", err
	}

	data.Exp = time.Now().Add(24 * time.Hour).Unix()
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	header64 := baseURLEncode(headerBytes)
	payload64 := baseURLEncode(payloadBytes)

	message := header64 + "." + payload64
	cnf := config.GetConfig()
	secretBytes := []byte(cnf.JWTSecretKey)

	h := hmac.New(sha256.New, secretBytes)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	signature64 := baseURLEncode(signature)

	jwt := header64 + "." + payload64 + "." + signature64
	return jwt, nil
}

func baseURLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}*/
