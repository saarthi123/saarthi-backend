package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/saarthi123/saarthi-backend/models"
)

// In-memory call store
var (
	calls     = make(map[string]*models.CallSession)
	callsLock sync.Mutex
)

// StartCall creates a new call session
func StartCall(c *gin.Context) {
	var req struct {
		CallerID   string `json:"caller_id"`
		ReceiverID string `json:"receiver_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	callID := uuid.New().String()
	callSession := &models.CallSession{
		ID:                  callID,
		CallerID:            req.CallerID,
		ReceiverID:          req.ReceiverID,
		IsMuted:             false,
		SpeakerOn:           true,
		AINoiseCancellation: true,
		DurationSeconds:     0,
		CreatedAt:           time.Now(),
	}

	callsLock.Lock()
	calls[callID] = callSession
	callsLock.Unlock()

	// TODO: Save callSession to PostgreSQL or DynamoDB
	c.JSON(http.StatusOK, gin.H{"call": callSession})
}

// EndCall deletes a call session
func EndCall(c *gin.Context) {
	callID := c.Query("id")
	if callID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing call ID"})
		return
	}

	callsLock.Lock()
	defer callsLock.Unlock()

	if _, exists := calls[callID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Call not found"})
		return
	}

	delete(calls, callID)

	// TODO: Delete from DB if using persistence
	c.Status(http.StatusNoContent)
}

// UpdateSettings updates mute/speaker/AI noise cancellation
func UpdateSettings(c *gin.Context) {
	callID := c.Query("id")
	if callID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing call ID"})
		return
	}

	var update struct {
		IsMuted             *bool `json:"is_muted,omitempty"`
		SpeakerOn           *bool `json:"speaker_on,omitempty"`
		AINoiseCancellation *bool `json:"ai_noise_cancellation,omitempty"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	callsLock.Lock()
	defer callsLock.Unlock()

	call, exists := calls[callID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Call not found"})
		return
	}

	if update.IsMuted != nil {
		call.IsMuted = *update.IsMuted
	}
	if update.SpeakerOn != nil {
		call.SpeakerOn = *update.SpeakerOn
	}
	if update.AINoiseCancellation != nil {
		call.AINoiseCancellation = *update.AINoiseCancellation
	}

	// TODO: Save updated call to DB
	c.JSON(http.StatusOK, gin.H{"call": call})
}
