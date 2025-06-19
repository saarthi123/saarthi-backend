package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/saarthi123/saarthi-backend/models"
)

func PlaceTradeOrder(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var order models.TradeOrder
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Here you could add logic to:
    // - Validate the order
    // - Save it to DB
    // - Send it to trade execution system

    // For demo, just return order received + a confirmation message
    response := map[string]interface{}{
        "message": "Order placed successfully",
        "order":   order,
        "orderId": "ORD123456789", // Example order id
    }

    json.NewEncoder(w).Encode(response)
}
