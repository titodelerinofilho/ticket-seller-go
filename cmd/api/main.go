package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/buy", BuyTicket)

	log.Println("API rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
