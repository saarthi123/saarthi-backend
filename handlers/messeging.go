// handlers/settings.go
package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/saarthi123/saarthi-backend/models"
	"sync"
)

var (
	currentSettings = models.UserSettings{
		ProfileName:         "User",
		Status:              "Available",
		Theme:               "neon",
		NotificationsEnabled: true,
		E2EEncryption:        true,
	}
	settingsMutex sync.Mutex
)

func GetUserSettings(w http.ResponseWriter, r *http.Request) {
	settingsMutex.Lock()
	defer settingsMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentSettings)
}

func UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	var updated models.UserSettings

	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	settingsMutex.Lock()
	currentSettings = updated
	settingsMutex.Unlock()

	w.WriteHeader(http.StatusOK)
}
