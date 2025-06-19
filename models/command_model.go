package models

type Command struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Phrase      string `gorm:"unique;not null" json:"phrase"`
    Action      string `json:"action"`
    Description string `json:"description"`
	Command     string `json:"command"`
}
