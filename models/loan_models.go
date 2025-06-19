package models

type Loan struct {
    ID           string  `json:"id"`
    Account      string  `json:"account"`
    Amount       float64 `json:"amount"`
    Remaining    float64 `json:"remaining"`
    EMIDate      string  `json:"emiDate"`
    Status       string  `json:"status"` // active, closed
}

type LoanSummary struct {
    ID            string  `json:"id"`
    AccountNumber string  `json:"accountNumber"`
    Principal     float64 `json:"principal"`
    Remaining     float64 `json:"remaining"`
    InterestRate  float64 `json:"interestRate"`
    TermMonths    int     `json:"termMonths"`
    StartDate     string  `json:"startDate"`
}

type EMISchedule struct {
    Month        string  `json:"month"`
    DueDate      string  `json:"dueDate"`
    Principal    float64 `json:"principal"`
    Interest     float64 `json:"interest"`
    TotalPayment float64 `json:"totalPayment"`
    Status       string  `json:"status"` // Paid / Due
}

