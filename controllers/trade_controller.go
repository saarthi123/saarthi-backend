package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// PlaceTrade handles creating a new trade
func PlaceTrade(c *gin.Context) {
	var trade models.Trade

	// Validate request
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save to DB
	if err := config.DB.Create(&trade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save trade",
		})
		return
	}

	// Respond with created trade
	c.JSON(http.StatusOK, gin.H{
		"message": "Trade placed successfully",
		"trade":   trade,
	})
}

// GetTrades returns all trades
func GetTrades(c *gin.Context) {
	var trades []models.Trade

	// Fetch from DB
	if err := config.DB.Find(&trades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve trades",
		})
		return
	}

	// Respond with trade list
	c.JSON(http.StatusOK, gin.H{
		"trades": trades,
	})
}
