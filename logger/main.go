package main

import (
	"bytes"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = ch.QueueDeclare("logs", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	msg, err := ch.Consume("logs", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalln(err.Error())
	}

	client.Indices.Create("ind")

	for m := range msg {
		log.Println(string(m.Body))
		client.Index("ind", bytes.NewReader(m.Body))
	}
}
