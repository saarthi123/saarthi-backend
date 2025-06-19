package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	simpleClients   = make(map[*websocket.Conn]bool)
	simpleClientsMu sync.Mutex
)

var simpleUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins
	},
}

func SimpleSocketHandler(c *gin.Context) {
	conn, err := simpleUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	registerSimpleClient(conn)
	defer unregisterSimpleClient(conn)

	for {
		var msg map[string]interface{}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}
		broadcastSimple(msg, conn)
	}
}

func registerSimpleClient(conn *websocket.Conn) {
	simpleClientsMu.Lock()
	defer simpleClientsMu.Unlock()
	simpleClients[conn] = true
}

func unregisterSimpleClient(conn *websocket.Conn) {
	simpleClientsMu.Lock()
	defer simpleClientsMu.Unlock()
	delete(simpleClients, conn)
}

func broadcastSimple(message map[string]interface{}, sender *websocket.Conn) {
	simpleClientsMu.Lock()
	defer simpleClientsMu.Unlock()

	for client := range simpleClients {
		if client == sender {
			continue
		}
		if err := client.WriteJSON(message); err != nil {
			log.Println("broadcast error:", err)
		}
	}
}
