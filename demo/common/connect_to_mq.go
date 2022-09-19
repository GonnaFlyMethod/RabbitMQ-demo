package common

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func TryToConnect() (*amqp.Connection, error) {
	const tries = 10

	for i := 0; i < tries; i++ {
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("can't connect to RabbitMQ")
}
