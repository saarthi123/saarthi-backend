package models

type LiveSession struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Speaker  string `json:"speaker"`
	DateTime string `json:"date_time"`
	Topic    string `json:"topic"`
	Instructor string `json:"instructor"`
	StartTime string `json:"startTime"`
}




