package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// SendMessage saves a new message to the database with basic validation
func SendMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message payload"})
		return
	}

	if msg.SenderID == "" || msg.ReceiverID == "" || msg.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	if err := config.DB.Create(&msg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "data": msg})
}

// GetMessages retrieves all messages in the system (admin/debug use)
func GetMessages(c *gin.Context) {
	var messages []models.Message
	if err := config.DB.Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

// GetInboxMessages retrieves all messages sent to a particular user
func GetInboxMessages(c *gin.Context) {
	receiverID := c.Query("receiverId")
	if receiverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receiver ID is required"})
		return
	}

	var messages []models.Message
	if err := config.DB.Where("receiver_id = ?", receiverID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inbox": messages})
}
