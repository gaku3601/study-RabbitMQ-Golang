package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var amqpURI string = "amqp://localhost:5672"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to MQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	sendMessage(channel, "my-queue", "{'test':'ababaab1'}")
	sendMessage(channel, "my-queue", "{'test':'ababaab2'}")
	sendMessage(channel, "my-queue", "{'test':'ababaab3'}")
}

func sendMessage(channel *amqp.Channel, queueName string, payload string) {
	err := channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		})
	failOnError(err, "Failed to publish a message")
}
