package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// DashboardHandler returns user dashboard details after authentication
func DashboardHandler(c *gin.Context) {
	// Try to get user ID from URL or JWT middleware
	id := c.Param("id")
	if id == "" {
		claims, _ := c.Get("claims")
		jwtData := claims.(map[string]interface{})
		id = jwtData["user_id"].(string)
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		log.Printf("Dashboard: User not found for ID %s", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"name":     user.Name,
			"email":    user.Email,
			"phone":    user.Phone,
			"profile":  user.ProfileExists,
		},
		"modules": gin.H{
			"messaging": true,
			"mail":      true,
			"trading":   true,
			"banking":   true,
		},
	})
}
