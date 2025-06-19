package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// SendOtpSMS sends the OTP SMS to a phone number using AWS SNS.
func SendOtpSMS(phone, otp string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("❌ [SendOtpSMS] AWS config load failed: %v", err)
		return err
	}

	client := sns.NewFromConfig(cfg)

	message := fmt.Sprintf("Your Saarthi OTP code is: %s", otp)
	input := &sns.PublishInput{
		Message:     &message,
		PhoneNumber: &phone,
	}

	_, err = client.Publish(context.TODO(), input)
	if err != nil {
		log.Printf("❌ [SendOtpSMS] Failed to send SMS: %v", err)
		return err
	}

	log.Printf("✅ [SendOtpSMS] OTP sent to %s", phone)
	return nil
}
