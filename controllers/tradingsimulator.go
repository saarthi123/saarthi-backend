package controllers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

var (
	userData = models.User{
		ID:      1,
		Balance: 100000,
		Trades:  []models.Trade{},
	}
	mutex = &sync.Mutex{}
)

// GetUserData returns the current user data including balance and trades
func GetUserData(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"user": userData, // ✅ wrap userData in a map
	})
}

// SimulateTrade handles trade simulation (buy action)
func SimulateTrade(c *gin.Context) {
	type TradeInput struct {
		Symbol   string  `json:"symbol" binding:"required"`
		Quantity int     `json:"quantity" binding:"required,gt=0"`
		Price    float64 `json:"price" binding:"required,gt=0"`
	}

	var input TradeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalCost := float64(input.Quantity) * input.Price

	mutex.Lock()
	defer mutex.Unlock()

	if totalCost > userData.Balance {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	// Deduct balance
	userData.Balance -= totalCost

	// Add trade
	trade := models.Trade{
		ID:        len(userData.Trades) + 1,
		Symbol:    input.Symbol,
		Quantity:  input.Quantity,
		Price:     input.Price,
		Timestamp: time.Now(),
	}
	userData.Trades = append(userData.Trades, trade)

	// Simulate score calculation - placeholder
	score := 80 + (time.Now().Nanosecond() % 20)

	c.JSON(http.StatusOK, gin.H{
		"message": "Trade simulated successfully",
		"balance": userData.Balance,
		"trade":   trade,
		"score":   score,
	})
}
