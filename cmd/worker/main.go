package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/titodelerinofilho/ingressos-vendas/config"
	"github.com/titodelerinofilho/ingressos-vendas/internal/services"
)

func main() {
	conn, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, _ := ch.QueueDeclare("ticket_orders", false, false, false, false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var order map[string]interface{}
			json.Unmarshal(msg.Body, &order)
			services.ProcessOrder(order)
		}
	}()

	fmt.Println("Worker esperando pedidos...")
	<-forever
}
