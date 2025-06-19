package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func DownloadPDF(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Statement PDF downloaded"})
}

func DownloadExcel(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Statement Excel downloaded"})
}
