package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
    "github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
	"github.com/saarthi123/saarthi-backend/utils"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}



// SendOtpHandler returns OTP (via SMS) and a JWT token with OTP embedded
func SendOtpHandler(c *gin.Context) {
	var request struct {
		Phone string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	// Generate OTP
	otp := generateOTP()

	// Send OTP via SMS
	if err := utils.SendOtpSMS(request.Phone, otp); err != nil {
		log.Printf("Error sending OTP SMS: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	// Store OTP in DB with expiry
	expiry := time.Now().Add(10 * time.Minute)
	if err := config.DB.Create(&models.Otp{
		Phone:     request.Phone,
		Code:       otp,
		ExpiresAt: expiry,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store OTP"})
		return
	}

	// Generate JWT token that includes OTP and expiry
	_, otpToken, err := utils.GenerateOtpToken(request.Phone)
	if err != nil {
		log.Printf("Error creating OTP token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP token"})
		return
	}

	// Check if user already exists
	var user models.User
	result := config.DB.Where("phone = ?", request.Phone).First(&user)
	exists := result.RowsAffected > 0

	log.Println("Generated OTP:", otp)

	c.JSON(http.StatusOK, gin.H{
		"message":      "OTP sent successfully",
		"otpToken":     otpToken,
		"existingUser": exists,
	})
}

// VerifyOtpHandler verifies the JWT token + OTP and logs user in or prompts profile creation
// VerifyOtpHandler validates OTP from PostgreSQL, deletes OTP, and logs in or creates token for profile creation
func VerifyOtpHandler(c *gin.Context) {
    var request struct {
        PhoneNumber string `json:"phone" binding:"required"`
        OTP         string `json:"otp" binding:"required"`
        OtpToken    string `json:"otpToken" binding:"required"`
    }
    fmt.Println("Received OTP verification request:", request.PhoneNumber, request.OTP)
fmt.Println("Received OTP token:", request.OtpToken)


    // Bind JSON and check required fields
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Phone, OTP, and otpToken are required"})
        return
    }

    // Validate OTP stored in DB
    var storedOtp models.Otp
    err := config.DB.Where("phone = ?", request.PhoneNumber).First(&storedOtp).Error
    if err != nil || storedOtp.Code != request.OTP || time.Now().After(storedOtp.ExpiresAt) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
        return
    }

    // Delete OTP after successful verification
    if err := config.DB.Delete(&storedOtp).Error; err != nil {
        log.Printf("Failed to delete OTP after verification: %v", err)
        // Not fatal, but log
    }

    // Check if user exists
    var user models.User
    result := config.DB.Where("phone_number = ?", request.PhoneNumber).First(&user)

    // Generate access token (3-day expiry) and refresh token (30 days)
    accessToken, err := utils.CreateJWTToken(request.PhoneNumber)
    if err != nil {
        log.Printf("Error generating access token: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    refreshToken, err := utils.CreateRefreshToken(request.PhoneNumber)
    if err != nil {
        log.Printf("Error generating refresh token: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
        return
    }

    // Set cookies for tokens (adjust domain and secure flags for production)
    c.SetCookie("accessToken", accessToken, 3600*24*3, "/", "localhost", false, true)      // 3 days expiry
    c.SetCookie("refreshToken", refreshToken, 3600*24*30, "/", "localhost", false, true)   // 30 days expiry

    if result.RowsAffected > 0 {
        // User exists — login success
        c.JSON(http.StatusOK, gin.H{
            "message": "Login successful",
            "user": gin.H{
                "id":    user.ID,
                "name":  user.Name,
                "phone": user.Phone,
                "email": user.Email,
            },
            "token":        accessToken,
            "refreshToken": refreshToken,
        })
        return
    }

    // User does NOT exist — prompt profile creation
    c.JSON(http.StatusOK, gin.H{
        "message":      "OTP verified. Please complete your profile.",
        "existingUser": false,
        "phone":        request.PhoneNumber,
        "token":        accessToken,
        "refreshToken": refreshToken,
    })

// Return auth token for frontend to use
c.SetCookie("accessToken", accessToken, 3600, "/", "localhost", false, true)
c.SetCookie("refreshToken", refreshToken, 3600*24*30, "/", "localhost", false, true)

}


var jwtKey = []byte("secret-key")

func Login(c *gin.Context) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // In production, verify password from DB
    if creds.Username != "test" || creds.Password != "1234" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": creds.Username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
