package handlers

import (
    "net/http"
    "github.com/saarthi123/saarthi-backend/models"

    "github.com/gin-gonic/gin"
)

func GetExchanges(c *gin.Context) {
    exchanges := []models.Exchange{
        {
            Name:         "National Stock Exchange (NSE)",
            Code:         "NSE",
            Country:      "India",
            TradingHours: "09:15 - 15:30 IST",
            Holidays:     []string{"Jan 26", "Aug 15", "Oct 2"},
            Region:       "Asia",
            LogoUrl:      "assets/logos/nse.png",
        },
        {
            Name:         "Bombay Stock Exchange (BSE)",
            Code:         "BSE",
            Country:      "India",
            TradingHours: "09:15 - 15:30 IST",
            Holidays:     []string{"Jan 26", "Aug 15", "Oct 2"},
            Region:       "Asia",
            LogoUrl:      "assets/logos/bse.png",
        },
        {
            Name:         "NASDAQ",
            Code:         "NASDAQ",
            Country:      "USA",
            TradingHours: "09:30 - 16:00 EST",
            Holidays:     []string{"Jan 1", "Jul 4", "Dec 25"},
            Region:       "North America",
            LogoUrl:      "assets/logos/nasdaq.png",
        },
        {
            Name:         "Dow Jones Industrial Average (DOW)",
            Code:         "DOW",
            Country:      "USA",
            TradingHours: "09:30 - 16:00 EST",
            Holidays:     []string{"Jan 1", "Jul 4", "Dec 25"},
            Region:       "North America",
            LogoUrl:      "assets/logos/dow.png",
        },
        // Add more as needed
    }

    c.JSON(http.StatusOK, gin.H{"exchanges": exchanges})
}
