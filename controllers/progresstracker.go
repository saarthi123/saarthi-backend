package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

// GET /progress?course=AI&completionStatus=completed&minCompletion=40
func GetProgress(c *gin.Context) {
	course := c.Query("course")
	status := c.Query("completionStatus")
	minStr := c.Query("minCompletion")

	var minCompletion float64
	if minStr != "" {
		if m, err := strconv.ParseFloat(minStr, 64); err == nil {
			minCompletion = m
		}
	}

	var progresses []models.StudentProgress
	query := models.DB.Model(&models.StudentProgress{})

	if course != "" {
		query = query.Where("course = ?", course)
	}
	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}
	if minStr != "" {
		query = query.Where("completion >= ?", minCompletion)
	}

	if err := query.Find(&progresses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch student progress data"})
		return
	}

c.JSON(http.StatusOK, gin.H{"data": progresses})
}

// GET /progress/:studentId
func GetStudentProgress(c *gin.Context) {
	idStr := c.Param("studentId")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var progress models.StudentProgress
	if err := models.DB.First(&progress, "student_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student progress not found"})
		return
	}

c.JSON(http.StatusOK, gin.H{"progress": progress})
}
