package models

import "time"

type Notification struct {
    Type      string    `json:"type"` 
    Title     string    `json:"title"`
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
    ID      string `json:"id"`
    Read    bool   `json:"read"`
    UserID  string `json:"user_id"` 

}



type NotificationPreferences struct {
    UserID            string `gorm:"primaryKey" json:"user_id"`
    LowBalanceAlerts  bool   `json:"low_balance_alerts"`
    SecurityAlerts    bool   `json:"security_alerts"`
    TransactionAlerts bool   `json:"transaction_alerts"`
    PaymentReminders  bool   `json:"payment_reminders"`
}