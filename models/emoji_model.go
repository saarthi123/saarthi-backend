package models

import "gorm.io/gorm"

type EmojiCategory struct {
	gorm.Model
	Name   string  `gorm:"unique;not null" json:"name"`
	Emojis []Emoji `gorm:"foreignKey:CategoryID"`
}

type Emoji struct {
	gorm.Model
	Symbol     string `gorm:"not null" json:"symbol"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}
