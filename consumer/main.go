package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var EXCHANGE_NAME = "exchange"
var QUEUE_NAME = "queue"

func main() {
	var conn *amqp.Connection
	var err error
	fmt.Println("CONSUMER Waiting rabbit mq setup")
	for {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
	}
	fmt.Println("CONSUMER started")

	defer conn.Close()

	rabbitmq, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer rabbitmq.Close()

	err = rabbitmq.ExchangeDeclare(
		EXCHANGE_NAME, // name
		"fanout",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	failOnError(err, "Failed to declare an exchange")

	_, err = rabbitmq.QueueDeclare(
		QUEUE_NAME, // name
		false,      // durable
		false,      // delete when unused
		true,       // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = rabbitmq.QueueBind(
		QUEUE_NAME,    // queue name
		"",            // routing key
		EXCHANGE_NAME, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := rabbitmq.Consume(
		QUEUE_NAME, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		name := string(msg.Body)
		fmt.Printf("Hello %s\n", name)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
