package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// SendOtpHandler sends a 6-digit OTP to the user's phone
func SendOtpHandler(c *gin.Context) {
	var req struct {
		Phone string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	// Generate 6-digit OTP
	otp := generateOtp()
	expiry := time.Now().Add(5 * time.Minute)

	// Store OTP in DB
	config.DB.Where("phone = ?", req.Phone).Delete(&models.Otp{})
	config.DB.Create(&models.Otp{
		Phone:     req.Phone,
		Code:      otp,
		ExpiresAt: expiry,
	})

	// Send OTP via AWS SNS
	msg := "Your Saarthi OTP is: " + otp
	sess := session.Must(session.NewSession())
	svc := sns.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	_, err := svc.Publish(&sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(req.Phone),
	})

	if err != nil {
		log.Printf("Failed to send OTP via SNS: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func generateOtp() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
