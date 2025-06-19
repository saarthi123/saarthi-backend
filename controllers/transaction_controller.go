package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionFilter struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
	Min  string `json:"min"`
	Max  string `json:"max"`
}

// GET filtered transactions
func GetFilteredTransactions(c *gin.Context) {
	var filter TransactionFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter format"})
		return
	}

	// You can add DB query logic here
	transactions := []gin.H{
		{"id": "txn001", "amount": 2000, "type": "credit", "date": "2025-05-01"},
		{"id": "txn002", "amount": 500, "type": "debit", "date": "2025-05-03"},
	}

c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// Download as PDF
func DownloadTransactionsPDF(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=transactions.pdf")
	c.Data(http.StatusOK, "application/pdf", []byte("Fake PDF binary"))
}

// Download as CSV
func DownloadTransactionsCSV(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=transactions.csv")
	c.Data(http.StatusOK, "text/csv", []byte("id,amount,type,date\ntxn001,2000,credit,2025-05-01\n"))
}
