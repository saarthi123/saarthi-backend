// 📁 controllers/attendance_controller.go
package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// GET /attendance?user_id=xyz
func GetAttendance(c *gin.Context) {
	userID := c.Query("user_id")
	var records []models.AttendanceRecord

	if err := config.DB.Where("user_id = ?", userID).Order("date desc").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": records}) // ✅ safer format
}

// POST /attendance
func MarkAttendance(c *gin.Context) {
	var input struct {
		UserID string `json:"user_id" binding:"required"`
		Mode   string `json:"mode" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	record := models.AttendanceRecord{
		UserID: input.UserID,
		Date:   time.Now(),
		Status: "Present",
		Mode:   input.Mode,
	}

	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"record": record}) // ✅ wrapped in gin.H
}
