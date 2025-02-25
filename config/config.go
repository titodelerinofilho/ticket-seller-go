package config

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rabbitmq/amqp091-go"
)

// Conexão com Redis
func ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
}

// Conexão com RabbitMQ
func ConnectRabbitMQ() (*amqp091.Connection, error) {
	return amqp091.Dial(os.Getenv("RABBITMQ_URL"))
}
