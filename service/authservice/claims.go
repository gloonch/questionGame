package authservice

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func (c Claims) Valid() error {
	return c.StandardClaims.Valid()
}
