package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

type CareerPathRequest struct {
	Interests []string `json:"interests"`
	Skills    []string `json:"skills"`
}

func GenerateCareerPaths(c *gin.Context) {
	var req CareerPathRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Interests) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or empty interests"})
		return
	}

	// Normalize interests for comparison
	normalized := make([]string, 0, len(req.Interests))
	for _, interest := range req.Interests {
		normalized = append(normalized, strings.ToLower(strings.TrimSpace(interest)))
	}

	// Dynamic Query: Find career paths matching ANY of the given interests
	var paths []models.CareerPath
	err := config.DB.Where("LOWER(interest) IN ?", normalized).Find(&paths).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch career paths from database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"career_paths": paths})
}
