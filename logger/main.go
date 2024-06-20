package main

import (
	"log"

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

	for m := range msg {
		log.Println(string(m.Body))
	}
}
