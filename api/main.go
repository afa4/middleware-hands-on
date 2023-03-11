package main

import (
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	var conn *amqp.Connection
	var err error
	fmt.Println("Waiting rabbit mq setup")
	for {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
	}

	defer conn.Close()

	rabbitmq, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer rabbitmq.Close()

	err = rabbitmq.ExchangeDeclare(
		"default", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	fmt.Println("Server started")
	StartServerWithRabbitmqInstance(rabbitmq)
}

func StartServerWithRabbitmqInstance(rabbitmq *amqp.Channel) {
	http.Handle("/", &HttpHandler{rabbitmq})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
