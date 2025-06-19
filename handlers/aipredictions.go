package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/saarthi123/saarthi-backend/models"
)

// This would ideally call some AI service or ML model.
// For now, returning dummy static data.
func GetPredictions(c *gin.Context) {
    predictions := []models.StockPrediction{
        {Symbol: "AAPL", Prediction: "📈 Bullish", Confidence: 89},
        {Symbol: "TSLA", Prediction: "📉 Bearish", Confidence: 73},
        {Symbol: "GOOGL", Prediction: "📈 Slight Bullish", Confidence: 81},
    }

    c.JSON(http.StatusOK, gin.H{"predictions": predictions})
}


type Transaction struct {
    ID     int
    Tag    string
    Date   string
    Stock  string
    Qty    int
    Price  int
    Amount int
}

var transactions = []Transaction{
    {ID: 1, Tag: "buy", Date: "2025-06-01", Stock: "RELIANCE", Qty: 10, Price: 2700, Amount: 27000},
    {ID: 2, Tag: "sell", Date: "2025-06-02", Stock: "TCS", Qty: 5, Price: 3500, Amount: 17500},
    {ID: 3, Tag: "buy", Date: "2025-06-03", Stock: "INFY", Qty: 15, Price: 1600, Amount: 24000},
}
