package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

var creditCards = []models.CreditCard{
	{
		ID:         "cc1",
		UserID:     "demo-user-id",
		BankName:   "HDFC",
		CardNumber: "1234567812345678",
		Due:        2500.50,
		DueDate:    "2025-06-10",
	},
}

var cardTransactions = map[string][]models.CreditCardTransaction{
	"cc1": {
		{ID: "t1", CardID: "cc1", Date: "2025-06-01", Description: "Amazon Purchase", Amount: 1299.99},
		{ID: "t2", CardID: "cc1", Date: "2025-05-28", Description: "Swiggy", Amount: 299.00},
	},
}

func GetAllCards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cards": creditCards})
}

func GetCardTransactions(c *gin.Context) {
	cardId := c.Param("cardId")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil || page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	cardTxns := cardTransactions[cardId]
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(cardTxns) {
		c.JSON(http.StatusOK, gin.H{"transactions": []models.CreditCardTransaction{}})
		return
	}
	if end > len(cardTxns) {
		end = len(cardTxns)
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": cardTxns[start:end],
		"page":         page,
		"pageSize":     pageSize,
		"total":        len(cardTxns),
	})
}

func PayDueAmount(c *gin.Context) {
	cardId := c.Param("cardId")
	var req models.PayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Pin != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid PIN"})
		return
	}

	for i, card := range creditCards {
		if card.ID == cardId {
			creditCards[i].Due -= req.Amount
			if creditCards[i].Due < 0 {
				creditCards[i].Due = 0
			}
			c.JSON(http.StatusOK, gin.H{"message": "Payment successful", "remainingDue": creditCards[i].Due})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
}

func DownloadStatement(c *gin.Context) {
	format := c.DefaultQuery("format", "pdf")
	cardId := c.Param("cardId")
	filename := fmt.Sprintf("statement_%s.%s", cardId, format)
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.String(http.StatusOK, fmt.Sprintf("Fake %s statement data for card %s", format, cardId))
}
