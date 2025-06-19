package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// ----------- Database Mail Operations -----------

func SendMail(c *gin.Context) {
	var mail models.Mail
	if err := c.ShouldBindJSON(&mail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	mail.ID = uuid.New().String()
	mail.Status = "sent"
	mail.Sender = "user@saarthi.com" // Should come from user context in production

	if err := config.DB.Create(&mail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save mail"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mail": mail})
}

func GetMails(c *gin.Context) {
	var mails []models.Mail
	if err := config.DB.Find(&mails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch mails"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mails": mails})
}

func GetMailByID(c *gin.Context) {
	id := c.Param("id")
	var mail models.Mail
	if err := config.DB.First(&mail, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mail not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mail": mail})
}

func GetMailsByStatus(status string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mails []models.Mail
		if err := config.DB.Where("status = ?", status).Find(&mails).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching mails"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"mails": mails})
	}
}

func SaveDraft(c *gin.Context) {
	var draft models.Mail
	if err := c.ShouldBindJSON(&draft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	draft.ID = uuid.New().String()
	draft.Status = "draft"
	if err := config.DB.Create(&draft).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save draft"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mail": draft})
}

func UpdateDraft(c *gin.Context) {
	id := c.Param("id")
	var input models.Mail
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var draft models.Mail
	if err := config.DB.First(&draft, "id = ? AND status = ?", id, "draft").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Draft not found"})
		return
	}

	draft.Subject = input.Subject
	draft.Body = input.Body
	draft.Recipient = input.Recipient

	if err := config.DB.Save(&draft).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mail": draft})
}

func DeleteDraft(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Where("id = ? AND status = ?", id, "draft").Delete(&models.Mail{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Draft not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateStarStatus(c *gin.Context) {
	id := c.Param("id")
	var payload struct {
		IsStarred bool `json:"isStarred"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if err := config.DB.Model(&models.Mail{}).Where("id = ?", id).Update("is_starred", payload.IsStarred).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mail not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Star status updated"})
}

func RestoreFromTrash(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Model(&models.Mail{}).Where("id = ? AND status = ?", id, "trash").Update("status", "inbox").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mail not found in trash"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Restored to inbox"})
}

func MarkAsNotSpam(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Model(&models.Mail{}).Where("id = ? AND status = ?", id, "spam").Update("status", "inbox").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mail not found in spam"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Marked as not spam"})
}

func PermanentlyDeleteMail(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Mail{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mail not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// ----------- AI Features -----------

func GetSmartReplies(c *gin.Context) {
	var req struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	replies := []string{
		"Thank you for your email.",
		"I will get back to you shortly.",
		"Please provide more details.",
	}
	c.JSON(http.StatusOK, gin.H{"replies": replies})
}

func RewriteTone(c *gin.Context) {
	var req struct {
		Text string `json:"text"`
		Tone string `json:"tone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	rewritten := req.Text + " (rewritten in " + req.Tone + " tone)"
	c.JSON(http.StatusOK, gin.H{"rewritten": rewritten})
}

func SummarizeMail(c *gin.Context) {
	var req struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	summary := "Summary: " + req.Text[:min(len(req.Text), 100)] + "..."
	c.JSON(http.StatusOK, gin.H{"summary": summary})
}

// ----------- Helpers -----------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
