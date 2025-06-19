package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

func GetNotifications(c *gin.Context) {
	// You can replace this with DB-based logic later
	notifications := []models.Notification{
		{
			Type:      "info",
			Title:     "Monthly Statement Ready",
			Message:   "Your May statement is available for download.",
			Timestamp: time.Now().Add(-time.Hour * 2),
		},
		{
			Type:      "warning",
			Title:     "Low Balance Alert",
			Message:   "Your balance has dropped below ₹500.",
			Timestamp: time.Now().Add(-time.Hour * 5),
		},
		{
			Type:      "success",
			Title:     "Loan EMI Paid",
			Message:   "Your EMI for May was successfully paid.",
			Timestamp: time.Now().Add(-time.Hour * 10),
		},
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications}) 
}
