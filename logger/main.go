package main

import (
	"bytes"
	"log"
	"logger/pool"
	"sync"

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

	p, err := pool.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	client := p.GetWorker()
	client.Indices.Create("ind")
	p.ReleaseWorker(client)

	// cancel := make(chan bool)
	forever := make(chan bool)

	go func() {
		wg := sync.WaitGroup{}
		for {
			m := <-msg
			wg.Add(1)
			go func() {
				worker := p.GetWorker()
				if worker == nil {
					log.Println("worker is nil")
				} else {
					log.Println(string(m.Body))
					_, err := worker.Index("ind", bytes.NewReader(m.Body))
					defer p.ReleaseWorker(worker)
					if err != nil {
						log.Println(err.Error())
					}
				}
				wg.Done()
			}()
		}
	}()
	<-forever
}
