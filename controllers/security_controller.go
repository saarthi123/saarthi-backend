package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PinChangeRequest struct {
	CurrentPin string `json:"current"`
	NewPin     string `json:"new"`
}

type ToggleRequest struct {
	Type  string `json:"type"`
	Value bool   `json:"value"`
}

// ───── Security Overview ─────────────────────────────────────

func GetSecurityStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "secure", "encryption": "AES-256", "auth": "2FA enabled"})
}

func GetDeviceActivity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"devices": []map[string]string{
			{"name": "iPhone 14", "location": "Delhi, India", "lastActive": "2025-06-15 19:45"},
			{"name": "MacBook Pro", "location": "AWS EC2, Mumbai", "lastActive": "2025-06-15 08:21"},
		},
	})
}

// ───── PIN & Alerts ──────────────────────────────────────────

func ChangePin(c *gin.Context) {
	var req PinChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.CurrentPin == "" || req.NewPin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PIN data"})
		return
	}

	// Simulate PIN validation (should be securely checked in DB)
	if req.CurrentPin != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Current PIN incorrect"})
		return
	}

	// Success (in real use: update DB)
	c.JSON(http.StatusOK, gin.H{"message": "PIN successfully changed"})
}

func LogoutDevice(c *gin.Context) {
	deviceID := c.Query("deviceId")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device ID required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Device " + deviceID + " logged out"})
}

func ToggleAlert(c *gin.Context) {
	var req ToggleRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid toggle data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Alert setting updated",
		"alertType":      req.Type,
		"alertsEnabled":  req.Value,
	})
}
