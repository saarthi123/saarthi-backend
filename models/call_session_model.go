package models
import "time"

type CallSession struct {
    ID                  string `json:"id"`
    CallerID            string `json:"caller_id"`
    ReceiverID          string `json:"receiver_id"`
    IsMuted             bool   `json:"is_muted"`
    SpeakerOn           bool   `json:"speaker_on"`
    AINoiseCancellation bool   `json:"ai_noise_cancellation"`
    DurationSeconds     int    `json:"duration_seconds"`
	CreatedAt           time.Time `json:"created_at"`
}
