package handlers

import (
    "net/http"
    "github.com/saarthi123/saarthi-backend/models"

    "github.com/gin-gonic/gin"
)

func GetMarketIndices(c *gin.Context) {
    // Example static data - replace with real data from DB or external API
    indices := []models.MarketIndex{
        {Symbol: "NSEI", Name: "Nifty 50", Price: 18300.25, Change: 120.5, PercentChange: 0.66},
        {Symbol: "BSESN", Name: "Sensex", Price: 61500.15, Change: -200.45, PercentChange: -0.32},
        {Symbol: "DJI", Name: "Dow Jones Industrial", Price: 34500.40, Change: 150.7, PercentChange: 0.44},
        {Symbol: "NASDAQ", Name: "Nasdaq Composite", Price: 13800.88, Change: -180.1, PercentChange: -1.29},
    }

    c.JSON(http.StatusOK, gin.H{"indices": indices})
}
