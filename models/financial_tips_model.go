package models

import "gorm.io/gorm"

type FinancialTip struct {
	gorm.Model
	UserID uint   `json:"user_id"` // Optional: relate to specific user
	Title  string `json:"title"`
	Body   string `json:"body"`
}
