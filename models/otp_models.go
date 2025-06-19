package models

import "time"

type Otp struct {
	ID        uint      `gorm:"primaryKey"`
	Phone     string    `gorm:"not null"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}

// Add this function to your utils package

var otpStore = make(map[string]string)

// StoreOTP stores the OTP for a phone number
func StoreOTP(phone, otp string) {
    otpStore[phone] = otp
}

// GenerateOTP generates a random OTP (dummy implementation)
func GenerateOTP() string {
    return "123456" // Replace with actual random generation
}

// VerifyOTP checks if the provided OTP matches the stored OTP for the phone number
func VerifyOTP(phone, otp string) bool {
    stored, ok := otpStore[phone]
    return ok && stored == otp
}