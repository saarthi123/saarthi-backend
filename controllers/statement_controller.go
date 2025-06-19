package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

// Define struct for filters
type StatementFilter struct {
	FromDate  string  `json:"fromDate"`
	ToDate    string  `json:"toDate"`
	Type      string  `json:"type"`     // e.g., "credit", "debit", etc.
	MinAmount float64 `json:"minAmount"`
	MaxAmount float64 `json:"maxAmount"`
}

// Sample transaction data (you can fetch from DB in real app)
var statements = []models.Transaction{
	{
		ID:          "txn001",
		Account:     "1234567890",
		Amount:      1200,
		Description: "Grocery",
		Date:        mustParseDate("2025-06-01"),
		Type:        "debit",
	},
	{
		ID:          "txn002",
		Account:     "1234567890",
		Amount:      900,
		Description: "Fuel",
		Date:        mustParseDate("2025-06-02"),
		Type:        "debit",
	},
	{
		ID:          "txn003",
		Account:     "1234567890",
		Amount:      5000,
		Description: "Salary",
		Date:        mustParseDate("2025-06-03"),
		Type:        "credit",
	},
}

func FilterStatements(c *gin.Context) {
	var filter StatementFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter data"})
		return
	}

	fromDate, _ := time.Parse("2006-01-02", filter.FromDate)
	toDate, _ := time.Parse("2006-01-02", filter.ToDate)

	var result []models.Transaction
	for _, txn := range statements {
		if !txn.Date.After(fromDate.Add(-time.Hour*24)) || !txn.Date.Before(toDate.Add(time.Hour*24)) {
			continue
		}
		if filter.Type != "" && strings.ToLower(txn.Type) != strings.ToLower(filter.Type) {
			continue
		}
		if txn.Amount < filter.MinAmount || txn.Amount > filter.MaxAmount {
			continue
		}
		result = append(result, txn)
	}

	c.JSON(http.StatusOK, gin.H{"statements": result})
}
