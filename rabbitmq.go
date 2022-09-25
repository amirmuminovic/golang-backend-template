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

func createChannel(rmq *amqp.Connection) *amqp.Channel {
	channel, err := rmq.Channel()

	if err != nil {
		log.Fatalln(err)
	}

	return channel
}

func createExchange(channel *amqp.Channel, rmqConfig RabbitMQConfig) {
	err := channel.ExchangeDeclare(
		rmqConfig.exchangeName,
		rmqConfig.exchangeType,
		rmqConfig.exchangeDurable,
		rmqConfig.exchangeAutoDelete,
		rmqConfig.exchangeInternal,
		rmqConfig.exchangeNoWait,
		rmqConfig.exchangeArguments,
	)

	if err != nil {
		log.Fatalln(err)
	}
}

func publishMessage(channel *amqp.Channel, rmqConfig RabbitMQConfig, body []byte) {
	publishingConfig := amqp.Publishing{
		ContentType: rmqConfig.messageContentType,
		Body:        body,
	}

	err := channel.Publish(
		rmqConfig.exchangeName,
		rmqConfig.routingKey,
		rmqConfig.messageMandatory,
		rmqConfig.messageImmediate,
		publishingConfig,
	)

	if err != nil {
		log.Fatalln(err)
	}
}
