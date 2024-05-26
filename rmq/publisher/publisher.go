package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// publisher
	conn, err := amqp.Dial("amqp://user:pass@localhost:5671/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
}
