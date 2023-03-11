package main

import (
	"encoding/json"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Response struct {
	Message string
}

type HttpHandler struct {
	rabbitmq *amqp.Channel
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := "from request"
	err := h.publishName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := Response{Message: "Message sent"}
	bytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(bytes)
}

func (h *HttpHandler) publishName(name string) error {
	return h.rabbitmq.Publish(
		"default", // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(name),
		})
}
