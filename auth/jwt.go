package auth

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}
