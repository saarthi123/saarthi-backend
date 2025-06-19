package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Mock user ID from context (assumes auth middleware)
func getUserID(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return fmt.Sprintf("%v", userID)
}

// ───── Portfolio ─────────────────────────────────────────────

func GetPortfolio(c *gin.Context) {
	userID := getUserID(c)

	portfolio := []gin.H{
		{"symbol": "TCS", "quantity": 10, "avgPrice": 3300.0, "currentPrice": 3450.0},
		{"symbol": "RELIANCE", "quantity": 5, "avgPrice": 2450.0, "currentPrice": 2480.0},
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":    userID,
		"portfolio": portfolio,
		"updatedAt": time.Now(),
	})
}

// ───── Exchange Info ─────────────────────────────────────────

func GetExchangeInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"exchange": "NSE",
		"location": "Mumbai, India",
		"status":   "Open",
		"opensAt":  "09:15",
		"closesAt": "15:30",
	})
}

// ───── Market News ───────────────────────────────────────────

func GetMarketNews(c *gin.Context) {
	news := []gin.H{
		{"headline": "Nifty crosses 20,000", "source": "Economic Times", "timestamp": time.Now().Add(-1 * time.Hour)},
		{"headline": "RBI keeps repo rate unchanged", "source": "Moneycontrol", "timestamp": time.Now().Add(-2 * time.Hour)},
	}

	c.JSON(http.StatusOK, gin.H{"news": news})
}

// ───── Trading History ───────────────────────────────────────

func GetTradingHistory(c *gin.Context) {
	userID := getUserID(c)

	history := []gin.H{
		{"symbol": "INFY", "action": "BUY", "quantity": 20, "price": 1460.0, "date": "2025-06-10"},
		{"symbol": "HDFC", "action": "SELL", "quantity": 10, "price": 2700.0, "date": "2025-06-12"},
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":  userID,
		"history": history,
	})
}

func ExportTradingHistory(c *gin.Context) {
	userID := getUserID(c)
	// Simulate AWS S3 signed export link
	s3ExportURL := fmt.Sprintf("https://s3.amazonaws.com/saarthi-trading-history/%s-history.csv?expires=1680000000", userID)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Trading history export link generated",
		"userId":      userID,
		"downloadUrl": s3ExportURL,
	})
}

// ───── Trading Simulation ───────────────────────────────────

func ResetSimulation(c *gin.Context) {
	userID := getUserID(c)
	// Simulate simulation data reset
	c.JSON(http.StatusOK, gin.H{
		"message":   "Simulation reset successful",
		"userId":    userID,
		"resetTime": time.Now(),
	})
}
