package models

import "gorm.io/gorm"

type Draft struct {
    gorm.Model
    UserID  string `json:"user_id" gorm:"not null"`
    To      string `json:"to"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
	ID      string `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

}
