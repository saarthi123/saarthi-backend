package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	ID          string       `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"unique;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`

}

type Permission struct {
	gorm.Model
	RoleID  uint   `gorm:"index" json:"role_id"` // Foreign key to Role
	Name    string `gorm:"not null" json:"name"`
	Allowed bool   `json:"allowed"`
	ID   string `gorm:"primaryKey" json:"id"`
}
