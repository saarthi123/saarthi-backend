package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
	"github.com/saarthi123/saarthi-backend/utils"
)

// VerifyOtpHandler validates OTP and returns auth token
func VerifyOtpHandler(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone" binding:"required"`
		OTP         string `json:"otp" binding:"required"`
		OtpToken    string `json:"otpToken" binding:"required"`
	}

	// Log raw request
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	fmt.Println("Request Body:", string(body))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone, OTP, and OTP token are required"})
		return
	}

	// Verify OTP token and values
	claims, err := utils.VerifyOtpToken(request.OtpToken)
	if err != nil || claims.PhoneNumber != request.PhoneNumber || claims.Otp != request.OTP {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	// Check if user exists
	var user models.User
	result := config.DB.Where("phone_number = ?", request.PhoneNumber).First(&user)

	// Create JWT token
	authToken, err := utils.CreateJWTToken(request.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate auth token"})
		return
	}

	if result.RowsAffected == 0 {
		// No profile yet
		c.JSON(http.StatusOK, gin.H{
			"message":      "OTP verified. Please complete your profile.",
			"existingUser": false,
			"token":        authToken,
		})
		return
	}

	// Return success login
	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"existingUser": true,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"phone": user.Phone,
			"email": user.Email,
		},
		"token": authToken,
	})
}
