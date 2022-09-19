package main

import (
	"context"
	"github.com/GonnaFlyMethod/RabbitMQ-demo/demo/common"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello world"

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("can't publish to RabbitMQ, err: %v", err)
	}

	log.Printf("successfully sent message %q to a queue %q", body, q.Name)
}
