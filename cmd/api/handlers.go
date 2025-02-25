package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/titodelerinofilho/ingressos-vendas/config"
)

var ctx = context.Background()
var redisDB = config.ConnectRedis()

func BuyTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	ticketsAvailable, _ := redisDB.Get(ctx, "tickets_available").Int()
	if ticketsAvailable <= 0 {
		http.Error(w, "No tickets available", http.StatusBadRequest)
		return
	}

	redisDB.Decr(ctx, "tickets_available")

	order := map[string]interface{}{"user_id": userID, "ticket": 1}
	PublishToQueue(order)

	fmt.Fprintf(w, "Ticket reserved for user %s", userID)
}
