package models

type Message struct {
    ID      uint   `json:"id" gorm:"primaryKey"`
    Content string `json:"content"`
    Sender  string `json:"sender"`
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
}



type SuggestionRequest struct {
    Text    string `json:"text"`
    Context string `json:"context"`
}

type SuggestionResponse struct {
    Suggestions []string `json:"suggestions"`
}



type UserSettings struct {
    ProfileName          string `json:"profileName"`
    Status               string `json:"status"`
    Theme                string `json:"theme"`
    NotificationsEnabled bool   `json:"notificationsEnabled"`
    E2EEncryption        bool   `json:"e2eEncryption"`
}


