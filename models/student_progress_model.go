package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentProgress struct {
    gorm.Model
    StudentID  uint    `json:"student_id"` // FK to User.ID
    Student    User    `gorm:"foreignKey:StudentID"`
    Name       string  `json:"name"`
    Course     string  `json:"course"`
    Status     string  `json:"status"`     // completed, in-progress, dropped
    Completion float64 `json:"completion"` // percentage 0-100
}


type AttendanceRecord struct {
    ID     uint      `gorm:"primaryKey" json:"id"`
    UserID string    `json:"user_id"`
    Date   time.Time `json:"date"`
    Status string    `json:"status"`
    Mode   string    `json:"mode"`
}

type CampusAttendance struct {
	gorm.Model
	StudentName string    `json:"student_name"`
	Course      string    `json:"course"`
	AttendDate  string    `json:"attend_date"` // store as "2006-01-02" format
	Status      string    `json:"status"`      // e.g., "Present" or "Absent"
}


type CareerPath struct {
	Title    string   `json:"title"`
	Skills   []string `json:"skills"`
	Courses  []string `json:"courses"`
	Colleges []string `json:"colleges"`
	Jobs     []string `json:"jobs"`
	Salary   string   `json:"salary"`
	Interest string   `json:"interest"` 
}


type Course struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Duration    string `json:"duration"` // e.g., "4 weeks"
	Level       string `json:"level"`    // e.g., "Beginner", "Advanced"
	ID       string `json:"id"`
	Content  string `json:"content"`
}

type StudentQuery struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Reply    string `json:"reply,omitempty"`
}


