// handlers/notification.go
package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/saarthi123/saarthi-backend/models"
	"sync"
)

var (
	notifications []models.Notification
	mutex         sync.Mutex
)

func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

func MarkNotificationRead(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mutex.Lock()
	defer mutex.Unlock()

	if id != "" {
		for i, n := range notifications {
			if n.ID == id {
				notifications[i].Read = true
				break
			}
		}
	} else {
		for i := range notifications {
			notifications[i].Read = true
		}
	}

	w.WriteHeader(http.StatusOK)
}
