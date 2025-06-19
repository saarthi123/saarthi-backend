package models

import (
    "gorm.io/gorm"
    "time"
)

type UpcomingClass struct {
    gorm.Model
    Title     string    `json:"title"`
    StartTime time.Time `json:"start_time"`
    Duration  int       `json:"duration"` // in minutes
    UserID    uint      `json:"user_id"`  // FK to User
	Mode  string    `json:"mode"`  // e.g., "Online" or "Offline"
	Date  time.Time `json:"date"`  // upcoming session date/time
}
