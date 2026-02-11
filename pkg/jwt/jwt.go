package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Jwt struct {
	secret string
}

func NewJwt(secret string) *Jwt {
	return &Jwt{secret: secret}
}

func (j Jwt) SigningString(id uuid.UUID) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id": id,
	})

	return t.SignedString([]byte(j.secret))
}

func (j Jwt) GetID(tokenStr string) (string, error) {
	t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("token parse fail")
	}

	if !t.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	if id, ok := t.Claims.(jwt.MapClaims)["id"]; ok {
		return id.(string), nil
	}

	return "", errors.New("token is invalid")
}
