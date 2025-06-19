package handlers

import (
	"log"
)

// Dummy User struct for demonstration; replace with your actual User struct if needed.
type User struct {
	Email string
}

// sendEmail sends an email to the specified address with the given subject and body.
// Replace this with your actual email sending implementation.
func sendEmail(to, subject, body string) error {
	// TODO: Implement actual email sending logic here.
	log.Printf("Sending email to %s: %s - %s", to, subject, body)
	return nil
}

func sendWelcomeEmail(user User) {
	go func() {
		err := sendEmail(user.Email, "Welcome to Saarthi", "Your Saarthi profile has been created successfully!")
		if err != nil {
			log.Printf("Failed to send welcome email: %v", err)
		}
	}()
}