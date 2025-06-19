package handlers

import (
    "net/http"
    "github.com/saarthi123/saarthi-backend/models"

    "github.com/gin-gonic/gin"
)

func GetNewsArticles(c *gin.Context) {
    // Dummy data - replace with real DB or external news API
    articles := []models.NewsArticle{
        {
            Title: "Sensex rallies 300 points amid positive earnings",
            Source: "Economic Times",
            Date: "2025-05-27",
            Sentiment: "Positive",
            SentimentEmoji: "📈",
        },
        {
            Title: "Rupee dips against dollar on global inflation worries",
            Source: "Business Standard",
            Date: "2025-05-27",
            Sentiment: "Negative",
            SentimentEmoji: "📉",
        },
        {
            Title: "Government proposes new tax reforms for startups",
            Source: "Mint",
            Date: "2025-05-26",
            Sentiment: "Neutral",
            SentimentEmoji: "⚖️",
        },
        {
            Title: "Reliance to expand green energy portfolio by 2026",
            Source: "MoneyControl",
            Date: "2025-05-25",
            Sentiment: "Positive",
            SentimentEmoji: "🌱",
        },
    }

    c.JSON(http.StatusOK, gin.H{
        "articles": articles,
    })
}
