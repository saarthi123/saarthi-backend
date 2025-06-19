package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

func GetDashboardData(c *gin.Context) {
	db := models.DB
	userID := c.Param("id")

	// Fetch user
	var user models.User
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch diplomas in progress
	var diplomas []models.Diploma
	db.Where("user_id = ?", userID).Find(&diplomas)

	// Fetch upcoming classes
	var classes []models.UpcomingClass
	db.Where("user_id = ?", userID).Order("date ASC").Find(&classes)

	// Generate dashboard response
	dashboard := models.Dashboard{
		UserName:           user.Username,
		DiplomasInProgress: diplomas,
		UpcomingClasses:    classes,
		AICareerSuggestion: "Take “Advanced Data Structures” to boost your Software Engineer path.", // Replace with AI engine logic
		AIStyleSuggestion:  "We suggest more video-rich courses for better engagement.",             // Replace with AI engine logic
		LastCourse:         diplomas[len(diplomas)-1].Title, // Last added diploma
		RecommendedPaths:   []string{"AI + Public Policy", "Entrepreneurship + Environmental Science", "Finance + Blockchain"}, // Replace with real logic
	}

	c.JSON(http.StatusOK, gin.H{"dashboard": dashboard})
}

func ContinueCourse(c *gin.Context) {
	db := models.DB
	userID := c.Param("id")

	// Optional: validate user exists
	var user models.User
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// TODO: Logic to update current course progress
	// Example: db.Model(&models.Diploma{}).Where("user_id = ? AND status = ?", userID, "in_progress").Update(...)

	c.JSON(http.StatusOK, gin.H{"message": "Course continued for user " + userID})
}
