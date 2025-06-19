package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

var accounts = map[string]models.Account{
	"1234567890": {AccountNumber: "1234567890", Balance: 50000.0},
}

var bankTransactions = map[string][]models.Transaction{
	"1234567890": {
		{
			ID:          "t1",
			Account:     "1234567890",
			Date:        mustParseDate("2025-06-01"),
			Description: "UPI to Rahul",
			Amount:      1500,
			Type:        "debit",
		},
		{
			ID:          "t2",
			Account:     "1234567890",
			Date:        mustParseDate("2025-05-30"),
			Description: "Salary",
			Amount:      60000,
			Type:        "credit",
		},
	},
}

var loans = map[string][]models.Loan{
	"1234567890": {
		{ID: "l1", Account: "1234567890", Amount: 100000, Remaining: 52000, EMIDate: "2025-06-10", Status: "active"},
	},
}

func mustParseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Now()
	}
	return t
}

func GetBalance(c *gin.Context) {
	acc := c.Param("accountNumber")
	account, ok := accounts[acc]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": account.Balance})
}

func GetRecentTransactions(c *gin.Context) {
	acc := c.Param("accountNumber")
	txns := bankTransactions[acc]
	c.JSON(http.StatusOK, gin.H{"transactions": txns})
}

func GetActiveLoans(c *gin.Context) {
	acc := c.Param("accountNumber")
	activeLoans := []models.Loan{}
	for _, loan := range loans[acc] {
		if strings.ToLower(loan.Status) == "active" {
			activeLoans = append(activeLoans, loan)
		}
	}
	c.JSON(http.StatusOK, gin.H{"loans": activeLoans})
}

func GetAiTips(c *gin.Context) {
	tips := []string{
		"💡 Set a monthly budget to track spending.",
		"📈 Invest 20% of your income in mutual funds.",
		"🛑 Avoid using more than 30% of your credit limit.",
	}
	c.JSON(http.StatusOK, gin.H{"tips": tips})
}

func GetTransactions(c *gin.Context) {
	category := c.Query("category")
	search := c.Query("search")

	dummyTxns := []models.Transaction{
		{
			ID:          "txn001",
			Account:     "1234567890",
			Amount:      1200,
			Description: "Grocery",
			Date:        mustParseDate("2025-06-01"),
			Type:        "food",
		},
		{
			ID:          "txn002",
			Account:     "1234567890",
			Amount:      900,
			Description: "Fuel",
			Date:        mustParseDate("2025-06-02"),
			Type:        "transport",
		},
		{
			ID:          "txn003",
			Account:     "1234567890",
			Amount:      5000,
			Description: "Shopping",
			Date:        mustParseDate("2025-06-03"),
			Type:        "shopping",
		},
	}

	filtered := []models.Transaction{}
	for _, t := range dummyTxns {
		if category != "" && category != "all" && t.Type != category {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(t.Description), strings.ToLower(search)) {
			continue
		}
		filtered = append(filtered, t)
	}

	c.JSON(http.StatusOK, gin.H{"transactions": filtered})
}

func ExportTransactionsPDF(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=transactions.pdf")
	c.Data(http.StatusOK, "application/pdf", []byte("%PDF-1.4 Fake PDF content"))
}

func ExportTransactionsExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=transactions.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte("Fake Excel content"))
}
