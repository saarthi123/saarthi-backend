package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// GET /courses?searchTerm=go&category=Programming
func GetCourses(c *gin.Context) {
	search := strings.ToLower(c.Query("searchTerm"))
	category := c.Query("category")

	var courses []models.Course

	query := config.DB
	if search != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+search+"%")
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

type EnrollmentRequest struct {
	UserID      string `json:"userId"`
	CourseTitle string `json:"courseTitle"`
}

// POST /enroll
func EnrollCourse(c *gin.Context) {
	var req EnrollmentRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.CourseTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment data"})
		return
	}

	// Log or store enrollment (future DB insert can go here)
	c.JSON(http.StatusOK, gin.H{
		"message": "Enrolled successfully",
		"course":  req.CourseTitle,
	})
}
