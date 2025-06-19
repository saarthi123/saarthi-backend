package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuggestionRequest struct {
	Text string `json:"text"`
}

type SuggestionResponse struct {
	Suggestions []string `json:"suggestions"`
}

func AISuggestionsHandler(c *gin.Context) {
	var req SuggestionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Simulated AI suggestions (you can later integrate real logic here)
	suggestions := generateSuggestions(req.Text)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Suggestions generated",
		"suggestions": suggestions,
	})
}

func generateSuggestions(input string) []string {
	if len(input) < 3 {
		return []string{}
	}
	return []string{
		input + " please",
		input + " now",
		input + " asap",
	}
}
