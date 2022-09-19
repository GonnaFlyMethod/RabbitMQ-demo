package main

import (
	"github.com/GonnaFlyMethod/RabbitMQ-demo/demo/common"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	conn, err := common.TryToConnect()
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *amqp.Connection) {
		if err := conn.Close(); err != nil {
			log.Printf("can't close connection with RabbitMQ, err: %v", err)
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("can't open channel, err: %v", err)
	}

	defer func(ch *amqp.Channel) {
		if err := ch.Close(); err != nil {
			log.Printf("can't close channel, err: %v", err)
		}
	}(ch)

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("can't declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("can't set up channel for consuming, err: %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("received message: %s", d.Body)
		}
	}()

	var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
