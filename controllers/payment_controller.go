package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// ───── Courses ──────────────────────────────────────────────

func UploadCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course data"})
		return
	}
	course.ID = time.Now().Format("20060102150405")
	config.DB.Create(&course)
	c.JSON(http.StatusCreated, gin.H{"message": "Course uploaded", "course": course})
}

// ───── Live Sessions ───────────────────────────────────────

func GetLiveSessions(c *gin.Context) {
	var sessions []models.LiveSession
	config.DB.Find(&sessions)
	c.JSON(http.StatusOK, gin.H{"sessions": sessions})
}

func ScheduleSession(c *gin.Context) {
	var session models.LiveSession
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session data"})
		return
	}
	session.ID = time.Now().Format("20060102150405")
	config.DB.Create(&session)
	c.JSON(http.StatusCreated, gin.H{"message": "Session scheduled", "session": session})
}

// ───── Student Queries ─────────────────────────────────────

func GetQueries(c *gin.Context) {
	var queries []models.StudentQuery
	config.DB.Find(&queries)
	c.JSON(http.StatusOK, gin.H{"queries": queries})
}

func ReplyQuery(c *gin.Context) {
	idStr := c.Param("id")
	var body struct {
		Reply string `json:"reply"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.Reply == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reply"})
		return
	}

	var query models.StudentQuery
	if err := config.DB.First(&query, "id = ?", idStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Query not found"})
		return
	}

	query.Reply = body.Reply
	config.DB.Save(&query)

	c.JSON(http.StatusOK, gin.H{"message": "Reply saved", "query": query})
}

// ───── Payments ────────────────────────────────────────────

func CreatePayment(c *gin.Context) {
	var paymentReq struct {
		Amount   float64 `json:"amount"`
		Method   string  `json:"method"`
		Currency string  `json:"currency"`
	}
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	payment := models.Payment{
		ID:        time.Now().Format("20060102150405"),
		Amount:    paymentReq.Amount,
		Method:    paymentReq.Method,
		Currency:  paymentReq.Currency,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	config.DB.Create(&payment)

	paymentURL := "https://aws-payment-gateway.com/pay/" + payment.ID
	c.JSON(http.StatusOK, gin.H{
		"message":    "Payment initiated",
		"paymentUrl": paymentURL,
		"payment":    payment,
	})
}

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	config.DB.Find(&payments)
	c.JSON(http.StatusOK, gin.H{"payments": payments})
}
