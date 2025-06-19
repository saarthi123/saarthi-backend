package models

import (
    "gorm.io/gorm"
     
)

type BankAccount struct {
    gorm.Model
    UserID        uint   `json:"userId" gorm:"not null"`
    BankName      string `json:"bankName" gorm:"not null"`
    AccountNumber string `json:"accountNumber" gorm:"not null;unique"`
    AccountType   string `json:"accountType" gorm:"not null"`
    IFSC          string `json:"ifsc" gorm:"not null"`
    Balance       float64 `json:"balance"`
}

type UpiTransaction struct {
    gorm.Model
    UserID uint    `json:"userId"`
    UpiID  string  `json:"upiId"`
    Amount float64 `json:"amount"`
    Type   string  `json:"type"` // sent or received
}





