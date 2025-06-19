package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/saarthi123/saarthi-backend/auth"
)

// JWTAuth middleware verifies JWT and sets `user_phone` into Gin context
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			abortUnauthorized(c, "Missing or malformed Authorization header")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			abortUnauthorized(c, "JWT_SECRET not configured in environment")
			return
		}

		// Parse token with auth.CustomClaims
		token, err := jwt.ParseWithClaims(tokenStr, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			abortUnauthorized(c, "Invalid or expired token")
			return
		}

		claims, ok := token.Claims.(*auth.CustomClaims)
		if !ok || claims.Phone == "" {
			abortUnauthorized(c, "Invalid token claims")
			return
		}

		// Set the phone number in context for use in handlers
		c.Set("user_phone", claims.Phone)
		c.Next()
	}
}

// abortUnauthorized returns a standardized 401 error response
func abortUnauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": msg,
	})
}
