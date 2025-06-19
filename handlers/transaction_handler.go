package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/saarthi123/saarthi-backend/models"
  "github.com/saarthi123/saarthi-backend/config"
  "net/http"
)

func GetTransactions(c *gin.Context) {
  var txns []models.Transaction
  config.DB.Order("created_at desc").Find(&txns)
c.JSON(http.StatusOK, gin.H{"txns": txns})
}
