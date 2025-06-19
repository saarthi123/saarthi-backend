package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

var bankAccounts = make(map[uint][]models.BankAccount)

func AddAccount(c *gin.Context) {
	userIdStr := c.Param("userId")
	uid, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var account models.BankAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account.UserID = uint(uid)
	bankAccounts[uint(uid)] = append(bankAccounts[uint(uid)], account)

	c.JSON(http.StatusOK, gin.H{"message": "Account added successfully"})
}

// PUT /api/bank/:userId
// PUT /api/bank/:userId
func UpdateAccount(c *gin.Context) {
	userIdStr := c.Param("userId")
	uid, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userId := uint(uid)

	var updated models.BankAccount
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accounts := bankAccounts[userId]
	for i, acc := range accounts {
		if acc.AccountNumber == updated.AccountNumber {
			updated.UserID = userId
			bankAccounts[userId][i] = updated
			c.JSON(http.StatusOK, gin.H{"message": "Account updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
}

// DELETE /api/bank/:userId/:accountNumber
func DeleteAccount(c *gin.Context) {
	userIdStr := c.Param("userId")
	uid, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userId := uint(uid)

	accountNumber := c.Param("accountNumber")
	accounts := bankAccounts[userId]
	for i, acc := range accounts {
		if acc.AccountNumber == accountNumber {
			bankAccounts[userId] = append(accounts[:i], accounts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
}
