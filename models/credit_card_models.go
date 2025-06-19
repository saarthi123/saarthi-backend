package models

type CreditCard struct {
    ID         string  `json:"id"`
    UserID     string  `json:"userId"`
    BankName   string  `json:"bankName"`
    CardNumber string  `json:"cardNumber"`
    Due        float64 `json:"due"`
    DueDate    string  `json:"dueDate"`
}

type CreditCardTransaction struct {
    ID          string  `json:"id"`
    CardID      string  `json:"cardId"`
    Date        string  `json:"date"`
    Description string  `json:"description"`
    Amount      float64 `json:"amount"`
}
