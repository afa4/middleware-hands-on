package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Response struct {
	Message string
}

type Request struct {
	Name string
}

type HttpHandler struct {
	rabbitmq *amqp.Channel
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := Request{}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.publishName(body.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := Response{Message: "Message sent"}
	bytes, err = json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(bytes)
}

func (h *HttpHandler) publishName(name string) error {
	return h.rabbitmq.Publish(
		EXCHANGE_NAME, // exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(name),
		})
}
