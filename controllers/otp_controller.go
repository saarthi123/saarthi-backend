package controllers

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Request body for sending OTP
type SendOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
}

// Request body for verifying OTP
type VerifyOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

// OTPEntry stores OTP and expiration time
type OTPEntry struct {
	OTP       string
	ExpiresAt time.Time
}

var (
	otpStore = make(map[string]OTPEntry)
	otpMutex sync.Mutex
)

// GenerateOTP creates a random 6-digit OTP
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// StoreOTP stores OTP with a 5-minute expiry
func storeOTP(phone, otp string) {
	otpMutex.Lock()
	defer otpMutex.Unlock()
	otpStore[phone] = OTPEntry{
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

// validateOTP checks if OTP is correct and not expired
func validateOTP(phone, input string) bool {
	otpMutex.Lock()
	defer otpMutex.Unlock()

	entry, exists := otpStore[phone]
	if !exists || time.Now().After(entry.ExpiresAt) || entry.OTP != input {
		return false
	}

	delete(otpStore, phone)
	return true
}

// 🔹 SendOTP - Generates and returns OTP (logically sends SMS)
func SendOTP(c *gin.Context) {
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	otp := generateOTP()
	storeOTP(req.Phone, otp)

	// TODO: Integrate AWS SNS/Pinpoint to send `otp` via SMS to `req.Phone`

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully",
		"otp":     otp, // ⛔ Remove this in production
	})
}

// 🔹 VerifyOTP - Verifies submitted OTP
func VerifyOTP(c *gin.Context) {
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone and OTP are required"})
		return
	}

	if !validateOTP(req.Phone, req.OTP) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
