package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/saarthi123/saarthi-backend/models"
	"github.com/saarthi123/saarthi-backend/utils"
)

// Client represents a single WebSocket connection
type Client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	UserID string // used for notification targeting
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	clients   = make(map[*Client]bool)
	broadcast = make(chan []byte)              // for public messages
	notify    = make(chan models.Notification) // for targeted push notifications
	mutex     sync.Mutex
)

const pingPeriod = 30 * time.Second

// ========================== WebSocket Entry Points ==========================

// Authenticated connection for notifications (token required)
func HandleNotificationWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	userID, err := utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}

	client := &Client{
		Conn:   conn,
		UserID: userID,
		Send:   make(chan []byte, 256),
	}

	registerClient(client)
	go handleOutgoing(client)
	go startHeartbeat(client)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	unregisterClient(client)
}

// Open public WebSocket (no token)
func HandlePublicWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	registerClient(client)
	go handleOutgoing(client)
	go startHeartbeat(client)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		broadcast <- msg
	}

	unregisterClient(client)
}

// ========================== Client Lifecycle ==========================

func registerClient(c *Client) {
	mutex.Lock()
	defer mutex.Unlock()
	clients[c] = true
}

func unregisterClient(c *Client) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, c)
	close(c.Send)
	c.Conn.Close()
}

// ========================== Heartbeat Ping ==========================

func startHeartbeat(client *Client) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := client.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				unregisterClient(client)
				return
			}
		}
	}
}

// ========================== Outgoing Messages ==========================

func handleOutgoing(client *Client) {
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
	client.Conn.Close()
}

// ========================== Broadcast Loop ==========================

func StartBroadcaster() {
	for {
		select {
		case msg := <-broadcast:
			mutex.Lock()
			for client := range clients {
				select {
				case client.Send <- msg:
				default:
					unregisterClient(client)
				}
			}
			mutex.Unlock()

		case notif := <-notify:
			data, _ := json.Marshal(notif)
			mutex.Lock()
			for client := range clients {
				// Replace 'UserID' with the actual field name in models.Notification that identifies the user
				if notif.UserID == "" || notif.UserID == client.UserID {
					select {
					case client.Send <- data:
					default:
						unregisterClient(client)
					}
				}
			}
			mutex.Unlock()
		}
	}
}

// ========================== External Push ==========================

func PushNotification(n models.Notification) {
	notify <- n
}
