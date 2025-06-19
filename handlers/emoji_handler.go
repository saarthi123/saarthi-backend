package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Call this once (in main or db initializer)
func InitEmojiHandlers(db *gorm.DB) {
	DB = db
}

// GET /api/emojis/categories
func GetEmojiCategories(c *gin.Context) {
	var categories []models.EmojiCategory
	if err := DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GET /api/emojis?category=Smileys
func GetEmojisByCategory(c *gin.Context) {
	categoryName := c.Query("category")
	if categoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	var category models.EmojiCategory
	if err := DB.Where("name = ?", categoryName).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var emojis []models.Emoji
	if err := DB.Where("category_id = ?", category.ID).Find(&emojis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch emojis"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"emojis": emojis})
}
