package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

func DeleteDraftHandler(c *gin.Context) {
	draftID := c.Param("id")
	if draftID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Draft ID is required"})
		return
	}

	// TODO: Add DB deletion logic here using draftID

	c.JSON(http.StatusOK, gin.H{
		"message": "Draft deleted successfully",
		"id":      draftID,
	})
}


// GetDraftsHandler fetches all drafts from AWS PostgreSQL
func GetDraftsHandler(c *gin.Context) {
	var drafts []models.Draft

	if err := config.DB.Find(&drafts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch drafts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"drafts": drafts})
}

// SaveDraftHandler handles POST /api/drafts
func SaveDraftHandler(c *gin.Context) {
	var draft models.Draft
	if err := c.ShouldBindJSON(&draft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid draft data"})
		return
	}

	// If ID is not provided, generate a new one
	if draft.ID == "" {
		draft.ID = time.Now().Format("20060102150405")
	}

	// Save or update the draft
	if err := config.DB.Save(&draft).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save draft"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Draft saved", "draft": draft})
}