package models

type ToastMessage struct {
    Message string `json:"message"`
    Type    string `json:"type"`
}