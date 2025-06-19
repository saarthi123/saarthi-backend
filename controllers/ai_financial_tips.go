package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
	"github.com/saarthi123/saarthi-backend/utils"

	"net/http"
	"strconv"
)

func GetFinancialTips(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	var tips []models.FinancialTip
	if err := models.DB.Where("user_id = ?", userId).Find(&tips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tips"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tips": tips})
}

func DownloadInsightReport(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	format := c.Query("format")

	var tips []models.FinancialTip
	if err := models.DB.Where("user_id = ?", userId).Find(&tips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tips"})
		return
	}

	filename := "Saarthi_AI_Tips." + format

	switch format {
	case "pdf":
		pdfBytes := utils.GeneratePDF(tips)
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(http.StatusOK, "application/pdf", pdfBytes)
	case "txt":
		txtContent := utils.GenerateText(tips)
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(http.StatusOK, "text/plain", []byte(txtContent))
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
	}
}
