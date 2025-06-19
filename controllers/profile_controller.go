package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

func CreateProfileController(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if phone/email already exists
	var existing models.User
	if err := config.DB.Where("phone_number = ? OR email = ?", user.Phone, user.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this phone or email already exists"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Profile created successfully",
		"user":    user,
	})
}

// In controllers/user.go or wherever suitable


// CheckProfileHandler checks if a user exists by phone number
func CheckProfileHandler(c *gin.Context) {
	phone := c.Param("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	exists, err := config.CheckUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
