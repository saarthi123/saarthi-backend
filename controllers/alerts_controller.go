// controllers/alerts_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func UpdateNotificationPreferences(c *gin.Context) {
	var toggle ToggleRequest
	if err := c.ShouldBindJSON(&toggle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// You can handle the logic dynamically using the toggle.Type and toggle.Value
	// Example: store in DB or update session config
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification preference updated",
		"type":    toggle.Type,
		"value":   toggle.Value,
	})
}
