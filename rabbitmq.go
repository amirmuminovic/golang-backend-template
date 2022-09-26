package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func connectToRMQ(ac AppConfig) *amqp.Connection {
	conn, err := amqp.Dial(ac.rabbitmqConnectionString)

	if err != nil {
		log.Fatalln(err)
	}

	return conn
}
