package models

import "gorm.io/gorm"

type Mail struct {
		gorm.Model
	ID        string `json:"id"`          // Unique identifier (use string for UUID)
	Sender    string `json:"sender"`      // Email of sender
	Recipient string `json:"recipient"`   // Email of recipient
	Subject   string `json:"subject"`     // Mail subject
	Body      string `json:"body"`        // Mail content
	Status    string `json:"status"`      // inbox, sent, draft, spam, trash
	IsStarred bool   `json:"isStarred"`   // Starred flag
	Receiver  string `json:"receiver"`
	SenderName string `json:"senderName"`
}
