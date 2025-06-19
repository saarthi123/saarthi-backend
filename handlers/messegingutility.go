package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
    "github.com/saarthi123/saarthi-backend/models"
)

var commands = []models.Command{
    {Command: "/help", Description: "Show help information"},
    {Command: "/clear", Description: "Clear chat screen"},
    {Command: "/gif", Description: "Search for GIFs"},
    {Command: "/mute", Description: "Mute notifications"},
    {Command: "/unmute", Description: "Unmute notifications"},
    // Add more commands here
}

// GetCommands returns list of commands matching the prefix query param "q"
func GetCommands(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    var matches []models.Command

    if query == "" {
        matches = commands
    } else {
        for _, cmd := range commands {
            if strings.HasPrefix(cmd.Command, query) {
                matches = append(matches, cmd)
            }
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(matches)
}
