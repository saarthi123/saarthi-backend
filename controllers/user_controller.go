package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Request body for user creation
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

// Helper function to generate a unique email from name
func generateEmail(name string) string {
	base := strings.ToLower(strings.ReplaceAll(name, " ", ""))
	randomNum := rand.Intn(10000)
	return fmt.Sprintf("%s%d@saarthi.com", base, randomNum)
}

// 🔹 CreateUser - Handles new user profile creation
// 🔹 CreateUser - Handles new user profile creation
func CreateUser(c *gin.Context) {
	db := config.DB  // global DB variable

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and phone required"})
		return
	}

	// Auto-generate email
	email := generateEmail(req.Name)

	// Check if user with same phone or email already exists
	var existingUser models.User
	if err := db.Where("phone_number = ? OR email = ?", req.Phone, email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone or email already exists"})
		return
	}

	// Create the new user
	user := models.User{
		Name:        req.Name,
		Phone: req.Phone,
		Email:       email,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"phone": user.Phone,
			"email": user.Email,
		},
	})
}


// 🔹 GetUserHandler - Fetch user by ID
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// GetProfile returns user info from phone in JWT
func GetProfile(c *gin.Context) {
	phone, _ := c.Get("user_phone")

	var user models.User
	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateProfile allows a user to change their name, email, etc.
func UpdateProfile(c *gin.Context) {
	phone, _ := c.Get("user_phone")

	var updateData struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Name = updateData.Name
	user.Email = updateData.Email

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated", "user": user})
}
// DeleteProfile allows// ChangePassword lets a user change their password (assuming password field exists in DB)
func ChangePassword(c *gin.Context) {
	phone, _ := c.Get("user_phone")

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// For simplicity, plain-text password match (should hash & compare securely in production)
	if user.Password != input.OldPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect old password"})
		return
	}

	user.Password = input.NewPassword
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}


func SecureHandler(c *gin.Context) {
	phone, _ := c.Get("user_phone")
	c.JSON(http.StatusOK, gin.H{
		"message": "Secure access granted",
		"user":    phone,
	})
}