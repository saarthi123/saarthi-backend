package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
)

// mockData simulates available dashboard features for search
var mockData = []models.SearchResult{
	{Title: "Messaging", Type: "messaging", Link: "/dashboard/messaging"},
	{Title: "Mail", Type: "mail", Link: "/dashboard/mail"},
	{Title: "Trading Panel", Type: "trading", Link: "/dashboard/trading"},
	{Title: "Banking Portal", Type: "banking", Link: "/dashboard/banking"},
}

func SearchHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	var results []models.SearchResult
	lowerQuery := strings.ToLower(query)

	for _, item := range mockData {
		if strings.Contains(strings.ToLower(item.Title), lowerQuery) {
			results = append(results, item)
		}
	}

c.JSON(http.StatusOK, gin.H{"results": results})
}
