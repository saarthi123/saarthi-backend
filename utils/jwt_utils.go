package utils

import (
	"crypto/rand"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

func init() {
	jwtKey = []byte(GetEnv("JWT_SECRET", ""))
	if len(jwtKey) == 0 {
		panic("JWT_SECRET not set")
	}
}

type OtpClaims struct {
	PhoneNumber string `json:"phone_number"`
	Otp         string `json:"otp"`
	jwt.RegisteredClaims
}

type CustomClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

// Secure numeric OTP generator
func GenerateSecureOTP(length int) (string, error) {
	const digits = "0123456789"
	if length <= 0 {
		return "", errors.New("invalid OTP length")
	}

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	otp := make([]byte, length)
	for i := range otp {
		otp[i] = digits[bytes[i]%10]
	}
	return string(otp), nil
}

// OTP token generator
func GenerateOtpToken(phone string) (string, string, error) {
	otp, err := GenerateSecureOTP(6)
	if err != nil {
		return "", "", err
	}

	claims := &OtpClaims{
		PhoneNumber: phone,
		Otp:         otp,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return otp, tokenStr, nil
}

// Verify OTP JWT Token
func VerifyOtpToken(tokenStr string) (*OtpClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &OtpClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*OtpClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired OTP token")
	}
	return claims, nil
}

// Validate user-entered OTP
func VerifyOTPWithToken(phone, inputOtp, tokenStr string) (bool, error) {
	claims, err := VerifyOtpToken(tokenStr)
	if err != nil {
		return false, err
	}
	if claims.PhoneNumber != phone || claims.Otp != inputOtp {
		return false, errors.New("OTP mismatch")
	}
	return true, nil
}

// Access JWT (3 days)
func CreateJWTToken(phone string) (string, error) {
	claims := &CustomClaims{
		PhoneNumber: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Verify access token
func VerifyJWTToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired access token")
	}
	return claims, nil
}

// Optional refresh token (30 days)
func CreateRefreshToken(phone string) (string, error) {
	claims := &RefreshClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyRefreshToken(tokenStr string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}
	return claims, nil
}


var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// ValidateJWT extracts UserID from token
func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("token expired or invalid")
	}

	return claims.UserID, nil
}