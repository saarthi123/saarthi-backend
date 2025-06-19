package services

import (
    "math/rand"
    "sync"
    "time"
    "errors"
)

type otpEntry struct {
    code      string
    expiresAt time.Time
}

var (
    otpStore = map[string]otpEntry{}
    mu       sync.Mutex
    otpTTL   = 5 * time.Minute
)

func GenerateOTP(phone string) (string, error) {
    mu.Lock()
    defer mu.Unlock()

    if entry, exists := otpStore[phone]; exists {
        if time.Now().Before(entry.expiresAt.Add(-4 * time.Minute)) {
            return "", errors.New("OTP recently sent, please wait before requesting again")
        }
    }

    otp := generateRandomOTP(6)
    otpStore[phone] = otpEntry{
        code:      otp,
        expiresAt: time.Now().Add(otpTTL),
    }
    return otp, nil
}

func generateRandomOTP(length int) string {
    rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    otp := make([]byte, length)
    for i := 0; i < length; i++ {
        otp[i] = digits[rand.Intn(len(digits))]
    }
    return string(otp)
}

func VerifyOTP(phone string, inputCode string) (bool, error) {
    mu.Lock()
    defer mu.Unlock()

    entry, exists := otpStore[phone]
    if !exists {
        return false, errors.New("OTP not found")
    }

    if time.Now().After(entry.expiresAt) {
        delete(otpStore, phone) // Clean up expired OTP
        return false, errors.New("OTP expired")
    }

    if entry.code != inputCode {
        return false, errors.New("invalid OTP")
    }

    delete(otpStore, phone) // Clean up after successful use
    return true, nil
}
