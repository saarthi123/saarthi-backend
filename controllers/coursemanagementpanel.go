package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

func GetInstructorCourses(c *gin.Context) {
	var courses []models.Course
	if err := config.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var updated models.Course

	courseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var existing models.Course
	if err := config.DB.First(&existing, courseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	if err := config.DB.Model(&existing).Updates(updated).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated"})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := config.DB.Delete(&models.Course{}, courseID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

type AnalyzeRequest struct {
	CourseID uint `json:"courseId"`
}

func AnalyzeCourseAI(c *gin.Context) {
	var req AnalyzeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Simulated AI response
	response := gin.H{
		"feedback": "This course could benefit from interactive projects and updated examples.",
	}
	c.JSON(http.StatusOK, response)
}
