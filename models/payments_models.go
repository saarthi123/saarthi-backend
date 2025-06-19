// (Assuming this file exists and contains other credit card related structs)

package models

import (
    "time"
"gorm.io/gorm"
)


type PayRequest struct {
    Amount float64 `json:"amount"`
    Pin    string  `json:"pin"`
}
type Account struct {
    AccountNumber string  `json:"accountNumber"`
    Balance       float64 `json:"balance"`
}


type Analytics struct {
    gorm.Model
    UserID uint   `json:"userId"`
    Action string `json:"action"`
    Count  int    `json:"count"`
}



type Bank struct {
    gorm.Model
    Name   string `json:"name"`
    Branch string `json:"branch"`
    IFSC   string `json:"ifsc"`
}

type Payment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`    // Add this
	Currency  string    `json:"currency"`  // Add this
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"` // Add this
}

