package main

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/titodelerinofilho/ingressos-vendas/config"
)

func PublishToQueue(order map[string]interface{}) {
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

	body, _ := json.Marshal(order)
	ch.Publish("", q.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
