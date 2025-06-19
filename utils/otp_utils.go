package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"sync"
	"time"
)

type otpEntry struct {
	OTP       string
	ExpiresAt time.Time
}

var (
	otpStore = make(map[string]otpEntry)
	mu       sync.Mutex
)

func SetOtp(phone, otp string) {
	mu.Lock()
	defer mu.Unlock()
	otpStore[phone] = otpEntry{
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func GetOtp(phone string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	entry, exists := otpStore[phone]
	if !exists {
		return "", fmt.Errorf("OTP not found")
	}
	if time.Now().After(entry.ExpiresAt) {
		delete(otpStore, phone)
		return "", fmt.Errorf("OTP expired")
	}
	return entry.OTP, nil
}

func DeleteOtp(phone string) {
	mu.Lock()
	defer mu.Unlock()
	delete(otpStore, phone)
}

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func IsValidPhone(phone string) bool {
	return regexp.MustCompile(`^\d{10}$`).MatchString(phone)
}
