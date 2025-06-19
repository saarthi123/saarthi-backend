// controllers/stock_controller.go
package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/saarthi123/saarthi-backend/models"
)

func GetWatchlist(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    watchlist := []models.Stock{
        {
            Symbol:  "RELIANCE",
            Company: "Reliance Industries",
            Price:   2748.20,
            Change:  1.42,
            Trend:   []float64{2700, 2710, 2725, 2740, 2748},
            AITag:   "🟢 Buy",
        },
        {
            Symbol:  "HDFCBANK",
            Company: "HDFC Bank",
            Price:   1620.75,
            Change:  -0.87,
            Trend:   []float64{1645, 1635, 1630, 1625, 1620},
            AITag:   "⚠️ Under Review",
        },
        {
            Symbol:  "TCS",
            Company: "Tata Consultancy",
            Price:   3680.15,
            Change:  0.21,
            Trend:   []float64{3660, 3665, 3670, 3675, 3680},
            AITag:   "🧠 Strong Fundamentals",
        },
    }

    json.NewEncoder(w).Encode(watchlist)
}
