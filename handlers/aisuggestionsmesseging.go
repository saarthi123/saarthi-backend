package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/saarthi123/saarthi-backend/models"
)

// Mock AI suggestions for demo. Replace with real AI integration.
func getSuggestionsAI(text, context string) []string {
    baseSuggestions := map[string][]string{
        "school":  {"Good job!", "Keep it up!", "Well done!"},
        "college": {"Interesting point.", "Could you elaborate?", "Thanks for sharing."},
        "office":  {"Let's schedule a meeting.", "Please review.", "Approved."},
    }

    suggestions, ok := baseSuggestions[context]
    if !ok {
        suggestions = []string{"Sorry, no suggestions available."}
    }
    // In reality, you’d call an AI or NLP service here with text+context.

    return suggestions
}

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
    var req models.SuggestionRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    suggestions := getSuggestionsAI(req.Text, req.Context)

    res := models.SuggestionResponse{Suggestions: suggestions}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
