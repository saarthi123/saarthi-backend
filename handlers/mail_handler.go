package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// ========== Models ==========

type DraftPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Mail struct {
	ID        string `json:"id"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	IsStarred bool   `json:"isStarred"`
}

// ========== In-Memory Stores ==========

var (
	drafts     = map[string]DraftPayload{}
	idCounter  = 0
	draftLock  sync.Mutex

	mailStore = make(map[string]*Mail)
	mailMutex sync.Mutex
)

// ========== Draft Handlers ==========

// SaveDraftHandler saves a new draft
func SaveDraftHandler(c *gin.Context) {
	var draft DraftPayload
	if err := c.ShouldBindJSON(&draft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid draft"})
		return
	}

	draftLock.Lock()
	defer draftLock.Unlock()

	idCounter++
	id := fmt.Sprintf("%d", idCounter)
	drafts[id] = draft

	c.JSON(http.StatusOK, gin.H{"message": "Draft saved", "id": id})
}

// GetDraftsHandler returns all saved drafts
func GetDraftsHandler(c *gin.Context) {
	draftLock.Lock()
	defer draftLock.Unlock()

	result := []map[string]interface{}{}
	for id, draft := range drafts {
		result = append(result, map[string]interface{}{
			"id":      id,
			"to":      draft.To,
			"subject": draft.Subject,
			"body":    draft.Body,
		})
	}
	c.JSON(http.StatusOK, gin.H{"drafts": result})
}

// DeleteDraftHandler deletes a draft by ID
func DeleteDraftHandler(c *gin.Context) {
	id := c.Param("id")

	draftLock.Lock()
	defer draftLock.Unlock()

	if _, exists := drafts[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Draft not found"})
		return
	}
	delete(drafts, id)
	c.JSON(http.StatusOK, gin.H{"message": "Draft deleted"})
}

// ========== Mail Handlers ==========

// UpdateStarStatusHandler updates the "starred" status of a mail by ID
func UpdateStarStatusHandler(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		IsStarred bool `json:"isStarred"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	mailMutex.Lock()
	defer mailMutex.Unlock()

	mail, exists := mailStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mail not found"})
		return
	}

	mail.IsStarred = req.IsStarred
	c.JSON(http.StatusOK, gin.H{"mail": mail})
}
