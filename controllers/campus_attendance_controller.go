// controllers/campus_attendance_controller.go
package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
  "github.com/saarthi123/saarthi-backend/config"
)

func GetCampusAttendance(c *gin.Context) {
	dateStr := c.Query("date")
	course := c.Query("course")

	if dateStr == "" || course == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date and course are required"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	var records []models.CampusAttendance
	if err := config.DB.
		Where("attend_date = ? AND course = ?", date, course).
		Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance"})
		return
	}

c.JSON(http.StatusOK, gin.H{"data": records})
}
