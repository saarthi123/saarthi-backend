package models

import (
    "time"
    "gorm.io/gorm"

)

type Transaction struct {
    	gorm.Model
    ID          string    `gorm:"primaryKey" json:"id"`
    AccountID   string    `json:"accountId"`
    Date        time.Time `json:"date"`
    Description string    `json:"description"`
    Amount      float64   `json:"amount"`
    Type        string    `json:"type"` // credit or debit
    Status      string    `json:"status"`
    Account     string    `json:"account"`         // ✅ Add this field
	UserID      uint      `json:"user_id"`
}