// auth/jwt.go
package controllers

import (
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtKey = []byte("your-secret-key")

// func GenerateJWTs(phone string) (string, string, error) {
// 	claims := &jwt.StandardClaims{
// 		Subject:   phone,
// 		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	accessToken, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	refreshClaims := &jwt.StandardClaims{
// 		Subject:   phone,
// 		ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
// 	}
// 	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
// 	refreshString, err := refreshToken.SignedString(jwtKey)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return accessToken, refreshString, nil
// }


func ExtractUserIDFromToken(r *http.Request) (string, error) {
	tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	return claims.Subject, nil
}