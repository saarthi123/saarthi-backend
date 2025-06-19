package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username     string `gorm:"uniqueIndex;not null" json:"username"`
    Name     string `json:"name"`     // ✅ Add this field if missing
    Email        string `gorm:"uniqueIndex;not null" json:"email"`
    PasswordHash string `gorm:"not null" json:"-"`
    Phone        string `gorm:"uniqueIndex" json:"phone"`
    RoleID       uint   `json:"role_id"`  // FK to Role
    Role         Role   `gorm:"foreignKey:RoleID"`
    UpiPin    string // Add this field for UPI PI
    ID        uint   `gorm:"primaryKey"`
    Balance  float64 `json:"balance"` // ✅ Add this
	Trades   []Trade `gorm:"foreignKey:UserID"` // ✅ Add this
    UserID    uint      `json:"user_id"` // Foreign key
	Password string `json:"password" gorm:"not null"` // ✅ Make sure this exists
    PhoneNumber   string
    // Relations
    Dashboards       []Dashboard      `gorm:"foreignKey:UserID"`
    Queries          []Query          `gorm:"foreignKey:StudentID"`
    StudentProgress  []StudentProgress `gorm:"foreignKey:StudentID"`
    CallSessions     []CallSession    `gorm:"foreignKey:CallerID"`
    
}

// Method to check if profile exists
func (u *User) ProfileExists() bool {
	return u.Name != "" && u.Email != ""
}
