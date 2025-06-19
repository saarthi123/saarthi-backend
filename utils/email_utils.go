package utils

import (
	"log"
	"net/smtp"
)

// sendEmail sends a basic email via SMTP
func sendEmail(to, subject, body string) error {
	from := "your-email@example.com"
	password := "your-password"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}

// SendWelcomeEmailAsync sends the welcome email in a background goroutine
func SendWelcomeEmailAsync(email string) {
	go func() {
		err := sendEmail(email, "Welcome to Saarthi", "Your Saarthi profile has been created successfully!")
		if err != nil {
			log.Printf("❌ Failed to send welcome email: %v", err)
		} else {
			log.Printf("📧 Welcome email sent to %s", email)
		}
	}()
}
