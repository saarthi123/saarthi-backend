package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID string
	Conn   *websocket.Conn
}

var (
	clients   = make(map[string]*Client)
	clientsMu sync.Mutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins (secure with auth in production)
	},
}

func VideoSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	var userID string
	for {
		var msg map[string]interface{}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}

		event := msg["event"].(string)
		data := msg["data"].(map[string]interface{})

		switch event {
		case "video-join":
			userID = data["userId"].(string)
			registerClient(userID, conn)
			broadcast("user-joined", data, userID)

		case "video-end":
			userID = data["userId"].(string)
			broadcast("user-left", data, userID)
			unregisterClient(userID)

		case "ai-assist-toggle":
			log.Println("AI Assist toggled by", data["userId"])
			broadcast("ai-assist-toggle", data, userID)

		default:
			log.Println("Unknown event:", event)
		}
	}

	unregisterClient(userID)
}

func registerClient(userID string, conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[userID] = &Client{UserID: userID, Conn: conn}
}

func unregisterClient(userID string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, userID)
}

func broadcast(event string, data interface{}, excludeUserID string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for id, client := range clients {
		if id == excludeUserID {
			continue
		}
		if err := client.Conn.WriteJSON(map[string]interface{}{
			"event": event,
			"data":  data,
		}); err != nil {
			log.Printf("broadcast error to user %s: %v", id, err)
		}
	}
}
