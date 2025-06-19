package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/saarthi123/saarthi-backend/models"
)

// GET /queries - fetch all submitted queries from DB
func GetAllQueries(c *gin.Context) {
	var queries []models.Query

	if err := models.DB.Order("time DESC").Find(&queries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch queries"})
		return
	}

c.JSON(http.StatusOK, gin.H{"queries": queries}) // ✅ Also valid (map form)
}

// POST /queries - submit a new query
func SubmitQuery(c *gin.Context) {
	var newQuery models.Query

	if err := c.ShouldBindJSON(&newQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newQuery.ID = uuid.NewString()
	newQuery.Time = time.Now()
	newQuery.Resolved = false

	if err := models.DB.Create(&newQuery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save query"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"query": newQuery})
}
