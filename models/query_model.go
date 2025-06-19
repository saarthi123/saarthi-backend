package models

import "time"

type Query struct {
    ID          string    `json:"id"`
    Question    string    `json:"question"`
    Subject     string    `json:"subject"`
    Attachments []string  `json:"attachments"` // URLs or filenames
    Student     string    `json:"student"`
    Course      string    `json:"course"`
    Message     string    `json:"message"`
    Time        time.Time `json:"time"`
    Resolved    bool      `json:"resolved"`
	UserID   uint      `json:"user_id"`
	
}
