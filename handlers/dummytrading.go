package handlers

import (
    "net/http"
    "github.com/saarthi123/saarthi-backend/models"
    "github.com/saarthi123/saarthi-backend/services"

    "github.com/gin-gonic/gin"
)

var tradingService = services.NewDummyTradingService()

func GetPortfolio(c *gin.Context) {
    portfolio := tradingService.GetPortfolio()
    balance := tradingService.GetBalance()
    c.JSON(http.StatusOK, gin.H{
        "balance":   balance,
        "portfolio": portfolio,
        "aiTip":     "Buy low, sell high! 🚀", // dummy AI tip
    })
}

func BuyStock(c *gin.Context) {
    var req models.TradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if req.Quantity <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be positive"})
        return
    }
    success := tradingService.Buy(req.Symbol, req.Quantity)
    if !success {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Stock bought"})
}

func SellStock(c *gin.Context) {
    var req models.TradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if req.Quantity <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be positive"})
        return
    }
    success := tradingService.Sell(req.Symbol, req.Quantity)
    if !success {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stocks to sell"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Stock sold"})
}

func ResetSimulation(c *gin.Context) {
    tradingService.Reset()
    c.JSON(http.StatusOK, gin.H{"message": "Simulation reset"})
}
