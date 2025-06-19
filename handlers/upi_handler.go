package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
	"github.com/saarthi123/saarthi-backend/config"
)

// ChangeUpiPin handles UPI PIN change requests
func ChangeUpiPin(c *gin.Context) {
	var req struct {
		Phone   string `json:"phone"`
		OldPin  string `json:"old_pin"`
		NewPin  string `json:"new_pin"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	if err := config.DB.Where("phone = ?", req.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Simulated PIN check — replace with secure hash comparison in real app
	if user.UpiPin != req.OldPin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect old PIN"})
		return
	}

	user.UpiPin = req.NewPin // Note: Use hashed storage in production!
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update PIN"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UPI PIN changed successfully"})
}





func AddBank(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bank linked"})
}

func GetUpiHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"history": []string{}})
}
