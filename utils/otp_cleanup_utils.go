package utils

import (
	"log"
	"time"

	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
)

// StartOtpCleanup runs a background goroutine to clean expired OTPs from DB.
func StartOtpCleanup(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			result := config.DB.Where("expires_at < ?", time.Now()).Delete(&models.Otp{})
			if result.Error != nil {
				log.Printf("❌ Failed to clean expired OTPs: %v", result.Error)
			} else if result.RowsAffected > 0 {
				log.Printf("🧹 Cleaned %d expired OTP(s)", result.RowsAffected)
			}
		}
	}()
}
