package models

import "gorm.io/gorm"

type Diploma struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Progress    int    `json:"progress"` // percentage completed
    UserID      uint   `json:"user_id"`  // FK to User

}

