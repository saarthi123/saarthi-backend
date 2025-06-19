// 📁 controllers/analytics_controller.go
package controllers

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/models"
	"github.com/sashabaranov/go-openai"
)

// --------- Summary Analytics ---------
func AnalyticsSummary(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	totalUsers := 1000 + rand.Intn(500)
	activeUsers := int(float64(totalUsers) * 0.7)
	completionRate := 60 + rand.Float64()*30
	videosWatched := 200 + rand.Intn(300)
	assignments := 80 + rand.Intn(100)
	avgSessionSecs := 600 + rand.Intn(900)

	summary := []models.AnalyticsData{
		{Label: "Total Users", Value: float64(totalUsers)},
		{Label: "Active Users", Value: float64(activeUsers)},
		{Label: "Course Completion (%)", Value: completionRate},
		{Label: "Videos Watched", Value: float64(videosWatched)},
		{Label: "Assignments Completed", Value: float64(assignments)},
		{Label: "Avg. Session Time (mins)", Value: float64(avgSessionSecs) / 60},
	}

	c.JSON(http.StatusOK, gin.H{
		"summary":                summary,
		"formatted_session_time": formatSeconds(avgSessionSecs),
	})
}

func formatSeconds(seconds int) string {
	min := seconds / 60
	sec := seconds % 60
	return formatTwoDigits(min) + "m " + formatTwoDigits(sec) + "s"
}

func formatTwoDigits(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

// --------- Engagement Data ---------
func GetEngagement(c *gin.Context) {
	data := []models.AnalyticsData{
		{Label: "Videos Watched", Value: 75},
		{Label: "Assignments Completed", Value: 60},
	}
	c.JSON(http.StatusOK, gin.H{"engagement": data})
}

// --------- Drop Off Points ---------
func GetDropOffPoints(c *gin.Context) {
	points := []string{"Module 3", "Quiz 2"}
	c.JSON(http.StatusOK, gin.H{"dropOffPoints": points})
}

// --------- AI Suggestions ---------
func GetAISuggestions(c *gin.Context) {
	var req models.AISuggestionRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Context == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Context required"})
		return
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		Prompt:    "Give educational improvement tips based on: " + req.Context,
		MaxTokens: 100,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"suggestion": resp.Choices[0].Text,
	})
}
