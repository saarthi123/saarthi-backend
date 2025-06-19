package middleware

import "github.com/gin-gonic/gin"

func RoleAuth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
