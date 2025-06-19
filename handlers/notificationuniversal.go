package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/saarthi123/saarthi-backend/models"
)

type NotificationHandler struct {
    DB *gorm.DB
}

// Get all notifications for authenticated user
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
    userID := c.GetString("userID")
    var notifications []models.Notification

    if err := h.DB.Where("user_id = ?", userID).Order("date desc").Find(&notifications).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

// Mark notification as read
func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
    userID := c.GetString("userID")
    notifID := c.Param("id")

    if err := h.DB.Model(&models.Notification{}).
        Where("id = ? AND user_id = ?", notifID, userID).
        Update("read", true).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "marked as read"})
}

// Mark notification as unread
func (h *NotificationHandler) MarkAsUnread(c *gin.Context) {
    userID := c.GetString("userID")
    notifID := c.Param("id")

    if err := h.DB.Model(&models.Notification{}).
        Where("id = ? AND user_id = ?", notifID, userID).
        Update("read", false).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as unread"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "marked as unread"})
}

// Get notification preferences
func (h *NotificationHandler) GetPreferences(c *gin.Context) {
    userID := c.GetString("userID")
    var prefs models.NotificationPreferences

    if err := h.DB.First(&prefs, "user_id = ?", userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Return default prefs if none set
            prefs = models.NotificationPreferences{
                UserID:            userID,
                LowBalanceAlerts:  true,
                SecurityAlerts:    true,
                TransactionAlerts: true,
                PaymentReminders:  true,
            }
            c.JSON(http.StatusOK, gin.H{"preferences": prefs})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch preferences"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"preferences": prefs})
}

// Update notification preferences
func (h *NotificationHandler) UpdatePreferences(c *gin.Context) {
    userID := c.GetString("userID")
    var prefs models.NotificationPreferences

    if err := c.BindJSON(&prefs); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    prefs.UserID = userID

    if err := h.DB.Save(&prefs).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update preferences"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"preferences": prefs})
}
